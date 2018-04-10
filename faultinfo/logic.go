package faultinfo

import (
	"context"

	"github.com/charakoba-com/everything/faultinfo/domain/repository"
	"github.com/charakoba-com/everything/faultinfo/input"
	"github.com/charakoba-com/everything/faultinfo/output"
	"github.com/charakoba-com/everything/faultinfo/schemas"
)

func GetInformationTemplateIDList(ctx context.Context) (out []string, err error) {
	return
}

func CreateInformationTemplate(ctx context.Context, req input.CreateTemplate) (err error) {
	return
}

func GetInformationTemplateDetail(ctx context.Context, id string) (out *schemas.Template, err error) {
	return
}

func DeleteInformationTemplate(ctx context.Context, id string) (err error) {
	return
}

func GetInformationList(ctx context.Context) (out []*schemas.Information, err error) {
	repo, err := repository.NewInformationRepository()
	if err != nil {
		return nil, err
	}
	list, err := repo.Listup(ctx)
	if err != nil {
		return nil, err
	}
	for _, info := range list {
		out = append(out, info.ToInfo())
	}
	return
}

func CreateInformation(ctx context.Context, req input.CreateInformation) (out *output.InformationCreated, err error) {
	return
}

func GetInformationDetail(ctx context.Context, id string) (out *schemas.InformationDetail, err error) {
	repo, err := repository.NewInformationRepository()
	if err != nil {
		return nil, err
	}
	info, err := repo.FindByPK(ctx, id)
	if err != nil {
		return nil, err
	}
	out = info.ToInformationDetail()
	return
}

func UpdateInformation(ctx context.Context, id string, req input.CreateInformation) (err error) {
	return
}

func DeleteInformation(ctx context.Context, id string) (err error) {
	return
}

func GetCommentList(ctx context.Context, id string) (out []*schemas.Comment, err error) {
	return
}

func CreateComment(ctx context.Context, id string, req input.CreateComment) (out *output.CommentCreated, err error) {
	return
}

func GetCommentDetail(ctx context.Context, infoId string, commentId string) (out *schemas.CommentDetail, err error) {
	return
}

func UpdateComment(ctx context.Context, infoId string, commentId string, req input.CreateComment) (err error) {
	return
}

func DeleteComment(ctx context.Context, infoId string, commentId string) (err error) {
	return
}

func GetInformationTypelist(ctx context.Context) (out []string, err error) {
	return
}

func CreateInformationType(ctx context.Context, req input.CreateInfoType) (err error) {
	return
}

func DeleteInformationType(ctx context.Context, typ string) (err error) {
	return
}
