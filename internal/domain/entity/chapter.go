package entity

import "time"

type Chapter struct {
	Name         string
	Num          string
	RegulationID uint64
	OrderNum     uint32
	UpdatedAt    time.Time
}
