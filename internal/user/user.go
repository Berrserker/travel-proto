package user

type Account struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"adress"`
	Token    string `json:"token"`
	ApiToken string `json:"api_token"`
	Role     string `json:"role"`
	Phone    string `json:"phone"`
	Manager  string `json:"manager"`
}

func (s *Account) Login() (*Account, error)
