package models

type Moneda struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Nombre     string `json:"nombre"`
	Cantidad   int    `json:"cantidad"`
	MonederoID uint   `json:"monederoID"` // Clave foránea a Monedero
}
