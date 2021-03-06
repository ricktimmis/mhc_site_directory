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

type User struct {
	ID                uuid.UUID  `json:"id" db:"id"`
	CreatedAt         time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at" db:"updated_at"`
	Username          string     `json:"username" db:"username"`
	Firstname         string     `json:"firstname" db:"firstname"`
	Lastname          string     `json:"lastname" db:"lastname"`
	Email             string     `json:"email" db:"email"`
	Birthday          nulls.Time `json:"birthday" db:"birthday"`
	Location          string     `json:"location" db:"location"`
	Membertype        string     `json:"membertype" db:"membertype"`
	LoginProvider     string     `json:"loginprovider" db:"loginprovider"`
	LoginProviderID   string     `json:"loginprovider_id" db:"loginprovider_id"`
	Membershipexpires nulls.Time `json:"membershipexpires" db:"membershipexpires"`
}

// String is not required by pop and may be deleted
func (u User) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Users is not required by pop and may be deleted
type Users []User

// String is not required by pop and may be deleted
func (u Users) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: u.Username, Name: "Username"},
		&validators.StringIsPresent{Field: u.Firstname, Name: "Firstname"},
		&validators.StringIsPresent{Field: u.Lastname, Name: "Lastname"},
		&validators.StringIsPresent{Field: u.Email, Name: "Email"},
		&validators.StringIsPresent{Field: u.Location, Name: "Location"},
		&validators.StringIsPresent{Field: u.Membertype, Name: "Membertype"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
