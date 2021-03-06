// Code generated by SQLBoiler 4.4.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// PortScan is an object representing the database table.
type PortScan struct {
	ID        int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	Done      int64     `boil:"done" json:"done" toml:"done" yaml:"done"`
	NodeID    int64     `boil:"node_id" json:"node_id" toml:"node_id" yaml:"node_id"`

	R *portScanR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L portScanL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PortScanColumns = struct {
	ID        string
	CreatedAt string
	Done      string
	NodeID    string
}{
	ID:        "id",
	CreatedAt: "created_at",
	Done:      "done",
	NodeID:    "node_id",
}

// Generated where

var PortScanWhere = struct {
	ID        whereHelperint64
	CreatedAt whereHelpertime_Time
	Done      whereHelperint64
	NodeID    whereHelperint64
}{
	ID:        whereHelperint64{field: "\"port_scans\".\"id\""},
	CreatedAt: whereHelpertime_Time{field: "\"port_scans\".\"created_at\""},
	Done:      whereHelperint64{field: "\"port_scans\".\"done\""},
	NodeID:    whereHelperint64{field: "\"port_scans\".\"node_id\""},
}

// PortScanRels is where relationship names are stored.
var PortScanRels = struct {
	Node  string
	Ports string
}{
	Node:  "Node",
	Ports: "Ports",
}

// portScanR is where relationships are stored.
type portScanR struct {
	Node  *Node     `boil:"Node" json:"Node" toml:"Node" yaml:"Node"`
	Ports PortSlice `boil:"Ports" json:"Ports" toml:"Ports" yaml:"Ports"`
}

// NewStruct creates a new relationship struct
func (*portScanR) NewStruct() *portScanR {
	return &portScanR{}
}

// portScanL is where Load methods for each relationship are stored.
type portScanL struct{}

var (
	portScanAllColumns            = []string{"id", "created_at", "done", "node_id"}
	portScanColumnsWithoutDefault = []string{"created_at", "done", "node_id"}
	portScanColumnsWithDefault    = []string{"id"}
	portScanPrimaryKeyColumns     = []string{"id"}
)

type (
	// PortScanSlice is an alias for a slice of pointers to PortScan.
	// This should generally be used opposed to []PortScan.
	PortScanSlice []*PortScan
	// PortScanHook is the signature for custom PortScan hook methods
	PortScanHook func(context.Context, boil.ContextExecutor, *PortScan) error

	portScanQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	portScanType                 = reflect.TypeOf(&PortScan{})
	portScanMapping              = queries.MakeStructMapping(portScanType)
	portScanPrimaryKeyMapping, _ = queries.BindMapping(portScanType, portScanMapping, portScanPrimaryKeyColumns)
	portScanInsertCacheMut       sync.RWMutex
	portScanInsertCache          = make(map[string]insertCache)
	portScanUpdateCacheMut       sync.RWMutex
	portScanUpdateCache          = make(map[string]updateCache)
	portScanUpsertCacheMut       sync.RWMutex
	portScanUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var portScanBeforeInsertHooks []PortScanHook
var portScanBeforeUpdateHooks []PortScanHook
var portScanBeforeDeleteHooks []PortScanHook
var portScanBeforeUpsertHooks []PortScanHook

var portScanAfterInsertHooks []PortScanHook
var portScanAfterSelectHooks []PortScanHook
var portScanAfterUpdateHooks []PortScanHook
var portScanAfterDeleteHooks []PortScanHook
var portScanAfterUpsertHooks []PortScanHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *PortScan) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range portScanBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *PortScan) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range portScanBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *PortScan) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range portScanBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *PortScan) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range portScanBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *PortScan) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range portScanAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *PortScan) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range portScanAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *PortScan) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range portScanAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *PortScan) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range portScanAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *PortScan) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range portScanAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPortScanHook registers your hook function for all future operations.
func AddPortScanHook(hookPoint boil.HookPoint, portScanHook PortScanHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		portScanBeforeInsertHooks = append(portScanBeforeInsertHooks, portScanHook)
	case boil.BeforeUpdateHook:
		portScanBeforeUpdateHooks = append(portScanBeforeUpdateHooks, portScanHook)
	case boil.BeforeDeleteHook:
		portScanBeforeDeleteHooks = append(portScanBeforeDeleteHooks, portScanHook)
	case boil.BeforeUpsertHook:
		portScanBeforeUpsertHooks = append(portScanBeforeUpsertHooks, portScanHook)
	case boil.AfterInsertHook:
		portScanAfterInsertHooks = append(portScanAfterInsertHooks, portScanHook)
	case boil.AfterSelectHook:
		portScanAfterSelectHooks = append(portScanAfterSelectHooks, portScanHook)
	case boil.AfterUpdateHook:
		portScanAfterUpdateHooks = append(portScanAfterUpdateHooks, portScanHook)
	case boil.AfterDeleteHook:
		portScanAfterDeleteHooks = append(portScanAfterDeleteHooks, portScanHook)
	case boil.AfterUpsertHook:
		portScanAfterUpsertHooks = append(portScanAfterUpsertHooks, portScanHook)
	}
}

