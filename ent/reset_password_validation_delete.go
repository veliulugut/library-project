// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"library/ent/predicate"
	"library/ent/reset_password_validation"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ResetPasswordValidationDelete is the builder for deleting a Reset_Password_Validation entity.
type ResetPasswordValidationDelete struct {
	config
	hooks    []Hook
	mutation *ResetPasswordValidationMutation
}

// Where appends a list predicates to the ResetPasswordValidationDelete builder.
func (rpvd *ResetPasswordValidationDelete) Where(ps ...predicate.Reset_Password_Validation) *ResetPasswordValidationDelete {
	rpvd.mutation.Where(ps...)
	return rpvd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rpvd *ResetPasswordValidationDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rpvd.sqlExec, rpvd.mutation, rpvd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rpvd *ResetPasswordValidationDelete) ExecX(ctx context.Context) int {
	n, err := rpvd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rpvd *ResetPasswordValidationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(reset_password_validation.Table, sqlgraph.NewFieldSpec(reset_password_validation.FieldID, field.TypeInt))
	if ps := rpvd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rpvd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rpvd.mutation.done = true
	return affected, err
}

// ResetPasswordValidationDeleteOne is the builder for deleting a single Reset_Password_Validation entity.
type ResetPasswordValidationDeleteOne struct {
	rpvd *ResetPasswordValidationDelete
}

// Where appends a list predicates to the ResetPasswordValidationDelete builder.
func (rpvdo *ResetPasswordValidationDeleteOne) Where(ps ...predicate.Reset_Password_Validation) *ResetPasswordValidationDeleteOne {
	rpvdo.rpvd.mutation.Where(ps...)
	return rpvdo
}

// Exec executes the deletion query.
func (rpvdo *ResetPasswordValidationDeleteOne) Exec(ctx context.Context) error {
	n, err := rpvdo.rpvd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{reset_password_validation.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rpvdo *ResetPasswordValidationDeleteOne) ExecX(ctx context.Context) {
	if err := rpvdo.Exec(ctx); err != nil {
		panic(err)
	}
}
