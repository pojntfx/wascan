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

// Service is an object representing the database table.
type Service struct {
	ID                int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt         time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	PortNumber        int64     `boil:"port_number" json:"port_number" toml:"port_number" yaml:"port_number"`
	TransportProtocol string    `boil:"transport_protocol" json:"transport_protocol" toml:"transport_protocol" yaml:"transport_protocol"`
	NodeScanID        int64     `boil:"node_scan_id" json:"node_scan_id" toml:"node_scan_id" yaml:"node_scan_id"`

	R *serviceR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L serviceL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ServiceColumns = struct {
	ID                string
	CreatedAt         string
	PortNumber        string
	TransportProtocol string
	NodeScanID        string
}{
	ID:                "id",
	CreatedAt:         "created_at",
	PortNumber:        "port_number",
	TransportProtocol: "transport_protocol",
	NodeScanID:        "node_scan_id",
}

// Generated where

var ServiceWhere = struct {
	ID                whereHelperint64
	CreatedAt         whereHelpertime_Time
	PortNumber        whereHelperint64
	TransportProtocol whereHelperstring
	NodeScanID        whereHelperint64
}{
	ID:                whereHelperint64{field: "\"services\".\"id\""},
	CreatedAt:         whereHelpertime_Time{field: "\"services\".\"created_at\""},
	PortNumber:        whereHelperint64{field: "\"services\".\"port_number\""},
	TransportProtocol: whereHelperstring{field: "\"services\".\"transport_protocol\""},
	NodeScanID:        whereHelperint64{field: "\"services\".\"node_scan_id\""},
}

// ServiceRels is where relationship names are stored.
var ServiceRels = struct {
	NodeScan string
}{
	NodeScan: "NodeScan",
}

// serviceR is where relationships are stored.
type serviceR struct {
	NodeScan *NodeScan
}

// NewStruct creates a new relationship struct
func (*serviceR) NewStruct() *serviceR {
	return &serviceR{}
}

// serviceL is where Load methods for each relationship are stored.
type serviceL struct{}

var (
	serviceAllColumns            = []string{"id", "created_at", "port_number", "transport_protocol", "node_scan_id"}
	serviceColumnsWithoutDefault = []string{"created_at", "port_number", "transport_protocol", "node_scan_id"}
	serviceColumnsWithDefault    = []string{"id"}
	servicePrimaryKeyColumns     = []string{"id"}
)

type (
	// ServiceSlice is an alias for a slice of pointers to Service.
	// This should generally be used opposed to []Service.
	ServiceSlice []*Service
	// ServiceHook is the signature for custom Service hook methods
	ServiceHook func(context.Context, boil.ContextExecutor, *Service) error

	serviceQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	serviceType                 = reflect.TypeOf(&Service{})
	serviceMapping              = queries.MakeStructMapping(serviceType)
	servicePrimaryKeyMapping, _ = queries.BindMapping(serviceType, serviceMapping, servicePrimaryKeyColumns)
	serviceInsertCacheMut       sync.RWMutex
	serviceInsertCache          = make(map[string]insertCache)
	serviceUpdateCacheMut       sync.RWMutex
	serviceUpdateCache          = make(map[string]updateCache)
	serviceUpsertCacheMut       sync.RWMutex
	serviceUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var serviceBeforeInsertHooks []ServiceHook
var serviceBeforeUpdateHooks []ServiceHook
var serviceBeforeDeleteHooks []ServiceHook
var serviceBeforeUpsertHooks []ServiceHook

var serviceAfterInsertHooks []ServiceHook
var serviceAfterSelectHooks []ServiceHook
var serviceAfterUpdateHooks []ServiceHook
var serviceAfterDeleteHooks []ServiceHook
var serviceAfterUpsertHooks []ServiceHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Service) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Service) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Service) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Service) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Service) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Service) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Service) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Service) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Service) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddServiceHook registers your hook function for all future operations.
func AddServiceHook(hookPoint boil.HookPoint, serviceHook ServiceHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		serviceBeforeInsertHooks = append(serviceBeforeInsertHooks, serviceHook)
	case boil.BeforeUpdateHook:
		serviceBeforeUpdateHooks = append(serviceBeforeUpdateHooks, serviceHook)
	case boil.BeforeDeleteHook:
		serviceBeforeDeleteHooks = append(serviceBeforeDeleteHooks, serviceHook)
	case boil.BeforeUpsertHook:
		serviceBeforeUpsertHooks = append(serviceBeforeUpsertHooks, serviceHook)
	case boil.AfterInsertHook:
		serviceAfterInsertHooks = append(serviceAfterInsertHooks, serviceHook)
	case boil.AfterSelectHook:
		serviceAfterSelectHooks = append(serviceAfterSelectHooks, serviceHook)
	case boil.AfterUpdateHook:
		serviceAfterUpdateHooks = append(serviceAfterUpdateHooks, serviceHook)
	case boil.AfterDeleteHook:
		serviceAfterDeleteHooks = append(serviceAfterDeleteHooks, serviceHook)
	case boil.AfterUpsertHook:
		serviceAfterUpsertHooks = append(serviceAfterUpsertHooks, serviceHook)
	}
}

