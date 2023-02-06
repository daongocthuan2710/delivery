// Code generated by SQLBoiler 4.14.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package entity

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
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

// Ward is an object representing the database table.
type Ward struct {
	ID         null.Int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name       null.String `boil:"name" json:"name" toml:"name" yaml:"name"`
	DistrictID null.Int    `boil:"district_id" json:"district_id" toml:"district_id" yaml:"district_id"`
	GHNCode    null.String `boil:"ghn_code" json:"ghn_code,omitempty" toml:"ghn_code" yaml:"ghn_code,omitempty"`

	R *wardR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L wardL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var WardColumns = struct {
	ID         string
	Name       string
	DistrictID string
	GHNCode    string
}{
	ID:         "id",
	Name:       "name",
	DistrictID: "district_id",
	GHNCode:    "ghn_code",
}

var WardTableColumns = struct {
	ID         string
	Name       string
	DistrictID string
	GHNCode    string
}{
	ID:         "wards.id",
	Name:       "wards.name",
	DistrictID: "wards.district_id",
	GHNCode:    "wards.ghn_code",
}

// Generated where

var WardWhere = struct {
	ID         whereHelpernull_Int
	Name       whereHelpernull_String
	DistrictID whereHelpernull_Int
	GHNCode    whereHelpernull_String
}{
	ID:         whereHelpernull_Int{field: "\"wards\".\"id\""},
	Name:       whereHelpernull_String{field: "\"wards\".\"name\""},
	DistrictID: whereHelpernull_Int{field: "\"wards\".\"district_id\""},
	GHNCode:    whereHelpernull_String{field: "\"wards\".\"ghn_code\""},
}

// WardRels is where relationship names are stored.
var WardRels = struct {
	District string
}{
	District: "District",
}

// wardR is where relationships are stored.
type wardR struct {
	District *District `boil:"District" json:"District" toml:"District" yaml:"District"`
}

// NewStruct creates a new relationship struct
func (*wardR) NewStruct() *wardR {
	return &wardR{}
}

func (r *wardR) GetDistrict() *District {
	if r == nil {
		return nil
	}
	return r.District
}

// wardL is where Load methods for each relationship are stored.
type wardL struct{}

var (
	wardAllColumns            = []string{"id", "name", "district_id", "ghn_code"}
	wardColumnsWithoutDefault = []string{"id", "name", "district_id"}
	wardColumnsWithDefault    = []string{"ghn_code"}
	wardPrimaryKeyColumns     = []string{"id"}
	wardGeneratedColumns      = []string{}
)

type (
	// WardSlice is an alias for a slice of pointers to Ward.
	// This should almost always be used instead of []Ward.
	WardSlice []*Ward
	// WardHook is the signature for custom Ward hook methods
	WardHook func(context.Context, boil.ContextExecutor, *Ward) error

	wardQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	wardType                 = reflect.TypeOf(&Ward{})
	wardMapping              = queries.MakeStructMapping(wardType)
	wardPrimaryKeyMapping, _ = queries.BindMapping(wardType, wardMapping, wardPrimaryKeyColumns)
	wardInsertCacheMut       sync.RWMutex
	wardInsertCache          = make(map[string]insertCache)
	wardUpdateCacheMut       sync.RWMutex
	wardUpdateCache          = make(map[string]updateCache)
	wardUpsertCacheMut       sync.RWMutex
	wardUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var wardAfterSelectHooks []WardHook

var wardBeforeInsertHooks []WardHook
var wardAfterInsertHooks []WardHook

var wardBeforeUpdateHooks []WardHook
var wardAfterUpdateHooks []WardHook

var wardBeforeDeleteHooks []WardHook
var wardAfterDeleteHooks []WardHook

var wardBeforeUpsertHooks []WardHook
var wardAfterUpsertHooks []WardHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Ward) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range wardAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Ward) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range wardBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Ward) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range wardAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Ward) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range wardBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Ward) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range wardAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Ward) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range wardBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Ward) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range wardAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Ward) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range wardBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Ward) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range wardAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddWardHook registers your hook function for all future operations.
func AddWardHook(hookPoint boil.HookPoint, wardHook WardHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		wardAfterSelectHooks = append(wardAfterSelectHooks, wardHook)
	case boil.BeforeInsertHook:
		wardBeforeInsertHooks = append(wardBeforeInsertHooks, wardHook)
	case boil.AfterInsertHook:
		wardAfterInsertHooks = append(wardAfterInsertHooks, wardHook)
	case boil.BeforeUpdateHook:
		wardBeforeUpdateHooks = append(wardBeforeUpdateHooks, wardHook)
	case boil.AfterUpdateHook:
		wardAfterUpdateHooks = append(wardAfterUpdateHooks, wardHook)
	case boil.BeforeDeleteHook:
		wardBeforeDeleteHooks = append(wardBeforeDeleteHooks, wardHook)
	case boil.AfterDeleteHook:
		wardAfterDeleteHooks = append(wardAfterDeleteHooks, wardHook)
	case boil.BeforeUpsertHook:
		wardBeforeUpsertHooks = append(wardBeforeUpsertHooks, wardHook)
	case boil.AfterUpsertHook:
		wardAfterUpsertHooks = append(wardAfterUpsertHooks, wardHook)
	}
}

// One returns a single ward record from the query.
func (q wardQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Ward, error) {
	o := &Ward{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: failed to execute a one query for wards")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Ward records from the query.
func (q wardQuery) All(ctx context.Context, exec boil.ContextExecutor) (WardSlice, error) {
	var o []*Ward

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "entity: failed to assign all query results to Ward slice")
	}

	if len(wardAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Ward records in the query.
func (q wardQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to count wards rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q wardQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "entity: failed to check if wards exists")
	}

	return count > 0, nil
}

