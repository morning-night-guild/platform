// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/morning-night-guild/platform/pkg/ent/article"
	"github.com/morning-night-guild/platform/pkg/ent/articletag"
	"github.com/morning-night-guild/platform/pkg/ent/predicate"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeArticle    = "Article"
	TypeArticleTag = "ArticleTag"
)

// ArticleMutation represents an operation that mutates the Article nodes in the graph.
type ArticleMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	title         *string
	url           *string
	description   *string
	thumbnail     *string
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	tags          map[uuid.UUID]struct{}
	removedtags   map[uuid.UUID]struct{}
	clearedtags   bool
	done          bool
	oldValue      func(context.Context) (*Article, error)
	predicates    []predicate.Article
}

var _ ent.Mutation = (*ArticleMutation)(nil)

// articleOption allows management of the mutation configuration using functional options.
type articleOption func(*ArticleMutation)

// newArticleMutation creates new mutation for the Article entity.
func newArticleMutation(c config, op Op, opts ...articleOption) *ArticleMutation {
	m := &ArticleMutation{
		config:        c,
		op:            op,
		typ:           TypeArticle,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withArticleID sets the ID field of the mutation.
func withArticleID(id uuid.UUID) articleOption {
	return func(m *ArticleMutation) {
		var (
			err   error
			once  sync.Once
			value *Article
		)
		m.oldValue = func(ctx context.Context) (*Article, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Article.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withArticle sets the old Article of the mutation.
func withArticle(node *Article) articleOption {
	return func(m *ArticleMutation) {
		m.oldValue = func(context.Context) (*Article, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ArticleMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ArticleMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Article entities.
func (m *ArticleMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ArticleMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ArticleMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Article.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetTitle sets the "title" field.
func (m *ArticleMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the value of the "title" field in the mutation.
func (m *ArticleMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old "title" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTitle is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

// ResetTitle resets all changes to the "title" field.
func (m *ArticleMutation) ResetTitle() {
	m.title = nil
}

// SetURL sets the "url" field.
func (m *ArticleMutation) SetURL(s string) {
	m.url = &s
}

// URL returns the value of the "url" field in the mutation.
func (m *ArticleMutation) URL() (r string, exists bool) {
	v := m.url
	if v == nil {
		return
	}
	return *v, true
}

// OldURL returns the old "url" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldURL(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldURL is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldURL: %w", err)
	}
	return oldValue.URL, nil
}

// ResetURL resets all changes to the "url" field.
func (m *ArticleMutation) ResetURL() {
	m.url = nil
}

// SetDescription sets the "description" field.
func (m *ArticleMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the value of the "description" field in the mutation.
func (m *ArticleMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// OldDescription returns the old "description" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldDescription(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDescription is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDescription requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescription: %w", err)
	}
	return oldValue.Description, nil
}

// ResetDescription resets all changes to the "description" field.
func (m *ArticleMutation) ResetDescription() {
	m.description = nil
}

// SetThumbnail sets the "thumbnail" field.
func (m *ArticleMutation) SetThumbnail(s string) {
	m.thumbnail = &s
}

// Thumbnail returns the value of the "thumbnail" field in the mutation.
func (m *ArticleMutation) Thumbnail() (r string, exists bool) {
	v := m.thumbnail
	if v == nil {
		return
	}
	return *v, true
}

// OldThumbnail returns the old "thumbnail" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldThumbnail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldThumbnail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldThumbnail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldThumbnail: %w", err)
	}
	return oldValue.Thumbnail, nil
}

// ResetThumbnail resets all changes to the "thumbnail" field.
func (m *ArticleMutation) ResetThumbnail() {
	m.thumbnail = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *ArticleMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *ArticleMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *ArticleMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *ArticleMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *ArticleMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *ArticleMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// AddTagIDs adds the "tags" edge to the ArticleTag entity by ids.
func (m *ArticleMutation) AddTagIDs(ids ...uuid.UUID) {
	if m.tags == nil {
		m.tags = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.tags[ids[i]] = struct{}{}
	}
}

// ClearTags clears the "tags" edge to the ArticleTag entity.
func (m *ArticleMutation) ClearTags() {
	m.clearedtags = true
}

// TagsCleared reports if the "tags" edge to the ArticleTag entity was cleared.
func (m *ArticleMutation) TagsCleared() bool {
	return m.clearedtags
}

// RemoveTagIDs removes the "tags" edge to the ArticleTag entity by IDs.
func (m *ArticleMutation) RemoveTagIDs(ids ...uuid.UUID) {
	if m.removedtags == nil {
		m.removedtags = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		delete(m.tags, ids[i])
		m.removedtags[ids[i]] = struct{}{}
	}
}

// RemovedTags returns the removed IDs of the "tags" edge to the ArticleTag entity.
func (m *ArticleMutation) RemovedTagsIDs() (ids []uuid.UUID) {
	for id := range m.removedtags {
		ids = append(ids, id)
	}
	return
}

// TagsIDs returns the "tags" edge IDs in the mutation.
func (m *ArticleMutation) TagsIDs() (ids []uuid.UUID) {
	for id := range m.tags {
		ids = append(ids, id)
	}
	return
}

// ResetTags resets all changes to the "tags" edge.
func (m *ArticleMutation) ResetTags() {
	m.tags = nil
	m.clearedtags = false
	m.removedtags = nil
}

// Where appends a list predicates to the ArticleMutation builder.
func (m *ArticleMutation) Where(ps ...predicate.Article) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the ArticleMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *ArticleMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Article, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *ArticleMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *ArticleMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Article).
func (m *ArticleMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ArticleMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.title != nil {
		fields = append(fields, article.FieldTitle)
	}
	if m.url != nil {
		fields = append(fields, article.FieldURL)
	}
	if m.description != nil {
		fields = append(fields, article.FieldDescription)
	}
	if m.thumbnail != nil {
		fields = append(fields, article.FieldThumbnail)
	}
	if m.created_at != nil {
		fields = append(fields, article.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, article.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ArticleMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case article.FieldTitle:
		return m.Title()
	case article.FieldURL:
		return m.URL()
	case article.FieldDescription:
		return m.Description()
	case article.FieldThumbnail:
		return m.Thumbnail()
	case article.FieldCreatedAt:
		return m.CreatedAt()
	case article.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ArticleMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case article.FieldTitle:
		return m.OldTitle(ctx)
	case article.FieldURL:
		return m.OldURL(ctx)
	case article.FieldDescription:
		return m.OldDescription(ctx)
	case article.FieldThumbnail:
		return m.OldThumbnail(ctx)
	case article.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case article.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Article field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ArticleMutation) SetField(name string, value ent.Value) error {
	switch name {
	case article.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case article.FieldURL:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetURL(v)
		return nil
	case article.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	case article.FieldThumbnail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetThumbnail(v)
		return nil
	case article.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case article.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Article field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ArticleMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ArticleMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ArticleMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Article numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ArticleMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ArticleMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ArticleMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Article nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ArticleMutation) ResetField(name string) error {
	switch name {
	case article.FieldTitle:
		m.ResetTitle()
		return nil
	case article.FieldURL:
		m.ResetURL()
		return nil
	case article.FieldDescription:
		m.ResetDescription()
		return nil
	case article.FieldThumbnail:
		m.ResetThumbnail()
		return nil
	case article.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case article.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Article field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ArticleMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.tags != nil {
		edges = append(edges, article.EdgeTags)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ArticleMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case article.EdgeTags:
		ids := make([]ent.Value, 0, len(m.tags))
		for id := range m.tags {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ArticleMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedtags != nil {
		edges = append(edges, article.EdgeTags)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ArticleMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case article.EdgeTags:
		ids := make([]ent.Value, 0, len(m.removedtags))
		for id := range m.removedtags {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ArticleMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedtags {
		edges = append(edges, article.EdgeTags)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ArticleMutation) EdgeCleared(name string) bool {
	switch name {
	case article.EdgeTags:
		return m.clearedtags
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ArticleMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Article unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ArticleMutation) ResetEdge(name string) error {
	switch name {
	case article.EdgeTags:
		m.ResetTags()
		return nil
	}
	return fmt.Errorf("unknown Article edge %s", name)
}

// ArticleTagMutation represents an operation that mutates the ArticleTag nodes in the graph.
type ArticleTagMutation struct {
	config
	op             Op
	typ            string
	id             *uuid.UUID
	tag            *string
	clearedFields  map[string]struct{}
	article        *uuid.UUID
	clearedarticle bool
	done           bool
	oldValue       func(context.Context) (*ArticleTag, error)
	predicates     []predicate.ArticleTag
}

var _ ent.Mutation = (*ArticleTagMutation)(nil)

// articletagOption allows management of the mutation configuration using functional options.
type articletagOption func(*ArticleTagMutation)

// newArticleTagMutation creates new mutation for the ArticleTag entity.
func newArticleTagMutation(c config, op Op, opts ...articletagOption) *ArticleTagMutation {
	m := &ArticleTagMutation{
		config:        c,
		op:            op,
		typ:           TypeArticleTag,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withArticleTagID sets the ID field of the mutation.
func withArticleTagID(id uuid.UUID) articletagOption {
	return func(m *ArticleTagMutation) {
		var (
			err   error
			once  sync.Once
			value *ArticleTag
		)
		m.oldValue = func(ctx context.Context) (*ArticleTag, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().ArticleTag.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withArticleTag sets the old ArticleTag of the mutation.
func withArticleTag(node *ArticleTag) articletagOption {
	return func(m *ArticleTagMutation) {
		m.oldValue = func(context.Context) (*ArticleTag, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ArticleTagMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ArticleTagMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of ArticleTag entities.
func (m *ArticleTagMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ArticleTagMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ArticleTagMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().ArticleTag.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetTag sets the "tag" field.
func (m *ArticleTagMutation) SetTag(s string) {
	m.tag = &s
}

// Tag returns the value of the "tag" field in the mutation.
func (m *ArticleTagMutation) Tag() (r string, exists bool) {
	v := m.tag
	if v == nil {
		return
	}
	return *v, true
}

// OldTag returns the old "tag" field's value of the ArticleTag entity.
// If the ArticleTag object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleTagMutation) OldTag(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTag is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTag requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTag: %w", err)
	}
	return oldValue.Tag, nil
}

// ResetTag resets all changes to the "tag" field.
func (m *ArticleTagMutation) ResetTag() {
	m.tag = nil
}

// SetArticleID sets the "article_id" field.
func (m *ArticleTagMutation) SetArticleID(u uuid.UUID) {
	m.article = &u
}

// ArticleID returns the value of the "article_id" field in the mutation.
func (m *ArticleTagMutation) ArticleID() (r uuid.UUID, exists bool) {
	v := m.article
	if v == nil {
		return
	}
	return *v, true
}

// OldArticleID returns the old "article_id" field's value of the ArticleTag entity.
// If the ArticleTag object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleTagMutation) OldArticleID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldArticleID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldArticleID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldArticleID: %w", err)
	}
	return oldValue.ArticleID, nil
}

// ResetArticleID resets all changes to the "article_id" field.
func (m *ArticleTagMutation) ResetArticleID() {
	m.article = nil
}

// ClearArticle clears the "article" edge to the Article entity.
func (m *ArticleTagMutation) ClearArticle() {
	m.clearedarticle = true
}

// ArticleCleared reports if the "article" edge to the Article entity was cleared.
func (m *ArticleTagMutation) ArticleCleared() bool {
	return m.clearedarticle
}

// ArticleIDs returns the "article" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// ArticleID instead. It exists only for internal usage by the builders.
func (m *ArticleTagMutation) ArticleIDs() (ids []uuid.UUID) {
	if id := m.article; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetArticle resets all changes to the "article" edge.
func (m *ArticleTagMutation) ResetArticle() {
	m.article = nil
	m.clearedarticle = false
}

// Where appends a list predicates to the ArticleTagMutation builder.
func (m *ArticleTagMutation) Where(ps ...predicate.ArticleTag) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the ArticleTagMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *ArticleTagMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.ArticleTag, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *ArticleTagMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *ArticleTagMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (ArticleTag).
func (m *ArticleTagMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ArticleTagMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.tag != nil {
		fields = append(fields, articletag.FieldTag)
	}
	if m.article != nil {
		fields = append(fields, articletag.FieldArticleID)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ArticleTagMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case articletag.FieldTag:
		return m.Tag()
	case articletag.FieldArticleID:
		return m.ArticleID()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ArticleTagMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case articletag.FieldTag:
		return m.OldTag(ctx)
	case articletag.FieldArticleID:
		return m.OldArticleID(ctx)
	}
	return nil, fmt.Errorf("unknown ArticleTag field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ArticleTagMutation) SetField(name string, value ent.Value) error {
	switch name {
	case articletag.FieldTag:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTag(v)
		return nil
	case articletag.FieldArticleID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetArticleID(v)
		return nil
	}
	return fmt.Errorf("unknown ArticleTag field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ArticleTagMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ArticleTagMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ArticleTagMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown ArticleTag numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ArticleTagMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ArticleTagMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ArticleTagMutation) ClearField(name string) error {
	return fmt.Errorf("unknown ArticleTag nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ArticleTagMutation) ResetField(name string) error {
	switch name {
	case articletag.FieldTag:
		m.ResetTag()
		return nil
	case articletag.FieldArticleID:
		m.ResetArticleID()
		return nil
	}
	return fmt.Errorf("unknown ArticleTag field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ArticleTagMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.article != nil {
		edges = append(edges, articletag.EdgeArticle)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ArticleTagMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case articletag.EdgeArticle:
		if id := m.article; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ArticleTagMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ArticleTagMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ArticleTagMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedarticle {
		edges = append(edges, articletag.EdgeArticle)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ArticleTagMutation) EdgeCleared(name string) bool {
	switch name {
	case articletag.EdgeArticle:
		return m.clearedarticle
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ArticleTagMutation) ClearEdge(name string) error {
	switch name {
	case articletag.EdgeArticle:
		m.ClearArticle()
		return nil
	}
	return fmt.Errorf("unknown ArticleTag unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ArticleTagMutation) ResetEdge(name string) error {
	switch name {
	case articletag.EdgeArticle:
		m.ResetArticle()
		return nil
	}
	return fmt.Errorf("unknown ArticleTag edge %s", name)
}
