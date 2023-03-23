package data

import (
	"encoding/json"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ticket struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Date       time.Time          `bson:"date" json:"date"`
	EndPlace   string             `bson:"endPlace" json:"endPlace"`
	StartPlace string             `bson:"startPlace" json:"startPlace"`
	Capacity   int                `bson:"capacity" json:"capacity"`
	Price      int                `bson:"price" json:"price"`
	FlightId   string             `bson:"flightId" json:"flightId"`
	UserID     string             `bson:"userId,omitempty" json:"userId"`
}

type Tickets []*Ticket

func (f *Tickets) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(f)
}

func (f *Ticket) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(f)
}

func (f *Ticket) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(f)
}
