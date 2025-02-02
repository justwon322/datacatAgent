// Code generated by ent, DO NOT EDIT.

package message

import (
	"datacatAgent/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Message {
	return predicate.Message(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Message {
	return predicate.Message(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Message {
	return predicate.Message(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Message {
	return predicate.Message(sql.FieldLTE(FieldID, id))
}

// Payload applies equality check predicate on the "payload" field. It's identical to PayloadEQ.
func Payload(v string) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldPayload, v))
}

// Sent applies equality check predicate on the "sent" field. It's identical to SentEQ.
func Sent(v int) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldSent, v))
}

// Resent applies equality check predicate on the "resent" field. It's identical to ResentEQ.
func Resent(v int) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldResent, v))
}

// SentAt applies equality check predicate on the "sentAt" field. It's identical to SentAtEQ.
func SentAt(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldSentAt, v))
}

// JobId applies equality check predicate on the "jobId" field. It's identical to JobIdEQ.
func JobId(v int) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldJobId, v))
}

// PayloadEQ applies the EQ predicate on the "payload" field.
func PayloadEQ(v string) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldPayload, v))
}

// PayloadNEQ applies the NEQ predicate on the "payload" field.
func PayloadNEQ(v string) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldPayload, v))
}

// PayloadIn applies the In predicate on the "payload" field.
func PayloadIn(vs ...string) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldPayload, vs...))
}

// PayloadNotIn applies the NotIn predicate on the "payload" field.
func PayloadNotIn(vs ...string) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldPayload, vs...))
}

// PayloadGT applies the GT predicate on the "payload" field.
func PayloadGT(v string) predicate.Message {
	return predicate.Message(sql.FieldGT(FieldPayload, v))
}

// PayloadGTE applies the GTE predicate on the "payload" field.
func PayloadGTE(v string) predicate.Message {
	return predicate.Message(sql.FieldGTE(FieldPayload, v))
}

// PayloadLT applies the LT predicate on the "payload" field.
func PayloadLT(v string) predicate.Message {
	return predicate.Message(sql.FieldLT(FieldPayload, v))
}

// PayloadLTE applies the LTE predicate on the "payload" field.
func PayloadLTE(v string) predicate.Message {
	return predicate.Message(sql.FieldLTE(FieldPayload, v))
}

// PayloadContains applies the Contains predicate on the "payload" field.
func PayloadContains(v string) predicate.Message {
	return predicate.Message(sql.FieldContains(FieldPayload, v))
}

// PayloadHasPrefix applies the HasPrefix predicate on the "payload" field.
func PayloadHasPrefix(v string) predicate.Message {
	return predicate.Message(sql.FieldHasPrefix(FieldPayload, v))
}

// PayloadHasSuffix applies the HasSuffix predicate on the "payload" field.
func PayloadHasSuffix(v string) predicate.Message {
	return predicate.Message(sql.FieldHasSuffix(FieldPayload, v))
}

// PayloadEqualFold applies the EqualFold predicate on the "payload" field.
func PayloadEqualFold(v string) predicate.Message {
	return predicate.Message(sql.FieldEqualFold(FieldPayload, v))
}

// PayloadContainsFold applies the ContainsFold predicate on the "payload" field.
func PayloadContainsFold(v string) predicate.Message {
	return predicate.Message(sql.FieldContainsFold(FieldPayload, v))
}

// SentEQ applies the EQ predicate on the "sent" field.
func SentEQ(v int) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldSent, v))
}

// SentNEQ applies the NEQ predicate on the "sent" field.
func SentNEQ(v int) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldSent, v))
}

// SentIn applies the In predicate on the "sent" field.
func SentIn(vs ...int) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldSent, vs...))
}

// SentNotIn applies the NotIn predicate on the "sent" field.
func SentNotIn(vs ...int) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldSent, vs...))
}

// SentGT applies the GT predicate on the "sent" field.
func SentGT(v int) predicate.Message {
	return predicate.Message(sql.FieldGT(FieldSent, v))
}

// SentGTE applies the GTE predicate on the "sent" field.
func SentGTE(v int) predicate.Message {
	return predicate.Message(sql.FieldGTE(FieldSent, v))
}

