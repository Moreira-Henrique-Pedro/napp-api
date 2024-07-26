package rules

import (
	"errors"

	"github.com/Moreira-Henrique-Pedro/napp-api/src/model"
)

type StockRules interface {
	ValidateStock(stock *model.Stock) error
	CalculateAvailableStock(stock *model.Stock) error
}

type stockRules struct{}

func NewStockRules() StockRules {
	return &stockRules{}
}

func (r *stockRules) ValidateStock(stock *model.Stock) error {

	if stock.PrecoDe < stock.PrecoPor {
		return errors.New("O Preço De não pode ser menor que Preço Por")
	}

	return nil
}

func (r *stockRules) CalculateAvailableStock(stock *model.Stock) error {
	stock.Estoque.EstoqueDisponivel = stock.Estoque.EstoqueTotal - stock.Estoque.EstoqueCorte
	return nil
}
