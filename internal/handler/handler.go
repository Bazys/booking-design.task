package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"applicationDesignTest/internal/entity"
	"applicationDesignTest/internal/logger"
	"applicationDesignTest/internal/storage"
)

func NewHandler(db storage.DB, log logger.Logger) Handler {
	return &handler{
		db:  db,
		log: log,
	}
}

type Handler interface {
	MakeOrder(w http.ResponseWriter, r *http.Request)
	GetOrders(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	db  storage.DB
	log logger.Logger
}

func (h *handler) MakeOrder(w http.ResponseWriter, r *http.Request) {
	userEmail := r.URL.Query().Get("email")
	if userEmail == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	room := r.URL.Query().Get("room")
	if room == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	var err error
	var R *entity.AvailableRoomsClass
	if R, err = entity.NewAvailableRoomsClass(room); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if _, isOK := entity.AvailableRooms[*R]; !isOK {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	from := r.URL.Query().Get("from")
	if from == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	fromTime, err := time.Parse("2006-01-02", from)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	to := r.URL.Query().Get("to")
	if to == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	toTime, err := time.Parse("2006-01-02", to)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	actualOrders := h.db.GetOrders()
	for _, order := range actualOrders {
		currentOrderFromTime, _ := time.Parse("2006-01-02", order.From)
		currentOrderToTime, _ := time.Parse("2006-01-02", order.To)
		if !(currentOrderToTime.Before(fromTime) || currentOrderFromTime.After(toTime)) {
			http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
			return
		}
	}

	newOrder := entity.Order{
		Room:      room,
		UserEmail: userEmail,
		From:      from,
		To:        to,
	}
	h.db.AddOrder(newOrder)

	w.WriteHeader(http.StatusCreated)
	h.log.Info("Method makeOrder was successfully done")
}

func (h *handler) GetOrders(w http.ResponseWriter, r *http.Request) {
	userEmail := r.URL.Query().Get("email")
	if userEmail == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(h.db.GetUsersOrders(userEmail))
	if err != nil {
		h.log.Errorf("error in getOrders method: %s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(b)

	h.log.Info("Method getOrders was successfully done")
}
