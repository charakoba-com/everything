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
	s.Router.HandleFunc("/template", gettemplate).Methods("GET")
	s.Router.HandleFunc("/template", posttemplate).Methods("POST")
	s.Router.HandleFunc("/template/{id}", gettemplateid).Methods("GET")
	s.Router.HandleFunc("/template/{id}", deletetemplateid).Methods("DELETE")
	s.Router.HandleFunc("/info", getinfo).Methods("GET")
	s.Router.HandleFunc("/info", postinfo).Methods("POST")
	s.Router.HandleFunc("/info/{id}", getinfoid).Methods("GET")
	s.Router.HandleFunc("/info/{id}", putinfoid).Methods("PUT")
	s.Router.HandleFunc("/info/{id}", deleteinfoid).Methods("DELETE")
	s.Router.HandleFunc("/info/{id}/comments", getinfoidcomments).Methods("GET")
	s.Router.HandleFunc("/info/{id}/comments", postinfoidcomments).Methods("POST")
	s.Router.HandleFunc("/info/{info_id}/comments/{comment_id}", getinfoinfo_idcommentscomment_id).Methods("GET")
	s.Router.HandleFunc("/info/{info_id}/comments/{comment_id}", putinfoinfo_idcommentscomment_id).Methods("PUT")
	s.Router.HandleFunc("/info/{info_id}/comments/{comment_id}", deleteinfoinfo_idcommentscomment_id).Methods("DELETE")
	s.Router.HandleFunc("/type", gettype).Methods("GET")
	s.Router.HandleFunc("/type", posttype).Methods("POST")
	s.Router.HandleFunc("/type", deletetype).Methods("DELETE")
}
