package model

import "time"

type Stock struct {
	ID        int64     `json: "id"`
	Nome      string    `json: "nome"`
	Estoque   Estoque   `json: "estoque"`
	PrecoDe   int64     `json: "preçoDe"`
	PrecoPor  int64     `json: "preçoPor"`
	CreatedAt time.Time `json: "createdAt"`
	UpdatedAt time.Time `json: "updatedAt"`
	DeletedAt time.Time `json: "deletedAt"`
}

type Estoque struct {
	EstoqueTotal      int64 `json: "estoqueTotal"`
	EstoqueCorte      int64 `json: "estoqueCorte"`
	EstoqueDisponivel int64 `json: "estoqueDisponivel"`
}
