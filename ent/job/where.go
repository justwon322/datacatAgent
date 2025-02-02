// Code generated by ent, DO NOT EDIT.

package job

import (
	"datacatAgent/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.Job {
	return predicate.Job(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.Job {
	return predicate.Job(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int32) predicate.Job {
	return predicate.Job(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int32) predicate.Job {
	return predicate.Job(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.Job {
	return predicate.Job(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.Job {
	return predicate.Job(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.Job {
	return predicate.Job(sql.FieldLTE(FieldID, id))
}

// Program applies equality check predicate on the "program" field. It's identical to ProgramEQ.
func Program(v string) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldProgram, v))
}

// UpdatedAt applies equality check predicate on the "updatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldCreatedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldName, v))
}

// UpdatedById applies equality check predicate on the "updatedById" field. It's identical to UpdatedByIdEQ.
func UpdatedById(v int) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldUpdatedById, v))
}

// CreatedById applies equality check predicate on the "createdById" field. It's identical to CreatedByIdEQ.
func CreatedById(v int) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldCreatedById, v))
}

// CategorySubId applies equality check predicate on the "categorySubId" field. It's identical to CategorySubIdEQ.
func CategorySubId(v int) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldCategorySubId, v))
}

// ProgramEQ applies the EQ predicate on the "program" field.
func ProgramEQ(v string) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldProgram, v))
}

// ProgramNEQ applies the NEQ predicate on the "program" field.
func ProgramNEQ(v string) predicate.Job {
	return predicate.Job(sql.FieldNEQ(FieldProgram, v))
}

// ProgramIn applies the In predicate on the "program" field.
func ProgramIn(vs ...string) predicate.Job {
	return predicate.Job(sql.FieldIn(FieldProgram, vs...))
}

// ProgramNotIn applies the NotIn predicate on the "program" field.
func ProgramNotIn(vs ...string) predicate.Job {
	return predicate.Job(sql.FieldNotIn(FieldProgram, vs...))
}

// ProgramGT applies the GT predicate on the "program" field.
func ProgramGT(v string) predicate.Job {
	return predicate.Job(sql.FieldGT(FieldProgram, v))
}

// ProgramGTE applies the GTE predicate on the "program" field.
func ProgramGTE(v string) predicate.Job {
	return predicate.Job(sql.FieldGTE(FieldProgram, v))
}

// ProgramLT applies the LT predicate on the "program" field.
func ProgramLT(v string) predicate.Job {
	return predicate.Job(sql.FieldLT(FieldProgram, v))
}

// ProgramLTE applies the LTE predicate on the "program" field.
func ProgramLTE(v string) predicate.Job {
	return predicate.Job(sql.FieldLTE(FieldProgram, v))
}

// ProgramContains applies the Contains predicate on the "program" field.
func ProgramContains(v string) predicate.Job {
	return predicate.Job(sql.FieldContains(FieldProgram, v))
}

// ProgramHasPrefix applies the HasPrefix predicate on the "program" field.
func ProgramHasPrefix(v string) predicate.Job {
	return predicate.Job(sql.FieldHasPrefix(FieldProgram, v))
}

// ProgramHasSuffix applies the HasSuffix predicate on the "program" field.
func ProgramHasSuffix(v string) predicate.Job {
	return predicate.Job(sql.FieldHasSuffix(FieldProgram, v))
}

// ProgramEqualFold applies the EqualFold predicate on the "program" field.
func ProgramEqualFold(v string) predicate.Job {
	return predicate.Job(sql.FieldEqualFold(FieldProgram, v))
}

