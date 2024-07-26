package controller

import (
	"net/http"
	"strconv"

	"github.com/Moreira-Henrique-Pedro/napp-api/src/model"
	"github.com/Moreira-Henrique-Pedro/napp-api/src/service"
	"github.com/gin-gonic/gin"
)

type StockController struct {
	service *service.StockService
}

func NewStockController(service *service.StockService) *StockController {
	return &StockController{
		service: service,
	}
}

func (s *StockController) InitRoutes() {
	app := gin.Default()
	api := app.Group("/api/stock-service")

	api.GET("/:id", s.findByID)
	api.GET("/", s.findAll)
	api.POST("/", s.create)
	api.PUT("/:id", s.update)
	api.DELETE("/:id", s.delete)

	app.Run(":3000")
}

func (s *StockController) create(ctx *gin.Context) {
	stock := new(model.Stock)
	if err := ctx.ShouldBindJSON(&stock); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err},
		)
		return
	}

	id, err := s.service.CreateStock(*stock)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err},
		)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}

func (s *StockController) findAll(ctx *gin.Context) {
	stocks, err := s.service.FindAllStocks()
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.JSON(http.StatusOK, stocks)
}

func (s *StockController) findByID(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Invalid stock ID"},
		)
		return
	}

	stock, err := s.service.FindStockByID(id)
	if err != nil {
		ctx.JSON(
			http.StatusNotFound,
			gin.H{"error": "Stock not found"},
		)
		return
	}

	ctx.JSON(http.StatusOK, stock)
}

func (s *StockController) update(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Invalid stock ID"},
		)
		return
	}

	stock := new(model.Stock)
	if err := ctx.ShouldBindJSON(&stock); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}
	stock.ID = id

	err = s.service.UpdateStock(*stock)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (s *StockController) delete(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Invalid stock ID"},
		)
		return
	}

	err = s.service.DeleteStockByID(id)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.Status(http.StatusNoContent)
}
