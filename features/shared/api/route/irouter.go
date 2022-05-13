package route

import "github.com/gorilla/mux"

type IRouter interface {
	WithRoute(route route) IRouter
	GetRouter() *mux.Router
}
