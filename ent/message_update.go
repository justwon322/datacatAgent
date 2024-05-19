// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"datacatAgent/ent/message"
	"datacatAgent/ent/predicate"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageUpdate is the builder for updating Message entities.
type MessageUpdate struct {
	config
	hooks    []Hook
	mutation *MessageMutation
}

// Where appends a list predicates to the MessageUpdate builder.
func (mu *MessageUpdate) Where(ps ...predicate.Message) *MessageUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetPayload sets the "payload" field.
func (mu *MessageUpdate) SetPayload(s string) *MessageUpdate {
	mu.mutation.SetPayload(s)
	return mu
}

// SetNillablePayload sets the "payload" field if the given value is not nil.
func (mu *MessageUpdate) SetNillablePayload(s *string) *MessageUpdate {
	if s != nil {
		mu.SetPayload(*s)
	}
	return mu
}

// SetSent sets the "sent" field.
func (mu *MessageUpdate) SetSent(i int) *MessageUpdate {
	mu.mutation.ResetSent()
	mu.mutation.SetSent(i)
	return mu
}

// SetNillableSent sets the "sent" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableSent(i *int) *MessageUpdate {
	if i != nil {
		mu.SetSent(*i)
	}
	return mu
}

// AddSent adds i to the "sent" field.
func (mu *MessageUpdate) AddSent(i int) *MessageUpdate {
	mu.mutation.AddSent(i)
	return mu
}

// SetResent sets the "resent" field.
func (mu *MessageUpdate) SetResent(i int) *MessageUpdate {
	mu.mutation.ResetResent()
	mu.mutation.SetResent(i)
	return mu
}

// SetNillableResent sets the "resent" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableResent(i *int) *MessageUpdate {
	if i != nil {
		mu.SetResent(*i)
	}
	return mu
}

// AddResent adds i to the "resent" field.
func (mu *MessageUpdate) AddResent(i int) *MessageUpdate {
	mu.mutation.AddResent(i)
	return mu
}

// SetSentAt sets the "sentAt" field.
func (mu *MessageUpdate) SetSentAt(t time.Time) *MessageUpdate {
	mu.mutation.SetSentAt(t)
	return mu
}

// SetJobId sets the "jobId" field.
func (mu *MessageUpdate) SetJobId(i int) *MessageUpdate {
	mu.mutation.ResetJobId()
	mu.mutation.SetJobId(i)
	return mu
}

// SetNillableJobId sets the "jobId" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableJobId(i *int) *MessageUpdate {
	if i != nil {
		mu.SetJobId(*i)
	}
	return mu
}

// AddJobId adds i to the "jobId" field.
func (mu *MessageUpdate) AddJobId(i int) *MessageUpdate {
	mu.mutation.AddJobId(i)
	return mu
}

// Mutation returns the MessageMutation object of the builder.
func (mu *MessageUpdate) Mutation() *MessageMutation {
	return mu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MessageUpdate) Save(ctx context.Context) (int, error) {
	mu.defaults()
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MessageUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MessageUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MessageUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MessageUpdate) defaults() {
	if _, ok := mu.mutation.SentAt(); !ok {
		v := message.UpdateDefaultSentAt()
		mu.mutation.SetSentAt(v)
	}
}

func (mu *MessageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(message.Table, message.Columns, sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt64))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Payload(); ok {
		_spec.SetField(message.FieldPayload, field.TypeString, value)
	}
	if value, ok := mu.mutation.Sent(); ok {
		_spec.SetField(message.FieldSent, field.TypeInt, value)
	}
	if value, ok := mu.mutation.AddedSent(); ok {
		_spec.AddField(message.FieldSent, field.TypeInt, value)
	}
	if value, ok := mu.mutation.Resent(); ok {
		_spec.SetField(message.FieldResent, field.TypeInt, value)
	}
	if value, ok := mu.mutation.AddedResent(); ok {
		_spec.AddField(message.FieldResent, field.TypeInt, value)
	}
	if value, ok := mu.mutation.SentAt(); ok {
		_spec.SetField(message.FieldSentAt, field.TypeTime, value)
	}
	if value, ok := mu.mutation.JobId(); ok {
		_spec.SetField(message.FieldJobId, field.TypeInt, value)
	}
	if value, ok := mu.mutation.AddedJobId(); ok {
		_spec.AddField(message.FieldJobId, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MessageUpdateOne is the builder for updating a single Message entity.
type MessageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MessageMutation
}

// SetPayload sets the "payload" field.
func (muo *MessageUpdateOne) SetPayload(s string) *MessageUpdateOne {
	muo.mutation.SetPayload(s)
	return muo
}

// SetNillablePayload sets the "payload" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillablePayload(s *string) *MessageUpdateOne {
	if s != nil {
		muo.SetPayload(*s)
	}
	return muo
}

