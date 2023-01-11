package router

import (
	"net/http"

	"applicationDesignTest/internal/handler"
	"applicationDesignTest/internal/logger"
	"applicationDesignTest/internal/storage"
)

func NewRouter(db storage.DB, log logger.Logger) http.Handler {
	h := handler.NewHandler(db, log)

	router := http.NewServeMux()
	router.HandleFunc("/makeOrder", h.MakeOrder)
	router.HandleFunc("/getOrders", h.GetOrders)

	return router
}
