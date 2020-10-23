// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/Piichet-3-/app/ent/gender"
	"github.com/Piichet-3-/app/ent/position"
	"github.com/Piichet-3-/app/ent/title"
	"github.com/Piichet-3-/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges       UserEdges `json:"edges"`
	dispense_id *int
	drug_id     *int
	gender_id   *int
	position_id *int
	title_id    *int
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Gender holds the value of the gender edge.
	Gender *Gender
	// Position holds the value of the position edge.
	Position *Position
	// Title holds the value of the title edge.
	Title *Title
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// GenderOrErr returns the Gender value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) GenderOrErr() (*Gender, error) {
	if e.loadedTypes[0] {
		if e.Gender == nil {
			// The edge gender was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: gender.Label}
		}
		return e.Gender, nil
	}
	return nil, &NotLoadedError{edge: "gender"}
}

// PositionOrErr returns the Position value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) PositionOrErr() (*Position, error) {
	if e.loadedTypes[1] {
		if e.Position == nil {
			// The edge position was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: position.Label}
		}
		return e.Position, nil
	}
	return nil, &NotLoadedError{edge: "position"}
}

// TitleOrErr returns the Title value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) TitleOrErr() (*Title, error) {
	if e.loadedTypes[2] {
		if e.Title == nil {
			// The edge title was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: title.Label}
		}
		return e.Title, nil
	}
	return nil, &NotLoadedError{edge: "title"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
		&sql.NullString{}, // email
		&sql.NullString{}, // password
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*User) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // dispense_id
		&sql.NullInt64{}, // drug_id
		&sql.NullInt64{}, // gender_id
		&sql.NullInt64{}, // position_id
		&sql.NullInt64{}, // title_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	u.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		u.Name = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email", values[1])
	} else if value.Valid {
		u.Email = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field password", values[2])
	} else if value.Valid {
		u.Password = value.String
	}
	values = values[3:]
	if len(values) == len(user.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field dispense_id", value)
		} else if value.Valid {
			u.dispense_id = new(int)
			*u.dispense_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field drug_id", value)
		} else if value.Valid {
			u.drug_id = new(int)
			*u.drug_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field gender_id", value)
		} else if value.Valid {
			u.gender_id = new(int)
			*u.gender_id = int(value.Int64)
		}
		if value, ok := values[3].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field position_id", value)
		} else if value.Valid {
			u.position_id = new(int)
			*u.position_id = int(value.Int64)
		}
		if value, ok := values[4].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field title_id", value)
		} else if value.Valid {
			u.title_id = new(int)
			*u.title_id = int(value.Int64)
		}
	}
	return nil
}

// QueryGender queries the gender edge of the User.
func (u *User) QueryGender() *GenderQuery {
	return (&UserClient{config: u.config}).QueryGender(u)
}

// QueryPosition queries the position edge of the User.
func (u *User) QueryPosition() *PositionQuery {
	return (&UserClient{config: u.config}).QueryPosition(u)
}

// QueryTitle queries the title edge of the User.
func (u *User) QueryTitle() *TitleQuery {
	return (&UserClient{config: u.config}).QueryTitle(u)
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", name=")
	builder.WriteString(u.Name)
	builder.WriteString(", email=")
	builder.WriteString(u.Email)
	builder.WriteString(", password=")
	builder.WriteString(u.Password)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
