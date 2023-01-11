package entity

import "fmt"

type AvailableRoomsClass string

const (
	Econom  AvailableRoomsClass = "econom"
	Standrt AvailableRoomsClass = "standart"
	Lux     AvailableRoomsClass = "lux"
)

var AvailableRooms = map[AvailableRoomsClass]struct{}{Econom: {}, Standrt: {}, Lux: {}}

func NewAvailableRoomsClass(in string) (*AvailableRoomsClass, error) {
	res := AvailableRoomsClass(in)
	if _, ok := AvailableRooms[res]; ok {
		return &res, nil
	}
	return nil, fmt.Errorf("invalid room class")
}
