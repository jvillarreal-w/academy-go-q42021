package icontext

type IContext interface {
	JSON(code int, i interface{}) error
	Bind(i interface{}) error
	Param(name string) string
	QueryParam(name string) string
}
