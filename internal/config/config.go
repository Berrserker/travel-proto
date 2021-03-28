package config

import (
	`time`
	
	`github.com/kkyr/fig`
	`github.com/pkg/errors`
)

type Config struct {
	
	StaticStorage string `fig:"StaticStorage,required"`
	StorageService string `fig:"StorageService,required"`
	
	JWTsecret string `fig:"JWTsecret,required"`
	Http struct {
		Host string `fig:"Host,required"`
		Port string `fig:"Port,required"`
	}
	
	DB struct {
		PostgresMaster string `fig:"PostgresMaster,required"`
		ReadTimeOut time.Duration `fig:"ReadTimeOut,required"`
		WriteTimeOut time.Duration `fig:"WriteTimeOut,required"`
	}
	
}

func (s *Config) Init() error {
	if err := fig.Load(s); err != nil {
		return errors.Wrap(err, "Cannot read a config")
	}
	
	return nil
}
