package main

import (
	`context`
	`fmt`
	`net/http`
	`os`
	`time`
	
	`github.com/form3tech-oss/jwt-go`
	`github.com/gorilla/mux`
	`github.com/gorilla/handlers`
	`github.com/sirupsen/logrus`
	
	`travel/internal/config`
	`travel/internal/db`
	
	"github.com/rs/cors"
	
	"github.com/auth0/go-jwt-middleware"
	httpHandler `travel/internal/http`
)

func main() {
	if err := bootstrap(); err != nil {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.Fatal(err)
	}
}

func bootstrap() error {
	globalCtx, globalCancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(globalCtx, time.Minute)
	defer cancel()
	
	var core config.Config
	
	if err := core.Init(); err != nil {
		return err
	}
	
	// init db service
	dba, err := db.New(globalCtx, core.DB.PostgresMaster)
	if err != nil {
		return err
	}
	
	// init storage
	// TODO: setup storage as microservice
	storage, err := storage.New()
	if err != nil {
		return err
	}
	
	// init http service
	httpService, err := httpHandler.New()
	if err != nil {
		return err
	}
	
	// cors
	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})
	
	// JWT
	var key jwt.Keyfunc
	key = func(token *jwt.Token) (interface{}, error) {
		return []byte(core.JWTsecret), nil
	}
	
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: key,
		SigningMethod: jwt.SigningMethodHS256,
	})
	
	// init router
	router := mux.NewRouter()
	router.Handle("/status", jwtMiddleware.Handler(http.HandlerFunc(httpService.Status))).Methods("GET")
	router.HandleFunc("/auth", httpService.Auth).Methods("POST")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(core.StaticStorage)))
	fmt.Println("Service start at port: ", core.Http.Port)
	err := http.ListenAndServe(":"+core.Http.Port, corsWrapper.Handler(handlers.LoggingHandler(os.Stdout, router)))
	if err != nil {
		fmt.Print(err)
	}
	return nil
}