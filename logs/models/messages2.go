// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Messages2 is an object representing the database table.
type Messages2 struct {
	ID             int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	GuildID        int64     `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	CreatedAt      time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt      time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	Deleted        bool      `boil:"deleted" json:"deleted" toml:"deleted" yaml:"deleted"`
	AuthorUsername string    `boil:"author_username" json:"author_username" toml:"author_username" yaml:"author_username"`
	AuthorID       int64     `boil:"author_id" json:"author_id" toml:"author_id" yaml:"author_id"`
	Content        string    `boil:"content" json:"content" toml:"content" yaml:"content"`

	R *messages2R `boil:"-" json:"-" toml:"-" yaml:"-"`
	L messages2L  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var Messages2Columns = struct {
	ID             string
	GuildID        string
	CreatedAt      string
	UpdatedAt      string
	Deleted        string
	AuthorUsername string
	AuthorID       string
	Content        string
}{
	ID:             "id",
	GuildID:        "guild_id",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	Deleted:        "deleted",
	AuthorUsername: "author_username",
	AuthorID:       "author_id",
	Content:        "content",
}

// Generated where

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var Messages2Where = struct {
	ID             whereHelperint64
	GuildID        whereHelperint64
	CreatedAt      whereHelpertime_Time
	UpdatedAt      whereHelpertime_Time
	Deleted        whereHelperbool
	AuthorUsername whereHelperstring
	AuthorID       whereHelperint64
	Content        whereHelperstring
}{
	ID:             whereHelperint64{field: "\"messages2\".\"id\""},
	GuildID:        whereHelperint64{field: "\"messages2\".\"guild_id\""},
	CreatedAt:      whereHelpertime_Time{field: "\"messages2\".\"created_at\""},
	UpdatedAt:      whereHelpertime_Time{field: "\"messages2\".\"updated_at\""},
	Deleted:        whereHelperbool{field: "\"messages2\".\"deleted\""},
	AuthorUsername: whereHelperstring{field: "\"messages2\".\"author_username\""},
	AuthorID:       whereHelperint64{field: "\"messages2\".\"author_id\""},
	Content:        whereHelperstring{field: "\"messages2\".\"content\""},
}

// Messages2Rels is where relationship names are stored.
var Messages2Rels = struct {
}{}

// messages2R is where relationships are stored.
type messages2R struct {
}

// NewStruct creates a new relationship struct
func (*messages2R) NewStruct() *messages2R {
	return &messages2R{}
}

// messages2L is where Load methods for each relationship are stored.
type messages2L struct{}

var (
	messages2AllColumns            = []string{"id", "guild_id", "created_at", "updated_at", "deleted", "author_username", "author_id", "content"}
	messages2ColumnsWithoutDefault = []string{"id", "guild_id", "created_at", "updated_at", "deleted", "author_username", "author_id", "content"}
	messages2ColumnsWithDefault    = []string{}
	messages2PrimaryKeyColumns     = []string{"id"}
)

type (
	// Messages2Slice is an alias for a slice of pointers to Messages2.
	// This should generally be used opposed to []Messages2.
	Messages2Slice []*Messages2

	messages2Query struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	messages2Type                 = reflect.TypeOf(&Messages2{})
	messages2Mapping              = queries.MakeStructMapping(messages2Type)
	messages2PrimaryKeyMapping, _ = queries.BindMapping(messages2Type, messages2Mapping, messages2PrimaryKeyColumns)
	messages2InsertCacheMut       sync.RWMutex
	messages2InsertCache          = make(map[string]insertCache)
	messages2UpdateCacheMut       sync.RWMutex
	messages2UpdateCache          = make(map[string]updateCache)
	messages2UpsertCacheMut       sync.RWMutex
	messages2UpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single messages2 record from the query using the global executor.
func (q messages2Query) OneG(ctx context.Context) (*Messages2, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single messages2 record from the query.
func (q messages2Query) One(ctx context.Context, exec boil.ContextExecutor) (*Messages2, error) {
	o := &Messages2{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for messages2")
	}

	return o, nil
}

// AllG returns all Messages2 records from the query using the global executor.
func (q messages2Query) AllG(ctx context.Context) (Messages2Slice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Messages2 records from the query.
func (q messages2Query) All(ctx context.Context, exec boil.ContextExecutor) (Messages2Slice, error) {
	var o []*Messages2

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Messages2 slice")
	}

	return o, nil
}

// CountG returns the count of all Messages2 records in the query, and panics on error.
func (q messages2Query) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Messages2 records in the query.
func (q messages2Query) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count messages2 rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q messages2Query) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q messages2Query) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if messages2 exists")
	}

	return count > 0, nil
}

// Messages2s retrieves all the records using an executor.
func Messages2s(mods ...qm.QueryMod) messages2Query {
	mods = append(mods, qm.From("\"messages2\""))
	return messages2Query{NewQuery(mods...)}
}

// FindMessages2G retrieves a single record by ID.
func FindMessages2G(ctx context.Context, iD int64, selectCols ...string) (*Messages2, error) {
	return FindMessages2(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindMessages2 retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMessages2(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Messages2, error) {
	messages2Obj := &Messages2{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"messages2\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, messages2Obj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from messages2")
	}

	return messages2Obj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Messages2) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Messages2) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no messages2 provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(messages2ColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	messages2InsertCacheMut.RLock()
	cache, cached := messages2InsertCache[key]
	messages2InsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			messages2AllColumns,
			messages2ColumnsWithDefault,
			messages2ColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(messages2Type, messages2Mapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(messages2Type, messages2Mapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"messages2\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"messages2\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into messages2")
	}

	if !cached {
		messages2InsertCacheMut.Lock()
		messages2InsertCache[key] = cache
		messages2InsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single Messages2 record using the global executor.
// See Update for more documentation.
func (o *Messages2) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Messages2.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Messages2) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	messages2UpdateCacheMut.RLock()
	cache, cached := messages2UpdateCache[key]
	messages2UpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			messages2AllColumns,
			messages2PrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update messages2, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"messages2\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, messages2PrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(messages2Type, messages2Mapping, append(wl, messages2PrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update messages2 row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for messages2")
	}

	if !cached {
		messages2UpdateCacheMut.Lock()
		messages2UpdateCache[key] = cache
		messages2UpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q messages2Query) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q messages2Query) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for messages2")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for messages2")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o Messages2Slice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o Messages2Slice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), messages2PrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"messages2\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, messages2PrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in messages2 slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all messages2")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Messages2) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Messages2) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no messages2 provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(messages2ColumnsWithDefault, o)

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

	messages2UpsertCacheMut.RLock()
	cache, cached := messages2UpsertCache[key]
	messages2UpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			messages2AllColumns,
			messages2ColumnsWithDefault,
			messages2ColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			messages2AllColumns,
			messages2PrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert messages2, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(messages2PrimaryKeyColumns))
			copy(conflict, messages2PrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"messages2\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(messages2Type, messages2Mapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(messages2Type, messages2Mapping, ret)
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
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert messages2")
	}

	if !cached {
		messages2UpsertCacheMut.Lock()
		messages2UpsertCache[key] = cache
		messages2UpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single Messages2 record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Messages2) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Messages2 record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Messages2) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Messages2 provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), messages2PrimaryKeyMapping)
	sql := "DELETE FROM \"messages2\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from messages2")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for messages2")
	}

	return rowsAff, nil
}

func (q messages2Query) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q messages2Query) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no messages2Query provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from messages2")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for messages2")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o Messages2Slice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o Messages2Slice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), messages2PrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"messages2\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, messages2PrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from messages2 slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for messages2")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Messages2) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no Messages2 provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Messages2) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMessages2(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *Messages2Slice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty Messages2Slice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *Messages2Slice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := Messages2Slice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), messages2PrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"messages2\".* FROM \"messages2\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, messages2PrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in Messages2Slice")
	}

	*o = slice

	return nil
}

// Messages2ExistsG checks if the Messages2 row exists.
func Messages2ExistsG(ctx context.Context, iD int64) (bool, error) {
	return Messages2Exists(ctx, boil.GetContextDB(), iD)
}

// Messages2Exists checks if the Messages2 row exists.
func Messages2Exists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"messages2\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if messages2 exists")
	}

	return exists, nil
}
