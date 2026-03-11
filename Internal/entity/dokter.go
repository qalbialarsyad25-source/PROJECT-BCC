package entity

type Dokter struct {
	ID           uint         `gorm:"primaryKey"`
	NamaDokter   string       `gorm:"type:varchar(225); not null"`
	DokterKonsul []Konsultasi `gorm:"foreignKey:DokterID"`
}
