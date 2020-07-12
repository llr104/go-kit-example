package service

type IApi interface {
	Hello(string) string
	Echo(string) string
}

type Api struct {

}

func (this* Api) Hello(h string) string {
	return h
}

func (this* Api) Echo(e string) string {
	return e
}