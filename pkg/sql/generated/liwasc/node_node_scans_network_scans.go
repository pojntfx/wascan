// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// NodeNodeScansNetworkScan is an object representing the database table.
type NodeNodeScansNetworkScan struct {
	ID            int64  `boil:"id" json:"id" toml:"id" yaml:"id"`
	NodeID        string `boil:"node_id" json:"node_id" toml:"node_id" yaml:"node_id"`
	NetworkScanID int64  `boil:"network_scan_id" json:"network_scan_id" toml:"network_scan_id" yaml:"network_scan_id"`
	NodeScanID    int64  `boil:"node_scan_id" json:"node_scan_id" toml:"node_scan_id" yaml:"node_scan_id"`

	R *nodeNodeScansNetworkScanR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L nodeNodeScansNetworkScanL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var NodeNodeScansNetworkScanColumns = struct {
	ID            string
	NodeID        string
	NetworkScanID string
	NodeScanID    string
}{
	ID:            "id",
	NodeID:        "node_id",
	NetworkScanID: "network_scan_id",
	NodeScanID:    "node_scan_id",
}

// Generated where

var NodeNodeScansNetworkScanWhere = struct {
	ID            whereHelperint64
	NodeID        whereHelperstring
	NetworkScanID whereHelperint64
	NodeScanID    whereHelperint64
}{
	ID:            whereHelperint64{field: "\"node_node_scans_network_scans\".\"id\""},
	NodeID:        whereHelperstring{field: "\"node_node_scans_network_scans\".\"node_id\""},
	NetworkScanID: whereHelperint64{field: "\"node_node_scans_network_scans\".\"network_scan_id\""},
	NodeScanID:    whereHelperint64{field: "\"node_node_scans_network_scans\".\"node_scan_id\""},
}

// NodeNodeScansNetworkScanRels is where relationship names are stored.
var NodeNodeScansNetworkScanRels = struct {
}{}

// nodeNodeScansNetworkScanR is where relationships are stored.
type nodeNodeScansNetworkScanR struct {
}

// NewStruct creates a new relationship struct
func (*nodeNodeScansNetworkScanR) NewStruct() *nodeNodeScansNetworkScanR {
	return &nodeNodeScansNetworkScanR{}
}

// nodeNodeScansNetworkScanL is where Load methods for each relationship are stored.
type nodeNodeScansNetworkScanL struct{}

var (
	nodeNodeScansNetworkScanAllColumns            = []string{"id", "node_id", "network_scan_id", "node_scan_id"}
	nodeNodeScansNetworkScanColumnsWithoutDefault = []string{"node_id", "network_scan_id", "node_scan_id"}
	nodeNodeScansNetworkScanColumnsWithDefault    = []string{"id"}
	nodeNodeScansNetworkScanPrimaryKeyColumns     = []string{"id"}
)

type (
	// NodeNodeScansNetworkScanSlice is an alias for a slice of pointers to NodeNodeScansNetworkScan.
	// This should generally be used opposed to []NodeNodeScansNetworkScan.
	NodeNodeScansNetworkScanSlice []*NodeNodeScansNetworkScan
	// NodeNodeScansNetworkScanHook is the signature for custom NodeNodeScansNetworkScan hook methods
	NodeNodeScansNetworkScanHook func(context.Context, boil.ContextExecutor, *NodeNodeScansNetworkScan) error

	nodeNodeScansNetworkScanQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	nodeNodeScansNetworkScanType                 = reflect.TypeOf(&NodeNodeScansNetworkScan{})
	nodeNodeScansNetworkScanMapping              = queries.MakeStructMapping(nodeNodeScansNetworkScanType)
	nodeNodeScansNetworkScanPrimaryKeyMapping, _ = queries.BindMapping(nodeNodeScansNetworkScanType, nodeNodeScansNetworkScanMapping, nodeNodeScansNetworkScanPrimaryKeyColumns)
	nodeNodeScansNetworkScanInsertCacheMut       sync.RWMutex
	nodeNodeScansNetworkScanInsertCache          = make(map[string]insertCache)
	nodeNodeScansNetworkScanUpdateCacheMut       sync.RWMutex
	nodeNodeScansNetworkScanUpdateCache          = make(map[string]updateCache)
	nodeNodeScansNetworkScanUpsertCacheMut       sync.RWMutex
	nodeNodeScansNetworkScanUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var nodeNodeScansNetworkScanBeforeInsertHooks []NodeNodeScansNetworkScanHook
var nodeNodeScansNetworkScanBeforeUpdateHooks []NodeNodeScansNetworkScanHook
var nodeNodeScansNetworkScanBeforeDeleteHooks []NodeNodeScansNetworkScanHook
var nodeNodeScansNetworkScanBeforeUpsertHooks []NodeNodeScansNetworkScanHook

var nodeNodeScansNetworkScanAfterInsertHooks []NodeNodeScansNetworkScanHook
var nodeNodeScansNetworkScanAfterSelectHooks []NodeNodeScansNetworkScanHook
var nodeNodeScansNetworkScanAfterUpdateHooks []NodeNodeScansNetworkScanHook
var nodeNodeScansNetworkScanAfterDeleteHooks []NodeNodeScansNetworkScanHook
var nodeNodeScansNetworkScanAfterUpsertHooks []NodeNodeScansNetworkScanHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *NodeNodeScansNetworkScan) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nodeNodeScansNetworkScanBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *NodeNodeScansNetworkScan) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nodeNodeScansNetworkScanBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *NodeNodeScansNetworkScan) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nodeNodeScansNetworkScanBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *NodeNodeScansNetworkScan) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nodeNodeScansNetworkScanBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *NodeNodeScansNetworkScan) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nodeNodeScansNetworkScanAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *NodeNodeScansNetworkScan) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nodeNodeScansNetworkScanAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *NodeNodeScansNetworkScan) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nodeNodeScansNetworkScanAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *NodeNodeScansNetworkScan) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nodeNodeScansNetworkScanAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *NodeNodeScansNetworkScan) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range nodeNodeScansNetworkScanAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddNodeNodeScansNetworkScanHook registers your hook function for all future operations.
func AddNodeNodeScansNetworkScanHook(hookPoint boil.HookPoint, nodeNodeScansNetworkScanHook NodeNodeScansNetworkScanHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		nodeNodeScansNetworkScanBeforeInsertHooks = append(nodeNodeScansNetworkScanBeforeInsertHooks, nodeNodeScansNetworkScanHook)
	case boil.BeforeUpdateHook:
		nodeNodeScansNetworkScanBeforeUpdateHooks = append(nodeNodeScansNetworkScanBeforeUpdateHooks, nodeNodeScansNetworkScanHook)
	case boil.BeforeDeleteHook:
		nodeNodeScansNetworkScanBeforeDeleteHooks = append(nodeNodeScansNetworkScanBeforeDeleteHooks, nodeNodeScansNetworkScanHook)
	case boil.BeforeUpsertHook:
		nodeNodeScansNetworkScanBeforeUpsertHooks = append(nodeNodeScansNetworkScanBeforeUpsertHooks, nodeNodeScansNetworkScanHook)
	case boil.AfterInsertHook:
		nodeNodeScansNetworkScanAfterInsertHooks = append(nodeNodeScansNetworkScanAfterInsertHooks, nodeNodeScansNetworkScanHook)
	case boil.AfterSelectHook:
		nodeNodeScansNetworkScanAfterSelectHooks = append(nodeNodeScansNetworkScanAfterSelectHooks, nodeNodeScansNetworkScanHook)
	case boil.AfterUpdateHook:
		nodeNodeScansNetworkScanAfterUpdateHooks = append(nodeNodeScansNetworkScanAfterUpdateHooks, nodeNodeScansNetworkScanHook)
	case boil.AfterDeleteHook:
		nodeNodeScansNetworkScanAfterDeleteHooks = append(nodeNodeScansNetworkScanAfterDeleteHooks, nodeNodeScansNetworkScanHook)
	case boil.AfterUpsertHook:
		nodeNodeScansNetworkScanAfterUpsertHooks = append(nodeNodeScansNetworkScanAfterUpsertHooks, nodeNodeScansNetworkScanHook)
	}
}

