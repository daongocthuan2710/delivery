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
	"github.com/shopspring/decimal"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Delivery is an object representing the database table.
type Delivery struct {
	ID                  null.String     `boil:"id" json:"id" toml:"id" yaml:"id"`
	Code                null.String     `boil:"code" json:"code" toml:"code" yaml:"code"`
	TrackingCode        null.String     `boil:"tracking_code" json:"tracking_code" toml:"tracking_code" yaml:"tracking_code"`
	Note                null.String     `boil:"note" json:"note" toml:"note" yaml:"note"`
	Status              null.String     `boil:"status" json:"status" toml:"status" yaml:"status"`
	CreatedAt           null.Time       `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt           null.Time       `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	Value               decimal.Decimal `boil:"value" json:"value" toml:"value" yaml:"value"`
	Cod                 decimal.Decimal `boil:"cod" json:"cod" toml:"cod" yaml:"cod"`
	Weight              decimal.Decimal `boil:"weight" json:"weight" toml:"weight" yaml:"weight"`
	ServiceCode         null.String     `boil:"service_code" json:"service_code" toml:"service_code" yaml:"service_code"`
	PartnerStatus       null.String     `boil:"partner_status" json:"partner_status" toml:"partner_status" yaml:"partner_status"`
	PartnerIdentityCode null.String     `boil:"partner_identity_code" json:"partner_identity_code" toml:"partner_identity_code" yaml:"partner_identity_code"`
	FromName            null.String     `boil:"from_name" json:"from_name" toml:"from_name" yaml:"from_name"`
	FromPhone           null.String     `boil:"from_phone" json:"from_phone" toml:"from_phone" yaml:"from_phone"`
	FromAddress         null.String     `boil:"from_address" json:"from_address" toml:"from_address" yaml:"from_address"`
	FromProvinceCode    null.Int        `boil:"from_province_code" json:"from_province_code" toml:"from_province_code" yaml:"from_province_code"`
	FromDistrictCode    null.Int        `boil:"from_district_code" json:"from_district_code" toml:"from_district_code" yaml:"from_district_code"`
	FromWardCode        null.Int        `boil:"from_ward_code" json:"from_ward_code" toml:"from_ward_code" yaml:"from_ward_code"`
	ToName              null.String     `boil:"to_name" json:"to_name" toml:"to_name" yaml:"to_name"`
	ToPhone             null.String     `boil:"to_phone" json:"to_phone" toml:"to_phone" yaml:"to_phone"`
	ToAddress           null.String     `boil:"to_address" json:"to_address" toml:"to_address" yaml:"to_address"`
	ToProvinceCode      null.Int        `boil:"to_province_code" json:"to_province_code" toml:"to_province_code" yaml:"to_province_code"`
	ToDistrictCode      null.Int        `boil:"to_district_code" json:"to_district_code" toml:"to_district_code" yaml:"to_district_code"`
	ToWardCode          null.Int        `boil:"to_ward_code" json:"to_ward_code" toml:"to_ward_code" yaml:"to_ward_code"`

	R *deliveryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L deliveryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DeliveryColumns = struct {
	ID                  string
	Code                string
	TrackingCode        string
	Note                string
	Status              string
	CreatedAt           string
	UpdatedAt           string
	Value               string
	Cod                 string
	Weight              string
	ServiceCode         string
	PartnerStatus       string
	PartnerIdentityCode string
	FromName            string
	FromPhone           string
	FromAddress         string
	FromProvinceCode    string
	FromDistrictCode    string
	FromWardCode        string
	ToName              string
	ToPhone             string
	ToAddress           string
	ToProvinceCode      string
	ToDistrictCode      string
	ToWardCode          string
}{
	ID:                  "id",
	Code:                "code",
	TrackingCode:        "tracking_code",
	Note:                "note",
	Status:              "status",
	CreatedAt:           "created_at",
	UpdatedAt:           "updated_at",
	Value:               "value",
	Cod:                 "cod",
	Weight:              "weight",
	ServiceCode:         "service_code",
	PartnerStatus:       "partner_status",
	PartnerIdentityCode: "partner_identity_code",
	FromName:            "from_name",
	FromPhone:           "from_phone",
	FromAddress:         "from_address",
	FromProvinceCode:    "from_province_code",
	FromDistrictCode:    "from_district_code",
	FromWardCode:        "from_ward_code",
	ToName:              "to_name",
	ToPhone:             "to_phone",
	ToAddress:           "to_address",
	ToProvinceCode:      "to_province_code",
	ToDistrictCode:      "to_district_code",
	ToWardCode:          "to_ward_code",
}

var DeliveryTableColumns = struct {
	ID                  string
	Code                string
	TrackingCode        string
	Note                string
	Status              string
	CreatedAt           string
	UpdatedAt           string
	Value               string
	Cod                 string
	Weight              string
	ServiceCode         string
	PartnerStatus       string
	PartnerIdentityCode string
	FromName            string
	FromPhone           string
	FromAddress         string
	FromProvinceCode    string
	FromDistrictCode    string
	FromWardCode        string
	ToName              string
	ToPhone             string
	ToAddress           string
	ToProvinceCode      string
	ToDistrictCode      string
	ToWardCode          string
}{
	ID:                  "deliveries.id",
	Code:                "deliveries.code",
	TrackingCode:        "deliveries.tracking_code",
	Note:                "deliveries.note",
	Status:              "deliveries.status",
	CreatedAt:           "deliveries.created_at",
	UpdatedAt:           "deliveries.updated_at",
	Value:               "deliveries.value",
	Cod:                 "deliveries.cod",
	Weight:              "deliveries.weight",
	ServiceCode:         "deliveries.service_code",
	PartnerStatus:       "deliveries.partner_status",
	PartnerIdentityCode: "deliveries.partner_identity_code",
	FromName:            "deliveries.from_name",
	FromPhone:           "deliveries.from_phone",
	FromAddress:         "deliveries.from_address",
	FromProvinceCode:    "deliveries.from_province_code",
	FromDistrictCode:    "deliveries.from_district_code",
	FromWardCode:        "deliveries.from_ward_code",
	ToName:              "deliveries.to_name",
	ToPhone:             "deliveries.to_phone",
	ToAddress:           "deliveries.to_address",
	ToProvinceCode:      "deliveries.to_province_code",
	ToDistrictCode:      "deliveries.to_district_code",
	ToWardCode:          "deliveries.to_ward_code",
}

// Generated where

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
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
func (w whereHelpernull_String) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_String) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperdecimal_Decimal struct{ field string }

