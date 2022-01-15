package core

type Resolver interface {
	Resolve(url string) (*Material, error)
}