// SentLT applies the LT predicate on the "sent" field.
func SentLT(v int) predicate.Message {
	return predicate.Message(sql.FieldLT(FieldSent, v))
}

// SentLTE applies the LTE predicate on the "sent" field.
func SentLTE(v int) predicate.Message {
	return predicate.Message(sql.FieldLTE(FieldSent, v))
}

// ResentEQ applies the EQ predicate on the "resent" field.
func ResentEQ(v int) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldResent, v))
}

// ResentNEQ applies the NEQ predicate on the "resent" field.
func ResentNEQ(v int) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldResent, v))
}

// ResentIn applies the In predicate on the "resent" field.
func ResentIn(vs ...int) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldResent, vs...))
}

// ResentNotIn applies the NotIn predicate on the "resent" field.
func ResentNotIn(vs ...int) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldResent, vs...))
}

// ResentGT applies the GT predicate on the "resent" field.
func ResentGT(v int) predicate.Message {
	return predicate.Message(sql.FieldGT(FieldResent, v))
}

// ResentGTE applies the GTE predicate on the "resent" field.
func ResentGTE(v int) predicate.Message {
	return predicate.Message(sql.FieldGTE(FieldResent, v))
}

// ResentLT applies the LT predicate on the "resent" field.
func ResentLT(v int) predicate.Message {
	return predicate.Message(sql.FieldLT(FieldResent, v))
}

// ResentLTE applies the LTE predicate on the "resent" field.
func ResentLTE(v int) predicate.Message {
	return predicate.Message(sql.FieldLTE(FieldResent, v))
}

// SentAtEQ applies the EQ predicate on the "sentAt" field.
func SentAtEQ(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldSentAt, v))
}

// SentAtNEQ applies the NEQ predicate on the "sentAt" field.
func SentAtNEQ(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldSentAt, v))
}

// SentAtIn applies the In predicate on the "sentAt" field.
func SentAtIn(vs ...time.Time) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldSentAt, vs...))
}

// SentAtNotIn applies the NotIn predicate on the "sentAt" field.
func SentAtNotIn(vs ...time.Time) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldSentAt, vs...))
}

// SentAtGT applies the GT predicate on the "sentAt" field.
func SentAtGT(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldGT(FieldSentAt, v))
}

// SentAtGTE applies the GTE predicate on the "sentAt" field.
func SentAtGTE(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldGTE(FieldSentAt, v))
}

// SentAtLT applies the LT predicate on the "sentAt" field.
func SentAtLT(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldLT(FieldSentAt, v))
}

// SentAtLTE applies the LTE predicate on the "sentAt" field.
func SentAtLTE(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldLTE(FieldSentAt, v))
}

// JobIdEQ applies the EQ predicate on the "jobId" field.
func JobIdEQ(v int) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldJobId, v))
}

// JobIdNEQ applies the NEQ predicate on the "jobId" field.
func JobIdNEQ(v int) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldJobId, v))
}

// JobIdIn applies the In predicate on the "jobId" field.
func JobIdIn(vs ...int) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldJobId, vs...))
}

// JobIdNotIn applies the NotIn predicate on the "jobId" field.
func JobIdNotIn(vs ...int) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldJobId, vs...))
}

// JobIdGT applies the GT predicate on the "jobId" field.
func JobIdGT(v int) predicate.Message {
	return predicate.Message(sql.FieldGT(FieldJobId, v))
}

// JobIdGTE applies the GTE predicate on the "jobId" field.
func JobIdGTE(v int) predicate.Message {
	return predicate.Message(sql.FieldGTE(FieldJobId, v))
}

// JobIdLT applies the LT predicate on the "jobId" field.
func JobIdLT(v int) predicate.Message {
	return predicate.Message(sql.FieldLT(FieldJobId, v))
}

// JobIdLTE applies the LTE predicate on the "jobId" field.
func JobIdLTE(v int) predicate.Message {
	return predicate.Message(sql.FieldLTE(FieldJobId, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Message) predicate.Message {
	return predicate.Message(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Message) predicate.Message {
	return predicate.Message(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Message) predicate.Message {
	return predicate.Message(sql.NotPredicates(p))
}
