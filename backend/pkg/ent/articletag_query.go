// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/morning-night-guild/platform/pkg/ent/article"
	"github.com/morning-night-guild/platform/pkg/ent/articletag"
	"github.com/morning-night-guild/platform/pkg/ent/predicate"
)

// ArticleTagQuery is the builder for querying ArticleTag entities.
type ArticleTagQuery struct {
	config
	limit       *int
	offset      *int
	unique      *bool
	order       []OrderFunc
	fields      []string
	inters      []Interceptor
	predicates  []predicate.ArticleTag
	withArticle *ArticleQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ArticleTagQuery builder.
func (atq *ArticleTagQuery) Where(ps ...predicate.ArticleTag) *ArticleTagQuery {
	atq.predicates = append(atq.predicates, ps...)
	return atq
}

// Limit the number of records to be returned by this query.
func (atq *ArticleTagQuery) Limit(limit int) *ArticleTagQuery {
	atq.limit = &limit
	return atq
}

// Offset to start from.
func (atq *ArticleTagQuery) Offset(offset int) *ArticleTagQuery {
	atq.offset = &offset
	return atq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (atq *ArticleTagQuery) Unique(unique bool) *ArticleTagQuery {
	atq.unique = &unique
	return atq
}

// Order specifies how the records should be ordered.
func (atq *ArticleTagQuery) Order(o ...OrderFunc) *ArticleTagQuery {
	atq.order = append(atq.order, o...)
	return atq
}

// QueryArticle chains the current query on the "article" edge.
func (atq *ArticleTagQuery) QueryArticle() *ArticleQuery {
	query := (&ArticleClient{config: atq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := atq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := atq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(articletag.Table, articletag.FieldID, selector),
			sqlgraph.To(article.Table, article.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, articletag.ArticleTable, articletag.ArticleColumn),
		)
		fromU = sqlgraph.SetNeighbors(atq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ArticleTag entity from the query.
// Returns a *NotFoundError when no ArticleTag was found.
func (atq *ArticleTagQuery) First(ctx context.Context) (*ArticleTag, error) {
	nodes, err := atq.Limit(1).All(newQueryContext(ctx, TypeArticleTag, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{articletag.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (atq *ArticleTagQuery) FirstX(ctx context.Context) *ArticleTag {
	node, err := atq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ArticleTag ID from the query.
// Returns a *NotFoundError when no ArticleTag ID was found.
func (atq *ArticleTagQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = atq.Limit(1).IDs(newQueryContext(ctx, TypeArticleTag, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{articletag.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (atq *ArticleTagQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := atq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ArticleTag entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ArticleTag entity is found.
// Returns a *NotFoundError when no ArticleTag entities are found.
func (atq *ArticleTagQuery) Only(ctx context.Context) (*ArticleTag, error) {
	nodes, err := atq.Limit(2).All(newQueryContext(ctx, TypeArticleTag, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{articletag.Label}
	default:
		return nil, &NotSingularError{articletag.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (atq *ArticleTagQuery) OnlyX(ctx context.Context) *ArticleTag {
	node, err := atq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ArticleTag ID in the query.
// Returns a *NotSingularError when more than one ArticleTag ID is found.
// Returns a *NotFoundError when no entities are found.
func (atq *ArticleTagQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = atq.Limit(2).IDs(newQueryContext(ctx, TypeArticleTag, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{articletag.Label}
	default:
		err = &NotSingularError{articletag.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (atq *ArticleTagQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := atq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ArticleTags.
func (atq *ArticleTagQuery) All(ctx context.Context) ([]*ArticleTag, error) {
	ctx = newQueryContext(ctx, TypeArticleTag, "All")
	if err := atq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ArticleTag, *ArticleTagQuery]()
	return withInterceptors[[]*ArticleTag](ctx, atq, qr, atq.inters)
}

// AllX is like All, but panics if an error occurs.
func (atq *ArticleTagQuery) AllX(ctx context.Context) []*ArticleTag {
	nodes, err := atq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ArticleTag IDs.
func (atq *ArticleTagQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = newQueryContext(ctx, TypeArticleTag, "IDs")
	if err := atq.Select(articletag.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (atq *ArticleTagQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := atq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (atq *ArticleTagQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeArticleTag, "Count")
	if err := atq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, atq, querierCount[*ArticleTagQuery](), atq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (atq *ArticleTagQuery) CountX(ctx context.Context) int {
	count, err := atq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (atq *ArticleTagQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeArticleTag, "Exist")
	switch _, err := atq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (atq *ArticleTagQuery) ExistX(ctx context.Context) bool {
	exist, err := atq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ArticleTagQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (atq *ArticleTagQuery) Clone() *ArticleTagQuery {
	if atq == nil {
		return nil
	}
	return &ArticleTagQuery{
		config:      atq.config,
		limit:       atq.limit,
		offset:      atq.offset,
		order:       append([]OrderFunc{}, atq.order...),
		inters:      append([]Interceptor{}, atq.inters...),
		predicates:  append([]predicate.ArticleTag{}, atq.predicates...),
		withArticle: atq.withArticle.Clone(),
		// clone intermediate query.
		sql:    atq.sql.Clone(),
		path:   atq.path,
		unique: atq.unique,
	}
}

// WithArticle tells the query-builder to eager-load the nodes that are connected to
// the "article" edge. The optional arguments are used to configure the query builder of the edge.
func (atq *ArticleTagQuery) WithArticle(opts ...func(*ArticleQuery)) *ArticleTagQuery {
	query := (&ArticleClient{config: atq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	atq.withArticle = query
	return atq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Tag string `json:"tag,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ArticleTag.Query().
//		GroupBy(articletag.FieldTag).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (atq *ArticleTagQuery) GroupBy(field string, fields ...string) *ArticleTagGroupBy {
	atq.fields = append([]string{field}, fields...)
	grbuild := &ArticleTagGroupBy{build: atq}
	grbuild.flds = &atq.fields
	grbuild.label = articletag.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Tag string `json:"tag,omitempty"`
//	}
//
//	client.ArticleTag.Query().
//		Select(articletag.FieldTag).
//		Scan(ctx, &v)
func (atq *ArticleTagQuery) Select(fields ...string) *ArticleTagSelect {
	atq.fields = append(atq.fields, fields...)
	sbuild := &ArticleTagSelect{ArticleTagQuery: atq}
	sbuild.label = articletag.Label
	sbuild.flds, sbuild.scan = &atq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ArticleTagSelect configured with the given aggregations.
func (atq *ArticleTagQuery) Aggregate(fns ...AggregateFunc) *ArticleTagSelect {
	return atq.Select().Aggregate(fns...)
}

func (atq *ArticleTagQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range atq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, atq); err != nil {
				return err
			}
		}
	}
	for _, f := range atq.fields {
		if !articletag.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if atq.path != nil {
		prev, err := atq.path(ctx)
		if err != nil {
			return err
		}
		atq.sql = prev
	}
	return nil
}

func (atq *ArticleTagQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ArticleTag, error) {
	var (
		nodes       = []*ArticleTag{}
		_spec       = atq.querySpec()
		loadedTypes = [1]bool{
			atq.withArticle != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ArticleTag).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ArticleTag{config: atq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, atq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := atq.withArticle; query != nil {
		if err := atq.loadArticle(ctx, query, nodes, nil,
			func(n *ArticleTag, e *Article) { n.Edges.Article = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (atq *ArticleTagQuery) loadArticle(ctx context.Context, query *ArticleQuery, nodes []*ArticleTag, init func(*ArticleTag), assign func(*ArticleTag, *Article)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ArticleTag)
	for i := range nodes {
		fk := nodes[i].ArticleID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(article.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "article_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (atq *ArticleTagQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := atq.querySpec()
	_spec.Node.Columns = atq.fields
	if len(atq.fields) > 0 {
		_spec.Unique = atq.unique != nil && *atq.unique
	}
	return sqlgraph.CountNodes(ctx, atq.driver, _spec)
}

func (atq *ArticleTagQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   articletag.Table,
			Columns: articletag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: articletag.FieldID,
			},
		},
		From:   atq.sql,
		Unique: true,
	}
	if unique := atq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := atq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, articletag.FieldID)
		for i := range fields {
			if fields[i] != articletag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := atq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := atq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := atq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := atq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (atq *ArticleTagQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(atq.driver.Dialect())
	t1 := builder.Table(articletag.Table)
	columns := atq.fields
	if len(columns) == 0 {
		columns = articletag.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if atq.sql != nil {
		selector = atq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if atq.unique != nil && *atq.unique {
		selector.Distinct()
	}
	for _, p := range atq.predicates {
		p(selector)
	}
	for _, p := range atq.order {
		p(selector)
	}
	if offset := atq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := atq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ArticleTagGroupBy is the group-by builder for ArticleTag entities.
type ArticleTagGroupBy struct {
	selector
	build *ArticleTagQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (atgb *ArticleTagGroupBy) Aggregate(fns ...AggregateFunc) *ArticleTagGroupBy {
	atgb.fns = append(atgb.fns, fns...)
	return atgb
}

// Scan applies the selector query and scans the result into the given value.
func (atgb *ArticleTagGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeArticleTag, "GroupBy")
	if err := atgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ArticleTagQuery, *ArticleTagGroupBy](ctx, atgb.build, atgb, atgb.build.inters, v)
}

func (atgb *ArticleTagGroupBy) sqlScan(ctx context.Context, root *ArticleTagQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(atgb.fns))
	for _, fn := range atgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*atgb.flds)+len(atgb.fns))
		for _, f := range *atgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*atgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := atgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ArticleTagSelect is the builder for selecting fields of ArticleTag entities.
type ArticleTagSelect struct {
	*ArticleTagQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ats *ArticleTagSelect) Aggregate(fns ...AggregateFunc) *ArticleTagSelect {
	ats.fns = append(ats.fns, fns...)
	return ats
}

// Scan applies the selector query and scans the result into the given value.
func (ats *ArticleTagSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeArticleTag, "Select")
	if err := ats.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ArticleTagQuery, *ArticleTagSelect](ctx, ats.ArticleTagQuery, ats, ats.inters, v)
}

func (ats *ArticleTagSelect) sqlScan(ctx context.Context, root *ArticleTagQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ats.fns))
	for _, fn := range ats.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ats.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ats.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
