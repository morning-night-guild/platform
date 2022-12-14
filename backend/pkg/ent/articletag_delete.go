// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/morning-night-guild/platform/pkg/ent/articletag"
	"github.com/morning-night-guild/platform/pkg/ent/predicate"
)

// ArticleTagDelete is the builder for deleting a ArticleTag entity.
type ArticleTagDelete struct {
	config
	hooks    []Hook
	mutation *ArticleTagMutation
}

// Where appends a list predicates to the ArticleTagDelete builder.
func (atd *ArticleTagDelete) Where(ps ...predicate.ArticleTag) *ArticleTagDelete {
	atd.mutation.Where(ps...)
	return atd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (atd *ArticleTagDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(atd.hooks) == 0 {
		affected, err = atd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ArticleTagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			atd.mutation = mutation
			affected, err = atd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(atd.hooks) - 1; i >= 0; i-- {
			if atd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = atd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, atd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (atd *ArticleTagDelete) ExecX(ctx context.Context) int {
	n, err := atd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (atd *ArticleTagDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: articletag.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: articletag.FieldID,
			},
		},
	}
	if ps := atd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, atd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// ArticleTagDeleteOne is the builder for deleting a single ArticleTag entity.
type ArticleTagDeleteOne struct {
	atd *ArticleTagDelete
}

// Exec executes the deletion query.
func (atdo *ArticleTagDeleteOne) Exec(ctx context.Context) error {
	n, err := atdo.atd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{articletag.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (atdo *ArticleTagDeleteOne) ExecX(ctx context.Context) {
	atdo.atd.ExecX(ctx)
}
