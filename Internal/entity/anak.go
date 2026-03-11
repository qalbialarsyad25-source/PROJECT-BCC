package entity

type Anak struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint
	Loganak    []LogMakanan `gorm:"foreignKey:AnakID"`
	Makanan    []Makanan    `gorm:"foreignKey:AnakID"`
	NamaAnak   string       `gorm:"type:varchar(225); not null"`
	TinggiAnak float64      `gorm:"type:decimal(10,2); not null"`
	BBAnak     float64      `gorm:"type:float; not null"`
	GenderAnak string       `gorm:"type:varchar(100); not null"`
	LingkarL   float64      `gorm:"type:float; not null"`
	LingkarK   float64      `gorm:"type:float; not null"`
}
