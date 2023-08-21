package types

type Response struct {
	Code int
	Body any // can be data in read apis or string messages in create, update and delete apis
}