// District pointed to by the foreign key.
func (o *Ward) District(mods ...qm.QueryMod) districtQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.DistrictID),
	}

	queryMods = append(queryMods, mods...)

	return Districts(queryMods...)
}

// LoadDistrict allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (wardL) LoadDistrict(ctx context.Context, e boil.ContextExecutor, singular bool, maybeWard interface{}, mods queries.Applicator) error {
	var slice []*Ward
	var object *Ward

	if singular {
		var ok bool
		object, ok = maybeWard.(*Ward)
		if !ok {
			object = new(Ward)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeWard)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeWard))
			}
		}
	} else {
		s, ok := maybeWard.(*[]*Ward)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeWard)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeWard))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &wardR{}
		}
		if !queries.IsNil(object.DistrictID) {
			args = append(args, object.DistrictID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &wardR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.DistrictID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.DistrictID) {
				args = append(args, obj.DistrictID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`districts`),
		qm.WhereIn(`districts.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load District")
	}

	var resultSlice []*District
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice District")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for districts")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for districts")
	}

	if len(districtAfterSelectHooks) != 0 {
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
		object.R.District = foreign
		if foreign.R == nil {
			foreign.R = &districtR{}
		}
		foreign.R.Wards = append(foreign.R.Wards, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.DistrictID, foreign.ID) {
				local.R.District = foreign
				if foreign.R == nil {
					foreign.R = &districtR{}
				}
				foreign.R.Wards = append(foreign.R.Wards, local)
				break
			}
		}
	}

	return nil
}

// SetDistrict of the ward to the related item.
// Sets o.R.District to related.
// Adds o to related.R.Wards.
func (o *Ward) SetDistrict(ctx context.Context, exec boil.ContextExecutor, insert bool, related *District) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"wards\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"district_id"}),
		strmangle.WhereClause("\"", "\"", 2, wardPrimaryKeyColumns),
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

	queries.Assign(&o.DistrictID, related.ID)
	if o.R == nil {
		o.R = &wardR{
			District: related,
		}
	} else {
		o.R.District = related
	}

	if related.R == nil {
		related.R = &districtR{
			Wards: WardSlice{o},
		}
	} else {
		related.R.Wards = append(related.R.Wards, o)
	}

	return nil
}

// Wards retrieves all the records using an executor.
func Wards(mods ...qm.QueryMod) wardQuery {
	mods = append(mods, qm.From("\"wards\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"wards\".*"})
	}

	return wardQuery{q}
}

// FindWard retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindWard(ctx context.Context, exec boil.ContextExecutor, iD null.Int, selectCols ...string) (*Ward, error) {
	wardObj := &Ward{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"wards\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, wardObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: unable to select from wards")
	}

	if err = wardObj.doAfterSelectHooks(ctx, exec); err != nil {
		return wardObj, err
	}

	return wardObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Ward) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("entity: no wards provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(wardColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	wardInsertCacheMut.RLock()
	cache, cached := wardInsertCache[key]
	wardInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			wardAllColumns,
			wardColumnsWithDefault,
			wardColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(wardType, wardMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(wardType, wardMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"wards\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"wards\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
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

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "entity: unable to insert into wards")
	}

	if !cached {
		wardInsertCacheMut.Lock()
		wardInsertCache[key] = cache
		wardInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Ward.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Ward) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	wardUpdateCacheMut.RLock()
	cache, cached := wardUpdateCache[key]
	wardUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			wardAllColumns,
			wardPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("entity: unable to update wards, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"wards\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, wardPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(wardType, wardMapping, append(wl, wardPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "entity: unable to update wards row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by update for wards")
	}

	if !cached {
		wardUpdateCacheMut.Lock()
		wardUpdateCache[key] = cache
		wardUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q wardQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all for wards")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected for wards")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o WardSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("entity: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), wardPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"wards\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, wardPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all in ward slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected all in update all ward")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Ward) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("entity: no wards provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(wardColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	wardUpsertCacheMut.RLock()
	cache, cached := wardUpsertCache[key]
	wardUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			wardAllColumns,
			wardColumnsWithDefault,
			wardColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			wardAllColumns,
			wardPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("entity: unable to upsert wards, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(wardPrimaryKeyColumns))
			copy(conflict, wardPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"wards\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(wardType, wardMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(wardType, wardMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "entity: unable to upsert wards")
	}

	if !cached {
		wardUpsertCacheMut.Lock()
		wardUpsertCache[key] = cache
		wardUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Ward record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Ward) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("entity: no Ward provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), wardPrimaryKeyMapping)
	sql := "DELETE FROM \"wards\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete from wards")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by delete for wards")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q wardQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("entity: no wardQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from wards")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for wards")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o WardSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(wardBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), wardPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"wards\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, wardPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from ward slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for wards")
	}

	if len(wardAfterDeleteHooks) != 0 {
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
func (o *Ward) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindWard(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *WardSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := WardSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), wardPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"wards\".* FROM \"wards\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, wardPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "entity: unable to reload all in WardSlice")
	}

	*o = slice

	return nil
}

// WardExists checks if the Ward row exists.
func WardExists(ctx context.Context, exec boil.ContextExecutor, iD null.Int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"wards\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "entity: unable to check if wards exists")
	}

	return exists, nil
}

// Exists checks if the Ward row exists.
func (o *Ward) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return WardExists(ctx, exec, o.ID)
}