func (w whereHelperdecimal_Decimal) EQ(x decimal.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelperdecimal_Decimal) NEQ(x decimal.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperdecimal_Decimal) LT(x decimal.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelperdecimal_Decimal) LTE(x decimal.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperdecimal_Decimal) GT(x decimal.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelperdecimal_Decimal) GTE(x decimal.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_Int struct{ field string }

func (w whereHelpernull_Int) EQ(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpernull_Int) NEQ(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpernull_Int) LT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int) LTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int) GT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int) GTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_Int) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_Int) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var DeliveryWhere = struct {
	ID                  whereHelpernull_String
	Code                whereHelpernull_String
	TrackingCode        whereHelpernull_String
	Note                whereHelpernull_String
	Status              whereHelpernull_String
	CreatedAt           whereHelpernull_Time
	UpdatedAt           whereHelpernull_Time
	Value               whereHelperdecimal_Decimal
	Cod                 whereHelperdecimal_Decimal
	Weight              whereHelperdecimal_Decimal
	ServiceCode         whereHelpernull_String
	PartnerStatus       whereHelpernull_String
	PartnerIdentityCode whereHelpernull_String
	FromName            whereHelpernull_String
	FromPhone           whereHelpernull_String
	FromAddress         whereHelpernull_String
	FromProvinceCode    whereHelpernull_Int
	FromDistrictCode    whereHelpernull_Int
	FromWardCode        whereHelpernull_Int
	ToName              whereHelpernull_String
	ToPhone             whereHelpernull_String
	ToAddress           whereHelpernull_String
	ToProvinceCode      whereHelpernull_Int
	ToDistrictCode      whereHelpernull_Int
	ToWardCode          whereHelpernull_Int
}{
	ID:                  whereHelpernull_String{field: "\"deliveries\".\"id\""},
	Code:                whereHelpernull_String{field: "\"deliveries\".\"code\""},
	TrackingCode:        whereHelpernull_String{field: "\"deliveries\".\"tracking_code\""},
	Note:                whereHelpernull_String{field: "\"deliveries\".\"note\""},
	Status:              whereHelpernull_String{field: "\"deliveries\".\"status\""},
	CreatedAt:           whereHelpernull_Time{field: "\"deliveries\".\"created_at\""},
	UpdatedAt:           whereHelpernull_Time{field: "\"deliveries\".\"updated_at\""},
	Value:               whereHelperdecimal_Decimal{field: "\"deliveries\".\"value\""},
	Cod:                 whereHelperdecimal_Decimal{field: "\"deliveries\".\"cod\""},
	Weight:              whereHelperdecimal_Decimal{field: "\"deliveries\".\"weight\""},
	ServiceCode:         whereHelpernull_String{field: "\"deliveries\".\"service_code\""},
	PartnerStatus:       whereHelpernull_String{field: "\"deliveries\".\"partner_status\""},
	PartnerIdentityCode: whereHelpernull_String{field: "\"deliveries\".\"partner_identity_code\""},
	FromName:            whereHelpernull_String{field: "\"deliveries\".\"from_name\""},
	FromPhone:           whereHelpernull_String{field: "\"deliveries\".\"from_phone\""},
	FromAddress:         whereHelpernull_String{field: "\"deliveries\".\"from_address\""},
	FromProvinceCode:    whereHelpernull_Int{field: "\"deliveries\".\"from_province_code\""},
	FromDistrictCode:    whereHelpernull_Int{field: "\"deliveries\".\"from_district_code\""},
	FromWardCode:        whereHelpernull_Int{field: "\"deliveries\".\"from_ward_code\""},
	ToName:              whereHelpernull_String{field: "\"deliveries\".\"to_name\""},
	ToPhone:             whereHelpernull_String{field: "\"deliveries\".\"to_phone\""},
	ToAddress:           whereHelpernull_String{field: "\"deliveries\".\"to_address\""},
	ToProvinceCode:      whereHelpernull_Int{field: "\"deliveries\".\"to_province_code\""},
	ToDistrictCode:      whereHelpernull_Int{field: "\"deliveries\".\"to_district_code\""},
	ToWardCode:          whereHelpernull_Int{field: "\"deliveries\".\"to_ward_code\""},
}

