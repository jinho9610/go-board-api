package router

import (
	"github.com/gorilla/mux"
	"go-board/handler"
)

var R *mux.Router

func Init() {
	R = mux.NewRouter()
	R.Methods("GET").Path("/").HandlerFunc(handler.TestHandler)
	R.Methods("GET").Path("/users").HandlerFunc(handler.GetUsers)
}
