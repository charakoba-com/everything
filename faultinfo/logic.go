package faultinfo

import (
	"github.com/charakoba-com/everything/faultinfo/input"
	"github.com/charakoba-com/everything/faultinfo/output"
	"github.com/charakoba-com/everything/faultinfo/schemas"
)

func GetInformationTemplateDetail(id string) (out schemas.Template, err error) {
	return
}

func DeleteInformationTemplate(id string) (err error) {
	return
}

func GetInformationList() (out []schemas.Information, err error) {
	return
}

func CreateInformation(req input.CreateInformation) (out output.InformationCreated, err error) {
	return
}

func GetInformationDetail(id string) (out schemas.InformationDetail, err error) {
	return
}

func UpdateInformation(id string, req input.CreateInformation) (err error) {
	return
}

func DeleteInformation(id string) (err error) {
	return
}

func GetCommentList(id string) (out []schemas.Comment, err error) {
	return
}

func CreateComment(id string, req input.CreateComment) (out output.CommentCreated, err error) {
	return
}

func GetCommentDetail(infoId string, commentId string) (out schemas.CommentDetail, err error) {
	return
}

func UpdateComment(infoId string, commentId string, req input.CreateComment) (err error) {
	return
}

func DeleteComment(infoId string, commentId string) (err error) {
	return
}

func GetInformationTypelist() (out []string, err error) {
	return
}

func CreateInformationType(req input.CreateInfoType) (err error) {
	return
}

func DeleteInformationType(typ string) (err error) {
	return
}

func GetInformationTemplateIDList() (out []string, err error) {
	return
}

func CreateInformationTemplate(req input.CreateTemplate) (err error) {
	return
}
