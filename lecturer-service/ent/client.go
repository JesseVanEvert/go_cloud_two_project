// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"lecturer/ent/migrate"

	"lecturer/ent/class"
	"lecturer/ent/lecturer"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Class is the client for interacting with the Class builders.
	Class *ClassClient
	// Lecturer is the client for interacting with the Lecturer builders.
	Lecturer *LecturerClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Class = NewClassClient(c.config)
	c.Lecturer = NewLecturerClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Class:    NewClassClient(cfg),
		Lecturer: NewLecturerClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Class:    NewClassClient(cfg),
		Lecturer: NewLecturerClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Class.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Class.Use(hooks...)
	c.Lecturer.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Class.Intercept(interceptors...)
	c.Lecturer.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *ClassMutation:
		return c.Class.mutate(ctx, m)
	case *LecturerMutation:
		return c.Lecturer.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// ClassClient is a client for the Class schema.
type ClassClient struct {
	config
}

// NewClassClient returns a client for the Class from the given config.
func NewClassClient(c config) *ClassClient {
	return &ClassClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `class.Hooks(f(g(h())))`.
func (c *ClassClient) Use(hooks ...Hook) {
	c.hooks.Class = append(c.hooks.Class, hooks...)
}

// Use adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `class.Intercept(f(g(h())))`.
func (c *ClassClient) Intercept(interceptors ...Interceptor) {
	c.inters.Class = append(c.inters.Class, interceptors...)
}

// Create returns a builder for creating a Class entity.
func (c *ClassClient) Create() *ClassCreate {
	mutation := newClassMutation(c.config, OpCreate)
	return &ClassCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Class entities.
func (c *ClassClient) CreateBulk(builders ...*ClassCreate) *ClassCreateBulk {
	return &ClassCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Class.
func (c *ClassClient) Update() *ClassUpdate {
	mutation := newClassMutation(c.config, OpUpdate)
	return &ClassUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ClassClient) UpdateOne(cl *Class) *ClassUpdateOne {
	mutation := newClassMutation(c.config, OpUpdateOne, withClass(cl))
	return &ClassUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ClassClient) UpdateOneID(id int) *ClassUpdateOne {
	mutation := newClassMutation(c.config, OpUpdateOne, withClassID(id))
	return &ClassUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Class.
func (c *ClassClient) Delete() *ClassDelete {
	mutation := newClassMutation(c.config, OpDelete)
	return &ClassDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ClassClient) DeleteOne(cl *Class) *ClassDeleteOne {
	return c.DeleteOneID(cl.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ClassClient) DeleteOneID(id int) *ClassDeleteOne {
	builder := c.Delete().Where(class.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ClassDeleteOne{builder}
}

// Query returns a query builder for Class.
func (c *ClassClient) Query() *ClassQuery {
	return &ClassQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeClass},
		inters: c.Interceptors(),
	}
}

// Get returns a Class entity by its id.
func (c *ClassClient) Get(ctx context.Context, id int) (*Class, error) {
	return c.Query().Where(class.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ClassClient) GetX(ctx context.Context, id int) *Class {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryLecturers queries the lecturers edge of a Class.
func (c *ClassClient) QueryLecturers(cl *Class) *LecturerQuery {
	query := (&LecturerClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(class.Table, class.FieldID, id),
			sqlgraph.To(lecturer.Table, lecturer.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, class.LecturersTable, class.LecturersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(cl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ClassClient) Hooks() []Hook {
	return c.hooks.Class
}

// Interceptors returns the client interceptors.
func (c *ClassClient) Interceptors() []Interceptor {
	return c.inters.Class
}

func (c *ClassClient) mutate(ctx context.Context, m *ClassMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ClassCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ClassUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ClassUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ClassDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Class mutation op: %q", m.Op())
	}
}

// LecturerClient is a client for the Lecturer schema.
type LecturerClient struct {
	config
}

// NewLecturerClient returns a client for the Lecturer from the given config.
func NewLecturerClient(c config) *LecturerClient {
	return &LecturerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `lecturer.Hooks(f(g(h())))`.
func (c *LecturerClient) Use(hooks ...Hook) {
	c.hooks.Lecturer = append(c.hooks.Lecturer, hooks...)
}

// Use adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `lecturer.Intercept(f(g(h())))`.
func (c *LecturerClient) Intercept(interceptors ...Interceptor) {
	c.inters.Lecturer = append(c.inters.Lecturer, interceptors...)
}

// Create returns a builder for creating a Lecturer entity.
func (c *LecturerClient) Create() *LecturerCreate {
	mutation := newLecturerMutation(c.config, OpCreate)
	return &LecturerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Lecturer entities.
func (c *LecturerClient) CreateBulk(builders ...*LecturerCreate) *LecturerCreateBulk {
	return &LecturerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Lecturer.
func (c *LecturerClient) Update() *LecturerUpdate {
	mutation := newLecturerMutation(c.config, OpUpdate)
	return &LecturerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LecturerClient) UpdateOne(l *Lecturer) *LecturerUpdateOne {
	mutation := newLecturerMutation(c.config, OpUpdateOne, withLecturer(l))
	return &LecturerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LecturerClient) UpdateOneID(id int) *LecturerUpdateOne {
	mutation := newLecturerMutation(c.config, OpUpdateOne, withLecturerID(id))
	return &LecturerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Lecturer.
func (c *LecturerClient) Delete() *LecturerDelete {
	mutation := newLecturerMutation(c.config, OpDelete)
	return &LecturerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *LecturerClient) DeleteOne(l *Lecturer) *LecturerDeleteOne {
	return c.DeleteOneID(l.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *LecturerClient) DeleteOneID(id int) *LecturerDeleteOne {
	builder := c.Delete().Where(lecturer.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LecturerDeleteOne{builder}
}

// Query returns a query builder for Lecturer.
func (c *LecturerClient) Query() *LecturerQuery {
	return &LecturerQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeLecturer},
		inters: c.Interceptors(),
	}
}

// Get returns a Lecturer entity by its id.
func (c *LecturerClient) Get(ctx context.Context, id int) (*Lecturer, error) {
	return c.Query().Where(lecturer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LecturerClient) GetX(ctx context.Context, id int) *Lecturer {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryClasses queries the classes edge of a Lecturer.
func (c *LecturerClient) QueryClasses(l *Lecturer) *ClassQuery {
	query := (&ClassClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := l.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(lecturer.Table, lecturer.FieldID, id),
			sqlgraph.To(class.Table, class.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, lecturer.ClassesTable, lecturer.ClassesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(l.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *LecturerClient) Hooks() []Hook {
	return c.hooks.Lecturer
}

// Interceptors returns the client interceptors.
func (c *LecturerClient) Interceptors() []Interceptor {
	return c.inters.Lecturer
}

func (c *LecturerClient) mutate(ctx context.Context, m *LecturerMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&LecturerCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&LecturerUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&LecturerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&LecturerDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Lecturer mutation op: %q", m.Op())
	}
}
