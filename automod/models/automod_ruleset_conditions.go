// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// AutomodRulesetCondition is an object representing the database table.
type AutomodRulesetCondition struct {
	ID        int64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	GuildID   int64      `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	RulesetID int64      `boil:"ruleset_id" json:"ruleset_id" toml:"ruleset_id" yaml:"ruleset_id"`
	Kind      int        `boil:"kind" json:"kind" toml:"kind" yaml:"kind"`
	TypeID    int        `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	Settings  types.JSON `boil:"settings" json:"settings" toml:"settings" yaml:"settings"`

	R *automodRulesetConditionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L automodRulesetConditionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AutomodRulesetConditionColumns = struct {
	ID        string
	GuildID   string
	RulesetID string
	Kind      string
	TypeID    string
	Settings  string
}{
	ID:        "id",
	GuildID:   "guild_id",
	RulesetID: "ruleset_id",
	Kind:      "kind",
	TypeID:    "type_id",
	Settings:  "settings",
}

var AutomodRulesetConditionTableColumns = struct {
	ID        string
	GuildID   string
	RulesetID string
	Kind      string
	TypeID    string
	Settings  string
}{
	ID:        "automod_ruleset_conditions.id",
	GuildID:   "automod_ruleset_conditions.guild_id",
	RulesetID: "automod_ruleset_conditions.ruleset_id",
	Kind:      "automod_ruleset_conditions.kind",
	TypeID:    "automod_ruleset_conditions.type_id",
	Settings:  "automod_ruleset_conditions.settings",
}

// Generated where

var AutomodRulesetConditionWhere = struct {
	ID        whereHelperint64
	GuildID   whereHelperint64
	RulesetID whereHelperint64
	Kind      whereHelperint
	TypeID    whereHelperint
	Settings  whereHelpertypes_JSON
}{
	ID:        whereHelperint64{field: "\"automod_ruleset_conditions\".\"id\""},
	GuildID:   whereHelperint64{field: "\"automod_ruleset_conditions\".\"guild_id\""},
	RulesetID: whereHelperint64{field: "\"automod_ruleset_conditions\".\"ruleset_id\""},
	Kind:      whereHelperint{field: "\"automod_ruleset_conditions\".\"kind\""},
	TypeID:    whereHelperint{field: "\"automod_ruleset_conditions\".\"type_id\""},
	Settings:  whereHelpertypes_JSON{field: "\"automod_ruleset_conditions\".\"settings\""},
}

// AutomodRulesetConditionRels is where relationship names are stored.
var AutomodRulesetConditionRels = struct {
	Ruleset string
}{
	Ruleset: "Ruleset",
}

// automodRulesetConditionR is where relationships are stored.
type automodRulesetConditionR struct {
	Ruleset *AutomodRuleset `boil:"Ruleset" json:"Ruleset" toml:"Ruleset" yaml:"Ruleset"`
}

// NewStruct creates a new relationship struct
func (*automodRulesetConditionR) NewStruct() *automodRulesetConditionR {
	return &automodRulesetConditionR{}
}

func (r *automodRulesetConditionR) GetRuleset() *AutomodRuleset {
	if r == nil {
		return nil
	}
	return r.Ruleset
}

// automodRulesetConditionL is where Load methods for each relationship are stored.
type automodRulesetConditionL struct{}

var (
	automodRulesetConditionAllColumns            = []string{"id", "guild_id", "ruleset_id", "kind", "type_id", "settings"}
	automodRulesetConditionColumnsWithoutDefault = []string{"guild_id", "ruleset_id", "kind", "type_id", "settings"}
	automodRulesetConditionColumnsWithDefault    = []string{"id"}
	automodRulesetConditionPrimaryKeyColumns     = []string{"id"}
	automodRulesetConditionGeneratedColumns      = []string{}
)

type (
	// AutomodRulesetConditionSlice is an alias for a slice of pointers to AutomodRulesetCondition.
	// This should almost always be used instead of []AutomodRulesetCondition.
	AutomodRulesetConditionSlice []*AutomodRulesetCondition

	automodRulesetConditionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	automodRulesetConditionType                 = reflect.TypeOf(&AutomodRulesetCondition{})
	automodRulesetConditionMapping              = queries.MakeStructMapping(automodRulesetConditionType)
	automodRulesetConditionPrimaryKeyMapping, _ = queries.BindMapping(automodRulesetConditionType, automodRulesetConditionMapping, automodRulesetConditionPrimaryKeyColumns)
	automodRulesetConditionInsertCacheMut       sync.RWMutex
	automodRulesetConditionInsertCache          = make(map[string]insertCache)
	automodRulesetConditionUpdateCacheMut       sync.RWMutex
	automodRulesetConditionUpdateCache          = make(map[string]updateCache)
	automodRulesetConditionUpsertCacheMut       sync.RWMutex
	automodRulesetConditionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single automodRulesetCondition record from the query using the global executor.
func (q automodRulesetConditionQuery) OneG(ctx context.Context) (*AutomodRulesetCondition, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single automodRulesetCondition record from the query.
func (q automodRulesetConditionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AutomodRulesetCondition, error) {
	o := &AutomodRulesetCondition{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for automod_ruleset_conditions")
	}

	return o, nil
}

// AllG returns all AutomodRulesetCondition records from the query using the global executor.
func (q automodRulesetConditionQuery) AllG(ctx context.Context) (AutomodRulesetConditionSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all AutomodRulesetCondition records from the query.
func (q automodRulesetConditionQuery) All(ctx context.Context, exec boil.ContextExecutor) (AutomodRulesetConditionSlice, error) {
	var o []*AutomodRulesetCondition

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AutomodRulesetCondition slice")
	}

	return o, nil
}

// CountG returns the count of all AutomodRulesetCondition records in the query using the global executor
func (q automodRulesetConditionQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all AutomodRulesetCondition records in the query.
func (q automodRulesetConditionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count automod_ruleset_conditions rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q automodRulesetConditionQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q automodRulesetConditionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if automod_ruleset_conditions exists")
	}

	return count > 0, nil
}

// Ruleset pointed to by the foreign key.
func (o *AutomodRulesetCondition) Ruleset(mods ...qm.QueryMod) automodRulesetQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.RulesetID),
	}

	queryMods = append(queryMods, mods...)

	return AutomodRulesets(queryMods...)
}