// ProgramContainsFold applies the ContainsFold predicate on the "program" field.
func ProgramContainsFold(v string) predicate.Job {
	return predicate.Job(sql.FieldContainsFold(FieldProgram, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.Job {
	return predicate.Job(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Job {
	return predicate.Job(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updatedAt" field.
func UpdatedAtGT(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updatedAt" field.
func UpdatedAtLT(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.Job {
	return predicate.Job(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Job {
	return predicate.Job(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldLTE(FieldCreatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Job {
	return predicate.Job(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Job {
	return predicate.Job(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Job {
	return predicate.Job(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Job {
	return predicate.Job(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Job {
	return predicate.Job(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Job {
	return predicate.Job(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Job {
	return predicate.Job(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Job {
	return predicate.Job(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Job {
	return predicate.Job(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Job {
	return predicate.Job(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Job {
	return predicate.Job(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Job {
	return predicate.Job(sql.FieldContainsFold(FieldName, v))
}

// UpdatedByIdEQ applies the EQ predicate on the "updatedById" field.
func UpdatedByIdEQ(v int) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldUpdatedById, v))
}

// UpdatedByIdNEQ applies the NEQ predicate on the "updatedById" field.
func UpdatedByIdNEQ(v int) predicate.Job {
	return predicate.Job(sql.FieldNEQ(FieldUpdatedById, v))
}

// UpdatedByIdIn applies the In predicate on the "updatedById" field.
func UpdatedByIdIn(vs ...int) predicate.Job {
	return predicate.Job(sql.FieldIn(FieldUpdatedById, vs...))
}

// UpdatedByIdNotIn applies the NotIn predicate on the "updatedById" field.
func UpdatedByIdNotIn(vs ...int) predicate.Job {
	return predicate.Job(sql.FieldNotIn(FieldUpdatedById, vs...))
}

// UpdatedByIdGT applies the GT predicate on the "updatedById" field.
func UpdatedByIdGT(v int) predicate.Job {
	return predicate.Job(sql.FieldGT(FieldUpdatedById, v))
}

// UpdatedByIdGTE applies the GTE predicate on the "updatedById" field.
func UpdatedByIdGTE(v int) predicate.Job {
	return predicate.Job(sql.FieldGTE(FieldUpdatedById, v))
}

// UpdatedByIdLT applies the LT predicate on the "updatedById" field.
func UpdatedByIdLT(v int) predicate.Job {
	return predicate.Job(sql.FieldLT(FieldUpdatedById, v))
}

// UpdatedByIdLTE applies the LTE predicate on the "updatedById" field.
func UpdatedByIdLTE(v int) predicate.Job {
	return predicate.Job(sql.FieldLTE(FieldUpdatedById, v))
}

// CreatedByIdEQ applies the EQ predicate on the "createdById" field.
func CreatedByIdEQ(v int) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldCreatedById, v))
}

// CreatedByIdNEQ applies the NEQ predicate on the "createdById" field.
func CreatedByIdNEQ(v int) predicate.Job {
	return predicate.Job(sql.FieldNEQ(FieldCreatedById, v))
}

// CreatedByIdIn applies the In predicate on the "createdById" field.
func CreatedByIdIn(vs ...int) predicate.Job {
	return predicate.Job(sql.FieldIn(FieldCreatedById, vs...))
}

// CreatedByIdNotIn applies the NotIn predicate on the "createdById" field.
func CreatedByIdNotIn(vs ...int) predicate.Job {
	return predicate.Job(sql.FieldNotIn(FieldCreatedById, vs...))
}

// CreatedByIdGT applies the GT predicate on the "createdById" field.
func CreatedByIdGT(v int) predicate.Job {
	return predicate.Job(sql.FieldGT(FieldCreatedById, v))
}

// CreatedByIdGTE applies the GTE predicate on the "createdById" field.
func CreatedByIdGTE(v int) predicate.Job {
	return predicate.Job(sql.FieldGTE(FieldCreatedById, v))
}

// CreatedByIdLT applies the LT predicate on the "createdById" field.
func CreatedByIdLT(v int) predicate.Job {
	return predicate.Job(sql.FieldLT(FieldCreatedById, v))
}

// CreatedByIdLTE applies the LTE predicate on the "createdById" field.
func CreatedByIdLTE(v int) predicate.Job {
	return predicate.Job(sql.FieldLTE(FieldCreatedById, v))
}

// CategorySubIdEQ applies the EQ predicate on the "categorySubId" field.
func CategorySubIdEQ(v int) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldCategorySubId, v))
}

// CategorySubIdNEQ applies the NEQ predicate on the "categorySubId" field.
func CategorySubIdNEQ(v int) predicate.Job {
	return predicate.Job(sql.FieldNEQ(FieldCategorySubId, v))
}

// CategorySubIdIn applies the In predicate on the "categorySubId" field.
func CategorySubIdIn(vs ...int) predicate.Job {
	return predicate.Job(sql.FieldIn(FieldCategorySubId, vs...))
}

// CategorySubIdNotIn applies the NotIn predicate on the "categorySubId" field.
func CategorySubIdNotIn(vs ...int) predicate.Job {
	return predicate.Job(sql.FieldNotIn(FieldCategorySubId, vs...))
}

// CategorySubIdGT applies the GT predicate on the "categorySubId" field.
func CategorySubIdGT(v int) predicate.Job {
	return predicate.Job(sql.FieldGT(FieldCategorySubId, v))
}

// CategorySubIdGTE applies the GTE predicate on the "categorySubId" field.
func CategorySubIdGTE(v int) predicate.Job {
	return predicate.Job(sql.FieldGTE(FieldCategorySubId, v))
}

// CategorySubIdLT applies the LT predicate on the "categorySubId" field.
func CategorySubIdLT(v int) predicate.Job {
	return predicate.Job(sql.FieldLT(FieldCategorySubId, v))
}

// CategorySubIdLTE applies the LTE predicate on the "categorySubId" field.
func CategorySubIdLTE(v int) predicate.Job {
	return predicate.Job(sql.FieldLTE(FieldCategorySubId, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Job) predicate.Job {
	return predicate.Job(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Job) predicate.Job {
	return predicate.Job(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Job) predicate.Job {
	return predicate.Job(sql.NotPredicates(p))
}