// One returns a single nodeNodeScansNetworkScan record from the query.
func (q nodeNodeScansNetworkScanQuery) One(ctx context.Context, exec boil.ContextExecutor) (*NodeNodeScansNetworkScan, error) {
	o := &NodeNodeScansNetworkScan{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for node_node_scans_network_scans")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all NodeNodeScansNetworkScan records from the query.
func (q nodeNodeScansNetworkScanQuery) All(ctx context.Context, exec boil.ContextExecutor) (NodeNodeScansNetworkScanSlice, error) {
	var o []*NodeNodeScansNetworkScan

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to NodeNodeScansNetworkScan slice")
	}

	if len(nodeNodeScansNetworkScanAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all NodeNodeScansNetworkScan records in the query.
func (q nodeNodeScansNetworkScanQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count node_node_scans_network_scans rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q nodeNodeScansNetworkScanQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if node_node_scans_network_scans exists")
	}

	return count > 0, nil
}

// NodeNodeScansNetworkScans retrieves all the records using an executor.
func NodeNodeScansNetworkScans(mods ...qm.QueryMod) nodeNodeScansNetworkScanQuery {
	mods = append(mods, qm.From("\"node_node_scans_network_scans\""))
	return nodeNodeScansNetworkScanQuery{NewQuery(mods...)}
}

// FindNodeNodeScansNetworkScan retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindNodeNodeScansNetworkScan(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*NodeNodeScansNetworkScan, error) {
	nodeNodeScansNetworkScanObj := &NodeNodeScansNetworkScan{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"node_node_scans_network_scans\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, nodeNodeScansNetworkScanObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from node_node_scans_network_scans")
	}

	return nodeNodeScansNetworkScanObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *NodeNodeScansNetworkScan) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no node_node_scans_network_scans provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(nodeNodeScansNetworkScanColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	nodeNodeScansNetworkScanInsertCacheMut.RLock()
	cache, cached := nodeNodeScansNetworkScanInsertCache[key]
	nodeNodeScansNetworkScanInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			nodeNodeScansNetworkScanAllColumns,
			nodeNodeScansNetworkScanColumnsWithDefault,
			nodeNodeScansNetworkScanColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(nodeNodeScansNetworkScanType, nodeNodeScansNetworkScanMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(nodeNodeScansNetworkScanType, nodeNodeScansNetworkScanMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"node_node_scans_network_scans\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"node_node_scans_network_scans\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"node_node_scans_network_scans\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, nodeNodeScansNetworkScanPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into node_node_scans_network_scans")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == nodeNodeScansNetworkScanMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for node_node_scans_network_scans")
	}

CacheNoHooks:
	if !cached {
		nodeNodeScansNetworkScanInsertCacheMut.Lock()
		nodeNodeScansNetworkScanInsertCache[key] = cache
		nodeNodeScansNetworkScanInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the NodeNodeScansNetworkScan.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *NodeNodeScansNetworkScan) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	nodeNodeScansNetworkScanUpdateCacheMut.RLock()
	cache, cached := nodeNodeScansNetworkScanUpdateCache[key]
	nodeNodeScansNetworkScanUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			nodeNodeScansNetworkScanAllColumns,
			nodeNodeScansNetworkScanPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update node_node_scans_network_scans, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"node_node_scans_network_scans\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, nodeNodeScansNetworkScanPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(nodeNodeScansNetworkScanType, nodeNodeScansNetworkScanMapping, append(wl, nodeNodeScansNetworkScanPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update node_node_scans_network_scans row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for node_node_scans_network_scans")
	}

	if !cached {
		nodeNodeScansNetworkScanUpdateCacheMut.Lock()
		nodeNodeScansNetworkScanUpdateCache[key] = cache
		nodeNodeScansNetworkScanUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q nodeNodeScansNetworkScanQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for node_node_scans_network_scans")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for node_node_scans_network_scans")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o NodeNodeScansNetworkScanSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), nodeNodeScansNetworkScanPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"node_node_scans_network_scans\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, nodeNodeScansNetworkScanPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in nodeNodeScansNetworkScan slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all nodeNodeScansNetworkScan")
	}
	return rowsAff, nil
}

