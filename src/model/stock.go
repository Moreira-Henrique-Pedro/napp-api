package model

import "time"

type Stock struct {
	ID        int64
	Nome      string
	Estoque   Estoque
	PrecoDe   int64
	PrecoPor  int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Estoque struct {
	EstoqueTotal      int64
	EstoqueCorte      int64
	EstoqueDisponivel int64
}