// One returns a single portScan record from the query.
func (q portScanQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PortScan, error) {
	o := &PortScan{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for port_scans")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all PortScan records from the query.
func (q portScanQuery) All(ctx context.Context, exec boil.ContextExecutor) (PortScanSlice, error) {
	var o []*PortScan

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to PortScan slice")
	}

	if len(portScanAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all PortScan records in the query.
func (q portScanQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count port_scans rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q portScanQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if port_scans exists")
	}

	return count > 0, nil
}

// Node pointed to by the foreign key.
func (o *PortScan) Node(mods ...qm.QueryMod) nodeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.NodeID),
	}

	queryMods = append(queryMods, mods...)

	query := Nodes(queryMods...)
	queries.SetFrom(query.Query, "\"nodes\"")

	return query
}

// Ports retrieves all the port's Ports with an executor.
func (o *PortScan) Ports(mods ...qm.QueryMod) portQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"ports\".\"port_scan_id\"=?", o.ID),
	)

	query := Ports(queryMods...)
	queries.SetFrom(query.Query, "\"ports\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"ports\".*"})
	}

	return query
}

// LoadNode allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (portScanL) LoadNode(ctx context.Context, e boil.ContextExecutor, singular bool, maybePortScan interface{}, mods queries.Applicator) error {
	var slice []*PortScan
	var object *PortScan

	if singular {
		object = maybePortScan.(*PortScan)
	} else {
		slice = *maybePortScan.(*[]*PortScan)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &portScanR{}
		}
		args = append(args, object.NodeID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &portScanR{}
			}

			for _, a := range args {
				if a == obj.NodeID {
					continue Outer
				}
			}

			args = append(args, obj.NodeID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`nodes`),
		qm.WhereIn(`nodes.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Node")
	}

	var resultSlice []*Node
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Node")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for nodes")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for nodes")
	}

	if len(portScanAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Node = foreign
		if foreign.R == nil {
			foreign.R = &nodeR{}
		}
		foreign.R.PortScans = append(foreign.R.PortScans, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.NodeID == foreign.ID {
				local.R.Node = foreign
				if foreign.R == nil {
					foreign.R = &nodeR{}
				}
				foreign.R.PortScans = append(foreign.R.PortScans, local)
				break
			}
		}
	}

	return nil
}

// LoadPorts allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (portScanL) LoadPorts(ctx context.Context, e boil.ContextExecutor, singular bool, maybePortScan interface{}, mods queries.Applicator) error {
	var slice []*PortScan
	var object *PortScan

	if singular {
		object = maybePortScan.(*PortScan)
	} else {
		slice = *maybePortScan.(*[]*PortScan)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &portScanR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &portScanR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`ports`),
		qm.WhereIn(`ports.port_scan_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load ports")
	}

	var resultSlice []*Port
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice ports")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on ports")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for ports")
	}

	if len(portAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Ports = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &portR{}
			}
			foreign.R.PortScan = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.PortScanID {
				local.R.Ports = append(local.R.Ports, foreign)
				if foreign.R == nil {
					foreign.R = &portR{}
				}
				foreign.R.PortScan = local
				break
			}
		}
	}

	return nil
}

