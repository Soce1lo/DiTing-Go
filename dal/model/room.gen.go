// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameRoom = "room"

// Room 房间表
type Room struct {
	ID           int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:id" json:"id"`                                                    // id
	Type         int32     `gorm:"column:type;not null;comment:房间类型 1群聊 2单聊" json:"type"`                                                           // 房间类型 1群聊 2单聊
	HotFlag      int32     `gorm:"column:hot_flag;comment:是否全员展示 0否 1是" json:"hot_flag"`                                                            // 是否全员展示 0否 1是
	ActiveTime   time.Time `gorm:"column:active_time;not null;default:CURRENT_TIMESTAMP(3);comment:群最后消息的更新时间（热点群不需要写扩散，只更新这里）" json:"active_time"` // 群最后消息的更新时间（热点群不需要写扩散，只更新这里）
	LastMsgID    int64     `gorm:"column:last_msg_id;comment:会话中的最后一条消息id" json:"last_msg_id"`                                                      // 会话中的最后一条消息id
	ExtJSON      string    `gorm:"column:ext_json;comment:额外信息（根据不同类型房间有不同存储的东西）" json:"ext_json"`                                                  // 额外信息（根据不同类型房间有不同存储的东西）
	CreateTime   time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP(3);comment:创建时间" json:"create_time"`                        // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP(3);comment:修改时间" json:"update_time"`                        // 修改时间
	DeleteStatus int32     `gorm:"column:delete_status;not null;comment:房间状态 0正常 1禁用(删好友了禁用)" json:"delete_status"`                                 // 房间状态 0正常 1禁用(删好友了禁用)
}

// TableName Room's table name
func (*Room) TableName() string {
	return TableNameRoom
}
