package faultinfo

// Code generated genserver.go. DO NOT EDIT.

import (
	"net/http"

	"github.com/gorilla/mux"
)

func bindroutes(router *mux.Router) {
	router.HandleFunc("/template/{id}", GetInformationTemplateDetailHandler).Methods(http.MethodGet)
	router.HandleFunc("/template/{id}", DeleteInformationTemplateHandler).Methods(http.MethodDelete)
	router.HandleFunc("/info", GetInformationListHandler).Methods(http.MethodGet)
	router.HandleFunc("/info", CreateInformationHandler).Methods(http.MethodPost)
	router.HandleFunc("/info/{id}", GetInformationDetailHandler).Methods(http.MethodGet)
	router.HandleFunc("/info/{id}", UpdateInformationHandler).Methods(http.MethodPut)
	router.HandleFunc("/info/{id}", DeleteInformationHandler).Methods(http.MethodDelete)
	router.HandleFunc("/info/{id}/comments", GetCommentListHandler).Methods(http.MethodGet)
	router.HandleFunc("/info/{id}/comments", CreateCommentHandler).Methods(http.MethodPost)
	router.HandleFunc("/info/{info_id}/comments/{comment_id}", GetCommentDetailHandler).Methods(http.MethodGet)
	router.HandleFunc("/info/{info_id}/comments/{comment_id}", UpdateCommentHandler).Methods(http.MethodPut)
	router.HandleFunc("/info/{info_id}/comments/{comment_id}", DeleteCommentHandler).Methods(http.MethodDelete)
	router.HandleFunc("/type", GetInformationTypelistHandler).Methods(http.MethodGet)
	router.HandleFunc("/type", CreateInformationTypeHandler).Methods(http.MethodPost)
	router.HandleFunc("/type/{typ}", DeleteInformationTypeHandler).Methods(http.MethodDelete)
	router.HandleFunc("/template", GetInformationTemplateIDListHandler).Methods(http.MethodGet)
	router.HandleFunc("/template", CreateInformationTemplateHandler).Methods(http.MethodPost)
}
