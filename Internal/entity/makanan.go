package entity

type Makanan struct {
	ID           uint   `gorm:"primaryKey"`
	NamaMakanan  string `gorm:"type:varchar(100); not null"`
	AnakID       uint
	LogMakananID uint
	EnergiFood   float64 `gorm:"type:float"`
	Protein      float64 `gorm:"type:float"`
	Lemak        float64 `gorm:"type:float"`
	Karbo        float64 `gorm:"type:float"`
}
