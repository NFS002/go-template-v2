package models

import (
	"time"
)

type Exhibition struct {
	ID        int              `db:"id" json:"-"`
	Title     string           `db:"title" json:"title"`
	HostID    int              `db:"host_id" json:"-"`
	Host      *Host            `db:"-" belongs_to:"host" json:"host" fk_id:"host_id"`
	ArtistID  int              `db:"artist_id" json:"-"`
	Artist    *Artist          `db:"-" belongs_to:"artist" json:"artist"`
	Notes     string           `db:"notes" json:"notes"`
	Dates     []ExhibitionDate `db:"-" json:"dates" has_many:"exhibition_dates" fk_id:"exhibition_id"`
	CreatedAt time.Time        `db:"created_at" json:"-"`
	UpdatedAt time.Time        `db:"updated_at" json:"-"`
}

func (e Exhibition) TableName() string {
	return "exhibitions"
}

func (e *Exhibition) StartDate() string {
	i := len(e.Dates)
	if i > 0 {
		dateFormat := "Mon, 02 Jan 2006"
		startTime := e.Dates[0].StartTime
		return startTime.Format(dateFormat)
	} else {
		return "N/A"
	}
}

func (e *Exhibition) EndDate() string {
	i := len(e.Dates)
	if i > 0 {
		dateFormat := "Mon, 02 Jan 2006"
		startTime := e.Dates[i-1].EndTime
		return startTime.Format(dateFormat)
	} else {
		return "N/A"
	}
}

func (e *Exhibition) Location() string {
	return e.Host.Address.Postcode[0:3]
}
