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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Vendordb is an object representing the database table.
type Vendordb struct {
	Mac          int64       `boil:"mac" json:"mac" toml:"mac" yaml:"mac"`
	Vendor       null.String `boil:"vendor" json:"vendor,omitempty" toml:"vendor" yaml:"vendor,omitempty"`
	Registry     string      `boil:"registry" json:"registry" toml:"registry" yaml:"registry"`
	Organization null.String `boil:"organization" json:"organization,omitempty" toml:"organization" yaml:"organization,omitempty"`
	Address      null.String `boil:"address" json:"address,omitempty" toml:"address" yaml:"address,omitempty"`
	Visibility   int64       `boil:"visibility" json:"visibility" toml:"visibility" yaml:"visibility"`

	R *vendordbR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L vendordbL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var VendordbColumns = struct {
	Mac          string
	Vendor       string
	Registry     string
	Organization string
	Address      string
	Visibility   string
}{
	Mac:          "mac",
	Vendor:       "vendor",
	Registry:     "registry",
	Organization: "organization",
	Address:      "address",
	Visibility:   "visibility",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var VendordbWhere = struct {
	Mac          whereHelperint64
	Vendor       whereHelpernull_String
	Registry     whereHelperstring
	Organization whereHelpernull_String
	Address      whereHelpernull_String
	Visibility   whereHelperint64
}{
	Mac:          whereHelperint64{field: "\"vendordb\".\"mac\""},
	Vendor:       whereHelpernull_String{field: "\"vendordb\".\"vendor\""},
	Registry:     whereHelperstring{field: "\"vendordb\".\"registry\""},
	Organization: whereHelpernull_String{field: "\"vendordb\".\"organization\""},
	Address:      whereHelpernull_String{field: "\"vendordb\".\"address\""},
	Visibility:   whereHelperint64{field: "\"vendordb\".\"visibility\""},
}

// VendordbRels is where relationship names are stored.
var VendordbRels = struct {
}{}

// vendordbR is where relationships are stored.
type vendordbR struct {
}

// NewStruct creates a new relationship struct
func (*vendordbR) NewStruct() *vendordbR {
	return &vendordbR{}
}

// vendordbL is where Load methods for each relationship are stored.
type vendordbL struct{}

var (
	vendordbAllColumns            = []string{"mac", "vendor", "registry", "organization", "address", "visibility"}
	vendordbColumnsWithoutDefault = []string{"vendor", "registry", "organization", "address", "visibility"}
	vendordbColumnsWithDefault    = []string{"mac"}
	vendordbPrimaryKeyColumns     = []string{"mac"}
)

type (
	// VendordbSlice is an alias for a slice of pointers to Vendordb.
	// This should generally be used opposed to []Vendordb.
	VendordbSlice []*Vendordb
	// VendordbHook is the signature for custom Vendordb hook methods
	VendordbHook func(context.Context, boil.ContextExecutor, *Vendordb) error

	vendordbQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	vendordbType                 = reflect.TypeOf(&Vendordb{})
	vendordbMapping              = queries.MakeStructMapping(vendordbType)
	vendordbPrimaryKeyMapping, _ = queries.BindMapping(vendordbType, vendordbMapping, vendordbPrimaryKeyColumns)
	vendordbInsertCacheMut       sync.RWMutex
	vendordbInsertCache          = make(map[string]insertCache)
	vendordbUpdateCacheMut       sync.RWMutex
	vendordbUpdateCache          = make(map[string]updateCache)
	vendordbUpsertCacheMut       sync.RWMutex
	vendordbUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var vendordbBeforeInsertHooks []VendordbHook
var vendordbBeforeUpdateHooks []VendordbHook
var vendordbBeforeDeleteHooks []VendordbHook
var vendordbBeforeUpsertHooks []VendordbHook

var vendordbAfterInsertHooks []VendordbHook
var vendordbAfterSelectHooks []VendordbHook
var vendordbAfterUpdateHooks []VendordbHook
var vendordbAfterDeleteHooks []VendordbHook
var vendordbAfterUpsertHooks []VendordbHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Vendordb) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vendordbBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Vendordb) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vendordbBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Vendordb) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vendordbBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Vendordb) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vendordbBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Vendordb) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vendordbAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Vendordb) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vendordbAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Vendordb) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vendordbAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Vendordb) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vendordbAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Vendordb) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vendordbAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddVendordbHook registers your hook function for all future operations.
func AddVendordbHook(hookPoint boil.HookPoint, vendordbHook VendordbHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		vendordbBeforeInsertHooks = append(vendordbBeforeInsertHooks, vendordbHook)
	case boil.BeforeUpdateHook:
		vendordbBeforeUpdateHooks = append(vendordbBeforeUpdateHooks, vendordbHook)
	case boil.BeforeDeleteHook:
		vendordbBeforeDeleteHooks = append(vendordbBeforeDeleteHooks, vendordbHook)
	case boil.BeforeUpsertHook:
		vendordbBeforeUpsertHooks = append(vendordbBeforeUpsertHooks, vendordbHook)
	case boil.AfterInsertHook:
		vendordbAfterInsertHooks = append(vendordbAfterInsertHooks, vendordbHook)
	case boil.AfterSelectHook:
		vendordbAfterSelectHooks = append(vendordbAfterSelectHooks, vendordbHook)
	case boil.AfterUpdateHook:
		vendordbAfterUpdateHooks = append(vendordbAfterUpdateHooks, vendordbHook)
	case boil.AfterDeleteHook:
		vendordbAfterDeleteHooks = append(vendordbAfterDeleteHooks, vendordbHook)
	case boil.AfterUpsertHook:
		vendordbAfterUpsertHooks = append(vendordbAfterUpsertHooks, vendordbHook)
	}
}