// SetNode of the portScan to the related item.
// Sets o.R.Node to related.
// Adds o to related.R.PortScans.
func (o *PortScan) SetNode(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Node) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"port_scans\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"node_id"}),
		strmangle.WhereClause("\"", "\"", 0, portScanPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.NodeID = related.ID
	if o.R == nil {
		o.R = &portScanR{
			Node: related,
		}
	} else {
		o.R.Node = related
	}

	if related.R == nil {
		related.R = &nodeR{
			PortScans: PortScanSlice{o},
		}
	} else {
		related.R.PortScans = append(related.R.PortScans, o)
	}

	return nil
}

// AddPorts adds the given related objects to the existing relationships
// of the port_scan, optionally inserting them as new records.
// Appends related to o.R.Ports.
// Sets related.R.PortScan appropriately.
func (o *PortScan) AddPorts(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Port) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.PortScanID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"ports\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 0, []string{"port_scan_id"}),
				strmangle.WhereClause("\"", "\"", 0, portPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.PortScanID = o.ID
		}
	}

	if o.R == nil {
		o.R = &portScanR{
			Ports: related,
		}
	} else {
		o.R.Ports = append(o.R.Ports, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &portR{
				PortScan: o,
			}
		} else {
			rel.R.PortScan = o
		}
	}
	return nil
}

// PortScans retrieves all the records using an executor.
func PortScans(mods ...qm.QueryMod) portScanQuery {
	mods = append(mods, qm.From("\"port_scans\""))
	return portScanQuery{NewQuery(mods...)}
}

// FindPortScan retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPortScan(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*PortScan, error) {
	portScanObj := &PortScan{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"port_scans\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, portScanObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from port_scans")
	}

	return portScanObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PortScan) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no port_scans provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(portScanColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	portScanInsertCacheMut.RLock()
	cache, cached := portScanInsertCache[key]
	portScanInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			portScanAllColumns,
			portScanColumnsWithDefault,
			portScanColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(portScanType, portScanMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(portScanType, portScanMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"port_scans\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"port_scans\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"port_scans\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, portScanPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into port_scans")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == portScanMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for port_scans")
	}

CacheNoHooks:
	if !cached {
		portScanInsertCacheMut.Lock()
		portScanInsertCache[key] = cache
		portScanInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the PortScan.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PortScan) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	portScanUpdateCacheMut.RLock()
	cache, cached := portScanUpdateCache[key]
	portScanUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			portScanAllColumns,
			portScanPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update port_scans, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"port_scans\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, portScanPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(portScanType, portScanMapping, append(wl, portScanPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update port_scans row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for port_scans")
	}

	if !cached {
		portScanUpdateCacheMut.Lock()
		portScanUpdateCache[key] = cache
		portScanUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q portScanQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for port_scans")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for port_scans")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PortScanSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), portScanPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"port_scans\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, portScanPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in portScan slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all portScan")
	}
	return rowsAff, nil
}

// Delete deletes a single PortScan record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PortScan) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PortScan provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), portScanPrimaryKeyMapping)
	sql := "DELETE FROM \"port_scans\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from port_scans")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for port_scans")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q portScanQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no portScanQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from port_scans")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for port_scans")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PortScanSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(portScanBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), portScanPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"port_scans\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, portScanPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from portScan slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for port_scans")
	}

	if len(portScanAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *PortScan) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPortScan(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PortScanSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PortScanSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), portScanPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"port_scans\".* FROM \"port_scans\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, portScanPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PortScanSlice")
	}

	*o = slice

	return nil
}

// PortScanExists checks if the PortScan row exists.
func PortScanExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"port_scans\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if port_scans exists")
	}

	return exists, nil
}
