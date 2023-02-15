// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/morning-night-guild/platform/pkg/ent/article"
	"github.com/morning-night-guild/platform/pkg/ent/articletag"
)

// ArticleTagCreate is the builder for creating a ArticleTag entity.
type ArticleTagCreate struct {
	config
	mutation *ArticleTagMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetTag sets the "tag" field.
func (atc *ArticleTagCreate) SetTag(s string) *ArticleTagCreate {
	atc.mutation.SetTag(s)
	return atc
}

// SetArticleID sets the "article_id" field.
func (atc *ArticleTagCreate) SetArticleID(u uuid.UUID) *ArticleTagCreate {
	atc.mutation.SetArticleID(u)
	return atc
}

// SetID sets the "id" field.
func (atc *ArticleTagCreate) SetID(u uuid.UUID) *ArticleTagCreate {
	atc.mutation.SetID(u)
	return atc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (atc *ArticleTagCreate) SetNillableID(u *uuid.UUID) *ArticleTagCreate {
	if u != nil {
		atc.SetID(*u)
	}
	return atc
}

// SetArticle sets the "article" edge to the Article entity.
func (atc *ArticleTagCreate) SetArticle(a *Article) *ArticleTagCreate {
	return atc.SetArticleID(a.ID)
}

// Mutation returns the ArticleTagMutation object of the builder.
func (atc *ArticleTagCreate) Mutation() *ArticleTagMutation {
	return atc.mutation
}

// Save creates the ArticleTag in the database.
func (atc *ArticleTagCreate) Save(ctx context.Context) (*ArticleTag, error) {
	atc.defaults()
	return withHooks[*ArticleTag, ArticleTagMutation](ctx, atc.sqlSave, atc.mutation, atc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (atc *ArticleTagCreate) SaveX(ctx context.Context) *ArticleTag {
	v, err := atc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atc *ArticleTagCreate) Exec(ctx context.Context) error {
	_, err := atc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atc *ArticleTagCreate) ExecX(ctx context.Context) {
	if err := atc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atc *ArticleTagCreate) defaults() {
	if _, ok := atc.mutation.ID(); !ok {
		v := articletag.DefaultID()
		atc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atc *ArticleTagCreate) check() error {
	if _, ok := atc.mutation.Tag(); !ok {
		return &ValidationError{Name: "tag", err: errors.New(`ent: missing required field "ArticleTag.tag"`)}
	}
	if _, ok := atc.mutation.ArticleID(); !ok {
		return &ValidationError{Name: "article_id", err: errors.New(`ent: missing required field "ArticleTag.article_id"`)}
	}
	if _, ok := atc.mutation.ArticleID(); !ok {
		return &ValidationError{Name: "article", err: errors.New(`ent: missing required edge "ArticleTag.article"`)}
	}
	return nil
}

func (atc *ArticleTagCreate) sqlSave(ctx context.Context) (*ArticleTag, error) {
	if err := atc.check(); err != nil {
		return nil, err
	}
	_node, _spec := atc.createSpec()
	if err := sqlgraph.CreateNode(ctx, atc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	atc.mutation.id = &_node.ID
	atc.mutation.done = true
	return _node, nil
}

func (atc *ArticleTagCreate) createSpec() (*ArticleTag, *sqlgraph.CreateSpec) {
	var (
		_node = &ArticleTag{config: atc.config}
		_spec = sqlgraph.NewCreateSpec(articletag.Table, sqlgraph.NewFieldSpec(articletag.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = atc.conflict
	if id, ok := atc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := atc.mutation.Tag(); ok {
		_spec.SetField(articletag.FieldTag, field.TypeString, value)
		_node.Tag = value
	}
	if nodes := atc.mutation.ArticleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   articletag.ArticleTable,
			Columns: []string{articletag.ArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: article.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ArticleID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ArticleTag.Create().
//		SetTag(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ArticleTagUpsert) {
//			SetTag(v+v).
//		}).
//		Exec(ctx)
func (atc *ArticleTagCreate) OnConflict(opts ...sql.ConflictOption) *ArticleTagUpsertOne {
	atc.conflict = opts
	return &ArticleTagUpsertOne{
		create: atc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ArticleTag.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (atc *ArticleTagCreate) OnConflictColumns(columns ...string) *ArticleTagUpsertOne {
	atc.conflict = append(atc.conflict, sql.ConflictColumns(columns...))
	return &ArticleTagUpsertOne{
		create: atc,
	}
}

type (
	// ArticleTagUpsertOne is the builder for "upsert"-ing
	//  one ArticleTag node.
	ArticleTagUpsertOne struct {
		create *ArticleTagCreate
	}

	// ArticleTagUpsert is the "OnConflict" setter.
	ArticleTagUpsert struct {
		*sql.UpdateSet
	}
)

// SetTag sets the "tag" field.
func (u *ArticleTagUpsert) SetTag(v string) *ArticleTagUpsert {
	u.Set(articletag.FieldTag, v)
	return u
}

// UpdateTag sets the "tag" field to the value that was provided on create.
func (u *ArticleTagUpsert) UpdateTag() *ArticleTagUpsert {
	u.SetExcluded(articletag.FieldTag)
	return u
}

// SetArticleID sets the "article_id" field.
func (u *ArticleTagUpsert) SetArticleID(v uuid.UUID) *ArticleTagUpsert {
	u.Set(articletag.FieldArticleID, v)
	return u
}

// UpdateArticleID sets the "article_id" field to the value that was provided on create.
func (u *ArticleTagUpsert) UpdateArticleID() *ArticleTagUpsert {
	u.SetExcluded(articletag.FieldArticleID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ArticleTag.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(articletag.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ArticleTagUpsertOne) UpdateNewValues() *ArticleTagUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(articletag.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ArticleTag.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ArticleTagUpsertOne) Ignore() *ArticleTagUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ArticleTagUpsertOne) DoNothing() *ArticleTagUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ArticleTagCreate.OnConflict
// documentation for more info.
func (u *ArticleTagUpsertOne) Update(set func(*ArticleTagUpsert)) *ArticleTagUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ArticleTagUpsert{UpdateSet: update})
	}))
	return u
}

// SetTag sets the "tag" field.
func (u *ArticleTagUpsertOne) SetTag(v string) *ArticleTagUpsertOne {
	return u.Update(func(s *ArticleTagUpsert) {
		s.SetTag(v)
	})
}

// UpdateTag sets the "tag" field to the value that was provided on create.
func (u *ArticleTagUpsertOne) UpdateTag() *ArticleTagUpsertOne {
	return u.Update(func(s *ArticleTagUpsert) {
		s.UpdateTag()
	})
}

// SetArticleID sets the "article_id" field.
func (u *ArticleTagUpsertOne) SetArticleID(v uuid.UUID) *ArticleTagUpsertOne {
	return u.Update(func(s *ArticleTagUpsert) {
		s.SetArticleID(v)
	})
}

// UpdateArticleID sets the "article_id" field to the value that was provided on create.
func (u *ArticleTagUpsertOne) UpdateArticleID() *ArticleTagUpsertOne {
	return u.Update(func(s *ArticleTagUpsert) {
		s.UpdateArticleID()
	})
}

// Exec executes the query.
func (u *ArticleTagUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ArticleTagCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ArticleTagUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ArticleTagUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ArticleTagUpsertOne.ID is not supported by MySQL driver. Use ArticleTagUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ArticleTagUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ArticleTagCreateBulk is the builder for creating many ArticleTag entities in bulk.
type ArticleTagCreateBulk struct {
	config
	builders []*ArticleTagCreate
	conflict []sql.ConflictOption
}

// Save creates the ArticleTag entities in the database.
func (atcb *ArticleTagCreateBulk) Save(ctx context.Context) ([]*ArticleTag, error) {
	specs := make([]*sqlgraph.CreateSpec, len(atcb.builders))
	nodes := make([]*ArticleTag, len(atcb.builders))
	mutators := make([]Mutator, len(atcb.builders))
	for i := range atcb.builders {
		func(i int, root context.Context) {
			builder := atcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ArticleTagMutation)
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
					_, err = mutators[i+1].Mutate(root, atcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = atcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, atcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, atcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (atcb *ArticleTagCreateBulk) SaveX(ctx context.Context) []*ArticleTag {
	v, err := atcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atcb *ArticleTagCreateBulk) Exec(ctx context.Context) error {
	_, err := atcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atcb *ArticleTagCreateBulk) ExecX(ctx context.Context) {
	if err := atcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ArticleTag.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ArticleTagUpsert) {
//			SetTag(v+v).
//		}).
//		Exec(ctx)
func (atcb *ArticleTagCreateBulk) OnConflict(opts ...sql.ConflictOption) *ArticleTagUpsertBulk {
	atcb.conflict = opts
	return &ArticleTagUpsertBulk{
		create: atcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ArticleTag.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (atcb *ArticleTagCreateBulk) OnConflictColumns(columns ...string) *ArticleTagUpsertBulk {
	atcb.conflict = append(atcb.conflict, sql.ConflictColumns(columns...))
	return &ArticleTagUpsertBulk{
		create: atcb,
	}
}

// ArticleTagUpsertBulk is the builder for "upsert"-ing
// a bulk of ArticleTag nodes.
type ArticleTagUpsertBulk struct {
	create *ArticleTagCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ArticleTag.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(articletag.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ArticleTagUpsertBulk) UpdateNewValues() *ArticleTagUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(articletag.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ArticleTag.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ArticleTagUpsertBulk) Ignore() *ArticleTagUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ArticleTagUpsertBulk) DoNothing() *ArticleTagUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ArticleTagCreateBulk.OnConflict
// documentation for more info.
func (u *ArticleTagUpsertBulk) Update(set func(*ArticleTagUpsert)) *ArticleTagUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ArticleTagUpsert{UpdateSet: update})
	}))
	return u
}

// SetTag sets the "tag" field.
func (u *ArticleTagUpsertBulk) SetTag(v string) *ArticleTagUpsertBulk {
	return u.Update(func(s *ArticleTagUpsert) {
		s.SetTag(v)
	})
}

// UpdateTag sets the "tag" field to the value that was provided on create.
func (u *ArticleTagUpsertBulk) UpdateTag() *ArticleTagUpsertBulk {
	return u.Update(func(s *ArticleTagUpsert) {
		s.UpdateTag()
	})
}

// SetArticleID sets the "article_id" field.
func (u *ArticleTagUpsertBulk) SetArticleID(v uuid.UUID) *ArticleTagUpsertBulk {
	return u.Update(func(s *ArticleTagUpsert) {
		s.SetArticleID(v)
	})
}

// UpdateArticleID sets the "article_id" field to the value that was provided on create.
func (u *ArticleTagUpsertBulk) UpdateArticleID() *ArticleTagUpsertBulk {
	return u.Update(func(s *ArticleTagUpsert) {
		s.UpdateArticleID()
	})
}

// Exec executes the query.
func (u *ArticleTagUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ArticleTagCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ArticleTagCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ArticleTagUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
