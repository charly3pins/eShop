package base

type Repository interface {
	NextIdentity() Identity
}
