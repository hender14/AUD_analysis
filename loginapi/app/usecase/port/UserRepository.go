package port

type UserRepository interface {
	QueryEmail() (err error)
	CreateAccoount() (err error)
}
