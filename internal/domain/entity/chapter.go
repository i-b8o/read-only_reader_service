package entity

import "time"

type Chapter struct {
	Name      string
	Num       string
	DocID     uint64
	OrderNum  uint32
	UpdatedAt time.Time
}
