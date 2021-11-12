// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/NpoolPlatform/kyc-management/pkg/db/ent/migrate"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/kyc-management/pkg/db/ent/kyc"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Kyc is the client for interacting with the Kyc builders.
	Kyc *KycClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Kyc = NewKycClient(c.config)
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
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Kyc:    NewKycClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
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
		config: cfg,
		Kyc:    NewKycClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Kyc.
//		Query().
//		Count(ctx)
//
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
	c.Kyc.Use(hooks...)
}

// KycClient is a client for the Kyc schema.
type KycClient struct {
	config
}

// NewKycClient returns a client for the Kyc from the given config.
func NewKycClient(c config) *KycClient {
	return &KycClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `kyc.Hooks(f(g(h())))`.
func (c *KycClient) Use(hooks ...Hook) {
	c.hooks.Kyc = append(c.hooks.Kyc, hooks...)
}

// Create returns a create builder for Kyc.
func (c *KycClient) Create() *KycCreate {
	mutation := newKycMutation(c.config, OpCreate)
	return &KycCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Kyc entities.
func (c *KycClient) CreateBulk(builders ...*KycCreate) *KycCreateBulk {
	return &KycCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Kyc.
func (c *KycClient) Update() *KycUpdate {
	mutation := newKycMutation(c.config, OpUpdate)
	return &KycUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *KycClient) UpdateOne(k *Kyc) *KycUpdateOne {
	mutation := newKycMutation(c.config, OpUpdateOne, withKyc(k))
	return &KycUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *KycClient) UpdateOneID(id uuid.UUID) *KycUpdateOne {
	mutation := newKycMutation(c.config, OpUpdateOne, withKycID(id))
	return &KycUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Kyc.
func (c *KycClient) Delete() *KycDelete {
	mutation := newKycMutation(c.config, OpDelete)
	return &KycDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *KycClient) DeleteOne(k *Kyc) *KycDeleteOne {
	return c.DeleteOneID(k.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *KycClient) DeleteOneID(id uuid.UUID) *KycDeleteOne {
	builder := c.Delete().Where(kyc.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &KycDeleteOne{builder}
}

// Query returns a query builder for Kyc.
func (c *KycClient) Query() *KycQuery {
	return &KycQuery{
		config: c.config,
	}
}

// Get returns a Kyc entity by its id.
func (c *KycClient) Get(ctx context.Context, id uuid.UUID) (*Kyc, error) {
	return c.Query().Where(kyc.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *KycClient) GetX(ctx context.Context, id uuid.UUID) *Kyc {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *KycClient) Hooks() []Hook {
	return c.hooks.Kyc
}
