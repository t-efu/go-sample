package handler

import (
	"html/template"
	"log"
	"net/http"
)

var path = "application/interfaces/web/handler/views/"

// IndexHandler - index handler
type IndexHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type indexHandler struct {
}

// NewIndexHandler - new user handler
func NewIndexHandler() IndexHandler {
	return &indexHandler{}
}

func (h *indexHandler) Get(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles(path+"base.html", path+"index.html"))
	message := "sample message"
	err := tpl.ExecuteTemplate(w, "base", map[string]interface{}{
		"Message": message,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		log.Println(err)
		return
	}
	return
}