// One returns a single vendordb record from the query.
func (q vendordbQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Vendordb, error) {
	o := &Vendordb{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for vendordb")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Vendordb records from the query.
func (q vendordbQuery) All(ctx context.Context, exec boil.ContextExecutor) (VendordbSlice, error) {
	var o []*Vendordb

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Vendordb slice")
	}

	if len(vendordbAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Vendordb records in the query.
func (q vendordbQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count vendordb rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q vendordbQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if vendordb exists")
	}

	return count > 0, nil
}

// Vendordbs retrieves all the records using an executor.
func Vendordbs(mods ...qm.QueryMod) vendordbQuery {
	mods = append(mods, qm.From("\"vendordb\""))
	return vendordbQuery{NewQuery(mods...)}
}

// FindVendordb retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindVendordb(ctx context.Context, exec boil.ContextExecutor, mac int64, selectCols ...string) (*Vendordb, error) {
	vendordbObj := &Vendordb{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"vendordb\" where \"mac\"=?", sel,
	)

	q := queries.Raw(query, mac)

	err := q.Bind(ctx, exec, vendordbObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from vendordb")
	}

	return vendordbObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Vendordb) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no vendordb provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(vendordbColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	vendordbInsertCacheMut.RLock()
	cache, cached := vendordbInsertCache[key]
	vendordbInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			vendordbAllColumns,
			vendordbColumnsWithDefault,
			vendordbColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(vendordbType, vendordbMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(vendordbType, vendordbMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"vendordb\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"vendordb\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"vendordb\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, vendordbPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into vendordb")
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

	o.Mac = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == vendordbMapping["mac"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.Mac,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for vendordb")
	}

CacheNoHooks:
	if !cached {
		vendordbInsertCacheMut.Lock()
		vendordbInsertCache[key] = cache
		vendordbInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Vendordb.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Vendordb) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	vendordbUpdateCacheMut.RLock()
	cache, cached := vendordbUpdateCache[key]
	vendordbUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			vendordbAllColumns,
			vendordbPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update vendordb, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"vendordb\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, vendordbPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(vendordbType, vendordbMapping, append(wl, vendordbPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update vendordb row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for vendordb")
	}

	if !cached {
		vendordbUpdateCacheMut.Lock()
		vendordbUpdateCache[key] = cache
		vendordbUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q vendordbQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for vendordb")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for vendordb")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o VendordbSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), vendordbPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"vendordb\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, vendordbPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in vendordb slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all vendordb")
	}
	return rowsAff, nil
}

// Delete deletes a single Vendordb record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Vendordb) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Vendordb provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), vendordbPrimaryKeyMapping)
	sql := "DELETE FROM \"vendordb\" WHERE \"mac\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from vendordb")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for vendordb")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q vendordbQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no vendordbQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from vendordb")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for vendordb")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o VendordbSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(vendordbBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), vendordbPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"vendordb\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, vendordbPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from vendordb slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for vendordb")
	}

	if len(vendordbAfterDeleteHooks) != 0 {
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
func (o *Vendordb) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindVendordb(ctx, exec, o.Mac)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *VendordbSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := VendordbSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), vendordbPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"vendordb\".* FROM \"vendordb\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, vendordbPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in VendordbSlice")
	}

	*o = slice

	return nil
}

// VendordbExists checks if the Vendordb row exists.
func VendordbExists(ctx context.Context, exec boil.ContextExecutor, mac int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"vendordb\" where \"mac\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, mac)
	}
	row := exec.QueryRowContext(ctx, sql, mac)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if vendordb exists")
	}

	return exists, nil
}