// DeliveryRels is where relationship names are stored.
var DeliveryRels = struct {
}{}

// deliveryR is where relationships are stored.
type deliveryR struct {
}

// NewStruct creates a new relationship struct
func (*deliveryR) NewStruct() *deliveryR {
	return &deliveryR{}
}

// deliveryL is where Load methods for each relationship are stored.
type deliveryL struct{}

var (
	deliveryAllColumns            = []string{"id", "code", "tracking_code", "note", "status", "created_at", "updated_at", "value", "cod", "weight", "service_code", "partner_status", "partner_identity_code", "from_name", "from_phone", "from_address", "from_province_code", "from_district_code", "from_ward_code", "to_name", "to_phone", "to_address", "to_province_code", "to_district_code", "to_ward_code"}
	deliveryColumnsWithoutDefault = []string{"id"}
	deliveryColumnsWithDefault    = []string{"code", "tracking_code", "note", "status", "created_at", "updated_at", "value", "cod", "weight", "service_code", "partner_status", "partner_identity_code", "from_name", "from_phone", "from_address", "from_province_code", "from_district_code", "from_ward_code", "to_name", "to_phone", "to_address", "to_province_code", "to_district_code", "to_ward_code"}
	deliveryPrimaryKeyColumns     = []string{"id"}
	deliveryGeneratedColumns      = []string{}
)

