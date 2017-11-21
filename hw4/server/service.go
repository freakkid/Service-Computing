package server

import (
	"net/http"

	"github.com/unrolled/render"
)

type TodoList struct {
	Username string
	Todos    []string
	Message  string
}

func apiTestHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var todoList TodoList = TodoList{
			Username: "Hello",
			Todos:    []string{"study", "sleep", "eat"},
			Message:  "success",
		}
		formatter.JSON(w, http.StatusOK, todoList)
	}
}
