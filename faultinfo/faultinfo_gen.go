package faultinfo
// DO NOT EDITS. Automatically generated.

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	*mux.Router
}

func New() *Server {
	s := &Server{
		Router: mux.NewRouter(),
	}
	s.bindRoutes()
	return s
}

func (s *Server) Run(l string) error {
	return http.ListenAndServe(l, s.Router)
}

func (s *Server) bindRoutes() {
	s.Router.HandleFunc("/template", gettemplate).Method("GET")
	s.Router.HandleFunc("/template", posttemplate).Method("POST")
	s.Router.HandleFunc("/template/{id}", gettemplateid).Method("GET")
	s.Router.HandleFunc("/template/{id}", deletetemplateid).Method("DELETE")
	s.Router.HandleFunc("/info", getinfo).Method("GET")
	s.Router.HandleFunc("/info", postinfo).Method("POST")
	s.Router.HandleFunc("/info/{id}", getinfoid).Method("GET")
	s.Router.HandleFunc("/info/{id}", putinfoid).Method("PUT")
	s.Router.HandleFunc("/info/{id}", deleteinfoid).Method("DELETE")
	s.Router.HandleFunc("/info/{id}/comments", getinfoidcomments).Method("GET")
	s.Router.HandleFunc("/info/{id}/comments", postinfoidcomments).Method("POST")
	s.Router.HandleFunc("/info/{info_id}/comments/{comment_id}", getinfoinfo_idcommentscomment_id).Method("GET")
	s.Router.HandleFunc("/info/{info_id}/comments/{comment_id}", putinfoinfo_idcommentscomment_id).Method("PUT")
	s.Router.HandleFunc("/info/{info_id}/comments/{comment_id}", deleteinfoinfo_idcommentscomment_id).Method("DELETE")
	s.Router.HandleFunc("/type", gettype).Method("GET")
	s.Router.HandleFunc("/type", posttype).Method("POST")
	s.Router.HandleFunc("/type", deletetype).Method("DELETE")
}
