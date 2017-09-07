package domain

import "time"

type ApoTask struct {
	ID           int           `json:"-" gorm:"primary_key;column:id;unique_index:devices_pkey"`
	AppID        string        `json:"app_id,omitempty"`
	AppName      string        `json:"app_name,omitempty"`
	BundleID     string        `json:"bundle_id,omitempty"`
	Level        int           `json:"level"`
	Total        int           `json:"total"`
	RealTotal    int           `json:"real_total"`
	DoneCount    int           `json:"done_count"`
	DoingCount   int           `json:"doing_count"`
	TimeoutCount int           `json:"timeount_count"`
	FailCount    int           `json:"fail_count,omitempty"`
	ApoKey       string        `json:"apo_key,omitempty"`
	AccountBrief string        `json:"account_brief,omitempty"`
	Cycle        int           `json:"cycel"`
	RemindCycle  int           `json:"remind_cycle"`
	UncatchDay   int           `json:"uncatch_day"`
	TypeModelID  int64         `json:"type_model_id,omitempty"`
	AmoutModelID int64         `json:"amount_model_id,omitempty"`
	Status       ApoTaskStatus `json:"status,omitempty"`
	PreaddCount  int           `json:"preadd_count"`
	PreaddTime   *time.Time    `json:"preadd_time,omitempty"`
	StartTime    *time.Time    `json:"start_time,omitempty" gorm:"column:start"`
	EndTime      *time.Time    `json:"end_time,omitempty" gorm:"column:end"`
	CreatedAt    *time.Time    `json:"create_time,omitempty" gorm:"column:create_time"`
	UpdatedAt    *time.Time    `json:"update_time,omitempty" gorm:"column:update_time"`
}

func (*ApoTask) TableName() string {
	return "apo_tasks"
}

type ApoTaskStatus int

const (
	_                    ApoTaskStatus = iota
	ApoTaskStatusStart                 // 开始
	ApoTaskStatusPause                 // 暂停
	ApoTaskStatusEnd                   // 完成
	ApoTaskStatusDeleted               // 删除
)

type ApoSubTask struct {
	ID       int        `json:"-" gorm:"primary_key;column:id;unique_index:devices_pkey"`
	ApoID    int        `json:"apo_id"`
	Count    int        `json:"count"`
	Status   int        `json:"status"`
	ExecTime *time.Time `json:"exec_time"`
	CreateAt *time.Time `json:"create_time" gorm:"column:create_time"`
}

func (*ApoSubTask) TableName() string {
	return "apo_sub_task"
}

type ApoSubTaskStatus int

const (
	ApoSubTaskDisable ApoSubTaskStatus = iota //  禁用
	ApoSubTaskEnable                          // 可用
)
