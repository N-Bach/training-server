package entity

import (
	"time"
)

type TimeStamp struct {
	CreateAt time.Time `gorethink:"createAt" json:"createAt"`
	UpdateAt time.Time `gorethink:"createAt" json:"createAt"`
}