package entities

import "time"

type Store struct {
	Id       int       `json:"id"`
	PetId    int       `json:"petId"`
	ShipDate time.Time `json:"shipDate"`
	Status   string    `json:"status"`
	Complete bool      `json:"complete"`
}
