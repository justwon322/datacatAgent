// Code generated by ent, DO NOT EDIT.

package ent

import (
	"datacatAgent/ent/message"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Message is the model entity for the Message schema.
type Message struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Payload holds the value of the "payload" field.
	Payload string `json:"payload,omitempty"`
	// Sent holds the value of the "sent" field.
	Sent int `json:"sent,omitempty"`
	// Resent holds the value of the "resent" field.
	Resent int `json:"resent,omitempty"`
	// SentAt holds the value of the "sentAt" field.
	SentAt time.Time `json:"sentAt,omitempty"`
	// JobId holds the value of the "jobId" field.
	JobId        int `json:"jobId,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Message) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case message.FieldID, message.FieldSent, message.FieldResent, message.FieldJobId:
			values[i] = new(sql.NullInt64)
		case message.FieldPayload:
			values[i] = new(sql.NullString)
		case message.FieldSentAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Message fields.
func (m *Message) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case message.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int64(value.Int64)
		case message.FieldPayload:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field payload", values[i])
			} else if value.Valid {
				m.Payload = value.String
			}
		case message.FieldSent:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sent", values[i])
			} else if value.Valid {
				m.Sent = int(value.Int64)
			}
		case message.FieldResent:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field resent", values[i])
			} else if value.Valid {
				m.Resent = int(value.Int64)
			}
		case message.FieldSentAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field sentAt", values[i])
			} else if value.Valid {
				m.SentAt = value.Time
			}
		case message.FieldJobId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field jobId", values[i])
			} else if value.Valid {
				m.JobId = int(value.Int64)
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Message.
// This includes values selected through modifiers, order, etc.
func (m *Message) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// Update returns a builder for updating this Message.
// Note that you need to call Message.Unwrap() before calling this method if this Message
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Message) Update() *MessageUpdateOne {
	return NewMessageClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Message entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Message) Unwrap() *Message {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Message is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Message) String() string {
	var builder strings.Builder
	builder.WriteString("Message(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("payload=")
	builder.WriteString(m.Payload)
	builder.WriteString(", ")
	builder.WriteString("sent=")
	builder.WriteString(fmt.Sprintf("%v", m.Sent))
	builder.WriteString(", ")
	builder.WriteString("resent=")
	builder.WriteString(fmt.Sprintf("%v", m.Resent))
	builder.WriteString(", ")
	builder.WriteString("sentAt=")
	builder.WriteString(m.SentAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("jobId=")
	builder.WriteString(fmt.Sprintf("%v", m.JobId))
	builder.WriteByte(')')
	return builder.String()
}

// Messages is a parsable slice of Message.
type Messages []*Message
