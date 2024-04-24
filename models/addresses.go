package models

import (
	"time"
)

type Address struct {
	ID        int       `db:"id" json:"-"`
	Number    int       `db:"number" json:"number"`
	Address1  string    `db:"address_1" json:"address_1"`
	Address2  string    `db:"address_2" json:"address_2"`
	City      string    `db:"city" json:"city"`
	Postcode  string    `db:"postcode" json:"postcode"`
	Host      *Host     `db:"-" has_one:"host" json:"-" fk_id:"address_id"`
	Artist    *Artist   `db:"-" has_one:"artist" json:"-" fk_id:"address_id"`
	CreatedAt time.Time `db:"created_at" json:"-"`
	UpdatedAt time.Time `db:"updated_at" json:"-"`
}

func (a Address) TableName() string {
	return "addresses"
}
