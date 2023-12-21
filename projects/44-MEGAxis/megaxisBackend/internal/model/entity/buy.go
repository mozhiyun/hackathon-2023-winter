// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Buy is the golang structure for table buy.
type Buy struct {
	Id         uint        `json:"id"         description:"pk"`
	Uid        string      `json:"uid"        description:"User ID"`
	Pid        string      `json:"pid"        description:"Prompt ID"`
	CreateTime *gtime.Time `json:"createTime" description:"Created Time"`
	UpdateTime *gtime.Time `json:"updateTime" description:"Updated Time"`
}