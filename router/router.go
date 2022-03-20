package router

import (
	"go-board/handler"

	"github.com/gorilla/mux"
)

var R *mux.Router

func Init() {
	R = mux.NewRouter()
	R.Methods("GET").Path("/").HandlerFunc(handler.TestHandler) // hello를 출력하는 테스트용 handler
	R.Methods("GET").Path("/articles").HandlerFunc(handler.GetArticles)
}
