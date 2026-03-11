package entity

type Informasi struct {
	ID        uint   `gorm:"primaryKey"`
	Ringkasan string `gorm:"type:longtext; not null"`
	Judul     string `gorm:"type:varchar(225); not null"`
	UserID    uint
}