type (
	// DeliverySlice is an alias for a slice of pointers to Delivery.
	// This should almost always be used instead of []Delivery.
	DeliverySlice []*Delivery
	// DeliveryHook is the signature for custom Delivery hook methods
	DeliveryHook func(context.Context, boil.ContextExecutor, *Delivery) error

	deliveryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	deliveryType                 = reflect.TypeOf(&Delivery{})
	deliveryMapping              = queries.MakeStructMapping(deliveryType)
	deliveryPrimaryKeyMapping, _ = queries.BindMapping(deliveryType, deliveryMapping, deliveryPrimaryKeyColumns)
	deliveryInsertCacheMut       sync.RWMutex
	deliveryInsertCache          = make(map[string]insertCache)
	deliveryUpdateCacheMut       sync.RWMutex
	deliveryUpdateCache          = make(map[string]updateCache)
	deliveryUpsertCacheMut       sync.RWMutex
	deliveryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var deliveryAfterSelectHooks []DeliveryHook

var deliveryBeforeInsertHooks []DeliveryHook
var deliveryAfterInsertHooks []DeliveryHook

var deliveryBeforeUpdateHooks []DeliveryHook
var deliveryAfterUpdateHooks []DeliveryHook

var deliveryBeforeDeleteHooks []DeliveryHook
var deliveryAfterDeleteHooks []DeliveryHook

var deliveryBeforeUpsertHooks []DeliveryHook
var deliveryAfterUpsertHooks []DeliveryHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Delivery) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deliveryAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Delivery) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deliveryBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Delivery) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deliveryAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Delivery) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deliveryBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Delivery) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deliveryAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Delivery) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deliveryBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Delivery) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deliveryAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Delivery) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deliveryBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Delivery) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deliveryAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDeliveryHook registers your hook function for all future operations.
func AddDeliveryHook(hookPoint boil.HookPoint, deliveryHook DeliveryHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		deliveryAfterSelectHooks = append(deliveryAfterSelectHooks, deliveryHook)
	case boil.BeforeInsertHook:
		deliveryBeforeInsertHooks = append(deliveryBeforeInsertHooks, deliveryHook)
	case boil.AfterInsertHook:
		deliveryAfterInsertHooks = append(deliveryAfterInsertHooks, deliveryHook)
	case boil.BeforeUpdateHook:
		deliveryBeforeUpdateHooks = append(deliveryBeforeUpdateHooks, deliveryHook)
	case boil.AfterUpdateHook:
		deliveryAfterUpdateHooks = append(deliveryAfterUpdateHooks, deliveryHook)
	case boil.BeforeDeleteHook:
		deliveryBeforeDeleteHooks = append(deliveryBeforeDeleteHooks, deliveryHook)
	case boil.AfterDeleteHook:
		deliveryAfterDeleteHooks = append(deliveryAfterDeleteHooks, deliveryHook)
	case boil.BeforeUpsertHook:
		deliveryBeforeUpsertHooks = append(deliveryBeforeUpsertHooks, deliveryHook)
	case boil.AfterUpsertHook:
		deliveryAfterUpsertHooks = append(deliveryAfterUpsertHooks, deliveryHook)
	}
}

// One returns a single delivery record from the query.
func (q deliveryQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Delivery, error) {
	o := &Delivery{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: failed to execute a one query for deliveries")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Delivery records from the query.
func (q deliveryQuery) All(ctx context.Context, exec boil.ContextExecutor) (DeliverySlice, error) {
	var o []*Delivery

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "entity: failed to assign all query results to Delivery slice")
	}

	if len(deliveryAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Delivery records in the query.
func (q deliveryQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to count deliveries rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q deliveryQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "entity: failed to check if deliveries exists")
	}

	return count > 0, nil
}