// LoadRuleset allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (automodRulesetConditionL) LoadRuleset(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAutomodRulesetCondition interface{}, mods queries.Applicator) error {
	var slice []*AutomodRulesetCondition
	var object *AutomodRulesetCondition

	if singular {
		var ok bool
		object, ok = maybeAutomodRulesetCondition.(*AutomodRulesetCondition)
		if !ok {
			object = new(AutomodRulesetCondition)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeAutomodRulesetCondition)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeAutomodRulesetCondition))
			}
		}
	} else {
		s, ok := maybeAutomodRulesetCondition.(*[]*AutomodRulesetCondition)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeAutomodRulesetCondition)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeAutomodRulesetCondition))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &automodRulesetConditionR{}
		}
		args[object.RulesetID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &automodRulesetConditionR{}
			}

			args[obj.RulesetID] = struct{}{}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`automod_rulesets`),
		qm.WhereIn(`automod_rulesets.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load AutomodRuleset")
	}

	var resultSlice []*AutomodRuleset
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice AutomodRuleset")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for automod_rulesets")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for automod_rulesets")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Ruleset = foreign
		if foreign.R == nil {
			foreign.R = &automodRulesetR{}
		}
		foreign.R.RulesetAutomodRulesetConditions = append(foreign.R.RulesetAutomodRulesetConditions, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.RulesetID == foreign.ID {
				local.R.Ruleset = foreign
				if foreign.R == nil {
					foreign.R = &automodRulesetR{}
				}
				foreign.R.RulesetAutomodRulesetConditions = append(foreign.R.RulesetAutomodRulesetConditions, local)
				break
			}
		}
	}

	return nil
}

// SetRulesetG of the automodRulesetCondition to the related item.
// Sets o.R.Ruleset to related.
// Adds o to related.R.RulesetAutomodRulesetConditions.
// Uses the global database handle.
func (o *AutomodRulesetCondition) SetRulesetG(ctx context.Context, insert bool, related *AutomodRuleset) error {
	return o.SetRuleset(ctx, boil.GetContextDB(), insert, related)
}

// SetRuleset of the automodRulesetCondition to the related item.
// Sets o.R.Ruleset to related.
// Adds o to related.R.RulesetAutomodRulesetConditions.
func (o *AutomodRulesetCondition) SetRuleset(ctx context.Context, exec boil.ContextExecutor, insert bool, related *AutomodRuleset) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"automod_ruleset_conditions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"ruleset_id"}),
		strmangle.WhereClause("\"", "\"", 2, automodRulesetConditionPrimaryKeyColumns),
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

	o.RulesetID = related.ID
	if o.R == nil {
		o.R = &automodRulesetConditionR{
			Ruleset: related,
		}
	} else {
		o.R.Ruleset = related
	}

	if related.R == nil {
		related.R = &automodRulesetR{
			RulesetAutomodRulesetConditions: AutomodRulesetConditionSlice{o},
		}
	} else {
		related.R.RulesetAutomodRulesetConditions = append(related.R.RulesetAutomodRulesetConditions, o)
	}

	return nil
}

// AutomodRulesetConditions retrieves all the records using an executor.
func AutomodRulesetConditions(mods ...qm.QueryMod) automodRulesetConditionQuery {
	mods = append(mods, qm.From("\"automod_ruleset_conditions\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"automod_ruleset_conditions\".*"})
	}

	return automodRulesetConditionQuery{q}
}

// FindAutomodRulesetConditionG retrieves a single record by ID.
func FindAutomodRulesetConditionG(ctx context.Context, iD int64, selectCols ...string) (*AutomodRulesetCondition, error) {
	return FindAutomodRulesetCondition(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindAutomodRulesetCondition retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAutomodRulesetCondition(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*AutomodRulesetCondition, error) {
	automodRulesetConditionObj := &AutomodRulesetCondition{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"automod_ruleset_conditions\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, automodRulesetConditionObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from automod_ruleset_conditions")
	}

	return automodRulesetConditionObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *AutomodRulesetCondition) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AutomodRulesetCondition) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no automod_ruleset_conditions provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(automodRulesetConditionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	automodRulesetConditionInsertCacheMut.RLock()
	cache, cached := automodRulesetConditionInsertCache[key]
	automodRulesetConditionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			automodRulesetConditionAllColumns,
			automodRulesetConditionColumnsWithDefault,
			automodRulesetConditionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(automodRulesetConditionType, automodRulesetConditionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(automodRulesetConditionType, automodRulesetConditionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"automod_ruleset_conditions\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"automod_ruleset_conditions\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into automod_ruleset_conditions")
	}

	if !cached {
		automodRulesetConditionInsertCacheMut.Lock()
		automodRulesetConditionInsertCache[key] = cache
		automodRulesetConditionInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single AutomodRulesetCondition record using the global executor.
// See Update for more documentation.
func (o *AutomodRulesetCondition) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the AutomodRulesetCondition.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AutomodRulesetCondition) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	automodRulesetConditionUpdateCacheMut.RLock()
	cache, cached := automodRulesetConditionUpdateCache[key]
	automodRulesetConditionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			automodRulesetConditionAllColumns,
			automodRulesetConditionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update automod_ruleset_conditions, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"automod_ruleset_conditions\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, automodRulesetConditionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(automodRulesetConditionType, automodRulesetConditionMapping, append(wl, automodRulesetConditionPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update automod_ruleset_conditions row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for automod_ruleset_conditions")
	}

	if !cached {
		automodRulesetConditionUpdateCacheMut.Lock()
		automodRulesetConditionUpdateCache[key] = cache
		automodRulesetConditionUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q automodRulesetConditionQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q automodRulesetConditionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for automod_ruleset_conditions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for automod_ruleset_conditions")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AutomodRulesetConditionSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AutomodRulesetConditionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), automodRulesetConditionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"automod_ruleset_conditions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, automodRulesetConditionPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in automodRulesetCondition slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all automodRulesetCondition")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *AutomodRulesetCondition) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns, opts...)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AutomodRulesetCondition) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no automod_ruleset_conditions provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(automodRulesetConditionColumnsWithDefault, o)

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

	automodRulesetConditionUpsertCacheMut.RLock()
	cache, cached := automodRulesetConditionUpsertCache[key]
	automodRulesetConditionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			automodRulesetConditionAllColumns,
			automodRulesetConditionColumnsWithDefault,
			automodRulesetConditionColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			automodRulesetConditionAllColumns,
			automodRulesetConditionPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert automod_ruleset_conditions, could not build update column list")
		}

		ret := strmangle.SetComplement(automodRulesetConditionAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(automodRulesetConditionPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert automod_ruleset_conditions, could not build conflict column list")
			}

			conflict = make([]string, len(automodRulesetConditionPrimaryKeyColumns))
			copy(conflict, automodRulesetConditionPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"automod_ruleset_conditions\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(automodRulesetConditionType, automodRulesetConditionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(automodRulesetConditionType, automodRulesetConditionMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert automod_ruleset_conditions")
	}

	if !cached {
		automodRulesetConditionUpsertCacheMut.Lock()
		automodRulesetConditionUpsertCache[key] = cache
		automodRulesetConditionUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single AutomodRulesetCondition record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *AutomodRulesetCondition) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single AutomodRulesetCondition record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AutomodRulesetCondition) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AutomodRulesetCondition provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), automodRulesetConditionPrimaryKeyMapping)
	sql := "DELETE FROM \"automod_ruleset_conditions\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from automod_ruleset_conditions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for automod_ruleset_conditions")
	}

	return rowsAff, nil
}

func (q automodRulesetConditionQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q automodRulesetConditionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no automodRulesetConditionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from automod_ruleset_conditions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for automod_ruleset_conditions")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o AutomodRulesetConditionSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AutomodRulesetConditionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), automodRulesetConditionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"automod_ruleset_conditions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, automodRulesetConditionPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from automodRulesetCondition slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for automod_ruleset_conditions")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *AutomodRulesetCondition) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no AutomodRulesetCondition provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AutomodRulesetCondition) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAutomodRulesetCondition(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AutomodRulesetConditionSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty AutomodRulesetConditionSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AutomodRulesetConditionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AutomodRulesetConditionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), automodRulesetConditionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"automod_ruleset_conditions\".* FROM \"automod_ruleset_conditions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, automodRulesetConditionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AutomodRulesetConditionSlice")
	}

	*o = slice

	return nil
}

// AutomodRulesetConditionExistsG checks if the AutomodRulesetCondition row exists.
func AutomodRulesetConditionExistsG(ctx context.Context, iD int64) (bool, error) {
	return AutomodRulesetConditionExists(ctx, boil.GetContextDB(), iD)
}

// AutomodRulesetConditionExists checks if the AutomodRulesetCondition row exists.
func AutomodRulesetConditionExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"automod_ruleset_conditions\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if automod_ruleset_conditions exists")
	}

	return exists, nil
}

// Exists checks if the AutomodRulesetCondition row exists.
func (o *AutomodRulesetCondition) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return AutomodRulesetConditionExists(ctx, exec, o.ID)
}