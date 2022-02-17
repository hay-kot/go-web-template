package services

type contextKeys struct {
	name string
}

var (
	ContextUser = &contextKeys{name: "user"}
)
