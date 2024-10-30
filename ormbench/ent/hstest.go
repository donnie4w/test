// Code generated by ent, DO NOT EDIT.

package ent

import (
	"enttest/ent/hstest"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Hstest is the model entity for the Hstest schema.
type Hstest struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Rowname holds the value of the "rowname" field.
	Rowname string `json:"rowname,omitempty"`
	// Value holds the value of the "value" field.
	Value string `json:"value,omitempty"`
	// Updatetime holds the value of the "updatetime" field.
	Updatetime time.Time `json:"updatetime,omitempty"`
	// Body holds the value of the "body" field.
	Body string `json:"body,omitempty"`
	// Floa holds the value of the "floa" field.
	Floa float64 `json:"floa,omitempty"`
	// Level holds the value of the "level" field.
	Level        int `json:"level,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Hstest) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case hstest.FieldFloa:
			values[i] = new(sql.NullFloat64)
		case hstest.FieldID, hstest.FieldAge, hstest.FieldLevel:
			values[i] = new(sql.NullInt64)
		case hstest.FieldRowname, hstest.FieldValue, hstest.FieldBody:
			values[i] = new(sql.NullString)
		case hstest.FieldUpdatetime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Hstest fields.
func (h *Hstest) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case hstest.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			h.ID = int(value.Int64)
		case hstest.FieldAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				h.Age = int(value.Int64)
			}
		case hstest.FieldRowname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field rowname", values[i])
			} else if value.Valid {
				h.Rowname = value.String
			}
		case hstest.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				h.Value = value.String
			}
		case hstest.FieldUpdatetime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatetime", values[i])
			} else if value.Valid {
				h.Updatetime = value.Time
			}
		case hstest.FieldBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field body", values[i])
			} else if value.Valid {
				h.Body = value.String
			}
		case hstest.FieldFloa:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field floa", values[i])
			} else if value.Valid {
				h.Floa = value.Float64
			}
		case hstest.FieldLevel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field level", values[i])
			} else if value.Valid {
				h.Level = int(value.Int64)
			}
		default:
			h.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// GetValue returns the ent.Value that was dynamically selected and assigned to the Hstest.
// This includes values selected through modifiers, order, etc.
func (h *Hstest) GetValue(name string) (ent.Value, error) {
	return h.selectValues.Get(name)
}

// Update returns a builder for updating this Hstest.
// Note that you need to call Hstest.Unwrap() before calling this method if this Hstest
// was returned from a transaction, and the transaction was committed or rolled back.
func (h *Hstest) Update() *HstestUpdateOne {
	return NewHstestClient(h.config).UpdateOne(h)
}

// Unwrap unwraps the Hstest entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (h *Hstest) Unwrap() *Hstest {
	_tx, ok := h.config.driver.(*txDriver)
	if !ok {
		panic("ent: Hstest is not a transactional entity")
	}
	h.config.driver = _tx.drv
	return h
}

// String implements the fmt.Stringer.
func (h *Hstest) String() string {
	var builder strings.Builder
	builder.WriteString("Hstest(")
	builder.WriteString(fmt.Sprintf("id=%v, ", h.ID))
	builder.WriteString("age=")
	builder.WriteString(fmt.Sprintf("%v", h.Age))
	builder.WriteString(", ")
	builder.WriteString("rowname=")
	builder.WriteString(h.Rowname)
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(h.Value)
	builder.WriteString(", ")
	builder.WriteString("updatetime=")
	builder.WriteString(h.Updatetime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("body=")
	builder.WriteString(h.Body)
	builder.WriteString(", ")
	builder.WriteString("floa=")
	builder.WriteString(fmt.Sprintf("%v", h.Floa))
	builder.WriteString(", ")
	builder.WriteString("level=")
	builder.WriteString(fmt.Sprintf("%v", h.Level))
	builder.WriteByte(')')
	return builder.String()
}

// Hstests is a parsable slice of Hstest.
type Hstests []*Hstest