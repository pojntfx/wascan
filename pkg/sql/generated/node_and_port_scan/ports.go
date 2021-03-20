// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// Port is an object representing the database table.
type Port struct {
	ID                int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt         time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	PortNumber        int64     `boil:"port_number" json:"port_number" toml:"port_number" yaml:"port_number"`
	TransportProtocol string    `boil:"transport_protocol" json:"transport_protocol" toml:"transport_protocol" yaml:"transport_protocol"`
	PortScanID        int64     `boil:"port_scan_id" json:"port_scan_id" toml:"port_scan_id" yaml:"port_scan_id"`

	R *portR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L portL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PortColumns = struct {
	ID                string
	CreatedAt         string
	PortNumber        string
	TransportProtocol string
	PortScanID        string
}{
	ID:                "id",
	CreatedAt:         "created_at",
	PortNumber:        "port_number",
	TransportProtocol: "transport_protocol",
	PortScanID:        "port_scan_id",
}

// Generated where

var PortWhere = struct {
	ID                whereHelperint64
	CreatedAt         whereHelpertime_Time
	PortNumber        whereHelperint64
	TransportProtocol whereHelperstring
	PortScanID        whereHelperint64
}{
	ID:                whereHelperint64{field: "\"ports\".\"id\""},
	CreatedAt:         whereHelpertime_Time{field: "\"ports\".\"created_at\""},
	PortNumber:        whereHelperint64{field: "\"ports\".\"port_number\""},
	TransportProtocol: whereHelperstring{field: "\"ports\".\"transport_protocol\""},
	PortScanID:        whereHelperint64{field: "\"ports\".\"port_scan_id\""},
}

// PortRels is where relationship names are stored.
var PortRels = struct {
	PortScan string
}{
	PortScan: "PortScan",
}

// portR is where relationships are stored.
type portR struct {
	PortScan *PortScan `boil:"PortScan" json:"PortScan" toml:"PortScan" yaml:"PortScan"`
}

// NewStruct creates a new relationship struct
func (*portR) NewStruct() *portR {
	return &portR{}
}

// portL is where Load methods for each relationship are stored.
type portL struct{}

var (
	portAllColumns            = []string{"id", "created_at", "port_number", "transport_protocol", "port_scan_id"}
	portColumnsWithoutDefault = []string{"created_at", "port_number", "transport_protocol", "port_scan_id"}
	portColumnsWithDefault    = []string{"id"}
	portPrimaryKeyColumns     = []string{"id"}
)

type (
	// PortSlice is an alias for a slice of pointers to Port.
	// This should generally be used opposed to []Port.
	PortSlice []*Port
	// PortHook is the signature for custom Port hook methods
	PortHook func(context.Context, boil.ContextExecutor, *Port) error

	portQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	portType                 = reflect.TypeOf(&Port{})
	portMapping              = queries.MakeStructMapping(portType)
	portPrimaryKeyMapping, _ = queries.BindMapping(portType, portMapping, portPrimaryKeyColumns)
	portInsertCacheMut       sync.RWMutex
	portInsertCache          = make(map[string]insertCache)
	portUpdateCacheMut       sync.RWMutex
	portUpdateCache          = make(map[string]updateCache)
	portUpsertCacheMut       sync.RWMutex
	portUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var portBeforeInsertHooks []PortHook
var portBeforeUpdateHooks []PortHook
var portBeforeDeleteHooks []PortHook
var portBeforeUpsertHooks []PortHook

var portAfterInsertHooks []PortHook
var portAfterSelectHooks []PortHook
var portAfterUpdateHooks []PortHook
var portAfterDeleteHooks []PortHook
var portAfterUpsertHooks []PortHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Port) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range portBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Port) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range portBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Port) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range portBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Port) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range portBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Port) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range portAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Port) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range portAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Port) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range portAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Port) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range portAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Port) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range portAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPortHook registers your hook function for all future operations.
func AddPortHook(hookPoint boil.HookPoint, portHook PortHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		portBeforeInsertHooks = append(portBeforeInsertHooks, portHook)
	case boil.BeforeUpdateHook:
		portBeforeUpdateHooks = append(portBeforeUpdateHooks, portHook)
	case boil.BeforeDeleteHook:
		portBeforeDeleteHooks = append(portBeforeDeleteHooks, portHook)
	case boil.BeforeUpsertHook:
		portBeforeUpsertHooks = append(portBeforeUpsertHooks, portHook)
	case boil.AfterInsertHook:
		portAfterInsertHooks = append(portAfterInsertHooks, portHook)
	case boil.AfterSelectHook:
		portAfterSelectHooks = append(portAfterSelectHooks, portHook)
	case boil.AfterUpdateHook:
		portAfterUpdateHooks = append(portAfterUpdateHooks, portHook)
	case boil.AfterDeleteHook:
		portAfterDeleteHooks = append(portAfterDeleteHooks, portHook)
	case boil.AfterUpsertHook:
		portAfterUpsertHooks = append(portAfterUpsertHooks, portHook)
	}
}

