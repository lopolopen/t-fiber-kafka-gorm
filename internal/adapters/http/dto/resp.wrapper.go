package dto

type Nothing struct{}

type Unknown struct{}

type RespNothing struct {
	Resp[Nothing]
}

type SwaggerUser struct {
	ID       uint
	Name     string
	Birthday string `example:"2006-01-02"`
	Age      int
}

type RespUsers struct {
	Resp[[]SwaggerUser]
}
