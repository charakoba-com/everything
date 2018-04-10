package model

import (
	"time"

	"github.com/charakoba-com/everything/faultinfo/schemas"
)

// Information is a main model of faultinfo service.
//+model table=information pkey=ID
type Information struct {
	ID         string
	Detail     string
	UpdatedAt  time.Time
	Updater    string
	Creator    string
	Type       string
	Begin      time.Time
	End        time.Time
	TemplateID string
	CreatedAt  time.Time
}

func (info *Information) ToInformationDetail() *schemas.InformationDetail {
	return &schemas.InformationDetail{
		ID:         info.ID,
		Detail:     info.Detail,
		UpdatedAt:  info.UpdatedAt,
		Updater:    info.Updater,
		Creator:    info.Creator,
		Type:       info.Type,
		Begin:      info.Begin,
		End:        info.End,
		TemplateID: info.TemplateID,
		CreatedAt:  info.CreatedAt,
	}
}

func (info *Information) ToInfo() *schemas.Information {
	return &schemas.Information{
		ID: info.ID,
	}
}

// Comment for Information
//+model
type Comment struct {
	ID            int
	InformationID int
	Creator       string
	Body          string
	CreatedAt     time.Time
}
