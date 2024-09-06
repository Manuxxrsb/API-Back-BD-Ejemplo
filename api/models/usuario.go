package models

type Usuario struct {
	UsuarioID uint     `json:"usuarioID" gorm:"primaryKey"`
	Nombre    string   `json:"nombre"`
	Monedero  Monedero `json:"monedero"` // Relación uno a uno con Monedero
}

func (Usuario) TableName() string {
	return "usuario"
}
