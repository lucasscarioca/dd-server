package port

type TokenProvider interface {
	Create(email string) (string, error)
}

type AuthService interface {
	Login(email, password string) (string, error)
	Register(name, email, password string) (string, error)
}
