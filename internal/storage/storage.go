package storage

import (
	"sync"

	"applicationDesignTest/internal/entity"
)

func NewDB() DB {
	return &db{
		actualOrders: make([]entity.Order, 0),
	}
}

type DB interface {
	AddOrder(order entity.Order)
	GetOrders() []entity.Order
	GetUsersOrders(userEmail string) []entity.Order
}

type db struct {
	m            sync.RWMutex
	actualOrders []entity.Order
}

func (d *db) AddOrder(order entity.Order) {
	d.m.Lock()
	defer d.m.Unlock()
	d.actualOrders = append(d.actualOrders, order)
}

func (d *db) GetOrders() []entity.Order {
	d.m.RLock()
	defer d.m.RUnlock()
	return d.actualOrders
}

func (d *db) GetUsersOrders(userEmail string) []entity.Order {
	d.m.RLock()
	defer d.m.RUnlock()
	var usersOrders []entity.Order
	for _, order := range d.actualOrders {
		if order.UserEmail == userEmail {
			usersOrders = append(usersOrders, order)
		}
	}
	return usersOrders
}
