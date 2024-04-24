package models

import (
	"time"
)

type Artist struct {
	ID        int       `db:"id" json:"-"`
	FirstName string    `db:"first_name" json:"first_name"`
	LastName  string    `db:"last_name" json:"last_name"`
	Email     string    `db:"email_address" json:"email"`
	Phone     string    `db:"phone_number" json:"phone"`
	AddressID int       `db:"address_id" json:"-"`
	Address   *Address  `db:"-" belongs_to:"address" json:"address"`
	CreatedAt time.Time `db:"created_at" json:"-"`
	UpdatedAt time.Time `db:"updated_at" json:"-"`
}

func (a Artist) TableName() string {
	return "artists"
}

func (a Artist) FullName() string {
	return a.FirstName + " " + a.LastName
}