// Delete deletes a single NodeNodeScansNetworkScan record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *NodeNodeScansNetworkScan) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no NodeNodeScansNetworkScan provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), nodeNodeScansNetworkScanPrimaryKeyMapping)
	sql := "DELETE FROM \"node_node_scans_network_scans\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from node_node_scans_network_scans")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for node_node_scans_network_scans")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q nodeNodeScansNetworkScanQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no nodeNodeScansNetworkScanQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from node_node_scans_network_scans")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for node_node_scans_network_scans")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o NodeNodeScansNetworkScanSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(nodeNodeScansNetworkScanBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), nodeNodeScansNetworkScanPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"node_node_scans_network_scans\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, nodeNodeScansNetworkScanPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from nodeNodeScansNetworkScan slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for node_node_scans_network_scans")
	}

	if len(nodeNodeScansNetworkScanAfterDeleteHooks) != 0 {
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
func (o *NodeNodeScansNetworkScan) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindNodeNodeScansNetworkScan(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *NodeNodeScansNetworkScanSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := NodeNodeScansNetworkScanSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), nodeNodeScansNetworkScanPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"node_node_scans_network_scans\".* FROM \"node_node_scans_network_scans\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, nodeNodeScansNetworkScanPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in NodeNodeScansNetworkScanSlice")
	}

	*o = slice

	return nil
}

// NodeNodeScansNetworkScanExists checks if the NodeNodeScansNetworkScan row exists.
func NodeNodeScansNetworkScanExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"node_node_scans_network_scans\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if node_node_scans_network_scans exists")
	}

	return exists, nil
}
