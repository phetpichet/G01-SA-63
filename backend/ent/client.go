// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/Piichet-3-/app/ent/migrate"

	"github.com/Piichet-3-/app/ent/dispense"
	"github.com/Piichet-3-/app/ent/drug"
	"github.com/Piichet-3-/app/ent/form"
	"github.com/Piichet-3-/app/ent/gender"
	"github.com/Piichet-3-/app/ent/position"
	"github.com/Piichet-3-/app/ent/title"
	"github.com/Piichet-3-/app/ent/unit"
	"github.com/Piichet-3-/app/ent/user"
	"github.com/Piichet-3-/app/ent/volume"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Dispense is the client for interacting with the Dispense builders.
	Dispense *DispenseClient
	// Drug is the client for interacting with the Drug builders.
	Drug *DrugClient
	// Form is the client for interacting with the Form builders.
	Form *FormClient
	// Gender is the client for interacting with the Gender builders.
	Gender *GenderClient
	// Position is the client for interacting with the Position builders.
	Position *PositionClient
	// Title is the client for interacting with the Title builders.
	Title *TitleClient
	// Unit is the client for interacting with the Unit builders.
	Unit *UnitClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// Volume is the client for interacting with the Volume builders.
	Volume *VolumeClient
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
	c.Dispense = NewDispenseClient(c.config)
	c.Drug = NewDrugClient(c.config)
	c.Form = NewFormClient(c.config)
	c.Gender = NewGenderClient(c.config)
	c.Position = NewPositionClient(c.config)
	c.Title = NewTitleClient(c.config)
	c.Unit = NewUnitClient(c.config)
	c.User = NewUserClient(c.config)
	c.Volume = NewVolumeClient(c.config)
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
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Dispense: NewDispenseClient(cfg),
		Drug:     NewDrugClient(cfg),
		Form:     NewFormClient(cfg),
		Gender:   NewGenderClient(cfg),
		Position: NewPositionClient(cfg),
		Title:    NewTitleClient(cfg),
		Unit:     NewUnitClient(cfg),
		User:     NewUserClient(cfg),
		Volume:   NewVolumeClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:   cfg,
		Dispense: NewDispenseClient(cfg),
		Drug:     NewDrugClient(cfg),
		Form:     NewFormClient(cfg),
		Gender:   NewGenderClient(cfg),
		Position: NewPositionClient(cfg),
		Title:    NewTitleClient(cfg),
		Unit:     NewUnitClient(cfg),
		User:     NewUserClient(cfg),
		Volume:   NewVolumeClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Dispense.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
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
	c.Dispense.Use(hooks...)
	c.Drug.Use(hooks...)
	c.Form.Use(hooks...)
	c.Gender.Use(hooks...)
	c.Position.Use(hooks...)
	c.Title.Use(hooks...)
	c.Unit.Use(hooks...)
	c.User.Use(hooks...)
	c.Volume.Use(hooks...)
}

// DispenseClient is a client for the Dispense schema.
type DispenseClient struct {
	config
}

