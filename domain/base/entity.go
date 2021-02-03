package base

// TODO check how to make the entities implement this interface
type Entity interface {
	Validate() error
}