// One returns a single service record from the query.
func (q serviceQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Service, error) {
	o := &Service{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for services")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Service records from the query.
func (q serviceQuery) All(ctx context.Context, exec boil.ContextExecutor) (ServiceSlice, error) {
	var o []*Service

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Service slice")
	}

	if len(serviceAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Service records in the query.
func (q serviceQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count services rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q serviceQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if services exists")
	}

	return count > 0, nil
}

// NodeScan pointed to by the foreign key.
func (o *Service) NodeScan(mods ...qm.QueryMod) nodeScanQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.NodeScanID),
	}

	queryMods = append(queryMods, mods...)

	query := NodeScans(queryMods...)
	queries.SetFrom(query.Query, "\"node_scans\"")

	return query
}

// LoadNodeScan allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (serviceL) LoadNodeScan(ctx context.Context, e boil.ContextExecutor, singular bool, maybeService interface{}, mods queries.Applicator) error {
	var slice []*Service
	var object *Service

	if singular {
		object = maybeService.(*Service)
	} else {
		slice = *maybeService.(*[]*Service)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &serviceR{}
		}
		args = append(args, object.NodeScanID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &serviceR{}
			}

			for _, a := range args {
				if a == obj.NodeScanID {
					continue Outer
				}
			}

			args = append(args, obj.NodeScanID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`node_scans`), qm.WhereIn(`node_scans.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load NodeScan")
	}

	var resultSlice []*NodeScan
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice NodeScan")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for node_scans")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for node_scans")
	}

	if len(serviceAfterSelectHooks) != 0 {
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
		object.R.NodeScan = foreign
		if foreign.R == nil {
			foreign.R = &nodeScanR{}
		}
		foreign.R.Services = append(foreign.R.Services, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.NodeScanID == foreign.ID {
				local.R.NodeScan = foreign
				if foreign.R == nil {
					foreign.R = &nodeScanR{}
				}
				foreign.R.Services = append(foreign.R.Services, local)
				break
			}
		}
	}

	return nil
}

// SetNodeScan of the service to the related item.
// Sets o.R.NodeScan to related.
// Adds o to related.R.Services.
func (o *Service) SetNodeScan(ctx context.Context, exec boil.ContextExecutor, insert bool, related *NodeScan) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"services\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"node_scan_id"}),
		strmangle.WhereClause("\"", "\"", 0, servicePrimaryKeyColumns),
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

	o.NodeScanID = related.ID
	if o.R == nil {
		o.R = &serviceR{
			NodeScan: related,
		}
	} else {
		o.R.NodeScan = related
	}

	if related.R == nil {
		related.R = &nodeScanR{
			Services: ServiceSlice{o},
		}
	} else {
		related.R.Services = append(related.R.Services, o)
	}

	return nil
}

// Services retrieves all the records using an executor.
func Services(mods ...qm.QueryMod) serviceQuery {
	mods = append(mods, qm.From("\"services\""))
	return serviceQuery{NewQuery(mods...)}
}

// FindService retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindService(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Service, error) {
	serviceObj := &Service{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"services\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, serviceObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from services")
	}

	return serviceObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Service) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no services provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(serviceColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	serviceInsertCacheMut.RLock()
	cache, cached := serviceInsertCache[key]
	serviceInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			serviceAllColumns,
			serviceColumnsWithDefault,
			serviceColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(serviceType, serviceMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(serviceType, serviceMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"services\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"services\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"services\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, servicePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into services")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == serviceMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for services")
	}

CacheNoHooks:
	if !cached {
		serviceInsertCacheMut.Lock()
		serviceInsertCache[key] = cache
		serviceInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Service.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Service) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	serviceUpdateCacheMut.RLock()
	cache, cached := serviceUpdateCache[key]
	serviceUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			serviceAllColumns,
			servicePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update services, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"services\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, servicePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(serviceType, serviceMapping, append(wl, servicePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update services row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for services")
	}

	if !cached {
		serviceUpdateCacheMut.Lock()
		serviceUpdateCache[key] = cache
		serviceUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q serviceQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for services")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for services")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ServiceSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), servicePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"services\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, servicePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in service slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all service")
	}
	return rowsAff, nil
}

// Delete deletes a single Service record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Service) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Service provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), servicePrimaryKeyMapping)
	sql := "DELETE FROM \"services\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from services")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for services")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q serviceQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no serviceQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from services")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for services")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ServiceSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(serviceBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), servicePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"services\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, servicePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from service slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for services")
	}

	if len(serviceAfterDeleteHooks) != 0 {
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
func (o *Service) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindService(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ServiceSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ServiceSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), servicePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"services\".* FROM \"services\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, servicePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ServiceSlice")
	}

	*o = slice

	return nil
}

// ServiceExists checks if the Service row exists.
func ServiceExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"services\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if services exists")
	}

	return exists, nil
}