// Deliveries retrieves all the records using an executor.
func Deliveries(mods ...qm.QueryMod) deliveryQuery {
	mods = append(mods, qm.From("\"deliveries\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"deliveries\".*"})
	}

	return deliveryQuery{q}
}

// FindDelivery retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDelivery(ctx context.Context, exec boil.ContextExecutor, iD null.String, selectCols ...string) (*Delivery, error) {
	deliveryObj := &Delivery{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"deliveries\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, deliveryObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: unable to select from deliveries")
	}

	if err = deliveryObj.doAfterSelectHooks(ctx, exec); err != nil {
		return deliveryObj, err
	}

	return deliveryObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Delivery) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("entity: no deliveries provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(deliveryColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	deliveryInsertCacheMut.RLock()
	cache, cached := deliveryInsertCache[key]
	deliveryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			deliveryAllColumns,
			deliveryColumnsWithDefault,
			deliveryColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(deliveryType, deliveryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(deliveryType, deliveryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"deliveries\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"deliveries\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "entity: unable to insert into deliveries")
	}

	if !cached {
		deliveryInsertCacheMut.Lock()
		deliveryInsertCache[key] = cache
		deliveryInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Delivery.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Delivery) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	deliveryUpdateCacheMut.RLock()
	cache, cached := deliveryUpdateCache[key]
	deliveryUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			deliveryAllColumns,
			deliveryPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("entity: unable to update deliveries, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"deliveries\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, deliveryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(deliveryType, deliveryMapping, append(wl, deliveryPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "entity: unable to update deliveries row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by update for deliveries")
	}

	if !cached {
		deliveryUpdateCacheMut.Lock()
		deliveryUpdateCache[key] = cache
		deliveryUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q deliveryQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all for deliveries")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected for deliveries")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DeliverySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), deliveryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"deliveries\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, deliveryPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all in delivery slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected all in update all delivery")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Delivery) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("entity: no deliveries provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(deliveryColumnsWithDefault, o)

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

	deliveryUpsertCacheMut.RLock()
	cache, cached := deliveryUpsertCache[key]
	deliveryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			deliveryAllColumns,
			deliveryColumnsWithDefault,
			deliveryColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			deliveryAllColumns,
			deliveryPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("entity: unable to upsert deliveries, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(deliveryPrimaryKeyColumns))
			copy(conflict, deliveryPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"deliveries\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(deliveryType, deliveryMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(deliveryType, deliveryMapping, ret)
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
		return errors.Wrap(err, "entity: unable to upsert deliveries")
	}

	if !cached {
		deliveryUpsertCacheMut.Lock()
		deliveryUpsertCache[key] = cache
		deliveryUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Delivery record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Delivery) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("entity: no Delivery provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), deliveryPrimaryKeyMapping)
	sql := "DELETE FROM \"deliveries\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete from deliveries")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by delete for deliveries")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q deliveryQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("entity: no deliveryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from deliveries")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for deliveries")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DeliverySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(deliveryBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), deliveryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"deliveries\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, deliveryPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from delivery slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for deliveries")
	}

	if len(deliveryAfterDeleteHooks) != 0 {
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
func (o *Delivery) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDelivery(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DeliverySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DeliverySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), deliveryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"deliveries\".* FROM \"deliveries\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, deliveryPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "entity: unable to reload all in DeliverySlice")
	}

	*o = slice

	return nil
}

// DeliveryExists checks if the Delivery row exists.
func DeliveryExists(ctx context.Context, exec boil.ContextExecutor, iD null.String) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"deliveries\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "entity: unable to check if deliveries exists")
	}

	return exists, nil
}

// Exists checks if the Delivery row exists.
func (o *Delivery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return DeliveryExists(ctx, exec, o.ID)
}
