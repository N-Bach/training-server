package entity

import (
	"time"
)

type TimeStamp struct {
	CreatedAt time.Time `gorethink:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorethink:"updatedAt" json:"updatedAt"`
}