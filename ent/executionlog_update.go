// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"datacatAgent/ent/executionlog"
	"datacatAgent/ent/predicate"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ExecutionLogUpdate is the builder for updating ExecutionLog entities.
type ExecutionLogUpdate struct {
	config
	hooks    []Hook
	mutation *ExecutionLogMutation
}

// Where appends a list predicates to the ExecutionLogUpdate builder.
func (elu *ExecutionLogUpdate) Where(ps ...predicate.ExecutionLog) *ExecutionLogUpdate {
	elu.mutation.Where(ps...)
	return elu
}

// SetStatus sets the "status" field.
func (elu *ExecutionLogUpdate) SetStatus(i int) *ExecutionLogUpdate {
	elu.mutation.ResetStatus()
	elu.mutation.SetStatus(i)
	return elu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (elu *ExecutionLogUpdate) SetNillableStatus(i *int) *ExecutionLogUpdate {
	if i != nil {
		elu.SetStatus(*i)
	}
	return elu
}

// AddStatus adds i to the "status" field.
func (elu *ExecutionLogUpdate) AddStatus(i int) *ExecutionLogUpdate {
	elu.mutation.AddStatus(i)
	return elu
}

// SetResult sets the "result" field.
func (elu *ExecutionLogUpdate) SetResult(s string) *ExecutionLogUpdate {
	elu.mutation.SetResult(s)
	return elu
}

// SetNillableResult sets the "result" field if the given value is not nil.
func (elu *ExecutionLogUpdate) SetNillableResult(s *string) *ExecutionLogUpdate {
	if s != nil {
		elu.SetResult(*s)
	}
	return elu
}

// SetScriptId sets the "scriptId" field.
func (elu *ExecutionLogUpdate) SetScriptId(i int) *ExecutionLogUpdate {
	elu.mutation.ResetScriptId()
	elu.mutation.SetScriptId(i)
	return elu
}

// SetNillableScriptId sets the "scriptId" field if the given value is not nil.
func (elu *ExecutionLogUpdate) SetNillableScriptId(i *int) *ExecutionLogUpdate {
	if i != nil {
		elu.SetScriptId(*i)
	}
	return elu
}

// AddScriptId adds i to the "scriptId" field.
func (elu *ExecutionLogUpdate) AddScriptId(i int) *ExecutionLogUpdate {
	elu.mutation.AddScriptId(i)
	return elu
}

// Mutation returns the ExecutionLogMutation object of the builder.
func (elu *ExecutionLogUpdate) Mutation() *ExecutionLogMutation {
	return elu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (elu *ExecutionLogUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, elu.sqlSave, elu.mutation, elu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (elu *ExecutionLogUpdate) SaveX(ctx context.Context) int {
	affected, err := elu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (elu *ExecutionLogUpdate) Exec(ctx context.Context) error {
	_, err := elu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (elu *ExecutionLogUpdate) ExecX(ctx context.Context) {
	if err := elu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (elu *ExecutionLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(executionlog.Table, executionlog.Columns, sqlgraph.NewFieldSpec(executionlog.FieldID, field.TypeInt64))
	if ps := elu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := elu.mutation.Status(); ok {
		_spec.SetField(executionlog.FieldStatus, field.TypeInt, value)
	}
	if value, ok := elu.mutation.AddedStatus(); ok {
		_spec.AddField(executionlog.FieldStatus, field.TypeInt, value)
	}
	if value, ok := elu.mutation.Result(); ok {
		_spec.SetField(executionlog.FieldResult, field.TypeString, value)
	}
	if value, ok := elu.mutation.ScriptId(); ok {
		_spec.SetField(executionlog.FieldScriptId, field.TypeInt, value)
	}
	if value, ok := elu.mutation.AddedScriptId(); ok {
		_spec.AddField(executionlog.FieldScriptId, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, elu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{executionlog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	elu.mutation.done = true
	return n, nil
}

// ExecutionLogUpdateOne is the builder for updating a single ExecutionLog entity.
type ExecutionLogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ExecutionLogMutation
}

// SetStatus sets the "status" field.
func (eluo *ExecutionLogUpdateOne) SetStatus(i int) *ExecutionLogUpdateOne {
	eluo.mutation.ResetStatus()
	eluo.mutation.SetStatus(i)
	return eluo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (eluo *ExecutionLogUpdateOne) SetNillableStatus(i *int) *ExecutionLogUpdateOne {
	if i != nil {
		eluo.SetStatus(*i)
	}
	return eluo
}

// AddStatus adds i to the "status" field.
func (eluo *ExecutionLogUpdateOne) AddStatus(i int) *ExecutionLogUpdateOne {
	eluo.mutation.AddStatus(i)
	return eluo
}

// SetResult sets the "result" field.
func (eluo *ExecutionLogUpdateOne) SetResult(s string) *ExecutionLogUpdateOne {
	eluo.mutation.SetResult(s)
	return eluo
}

// SetNillableResult sets the "result" field if the given value is not nil.
func (eluo *ExecutionLogUpdateOne) SetNillableResult(s *string) *ExecutionLogUpdateOne {
	if s != nil {
		eluo.SetResult(*s)
	}
	return eluo
}

// SetScriptId sets the "scriptId" field.
func (eluo *ExecutionLogUpdateOne) SetScriptId(i int) *ExecutionLogUpdateOne {
	eluo.mutation.ResetScriptId()
	eluo.mutation.SetScriptId(i)
	return eluo
}

// SetNillableScriptId sets the "scriptId" field if the given value is not nil.
func (eluo *ExecutionLogUpdateOne) SetNillableScriptId(i *int) *ExecutionLogUpdateOne {
	if i != nil {
		eluo.SetScriptId(*i)
	}
	return eluo
}

// AddScriptId adds i to the "scriptId" field.
func (eluo *ExecutionLogUpdateOne) AddScriptId(i int) *ExecutionLogUpdateOne {
	eluo.mutation.AddScriptId(i)
	return eluo
}

// Mutation returns the ExecutionLogMutation object of the builder.
func (eluo *ExecutionLogUpdateOne) Mutation() *ExecutionLogMutation {
	return eluo.mutation
}

// Where appends a list predicates to the ExecutionLogUpdate builder.
func (eluo *ExecutionLogUpdateOne) Where(ps ...predicate.ExecutionLog) *ExecutionLogUpdateOne {
	eluo.mutation.Where(ps...)
	return eluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (eluo *ExecutionLogUpdateOne) Select(field string, fields ...string) *ExecutionLogUpdateOne {
	eluo.fields = append([]string{field}, fields...)
	return eluo
}

// Save executes the query and returns the updated ExecutionLog entity.
func (eluo *ExecutionLogUpdateOne) Save(ctx context.Context) (*ExecutionLog, error) {
	return withHooks(ctx, eluo.sqlSave, eluo.mutation, eluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eluo *ExecutionLogUpdateOne) SaveX(ctx context.Context) *ExecutionLog {
	node, err := eluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (eluo *ExecutionLogUpdateOne) Exec(ctx context.Context) error {
	_, err := eluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eluo *ExecutionLogUpdateOne) ExecX(ctx context.Context) {
	if err := eluo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (eluo *ExecutionLogUpdateOne) sqlSave(ctx context.Context) (_node *ExecutionLog, err error) {
	_spec := sqlgraph.NewUpdateSpec(executionlog.Table, executionlog.Columns, sqlgraph.NewFieldSpec(executionlog.FieldID, field.TypeInt64))
	id, ok := eluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ExecutionLog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := eluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, executionlog.FieldID)
		for _, f := range fields {
			if !executionlog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != executionlog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := eluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eluo.mutation.Status(); ok {
		_spec.SetField(executionlog.FieldStatus, field.TypeInt, value)
	}
	if value, ok := eluo.mutation.AddedStatus(); ok {
		_spec.AddField(executionlog.FieldStatus, field.TypeInt, value)
	}
	if value, ok := eluo.mutation.Result(); ok {
		_spec.SetField(executionlog.FieldResult, field.TypeString, value)
	}
	if value, ok := eluo.mutation.ScriptId(); ok {
		_spec.SetField(executionlog.FieldScriptId, field.TypeInt, value)
	}
	if value, ok := eluo.mutation.AddedScriptId(); ok {
		_spec.AddField(executionlog.FieldScriptId, field.TypeInt, value)
	}
	_node = &ExecutionLog{config: eluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, eluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{executionlog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	eluo.mutation.done = true
	return _node, nil
}
