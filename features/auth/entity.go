package auth

type User struct {
	ID       uint
	Email    string
	Password string
}

type Business interface {
	LoginUsers(data User, input string) (interface{}, error)
}

type Data interface {
	SelectLogin(data User, input string) (interface{}, error)
}
