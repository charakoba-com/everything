package faultinfo

// Code generated genserver.go. DO NOT EDIT.

import (
	"net/http"

	"github.com/gorilla/mux"
)

func bindroutes(router *mux.Router) {
	router.HandlerFunc("/type/{typ}", DeleteInformationType).Methods(http.MethodDELETE)
	router.HandlerFunc("/template", GetInformationTemplateIDList).Methods(http.MethodGET)
	router.HandlerFunc("/template", CreateInformationTemplate).Methods(http.MethodPOST)
	router.HandlerFunc("/template/{id}", GetInformationTemplateDetail).Methods(http.MethodGET)
	router.HandlerFunc("/template/{id}", DeleteInformationTemplate).Methods(http.MethodDELETE)
	router.HandlerFunc("/info", GetInformationList).Methods(http.MethodGET)
	router.HandlerFunc("/info", CreateInformation).Methods(http.MethodPOST)
	router.HandlerFunc("/info/{id}", GetInformationDetail).Methods(http.MethodGET)
	router.HandlerFunc("/info/{id}", UpdateInformation).Methods(http.MethodPUT)
	router.HandlerFunc("/info/{id}", DeleteInformation).Methods(http.MethodDELETE)
	router.HandlerFunc("/info/{id}/comments", GetCommentList).Methods(http.MethodGET)
	router.HandlerFunc("/info/{id}/comments", CreateComment).Methods(http.MethodPOST)
	router.HandlerFunc("/info/{info_id}/comments/{comment_id}", GetCommentDetail).Methods(http.MethodGET)
	router.HandlerFunc("/info/{info_id}/comments/{comment_id}", UpdateComment).Methods(http.MethodPUT)
	router.HandlerFunc("/info/{info_id}/comments/{comment_id}", DeleteComment).Methods(http.MethodDELETE)
	router.HandlerFunc("/type", GetInformationTypelist).Methods(http.MethodGET)
	router.HandlerFunc("/type", CreateInformationType).Methods(http.MethodPOST)
}
