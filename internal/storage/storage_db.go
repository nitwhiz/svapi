package storage

import (
	"errors"
	"github.com/hashicorp/go-memdb"
	"reflect"
	"strings"
)

var ModelByTable = map[string]Model{}
var ResourceByModel = map[Model]any{}

var Database *memdb.MemDB

func RegisterModelAndResource(m Model, r any) {
	ModelByTable[m.TableName()] = m
	ResourceByModel[m] = r
}

func newDbSchema() *memdb.DBSchema {
	tables := map[string]*memdb.TableSchema{}

	schema := &memdb.DBSchema{
		Tables: tables,
	}

	for t, m := range ModelByTable {
		indexes := m.Indexes()

		indexes["id"] = &memdb.IndexSchema{
			Name:    "id",
			Unique:  true,
			Indexer: &memdb.UUIDFieldIndex{Field: "ID"},
		}

		indexes["search"] = &memdb.IndexSchema{
			Name:    "search",
			Indexer: &SearchIndex{},
		}

		tables[t] = &memdb.TableSchema{
			Name:    t,
			Indexes: indexes,
		}
	}

	return schema
}

func Init() error {
	if Database != nil {
		return errors.New("already initialized")
	}

	db, err := memdb.NewMemDB(newDbSchema())

	if err != nil {
		return err
	}

	Database = db

	return nil
}

func SearchAll(txn *memdb.Txn, search Model) ([]Model, error) {
	it, err := txn.Get(search.TableName(), "search", strings.Join(search.SearchIndexContents(), "_"))

	if err != nil {
		return nil, err
	}

	var res []Model

	for obj := it.Next(); obj != nil; obj = it.Next() {
		res = append(res, obj.(Model))
	}

	return res, nil
}

func FindAll(txn *memdb.Txn, tableName string, index string, args ...interface{}) ([]Model, error) {
	it, err := txn.Get(tableName, index, args...)

	if err != nil {
		return nil, err
	}

	var res []Model

	for obj := it.Next(); obj != nil; obj = it.Next() {
		res = append(res, obj.(Model))
	}

	return res, nil
}

func FindAllIn(txn *memdb.Txn, tableName string, index string, fieldName string, args ...interface{}) (any, error) {
	it, err := txn.Get(tableName, index, args...)

	if err != nil {
		return nil, err
	}

	var res []any
	isSingleResult := false
	findCount := 0

	for obj := it.Next(); obj != nil; obj = it.Next() {
		findCount++

		resVal := reflect.ValueOf(obj).Elem()
		resType := resVal.Type()

		var resField reflect.Value

		for i := 0; i < resVal.NumField(); i++ {
			if strings.ToLower(resType.Field(i).Name) == fieldName {
				resField = resVal.Field(i)
				break
			}
		}

		if resField.IsValid() {
			if resField.Kind() == reflect.Slice {
				for i := 0; i < resField.Len(); i++ {
					res = append(res, resField.Index(i).Interface())
				}
			} else {
				isSingleResult = true
				res = append(res, resField.Interface())
			}
		}
	}

	if findCount == 1 && isSingleResult && len(res) == 1 {
		return res[0], nil
	}

	return res, nil
}

func First(txn *memdb.Txn, tableName string, index string, args ...interface{}) (Model, error) {
	raw, err := txn.First(tableName, index, args...)

	if err != nil {
		return nil, err
	}

	if raw == nil {
		return nil, nil
	}

	return raw.(Model), nil
}

func Insert[ModelType Model](txn *memdb.Txn, m ModelType) error {
	return txn.Insert(m.TableName(), m)
}
