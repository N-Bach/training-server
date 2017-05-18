package entity

import (
	"time"
)

type Period struct {
	From time.Time `gorethink:"from" json:"from"`
	To time.Time `gorethink:"to" json:"to"`
}