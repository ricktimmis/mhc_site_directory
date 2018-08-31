package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Campsite struct {
	ID             uuid.UUID     `json:"id" db:"id"`
	CreatedAt      time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at" db:"updated_at"`
	Name           string        `json:"name" db:"name"`
	Name2          string        `json: "name" db: "`
	Description    nulls.String  `json:"description" db:"description"`
	Attractions    nulls.String  `json:"attraction" db:"attraction"`
	Email          string        `json:"email" db:"email"`
	Telephone      string        `json:"telephone" db:"telephone"`
	Website        nulls.String  `json:"website" db:"website"`
	Facilities     Facility		 `has_many:"facility" order_by:"name asc"`
	Owninguser     User          `belongs_to:"User"`
	Images         Images        `has_many:"image" db:"_"`
	Longitude      float64       `json:"longitude" db:"longitude"`
	Latitude       float64       `json:"latitude" db:"latitude"`
	Addressstreet  nulls.String  `json:"addressstreet" db:"addressstreet"`
	Locality       nulls.String  `json:"locality" db:"locality"`
	Town           nulls.String  `json:"town" db:"town"`
	County         nulls.String  `json:"county" db:"county"`
	Postcode       nulls.String  `json:"postcode" db:"postcode"`
	Listingexpires time.Time     `json:"listingexpire" db:"listingexpire"`
}

// String is not required by pop and may be deleted
func (c Campsite) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Campsites is not required by pop and may be deleted
type Campsites []Campsite

// String is not required by pop and may be deleted
func (c Campsites) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *Campsite) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.TimeIsPresent{Field: c.Listingexpires, Name: "Listingexpires"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *Campsite) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *Campsite) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
