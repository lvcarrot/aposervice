package domain

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type ApoComment struct {
	ID       bson.ObjectId    `json:"id,omitempty" bson:"_id"`
	AppID    string           `json:"app_id,omitempty" bson:"app_id,omitempty"`
	ApoID    string           `json:"apo_id,omitempty" bson:"app_od,omitempty"`
	Content  interface{}      `json:"content,omitempty" bson:"content,omitempty"`
	MD5      string           `json:"md5,omitempty" bson:"md5,omitempty"`
	IP       string           `json:"ip,omitempty" bson:"ip,omitempty"`
	Errno    int              `json:"errno,omitempty" bson:"-"`
	Status   ApoCommentStatus `json:"status,omitempty" bson:"status,omitempty"`
	Action   string           `json:"action,omitempty" bson:"-"`
	CreateAt *time.Time       `json:"create_time,omitempty" bson:"create_time,omitempty"`
	UpdateAt *time.Time       `json:"update_time,omitempty" bson:"update_time,omitempty"`
}

func (*ApoComment) TableName() string {
	return "apo_comments"
}

type ApoCommentStatus int

const (
	_                      ApoCommentStatus = iota
	ApoCommentStatusFree                    // 可用
	ApoCommentStatusLocked                  // 锁定
	ApoCommentStatusUsed                    // 已用
)

type AppCommentRecord struct {
	AppID      string     `json:"app_id,omitempty" bson:"app_id,omitempty"`
	IP         string     `json:"ip,omitempty" bson:"ip,omitempty"`
	Count      int        `json:"count,omitempty" bson:"count,omitempty"`
	CreateTime *time.Time `json:"create_time,omitempty" bson:"create_time,omitempty"`
	UpdateTime *time.Time `json:"update_time,omitempty" bson:"update_time,omitempty"`
}

func (*AppCommentRecord) TableName() string {
	return "app_comments"
}
