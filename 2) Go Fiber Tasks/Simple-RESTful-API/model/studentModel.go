package models

import "time"

type Student struct {
	Id        uint      `json: "id"`
	Name      string    `json: "name"`
	Passcode  string    `json:"passcode"`
	createdAt time.Time `json:"created_at"`
	updatedAt time.Time `json:"updated_at"`
}
