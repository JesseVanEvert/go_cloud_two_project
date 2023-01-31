// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"lecturer/ent/class"
	"lecturer/ent/classlecturer"
	"lecturer/ent/lecturer"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ClassLecturerCreate is the builder for creating a ClassLecturer entity.
type ClassLecturerCreate struct {
	config
	mutation *ClassLecturerMutation
	hooks    []Hook
}

// SetDeletedAt sets the "deleted_at" field.
func (clc *ClassLecturerCreate) SetDeletedAt(t time.Time) *ClassLecturerCreate {
	clc.mutation.SetDeletedAt(t)
	return clc
}

// SetClassID sets the "class" edge to the Class entity by ID.
func (clc *ClassLecturerCreate) SetClassID(id int) *ClassLecturerCreate {
	clc.mutation.SetClassID(id)
	return clc
}

// SetNillableClassID sets the "class" edge to the Class entity by ID if the given value is not nil.
func (clc *ClassLecturerCreate) SetNillableClassID(id *int) *ClassLecturerCreate {
	if id != nil {
		clc = clc.SetClassID(*id)
	}
	return clc
}

// SetClass sets the "class" edge to the Class entity.
func (clc *ClassLecturerCreate) SetClass(c *Class) *ClassLecturerCreate {
	return clc.SetClassID(c.ID)
}

// SetLecturerID sets the "lecturer" edge to the Lecturer entity by ID.
func (clc *ClassLecturerCreate) SetLecturerID(id int) *ClassLecturerCreate {
	clc.mutation.SetLecturerID(id)
	return clc
}

// SetNillableLecturerID sets the "lecturer" edge to the Lecturer entity by ID if the given value is not nil.
func (clc *ClassLecturerCreate) SetNillableLecturerID(id *int) *ClassLecturerCreate {
	if id != nil {
		clc = clc.SetLecturerID(*id)
	}
	return clc
}

// SetLecturer sets the "lecturer" edge to the Lecturer entity.
func (clc *ClassLecturerCreate) SetLecturer(l *Lecturer) *ClassLecturerCreate {
	return clc.SetLecturerID(l.ID)
}

// Mutation returns the ClassLecturerMutation object of the builder.
func (clc *ClassLecturerCreate) Mutation() *ClassLecturerMutation {
	return clc.mutation
}

// Save creates the ClassLecturer in the database.
func (clc *ClassLecturerCreate) Save(ctx context.Context) (*ClassLecturer, error) {
	return withHooks[*ClassLecturer, ClassLecturerMutation](ctx, clc.sqlSave, clc.mutation, clc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (clc *ClassLecturerCreate) SaveX(ctx context.Context) *ClassLecturer {
	v, err := clc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (clc *ClassLecturerCreate) Exec(ctx context.Context) error {
	_, err := clc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (clc *ClassLecturerCreate) ExecX(ctx context.Context) {
	if err := clc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (clc *ClassLecturerCreate) check() error {
	if _, ok := clc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "ClassLecturer.deleted_at"`)}
	}
	return nil
}

func (clc *ClassLecturerCreate) sqlSave(ctx context.Context) (*ClassLecturer, error) {
	if err := clc.check(); err != nil {
		return nil, err
	}
	_node, _spec := clc.createSpec()
	if err := sqlgraph.CreateNode(ctx, clc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	clc.mutation.id = &_node.ID
	clc.mutation.done = true
	return _node, nil
}

func (clc *ClassLecturerCreate) createSpec() (*ClassLecturer, *sqlgraph.CreateSpec) {
	var (
		_node = &ClassLecturer{config: clc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: classlecturer.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: classlecturer.FieldID,
			},
		}
	)
	if value, ok := clc.mutation.DeletedAt(); ok {
		_spec.SetField(classlecturer.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if nodes := clc.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   classlecturer.ClassTable,
			Columns: []string{classlecturer.ClassColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: class.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.class_class_lecturers = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := clc.mutation.LecturerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   classlecturer.LecturerTable,
			Columns: []string{classlecturer.LecturerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: lecturer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.lecturer_class_lecturers = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ClassLecturerCreateBulk is the builder for creating many ClassLecturer entities in bulk.
type ClassLecturerCreateBulk struct {
	config
	builders []*ClassLecturerCreate
}

// Save creates the ClassLecturer entities in the database.
func (clcb *ClassLecturerCreateBulk) Save(ctx context.Context) ([]*ClassLecturer, error) {
	specs := make([]*sqlgraph.CreateSpec, len(clcb.builders))
	nodes := make([]*ClassLecturer, len(clcb.builders))
	mutators := make([]Mutator, len(clcb.builders))
	for i := range clcb.builders {
		func(i int, root context.Context) {
			builder := clcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ClassLecturerMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, clcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, clcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, clcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (clcb *ClassLecturerCreateBulk) SaveX(ctx context.Context) []*ClassLecturer {
	v, err := clcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (clcb *ClassLecturerCreateBulk) Exec(ctx context.Context) error {
	_, err := clcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (clcb *ClassLecturerCreateBulk) ExecX(ctx context.Context) {
	if err := clcb.Exec(ctx); err != nil {
		panic(err)
	}
}
