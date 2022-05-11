package models

import (
	"time"

	"gorm.io/gorm"
)

type Games struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	Nome            string         `json:"nome"`
	Desenvolvedor   string         `json:"desenvolvedor"`
	AnoDeLancamento uint           `json:"anoDeLancamento"`
	Preco           float64        `json:"preco"`
	NotaPessoal     uint           `json:"notaPessoal"`
	UrlDaImagem     string         `json:"img_url"`
	CreatedAt       time.Time      `json:"created"`
	UpdatedAt       time.Time      `json:"updated"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted"`
}
