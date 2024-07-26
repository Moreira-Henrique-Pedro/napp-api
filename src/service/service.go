package service

import (
	"context"
	"log/slog"
	"time"

	"github.com/Moreira-Henrique-Pedro/napp-api/src/infra"
	"github.com/Moreira-Henrique-Pedro/napp-api/src/model"
	"github.com/Moreira-Henrique-Pedro/napp-api/src/rules"
	"github.com/jackc/pgx/v4"
)

type StockService struct {
	db    *pgx.Conn
	rules rules.StockRules
}

func NewStockService() *StockService {
	return &StockService{
		db:    infra.Conn,
		rules: rules.NewStockRules(),
	}
}

func (s *StockService) CreateStock(stock model.Stock) (int64, error) {
	slog.Info("Creating Stock")

	if err := s.rules.ValidateStock(&stock); err != nil {
		return 0, err
	}

	if err := s.rules.CalculateAvailableStock(&stock); err != nil {
		return 0, err
	}

	query := `
		INSERT INTO stocks (ID, nome, estoque_total, estoque_corte, estoque_disponivel, preco_de, preco_por, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id;
	`

	var id int64
	err := s.db.QueryRow(context.Background(), query,
		stock.ID,
		stock.Nome,
		stock.Estoque.EstoqueTotal,
		stock.Estoque.EstoqueCorte,
		stock.Estoque.EstoqueDisponivel,
		stock.PrecoDe,
		stock.PrecoPor,
		time.Now(),
		time.Now(),
	).Scan(&id)

	if err != nil {
		slog.Error("Failed to create stock: %v", err)
		return 0, err
	}

	return id, nil
}

func (s *StockService) UpdateStock(stock model.Stock) error {
	slog.Info("Updating Stock")

	if err := s.rules.ValidateStock(&stock); err != nil {
		return err
	}

	if err := s.rules.CalculateAvailableStock(&stock); err != nil {
		return err
	}

	query := `
		UPDATE stocks
		SET nome = $1, estoque_total = $2, estoque_corte = $3, estoque_disponivel = $4, preco_de = $5, preco_por = $6, updated_at = $7, deleted_at = $8
		WHERE id = $9;
	`

	_, err := s.db.Exec(context.Background(), query,
		stock.Nome,
		stock.Estoque.EstoqueTotal,
		stock.Estoque.EstoqueCorte,
		stock.Estoque.EstoqueDisponivel,
		stock.PrecoDe,
		stock.PrecoPor,
		time.Now(),
		stock.DeletedAt,
		stock.ID,
	)

	if err != nil {
		slog.Error("Failed to update stock: %v", err)
		return err
	}

	return nil
}

func (s *StockService) FindStockByID(id int64) (model.Stock, error) {
	slog.Info("Finding Stock by ID")

	query := `SELECT id, nome, estoque_total, estoque_corte, preco_de, preco_por, created_at, updated_at, deleted_at FROM stocks WHERE id = $1`
	var stock model.Stock

	err := s.db.QueryRow(context.Background(), query, id).Scan(
		&stock.ID,
		&stock.Nome,
		&stock.Estoque.EstoqueTotal,
		&stock.Estoque.EstoqueCorte,
		&stock.PrecoDe,
		&stock.PrecoPor,
		&stock.CreatedAt,
		&stock.UpdatedAt,
		&stock.DeletedAt,
	)

	if err != nil {
		slog.Error("Failed to find stock: %v", err)
		return model.Stock{}, err
	}

	return stock, nil
}

func (s *StockService) FindAllStocks() ([]model.Stock, error) {
	slog.Info("Finding all stocks")

	query := `SELECT id, nome, estoque_total, estoque_corte, preco_de, preco_por, created_at, updated_at, deleted_at FROM stocks`

	rows, err := s.db.Query(context.Background(), query)
	if err != nil {
		slog.Error("Failed to find stocks: %v", err)
		return nil, err
	}
	defer rows.Close()

	var stocks []model.Stock
	for rows.Next() {
		var stock model.Stock
		err := rows.Scan(
			&stock.ID,
			&stock.Nome,
			&stock.Estoque.EstoqueTotal,
			&stock.Estoque.EstoqueCorte,
			&stock.PrecoDe,
			&stock.PrecoPor,
			&stock.CreatedAt,
			&stock.UpdatedAt,
			&stock.DeletedAt,
		)
		if err != nil {
			slog.Error("Failed to scan stock: %v", err)
			return nil, err
		}
		stocks = append(stocks, stock)
	}

	if err = rows.Err(); err != nil {
		slog.Error("Rows error: %v", err)
		return nil, err
	}

	return stocks, nil
}

func (s *StockService) DeleteStockByID(id int64) error {
	slog.Info("Deleting Stock by ID")

	query := `DELETE FROM stocks WHERE id = $1`
	_, err := s.db.Exec(context.Background(), query, id)

	if err != nil {
		slog.Error("Failed to delete stock: %v", err)
		return err
	}

	return nil
}