// NewDispenseClient returns a client for the Dispense from the given config.
func NewDispenseClient(c config) *DispenseClient {
	return &DispenseClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `dispense.Hooks(f(g(h())))`.
func (c *DispenseClient) Use(hooks ...Hook) {
	c.hooks.Dispense = append(c.hooks.Dispense, hooks...)
}

// Create returns a create builder for Dispense.
func (c *DispenseClient) Create() *DispenseCreate {
	mutation := newDispenseMutation(c.config, OpCreate)
	return &DispenseCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Dispense.
func (c *DispenseClient) Update() *DispenseUpdate {
	mutation := newDispenseMutation(c.config, OpUpdate)
	return &DispenseUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DispenseClient) UpdateOne(d *Dispense) *DispenseUpdateOne {
	mutation := newDispenseMutation(c.config, OpUpdateOne, withDispense(d))
	return &DispenseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DispenseClient) UpdateOneID(id int) *DispenseUpdateOne {
	mutation := newDispenseMutation(c.config, OpUpdateOne, withDispenseID(id))
	return &DispenseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Dispense.
func (c *DispenseClient) Delete() *DispenseDelete {
	mutation := newDispenseMutation(c.config, OpDelete)
	return &DispenseDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DispenseClient) DeleteOne(d *Dispense) *DispenseDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DispenseClient) DeleteOneID(id int) *DispenseDeleteOne {
	builder := c.Delete().Where(dispense.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DispenseDeleteOne{builder}
}

// Create returns a query builder for Dispense.
func (c *DispenseClient) Query() *DispenseQuery {
	return &DispenseQuery{config: c.config}
}

// Get returns a Dispense entity by its id.
func (c *DispenseClient) Get(ctx context.Context, id int) (*Dispense, error) {
	return c.Query().Where(dispense.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DispenseClient) GetX(ctx context.Context, id int) *Dispense {
	d, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return d
}

// QueryDrug queries the drug edge of a Dispense.
func (c *DispenseClient) QueryDrug(d *Dispense) *DrugQuery {
	query := &DrugQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(dispense.Table, dispense.FieldID, id),
			sqlgraph.To(drug.Table, drug.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, dispense.DrugTable, dispense.DrugColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUsers queries the users edge of a Dispense.
func (c *DispenseClient) QueryUsers(d *Dispense) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(dispense.Table, dispense.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, dispense.UsersTable, dispense.UsersColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDrugs queries the drugs edge of a Dispense.
func (c *DispenseClient) QueryDrugs(d *Dispense) *DrugQuery {
	query := &DrugQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(dispense.Table, dispense.FieldID, id),
			sqlgraph.To(drug.Table, drug.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, dispense.DrugsTable, dispense.DrugsColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DispenseClient) Hooks() []Hook {
	return c.hooks.Dispense
}

// DrugClient is a client for the Drug schema.
type DrugClient struct {
	config
}

// NewDrugClient returns a client for the Drug from the given config.
func NewDrugClient(c config) *DrugClient {
	return &DrugClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `drug.Hooks(f(g(h())))`.
func (c *DrugClient) Use(hooks ...Hook) {
	c.hooks.Drug = append(c.hooks.Drug, hooks...)
}

// Create returns a create builder for Drug.
func (c *DrugClient) Create() *DrugCreate {
	mutation := newDrugMutation(c.config, OpCreate)
	return &DrugCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Drug.
func (c *DrugClient) Update() *DrugUpdate {
	mutation := newDrugMutation(c.config, OpUpdate)
	return &DrugUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DrugClient) UpdateOne(d *Drug) *DrugUpdateOne {
	mutation := newDrugMutation(c.config, OpUpdateOne, withDrug(d))
	return &DrugUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DrugClient) UpdateOneID(id int) *DrugUpdateOne {
	mutation := newDrugMutation(c.config, OpUpdateOne, withDrugID(id))
	return &DrugUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Drug.
func (c *DrugClient) Delete() *DrugDelete {
	mutation := newDrugMutation(c.config, OpDelete)
	return &DrugDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DrugClient) DeleteOne(d *Drug) *DrugDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DrugClient) DeleteOneID(id int) *DrugDeleteOne {
	builder := c.Delete().Where(drug.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DrugDeleteOne{builder}
}

// Create returns a query builder for Drug.
func (c *DrugClient) Query() *DrugQuery {
	return &DrugQuery{config: c.config}
}

// Get returns a Drug entity by its id.
func (c *DrugClient) Get(ctx context.Context, id int) (*Drug, error) {
	return c.Query().Where(drug.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DrugClient) GetX(ctx context.Context, id int) *Drug {
	d, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return d
}

// QueryForm queries the form edge of a Drug.
func (c *DrugClient) QueryForm(d *Drug) *FormQuery {
	query := &FormQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(drug.Table, drug.FieldID, id),
			sqlgraph.To(form.Table, form.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, drug.FormTable, drug.FormColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUnit queries the unit edge of a Drug.
func (c *DrugClient) QueryUnit(d *Drug) *UnitQuery {
	query := &UnitQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(drug.Table, drug.FieldID, id),
			sqlgraph.To(unit.Table, unit.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, drug.UnitTable, drug.UnitColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryVolume queries the volume edge of a Drug.
func (c *DrugClient) QueryVolume(d *Drug) *VolumeQuery {
	query := &VolumeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(drug.Table, drug.FieldID, id),
			sqlgraph.To(volume.Table, volume.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, drug.VolumeTable, drug.VolumeColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUsers queries the users edge of a Drug.
func (c *DrugClient) QueryUsers(d *Drug) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(drug.Table, drug.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, drug.UsersTable, drug.UsersColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDispenses queries the dispenses edge of a Drug.
func (c *DrugClient) QueryDispenses(d *Drug) *DispenseQuery {
	query := &DispenseQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(drug.Table, drug.FieldID, id),
			sqlgraph.To(dispense.Table, dispense.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, drug.DispensesTable, drug.DispensesColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DrugClient) Hooks() []Hook {
	return c.hooks.Drug
}

// FormClient is a client for the Form schema.
type FormClient struct {
	config
}

// NewFormClient returns a client for the Form from the given config.
func NewFormClient(c config) *FormClient {
	return &FormClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `form.Hooks(f(g(h())))`.
func (c *FormClient) Use(hooks ...Hook) {
	c.hooks.Form = append(c.hooks.Form, hooks...)
}

// Create returns a create builder for Form.
func (c *FormClient) Create() *FormCreate {
	mutation := newFormMutation(c.config, OpCreate)
	return &FormCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Form.
func (c *FormClient) Update() *FormUpdate {
	mutation := newFormMutation(c.config, OpUpdate)
	return &FormUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FormClient) UpdateOne(f *Form) *FormUpdateOne {
	mutation := newFormMutation(c.config, OpUpdateOne, withForm(f))
	return &FormUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FormClient) UpdateOneID(id int) *FormUpdateOne {
	mutation := newFormMutation(c.config, OpUpdateOne, withFormID(id))
	return &FormUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Form.
func (c *FormClient) Delete() *FormDelete {
	mutation := newFormMutation(c.config, OpDelete)
	return &FormDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *FormClient) DeleteOne(f *Form) *FormDeleteOne {
	return c.DeleteOneID(f.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *FormClient) DeleteOneID(id int) *FormDeleteOne {
	builder := c.Delete().Where(form.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FormDeleteOne{builder}
}

// Create returns a query builder for Form.
func (c *FormClient) Query() *FormQuery {
	return &FormQuery{config: c.config}
}

// Get returns a Form entity by its id.
func (c *FormClient) Get(ctx context.Context, id int) (*Form, error) {
	return c.Query().Where(form.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FormClient) GetX(ctx context.Context, id int) *Form {
	f, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return f
}

// QueryDrugs queries the drugs edge of a Form.
func (c *FormClient) QueryDrugs(f *Form) *DrugQuery {
	query := &DrugQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(form.Table, form.FieldID, id),
			sqlgraph.To(drug.Table, drug.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, form.DrugsTable, form.DrugsColumn),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *FormClient) Hooks() []Hook {
	return c.hooks.Form
}

// GenderClient is a client for the Gender schema.
type GenderClient struct {
	config
}

// NewGenderClient returns a client for the Gender from the given config.
func NewGenderClient(c config) *GenderClient {
	return &GenderClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `gender.Hooks(f(g(h())))`.
func (c *GenderClient) Use(hooks ...Hook) {
	c.hooks.Gender = append(c.hooks.Gender, hooks...)
}

// Create returns a create builder for Gender.
func (c *GenderClient) Create() *GenderCreate {
	mutation := newGenderMutation(c.config, OpCreate)
	return &GenderCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Gender.
func (c *GenderClient) Update() *GenderUpdate {
	mutation := newGenderMutation(c.config, OpUpdate)
	return &GenderUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GenderClient) UpdateOne(ge *Gender) *GenderUpdateOne {
	mutation := newGenderMutation(c.config, OpUpdateOne, withGender(ge))
	return &GenderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GenderClient) UpdateOneID(id int) *GenderUpdateOne {
	mutation := newGenderMutation(c.config, OpUpdateOne, withGenderID(id))
	return &GenderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Gender.
func (c *GenderClient) Delete() *GenderDelete {
	mutation := newGenderMutation(c.config, OpDelete)
	return &GenderDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *GenderClient) DeleteOne(ge *Gender) *GenderDeleteOne {
	return c.DeleteOneID(ge.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *GenderClient) DeleteOneID(id int) *GenderDeleteOne {
	builder := c.Delete().Where(gender.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GenderDeleteOne{builder}
}

// Create returns a query builder for Gender.
func (c *GenderClient) Query() *GenderQuery {
	return &GenderQuery{config: c.config}
}

// Get returns a Gender entity by its id.
func (c *GenderClient) Get(ctx context.Context, id int) (*Gender, error) {
	return c.Query().Where(gender.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GenderClient) GetX(ctx context.Context, id int) *Gender {
	ge, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return ge
}

// QueryUsers queries the users edge of a Gender.
func (c *GenderClient) QueryUsers(ge *Gender) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ge.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(gender.Table, gender.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, gender.UsersTable, gender.UsersColumn),
		)
		fromV = sqlgraph.Neighbors(ge.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GenderClient) Hooks() []Hook {
	return c.hooks.Gender
}

// PositionClient is a client for the Position schema.
type PositionClient struct {
	config
}

// NewPositionClient returns a client for the Position from the given config.
func NewPositionClient(c config) *PositionClient {
	return &PositionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `position.Hooks(f(g(h())))`.
func (c *PositionClient) Use(hooks ...Hook) {
	c.hooks.Position = append(c.hooks.Position, hooks...)
}

// Create returns a create builder for Position.
func (c *PositionClient) Create() *PositionCreate {
	mutation := newPositionMutation(c.config, OpCreate)
	return &PositionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Position.
func (c *PositionClient) Update() *PositionUpdate {
	mutation := newPositionMutation(c.config, OpUpdate)
	return &PositionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PositionClient) UpdateOne(po *Position) *PositionUpdateOne {
	mutation := newPositionMutation(c.config, OpUpdateOne, withPosition(po))
	return &PositionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PositionClient) UpdateOneID(id int) *PositionUpdateOne {
	mutation := newPositionMutation(c.config, OpUpdateOne, withPositionID(id))
	return &PositionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Position.
func (c *PositionClient) Delete() *PositionDelete {
	mutation := newPositionMutation(c.config, OpDelete)
	return &PositionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PositionClient) DeleteOne(po *Position) *PositionDeleteOne {
	return c.DeleteOneID(po.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PositionClient) DeleteOneID(id int) *PositionDeleteOne {
	builder := c.Delete().Where(position.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PositionDeleteOne{builder}
}

// Create returns a query builder for Position.
func (c *PositionClient) Query() *PositionQuery {
	return &PositionQuery{config: c.config}
}

// Get returns a Position entity by its id.
func (c *PositionClient) Get(ctx context.Context, id int) (*Position, error) {
	return c.Query().Where(position.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PositionClient) GetX(ctx context.Context, id int) *Position {
	po, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return po
}

// QueryUsers queries the users edge of a Position.
func (c *PositionClient) QueryUsers(po *Position) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(position.Table, position.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, position.UsersTable, position.UsersColumn),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PositionClient) Hooks() []Hook {
	return c.hooks.Position
}

// TitleClient is a client for the Title schema.
type TitleClient struct {
	config
}

// NewTitleClient returns a client for the Title from the given config.
func NewTitleClient(c config) *TitleClient {
	return &TitleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `title.Hooks(f(g(h())))`.
func (c *TitleClient) Use(hooks ...Hook) {
	c.hooks.Title = append(c.hooks.Title, hooks...)
}

// Create returns a create builder for Title.
func (c *TitleClient) Create() *TitleCreate {
	mutation := newTitleMutation(c.config, OpCreate)
	return &TitleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Title.
func (c *TitleClient) Update() *TitleUpdate {
	mutation := newTitleMutation(c.config, OpUpdate)
	return &TitleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TitleClient) UpdateOne(t *Title) *TitleUpdateOne {
	mutation := newTitleMutation(c.config, OpUpdateOne, withTitle(t))
	return &TitleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TitleClient) UpdateOneID(id int) *TitleUpdateOne {
	mutation := newTitleMutation(c.config, OpUpdateOne, withTitleID(id))
	return &TitleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Title.
func (c *TitleClient) Delete() *TitleDelete {
	mutation := newTitleMutation(c.config, OpDelete)
	return &TitleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TitleClient) DeleteOne(t *Title) *TitleDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TitleClient) DeleteOneID(id int) *TitleDeleteOne {
	builder := c.Delete().Where(title.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TitleDeleteOne{builder}
}

// Create returns a query builder for Title.
func (c *TitleClient) Query() *TitleQuery {
	return &TitleQuery{config: c.config}
}

// Get returns a Title entity by its id.
func (c *TitleClient) Get(ctx context.Context, id int) (*Title, error) {
	return c.Query().Where(title.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TitleClient) GetX(ctx context.Context, id int) *Title {
	t, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return t
}

// QueryUsers queries the users edge of a Title.
func (c *TitleClient) QueryUsers(t *Title) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(title.Table, title.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, title.UsersTable, title.UsersColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TitleClient) Hooks() []Hook {
	return c.hooks.Title
}

// UnitClient is a client for the Unit schema.
type UnitClient struct {
	config
}

// NewUnitClient returns a client for the Unit from the given config.
func NewUnitClient(c config) *UnitClient {
	return &UnitClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `unit.Hooks(f(g(h())))`.
func (c *UnitClient) Use(hooks ...Hook) {
	c.hooks.Unit = append(c.hooks.Unit, hooks...)
}

// Create returns a create builder for Unit.
func (c *UnitClient) Create() *UnitCreate {
	mutation := newUnitMutation(c.config, OpCreate)
	return &UnitCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Unit.
func (c *UnitClient) Update() *UnitUpdate {
	mutation := newUnitMutation(c.config, OpUpdate)
	return &UnitUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UnitClient) UpdateOne(u *Unit) *UnitUpdateOne {
	mutation := newUnitMutation(c.config, OpUpdateOne, withUnit(u))
	return &UnitUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UnitClient) UpdateOneID(id int) *UnitUpdateOne {
	mutation := newUnitMutation(c.config, OpUpdateOne, withUnitID(id))
	return &UnitUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Unit.
func (c *UnitClient) Delete() *UnitDelete {
	mutation := newUnitMutation(c.config, OpDelete)
	return &UnitDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UnitClient) DeleteOne(u *Unit) *UnitDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UnitClient) DeleteOneID(id int) *UnitDeleteOne {
	builder := c.Delete().Where(unit.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UnitDeleteOne{builder}
}

// Create returns a query builder for Unit.
func (c *UnitClient) Query() *UnitQuery {
	return &UnitQuery{config: c.config}
}

// Get returns a Unit entity by its id.
func (c *UnitClient) Get(ctx context.Context, id int) (*Unit, error) {
	return c.Query().Where(unit.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UnitClient) GetX(ctx context.Context, id int) *Unit {
	u, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return u
}

// QueryDrugs queries the drugs edge of a Unit.
func (c *UnitClient) QueryDrugs(u *Unit) *DrugQuery {
	query := &DrugQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(unit.Table, unit.FieldID, id),
			sqlgraph.To(drug.Table, drug.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, unit.DrugsTable, unit.DrugsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UnitClient) Hooks() []Hook {
	return c.hooks.Unit
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Create returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	u, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return u
}

// QueryGender queries the gender edge of a User.
func (c *UserClient) QueryGender(u *User) *GenderQuery {
	query := &GenderQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(gender.Table, gender.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, user.GenderTable, user.GenderColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPosition queries the position edge of a User.
func (c *UserClient) QueryPosition(u *User) *PositionQuery {
	query := &PositionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(position.Table, position.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, user.PositionTable, user.PositionColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTitle queries the title edge of a User.
func (c *UserClient) QueryTitle(u *User) *TitleQuery {
	query := &TitleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(title.Table, title.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, user.TitleTable, user.TitleColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// VolumeClient is a client for the Volume schema.
type VolumeClient struct {
	config
}

// NewVolumeClient returns a client for the Volume from the given config.
func NewVolumeClient(c config) *VolumeClient {
	return &VolumeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `volume.Hooks(f(g(h())))`.
func (c *VolumeClient) Use(hooks ...Hook) {
	c.hooks.Volume = append(c.hooks.Volume, hooks...)
}

// Create returns a create builder for Volume.
func (c *VolumeClient) Create() *VolumeCreate {
	mutation := newVolumeMutation(c.config, OpCreate)
	return &VolumeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Volume.
func (c *VolumeClient) Update() *VolumeUpdate {
	mutation := newVolumeMutation(c.config, OpUpdate)
	return &VolumeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *VolumeClient) UpdateOne(v *Volume) *VolumeUpdateOne {
	mutation := newVolumeMutation(c.config, OpUpdateOne, withVolume(v))
	return &VolumeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *VolumeClient) UpdateOneID(id int) *VolumeUpdateOne {
	mutation := newVolumeMutation(c.config, OpUpdateOne, withVolumeID(id))
	return &VolumeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Volume.
func (c *VolumeClient) Delete() *VolumeDelete {
	mutation := newVolumeMutation(c.config, OpDelete)
	return &VolumeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *VolumeClient) DeleteOne(v *Volume) *VolumeDeleteOne {
	return c.DeleteOneID(v.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *VolumeClient) DeleteOneID(id int) *VolumeDeleteOne {
	builder := c.Delete().Where(volume.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &VolumeDeleteOne{builder}
}

// Create returns a query builder for Volume.
func (c *VolumeClient) Query() *VolumeQuery {
	return &VolumeQuery{config: c.config}
}

// Get returns a Volume entity by its id.
func (c *VolumeClient) Get(ctx context.Context, id int) (*Volume, error) {
	return c.Query().Where(volume.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *VolumeClient) GetX(ctx context.Context, id int) *Volume {
	v, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return v
}

// QueryDrugs queries the drugs edge of a Volume.
func (c *VolumeClient) QueryDrugs(v *Volume) *DrugQuery {
	query := &DrugQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(volume.Table, volume.FieldID, id),
			sqlgraph.To(drug.Table, drug.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, volume.DrugsTable, volume.DrugsColumn),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *VolumeClient) Hooks() []Hook {
	return c.hooks.Volume
}
