package model

import (
	"time"
)

type BuatPesan struct {
	Pesan      string    `json:"pesan"`
	WaktuPesan time.Time `json:"waktu_pesan"`
}

type EditPesan struct {
	Pesan      string    `json:"pesan"`
	WaktuPesan time.Time `json:"waktu_pesan"`
}

func (e *EditPesan) ToMap() map[string]any {
	updates := map[string]any{}

	if e.Pesan != "" {
		updates["pesan"] = e.Pesan
	}
	if !e.WaktuPesan.IsZero() {
		updates["waktu_pesan"] = e.WaktuPesan
	}

	return updates
}
