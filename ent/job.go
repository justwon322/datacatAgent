// Code generated by ent, DO NOT EDIT.

package ent

import (
	"datacatAgent/ent/job"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Job is the model entity for the Job schema.
type Job struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// Program holds the value of the "program" field.
	Program string `json:"program,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// UpdatedById holds the value of the "updatedById" field.
	UpdatedById int `json:"updatedById,omitempty"`
	// CreatedById holds the value of the "createdById" field.
	CreatedById int `json:"createdById,omitempty"`
	// CategorySubId holds the value of the "categorySubId" field.
	CategorySubId int `json:"categorySubId,omitempty"`
	selectValues  sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Job) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case job.FieldID, job.FieldUpdatedById, job.FieldCreatedById, job.FieldCategorySubId:
			values[i] = new(sql.NullInt64)
		case job.FieldProgram, job.FieldName:
			values[i] = new(sql.NullString)
		case job.FieldUpdatedAt, job.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Job fields.
func (j *Job) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case job.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			j.ID = int32(value.Int64)
		case job.FieldProgram:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field program", values[i])
			} else if value.Valid {
				j.Program = value.String
			}
		case job.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				j.UpdatedAt = value.Time
			}
		case job.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				j.CreatedAt = value.Time
			}
		case job.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				j.Name = value.String
			}
		case job.FieldUpdatedById:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updatedById", values[i])
			} else if value.Valid {
				j.UpdatedById = int(value.Int64)
			}
		case job.FieldCreatedById:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field createdById", values[i])
			} else if value.Valid {
				j.CreatedById = int(value.Int64)
			}
		case job.FieldCategorySubId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field categorySubId", values[i])
			} else if value.Valid {
				j.CategorySubId = int(value.Int64)
			}
		default:
			j.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Job.
// This includes values selected through modifiers, order, etc.
func (j *Job) Value(name string) (ent.Value, error) {
	return j.selectValues.Get(name)
}

// Update returns a builder for updating this Job.
// Note that you need to call Job.Unwrap() before calling this method if this Job
// was returned from a transaction, and the transaction was committed or rolled back.
func (j *Job) Update() *JobUpdateOne {
	return NewJobClient(j.config).UpdateOne(j)
}

// Unwrap unwraps the Job entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (j *Job) Unwrap() *Job {
	_tx, ok := j.config.driver.(*txDriver)
	if !ok {
		panic("ent: Job is not a transactional entity")
	}
	j.config.driver = _tx.drv
	return j
}

// String implements the fmt.Stringer.
func (j *Job) String() string {
	var builder strings.Builder
	builder.WriteString("Job(")
	builder.WriteString(fmt.Sprintf("id=%v, ", j.ID))
	builder.WriteString("program=")
	builder.WriteString(j.Program)
	builder.WriteString(", ")
	builder.WriteString("updatedAt=")
	builder.WriteString(j.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("createdAt=")
	builder.WriteString(j.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(j.Name)
	builder.WriteString(", ")
	builder.WriteString("updatedById=")
	builder.WriteString(fmt.Sprintf("%v", j.UpdatedById))
	builder.WriteString(", ")
	builder.WriteString("createdById=")
	builder.WriteString(fmt.Sprintf("%v", j.CreatedById))
	builder.WriteString(", ")
	builder.WriteString("categorySubId=")
	builder.WriteString(fmt.Sprintf("%v", j.CategorySubId))
	builder.WriteByte(')')
	return builder.String()
}

// Jobs is a parsable slice of Job.
type Jobs []*Job