// SetSent sets the "sent" field.
func (muo *MessageUpdateOne) SetSent(i int) *MessageUpdateOne {
	muo.mutation.ResetSent()
	muo.mutation.SetSent(i)
	return muo
}

// SetNillableSent sets the "sent" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableSent(i *int) *MessageUpdateOne {
	if i != nil {
		muo.SetSent(*i)
	}
	return muo
}

// AddSent adds i to the "sent" field.
func (muo *MessageUpdateOne) AddSent(i int) *MessageUpdateOne {
	muo.mutation.AddSent(i)
	return muo
}

// SetResent sets the "resent" field.
func (muo *MessageUpdateOne) SetResent(i int) *MessageUpdateOne {
	muo.mutation.ResetResent()
	muo.mutation.SetResent(i)
	return muo
}

// SetNillableResent sets the "resent" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableResent(i *int) *MessageUpdateOne {
	if i != nil {
		muo.SetResent(*i)
	}
	return muo
}

// AddResent adds i to the "resent" field.
func (muo *MessageUpdateOne) AddResent(i int) *MessageUpdateOne {
	muo.mutation.AddResent(i)
	return muo
}

// SetSentAt sets the "sentAt" field.
func (muo *MessageUpdateOne) SetSentAt(t time.Time) *MessageUpdateOne {
	muo.mutation.SetSentAt(t)
	return muo
}

// SetJobId sets the "jobId" field.
func (muo *MessageUpdateOne) SetJobId(i int) *MessageUpdateOne {
	muo.mutation.ResetJobId()
	muo.mutation.SetJobId(i)
	return muo
}

// SetNillableJobId sets the "jobId" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableJobId(i *int) *MessageUpdateOne {
	if i != nil {
		muo.SetJobId(*i)
	}
	return muo
}

// AddJobId adds i to the "jobId" field.
func (muo *MessageUpdateOne) AddJobId(i int) *MessageUpdateOne {
	muo.mutation.AddJobId(i)
	return muo
}

// Mutation returns the MessageMutation object of the builder.
func (muo *MessageUpdateOne) Mutation() *MessageMutation {
	return muo.mutation
}

// Where appends a list predicates to the MessageUpdate builder.
func (muo *MessageUpdateOne) Where(ps ...predicate.Message) *MessageUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MessageUpdateOne) Select(field string, fields ...string) *MessageUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Message entity.
func (muo *MessageUpdateOne) Save(ctx context.Context) (*Message, error) {
	muo.defaults()
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MessageUpdateOne) SaveX(ctx context.Context) *Message {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MessageUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MessageUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MessageUpdateOne) defaults() {
	if _, ok := muo.mutation.SentAt(); !ok {
		v := message.UpdateDefaultSentAt()
		muo.mutation.SetSentAt(v)
	}
}

func (muo *MessageUpdateOne) sqlSave(ctx context.Context) (_node *Message, err error) {
	_spec := sqlgraph.NewUpdateSpec(message.Table, message.Columns, sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt64))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Message.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, message.FieldID)
		for _, f := range fields {
			if !message.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != message.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Payload(); ok {
		_spec.SetField(message.FieldPayload, field.TypeString, value)
	}
	if value, ok := muo.mutation.Sent(); ok {
		_spec.SetField(message.FieldSent, field.TypeInt, value)
	}
	if value, ok := muo.mutation.AddedSent(); ok {
		_spec.AddField(message.FieldSent, field.TypeInt, value)
	}
	if value, ok := muo.mutation.Resent(); ok {
		_spec.SetField(message.FieldResent, field.TypeInt, value)
	}
	if value, ok := muo.mutation.AddedResent(); ok {
		_spec.AddField(message.FieldResent, field.TypeInt, value)
	}
	if value, ok := muo.mutation.SentAt(); ok {
		_spec.SetField(message.FieldSentAt, field.TypeTime, value)
	}
	if value, ok := muo.mutation.JobId(); ok {
		_spec.SetField(message.FieldJobId, field.TypeInt, value)
	}
	if value, ok := muo.mutation.AddedJobId(); ok {
		_spec.AddField(message.FieldJobId, field.TypeInt, value)
	}
	_node = &Message{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
