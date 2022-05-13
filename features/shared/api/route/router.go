package route

import "github.com/gorilla/mux"

type router struct {
	muxRouter *mux.Router
}

func NewRouter() IRouter {
	muxRouter := mux.NewRouter().StrictSlash(true)
	return router{
		muxRouter: muxRouter,
	}
}

func (r router) WithRoute(route route) IRouter {
	r.muxRouter.HandleFunc(route.name, route.handler)
	return r
}

func (r router) GetRouter() *mux.Router {
	return r.muxRouter
}
