// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testNicknameListings(t *testing.T) {
	t.Parallel()

	query := NicknameListings()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testNicknameListingsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NicknameListing{}
	if err = randomize.Struct(seed, o, nicknameListingDBTypes, true, nicknameListingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NicknameListing struct: %s", err)
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

	count, err := NicknameListings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNicknameListingsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NicknameListing{}
	if err = randomize.Struct(seed, o, nicknameListingDBTypes, true, nicknameListingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NicknameListing struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := NicknameListings().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := NicknameListings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNicknameListingsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NicknameListing{}
	if err = randomize.Struct(seed, o, nicknameListingDBTypes, true, nicknameListingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NicknameListing struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NicknameListingSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := NicknameListings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNicknameListingsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NicknameListing{}
	if err = randomize.Struct(seed, o, nicknameListingDBTypes, true, nicknameListingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NicknameListing struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := NicknameListingExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if NicknameListing exists: %s", err)
	}
	if !e {
		t.Errorf("Expected NicknameListingExists to return true, but got false.")
	}
}

func testNicknameListingsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NicknameListing{}
	if err = randomize.Struct(seed, o, nicknameListingDBTypes, true, nicknameListingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NicknameListing struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	nicknameListingFound, err := FindNicknameListing(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if nicknameListingFound == nil {
		t.Error("want a record, got nil")
	}
}

func testNicknameListingsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NicknameListing{}
	if err = randomize.Struct(seed, o, nicknameListingDBTypes, true, nicknameListingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NicknameListing struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = NicknameListings().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testNicknameListingsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NicknameListing{}
	if err = randomize.Struct(seed, o, nicknameListingDBTypes, true, nicknameListingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NicknameListing struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := NicknameListings().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testNicknameListingsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	nicknameListingOne := &NicknameListing{}
	nicknameListingTwo := &NicknameListing{}
	if err = randomize.Struct(seed, nicknameListingOne, nicknameListingDBTypes, false, nicknameListingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NicknameListing struct: %s", err)
	}
	if err = randomize.Struct(seed, nicknameListingTwo, nicknameListingDBTypes, false, nicknameListingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NicknameListing struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = nicknameListingOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = nicknameListingTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := NicknameListings().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testNicknameListingsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	nicknameListingOne := &NicknameListing{}
	nicknameListingTwo := &NicknameListing{}
	if err = randomize.Struct(seed, nicknameListingOne, nicknameListingDBTypes, false, nicknameListingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NicknameListing struct: %s", err)
	}
	if err = randomize.Struct(seed, nicknameListingTwo, nicknameListingDBTypes, false, nicknameListingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NicknameListing struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = nicknameListingOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = nicknameListingTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NicknameListings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testNicknameListingsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NicknameListing{}
	if err = randomize.Struct(seed, o, nicknameListingDBTypes, true, nicknameListingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NicknameListing struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NicknameListings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNicknameListingsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NicknameListing{}
	if err = randomize.Struct(seed, o, nicknameListingDBTypes, true); err != nil {
		t.Errorf("Unable to randomize NicknameListing struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(nicknameListingColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := NicknameListings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNicknameListingsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NicknameListing{}
	if err = randomize.Struct(seed, o, nicknameListingDBTypes, true, nicknameListingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NicknameListing struct: %s", err)
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

func testNicknameListingsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NicknameListing{}
	if err = randomize.Struct(seed, o, nicknameListingDBTypes, true, nicknameListingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NicknameListing struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NicknameListingSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testNicknameListingsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NicknameListing{}
	if err = randomize.Struct(seed, o, nicknameListingDBTypes, true, nicknameListingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NicknameListing struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := NicknameListings().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	nicknameListingDBTypes = map[string]string{`ID`: `integer`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`, `DeletedAt`: `timestamp with time zone`, `UserID`: `bigint`, `GuildID`: `text`, `Nickname`: `text`}
	_                      = bytes.MinRead
)

func testNicknameListingsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(nicknameListingPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(nicknameListingColumns) == len(nicknameListingPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &NicknameListing{}
	if err = randomize.Struct(seed, o, nicknameListingDBTypes, true, nicknameListingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NicknameListing struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NicknameListings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, nicknameListingDBTypes, true, nicknameListingPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize NicknameListing struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testNicknameListingsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(nicknameListingColumns) == len(nicknameListingPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &NicknameListing{}
	if err = randomize.Struct(seed, o, nicknameListingDBTypes, true, nicknameListingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NicknameListing struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NicknameListings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, nicknameListingDBTypes, true, nicknameListingPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize NicknameListing struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(nicknameListingColumns, nicknameListingPrimaryKeyColumns) {
		fields = nicknameListingColumns
	} else {
		fields = strmangle.SetComplement(
			nicknameListingColumns,
			nicknameListingPrimaryKeyColumns,
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

	slice := NicknameListingSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testNicknameListingsUpsert(t *testing.T) {
	t.Parallel()

	if len(nicknameListingColumns) == len(nicknameListingPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := NicknameListing{}
	if err = randomize.Struct(seed, &o, nicknameListingDBTypes, true); err != nil {
		t.Errorf("Unable to randomize NicknameListing struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert NicknameListing: %s", err)
	}

	count, err := NicknameListings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, nicknameListingDBTypes, false, nicknameListingPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize NicknameListing struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert NicknameListing: %s", err)
	}

	count, err = NicknameListings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
