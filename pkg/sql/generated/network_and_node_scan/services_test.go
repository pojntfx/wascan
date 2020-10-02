// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testServices(t *testing.T) {
	t.Parallel()

	query := Services()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testServicesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Service{}
	if err = randomize.Struct(seed, o, serviceDBTypes, true, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Services().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServicesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Service{}
	if err = randomize.Struct(seed, o, serviceDBTypes, true, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Services().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Services().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServicesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Service{}
	if err = randomize.Struct(seed, o, serviceDBTypes, true, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Services().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServicesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Service{}
	if err = randomize.Struct(seed, o, serviceDBTypes, true, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ServiceExists(ctx, tx, o.PortNumber)
	if err != nil {
		t.Errorf("Unable to check if Service exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ServiceExists to return true, but got false.")
	}
}

func testServicesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Service{}
	if err = randomize.Struct(seed, o, serviceDBTypes, true, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	serviceFound, err := FindService(ctx, tx, o.PortNumber)
	if err != nil {
		t.Error(err)
	}

	if serviceFound == nil {
		t.Error("want a record, got nil")
	}
}

func testServicesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Service{}
	if err = randomize.Struct(seed, o, serviceDBTypes, true, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Services().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testServicesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Service{}
	if err = randomize.Struct(seed, o, serviceDBTypes, true, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Services().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testServicesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	serviceOne := &Service{}
	serviceTwo := &Service{}
	if err = randomize.Struct(seed, serviceOne, serviceDBTypes, false, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceTwo, serviceDBTypes, false, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Services().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testServicesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	serviceOne := &Service{}
	serviceTwo := &Service{}
	if err = randomize.Struct(seed, serviceOne, serviceDBTypes, false, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceTwo, serviceDBTypes, false, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Services().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func serviceBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Service) error {
	*o = Service{}
	return nil
}

func serviceAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Service) error {
	*o = Service{}
	return nil
}

func serviceAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Service) error {
	*o = Service{}
	return nil
}

func serviceBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Service) error {
	*o = Service{}
	return nil
}

func serviceAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Service) error {
	*o = Service{}
	return nil
}

func serviceBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Service) error {
	*o = Service{}
	return nil
}

func serviceAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Service) error {
	*o = Service{}
	return nil
}

func serviceBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Service) error {
	*o = Service{}
	return nil
}

func serviceAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Service) error {
	*o = Service{}
	return nil
}

func testServicesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Service{}
	o := &Service{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, serviceDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Service object: %s", err)
	}

	AddServiceHook(boil.BeforeInsertHook, serviceBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	serviceBeforeInsertHooks = []ServiceHook{}

	AddServiceHook(boil.AfterInsertHook, serviceAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	serviceAfterInsertHooks = []ServiceHook{}

	AddServiceHook(boil.AfterSelectHook, serviceAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	serviceAfterSelectHooks = []ServiceHook{}

	AddServiceHook(boil.BeforeUpdateHook, serviceBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	serviceBeforeUpdateHooks = []ServiceHook{}

	AddServiceHook(boil.AfterUpdateHook, serviceAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	serviceAfterUpdateHooks = []ServiceHook{}

	AddServiceHook(boil.BeforeDeleteHook, serviceBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	serviceBeforeDeleteHooks = []ServiceHook{}

	AddServiceHook(boil.AfterDeleteHook, serviceAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	serviceAfterDeleteHooks = []ServiceHook{}

	AddServiceHook(boil.BeforeUpsertHook, serviceBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	serviceBeforeUpsertHooks = []ServiceHook{}

	AddServiceHook(boil.AfterUpsertHook, serviceAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	serviceAfterUpsertHooks = []ServiceHook{}
}

func testServicesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Service{}
	if err = randomize.Struct(seed, o, serviceDBTypes, true, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Services().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServicesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Service{}
	if err = randomize.Struct(seed, o, serviceDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(serviceColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Services().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServicesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Service{}
	if err = randomize.Struct(seed, o, serviceDBTypes, true, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testServicesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Service{}
	if err = randomize.Struct(seed, o, serviceDBTypes, true, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testServicesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Service{}
	if err = randomize.Struct(seed, o, serviceDBTypes, true, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Services().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	serviceDBTypes = map[string]string{`ServiceName`: `TEXT`, `PortNumber`: `INTEGER`, `TransportProtocol`: `TEXT`, `Description`: `TEXT`, `Assignee`: `TEXT`, `Contact`: `TEXT`, `RegistrationDate`: `TEXT`, `ModificationDate`: `TEXT`, `Reference`: `TEXT`, `ServiceCode`: `TEXT`, `UnauthorizedUseReported`: `TEXT`, `AssignmentNotes`: `TEXT`}
	_              = bytes.MinRead
)

func testServicesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(servicePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(serviceAllColumns) == len(servicePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Service{}
	if err = randomize.Struct(seed, o, serviceDBTypes, true, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Services().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceDBTypes, true, servicePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testServicesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(serviceAllColumns) == len(servicePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Service{}
	if err = randomize.Struct(seed, o, serviceDBTypes, true, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Services().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceDBTypes, true, servicePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(serviceAllColumns, servicePrimaryKeyColumns) {
		fields = serviceAllColumns
	} else {
		fields = strmangle.SetComplement(
			serviceAllColumns,
			servicePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := ServiceSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}