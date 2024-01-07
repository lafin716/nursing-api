package api

type UserHttpApi interface {
}

type userHttpApi struct {
}

func NewUserHttpApi() UserHttpApi {
	return &userHttpApi{}
}
