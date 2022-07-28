package storage

import (
	"fmt"
	"github.com/nitwhiz/stardew-valley-guide-api/internal/importer"
	"github.com/nitwhiz/stardew-valley-guide-api/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"reflect"
	"strings"
	"time"
)

func InitDB(isRelease bool) (*gorm.DB, error) {
	dbConfig := &gorm.Config{}

	if !isRelease {
		dbConfig.Logger = logger.New(
			log.New(os.Stdout, "\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				Colorful:                  true,
				IgnoreRecordNotFoundError: false,
				LogLevel:                  logger.Info,
			},
		)
	}

	db, err := gorm.Open(postgres.Open(strings.Trim(fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s %s",
		os.Getenv("API_SERVER_DB_HOST"),
		os.Getenv("API_SERVER_DB_PORT"),
		os.Getenv("API_SERVER_DB_USER"),
		os.Getenv("API_SERVER_DB_PASSWORD"),
		os.Getenv("API_SERVER_DB_DATABASE"),
		os.Getenv("API_SERVER_DB_DSN_OPTS"),
	), " ")), dbConfig)

	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(
		&model.Item{},
		&model.Npc{},
		&model.ItemName{},
		&model.NpcName{},
		&model.GiftTaste{},
	); err != nil {
		return nil, err
	}

	itemCount := int64(0)

	db.Model(&model.Item{}).Count(&itemCount)

	if itemCount == 0 {
		if err := importer.ImportItems(db); err != nil {
			return nil, err
		}

		if err := importer.ImportNpcs(db); err != nil {
			return nil, err
		}

		if err := importer.ImportGiftTastes(db); err != nil {
			return nil, err
		}
	}

	return db, nil
}

const (
	OrderDirectionAsc = iota
	OrderDirectionDesc
)

type Order struct {
	Field     string
	Direction int
}

type Search struct {
	Query  string
	Column string
}

type QueryOptions struct {
	WhereColumns map[string]any
	Join         []string
	Preload      []string
	Offset       uint64
	Limit        uint64
	Orders       []Order
	Search       *Search
}

func applyWhere(tx *gorm.DB, opts *QueryOptions) {
	if opts.WhereColumns != nil && len(opts.WhereColumns) > 0 {
		for colName, colValue := range opts.WhereColumns {
			switch reflect.TypeOf(colValue).Kind() {
			case reflect.Slice:
				tx.Where(tx.NamingStrategy.ColumnName("", colName)+" in ?", colValue)
				break
			default:
				tx.Where(tx.NamingStrategy.ColumnName("", colName)+" = ?", colValue)
				break
			}
		}
	}
}

func applyPreload(tx *gorm.DB, opts *QueryOptions) {
	if opts.Preload != nil && len(opts.Preload) > 0 {
		for _, preload := range opts.Preload {
			tx.Preload(preload)
		}
	}
}

func applyJoin(tx *gorm.DB, opts *QueryOptions) {
	if opts.Join != nil && len(opts.Join) > 0 {
		for _, join := range opts.Join {
			tx.Joins(join)
		}
	}
}

func applyOffsetAndLimit(tx *gorm.DB, opts *QueryOptions) {
	if opts.Offset != 0 {
		tx.Offset(int(opts.Offset))
	}

	if opts.Limit != 0 {
		tx.Limit(int(opts.Limit))
	}
}

func applySearchQuery(tx *gorm.DB, opts *QueryOptions) {
	if opts.Search == nil {
		return
	}

	tx.Where(opts.Search.Column + " ILIKE '%" + opts.Search.Query + "%'")

	if opts.Limit == 0 {
		tx.Limit(5)
	}
}

func applyOrder(tx *gorm.DB, opts *QueryOptions) {
	if opts.Orders == nil || len(opts.Orders) == 0 {
		tx.Order(tx.NamingStrategy.ColumnName("", "id") + " ASC")
		return
	}

	for _, order := range opts.Orders {
		orderString := ""

		if order.Field != "" {
			orderString = tx.NamingStrategy.ColumnName("", order.Field)
		}

		if orderString != "" {
			if order.Direction == OrderDirectionDesc {
				orderString += " DESC"
			} else {
				orderString += " ASC"
			}
		}

		tx.Order(orderString)
	}
}

func QueryTotalCount[ModelType any](db *gorm.DB, opts *QueryOptions) uint {
	tx := db.Model(new(ModelType))

	if opts != nil {
		applyPreload(tx, opts)
		applyWhere(tx, opts)
		applyJoin(tx, opts)
		applySearchQuery(tx, opts)
	}

	var res int64

	tx.Preload(clause.Associations).Count(&res)

	return uint(res)
}

func QueryAll[ResultModelType any](db *gorm.DB, opts *QueryOptions) []ResultModelType {
	var res []ResultModelType

	tx := db.Model(new(ResultModelType))

	if opts != nil {
		applyPreload(tx, opts)
		applyWhere(tx, opts)
		applyJoin(tx, opts)
		applySearchQuery(tx, opts)
		applyOrder(tx, opts)
		applyOffsetAndLimit(tx, opts)
	}

	tx.Preload(clause.Associations).Find(&res)

	return res
}

func QueryOne[ResultModelType any](db *gorm.DB, id string, opts *QueryOptions) (ResultModelType, error) {
	var res ResultModelType

	tx := db.Model(new(ResultModelType))

	if opts != nil {
		applyPreload(tx, opts)
		applyWhere(tx, opts)
		applyJoin(tx, opts)
		applySearchQuery(tx, opts)
		applyOrder(tx, opts)
		applyOffsetAndLimit(tx, opts)
	}

	tx.Preload(clause.Associations).First(&res, "id = ?", id)

	if err := db.Error; err == gorm.ErrRecordNotFound {
		return res, fmt.Errorf("element for id %s not found", id)
	}

	return res, nil
}