// One returns a single port record from the query.
func (q portQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Port, error) {
	o := &Port{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for ports")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Port records from the query.
func (q portQuery) All(ctx context.Context, exec boil.ContextExecutor) (PortSlice, error) {
	var o []*Port

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Port slice")
	}

	if len(portAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Port records in the query.
func (q portQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count ports rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q portQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if ports exists")
	}

	return count > 0, nil
}

// PortScan pointed to by the foreign key.
func (o *Port) PortScan(mods ...qm.QueryMod) portScanQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.PortScanID),
	}

	queryMods = append(queryMods, mods...)

	query := PortScans(queryMods...)
	queries.SetFrom(query.Query, "\"port_scans\"")

	return query
}

// LoadPortScan allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (portL) LoadPortScan(ctx context.Context, e boil.ContextExecutor, singular bool, maybePort interface{}, mods queries.Applicator) error {
	var slice []*Port
	var object *Port

	if singular {
		object = maybePort.(*Port)
	} else {
		slice = *maybePort.(*[]*Port)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &portR{}
		}
		args = append(args, object.PortScanID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &portR{}
			}

			for _, a := range args {
				if a == obj.PortScanID {
					continue Outer
				}
			}

			args = append(args, obj.PortScanID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`port_scans`),
		qm.WhereIn(`port_scans.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load PortScan")
	}

	var resultSlice []*PortScan
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice PortScan")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for port_scans")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for port_scans")
	}

	if len(portAfterSelectHooks) != 0 {
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
		object.R.PortScan = foreign
		if foreign.R == nil {
			foreign.R = &portScanR{}
		}
		foreign.R.Ports = append(foreign.R.Ports, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PortScanID == foreign.ID {
				local.R.PortScan = foreign
				if foreign.R == nil {
					foreign.R = &portScanR{}
				}
				foreign.R.Ports = append(foreign.R.Ports, local)
				break
			}
		}
	}

	return nil
}

// SetPortScan of the port to the related item.
// Sets o.R.PortScan to related.
// Adds o to related.R.Ports.
func (o *Port) SetPortScan(ctx context.Context, exec boil.ContextExecutor, insert bool, related *PortScan) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"ports\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"port_scan_id"}),
		strmangle.WhereClause("\"", "\"", 0, portPrimaryKeyColumns),
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

	o.PortScanID = related.ID
	if o.R == nil {
		o.R = &portR{
			PortScan: related,
		}
	} else {
		o.R.PortScan = related
	}

	if related.R == nil {
		related.R = &portScanR{
			Ports: PortSlice{o},
		}
	} else {
		related.R.Ports = append(related.R.Ports, o)
	}

	return nil
}

// Ports retrieves all the records using an executor.
func Ports(mods ...qm.QueryMod) portQuery {
	mods = append(mods, qm.From("\"ports\""))
	return portQuery{NewQuery(mods...)}
}

// FindPort retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPort(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Port, error) {
	portObj := &Port{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"ports\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, portObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from ports")
	}

	return portObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Port) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no ports provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(portColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	portInsertCacheMut.RLock()
	cache, cached := portInsertCache[key]
	portInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			portAllColumns,
			portColumnsWithDefault,
			portColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(portType, portMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(portType, portMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"ports\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"ports\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"ports\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, portPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into ports")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == portMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for ports")
	}

CacheNoHooks:
	if !cached {
		portInsertCacheMut.Lock()
		portInsertCache[key] = cache
		portInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Port.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Port) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	portUpdateCacheMut.RLock()
	cache, cached := portUpdateCache[key]
	portUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			portAllColumns,
			portPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update ports, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"ports\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, portPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(portType, portMapping, append(wl, portPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update ports row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for ports")
	}

	if !cached {
		portUpdateCacheMut.Lock()
		portUpdateCache[key] = cache
		portUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q portQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for ports")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for ports")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PortSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), portPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"ports\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, portPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in port slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all port")
	}
	return rowsAff, nil
}

// Delete deletes a single Port record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Port) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Port provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), portPrimaryKeyMapping)
	sql := "DELETE FROM \"ports\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from ports")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for ports")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q portQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no portQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from ports")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for ports")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PortSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(portBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), portPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"ports\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, portPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from port slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for ports")
	}

	if len(portAfterDeleteHooks) != 0 {
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
func (o *Port) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPort(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PortSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PortSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), portPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"ports\".* FROM \"ports\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, portPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PortSlice")
	}

	*o = slice

	return nil
}

// PortExists checks if the Port row exists.
func PortExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"ports\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if ports exists")
	}

	return exists, nil
}
