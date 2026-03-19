package model

type BuatUserDokter struct {
	Nama string `json:"nama"`
}

type EditDokter struct {
	Nama string `json:"nama"`
}

func (e *EditDokter) ToMap() map[string]any {
	updates := map[string]any{}

	if e.Nama != "" {
		updates["nama"] = e.Nama
	}
	return updates
}
