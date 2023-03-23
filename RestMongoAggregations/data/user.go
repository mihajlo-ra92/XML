package data

import (
	"encoding/json"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Patient struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name         string             `bson:"name" json:"name"`
	Surname      string             `bson:"surname,omitempty" json:"surname"`
	PhoneNumbers []string           `bson:"phoneNumbers,omitempty" json:"phoneNumbers"`
	Address      Address            `bson:"address,omitempty" json:"address"`
	Anamnesis    []Anamnesis        `bson:"anamnesis,omitempty" json:"anamnesis"`
	Therapy      []Therapy          `bson:"therapy,omitempty" json:"therapy"`
}

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
	UserType string             `bson:"userType" json:"userType"`
	FirstName	string			`bson:"firstName" json:"firstName"`
	LastName 	string			`bson:"lastName" json:"lastName"`
	Gender		string			`bson:"gender" json:"gender"`
	BirthDate	time.Time		`bson:"birthDate" json:"birthDate"`
	Email		string			`bson:"email" json:"email"`
	GovernmentId	string		`bson:"governmentId" json:"governmentId"`

}

type Address struct {
	Street  string `bson:"street,omitempty" json:"street"`
	City    string `bson:"city,omitempty" json:"city"`
	Country string `bson:"country,omitempty" json:"country"`
}

type Anamnesis struct {
	Symptom   string    `bson:"symptom,omitempty" json:"symptom"`
	Intensity string    `bson:"intensity,omitempty" json:"intensity"`
	StartDate time.Time `bson:"startDate,omitempty" json:"startDate"`
}

type Therapy struct {
	Name  string  `bson:"name,omitempty" json:"name"`
	Price float32 `bson:"price,omitempty" json:"price"`
}

type Patients []*Patient

type Users []*User

func (u *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *User) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(u)
}

func (u *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}
