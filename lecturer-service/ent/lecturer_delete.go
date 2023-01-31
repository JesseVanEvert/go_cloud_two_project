// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"lecturer/ent/lecturer"
	"lecturer/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LecturerDelete is the builder for deleting a Lecturer entity.
type LecturerDelete struct {
	config
	hooks    []Hook
	mutation *LecturerMutation
}

// Where appends a list predicates to the LecturerDelete builder.
func (ld *LecturerDelete) Where(ps ...predicate.Lecturer) *LecturerDelete {
	ld.mutation.Where(ps...)
	return ld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ld *LecturerDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, LecturerMutation](ctx, ld.sqlExec, ld.mutation, ld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ld *LecturerDelete) ExecX(ctx context.Context) int {
	n, err := ld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ld *LecturerDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: lecturer.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lecturer.FieldID,
			},
		},
	}
	if ps := ld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ld.mutation.done = true
	return affected, err
}

// LecturerDeleteOne is the builder for deleting a single Lecturer entity.
type LecturerDeleteOne struct {
	ld *LecturerDelete
}

// Where appends a list predicates to the LecturerDelete builder.
func (ldo *LecturerDeleteOne) Where(ps ...predicate.Lecturer) *LecturerDeleteOne {
	ldo.ld.mutation.Where(ps...)
	return ldo
}

// Exec executes the deletion query.
func (ldo *LecturerDeleteOne) Exec(ctx context.Context) error {
	n, err := ldo.ld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{lecturer.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ldo *LecturerDeleteOne) ExecX(ctx context.Context) {
	if err := ldo.Exec(ctx); err != nil {
		panic(err)
	}
}
