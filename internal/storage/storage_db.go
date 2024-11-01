package storage

import (
	"errors"
	"github.com/hashicorp/go-memdb"
	"reflect"
	"strings"
)

var ModelByTable = map[string]Model{}
var ResourceByModel = map[Model]any{}
var FieldsByModelTableName = map[string]map[string]reflect.StructField{}

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

		modelType := reflect.TypeOf(m).Elem()

		fs, ok := FieldsByModelTableName[t]

		if !ok {
			fs = map[string]reflect.StructField{}
			FieldsByModelTableName[t] = fs
		}

		for i := 0; i < modelType.NumField(); i++ {
			f := modelType.Field(i)
			fs[strings.ToLower(f.Name)] = f
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
	return FindAll(txn, search.TableName(), "search", strings.Join(search.SearchIndexContents(), "_"))
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

func FindAllInField(txn *memdb.Txn, tableName string, index string, fieldName string, args ...interface{}) (any, error) {
	it, err := txn.Get(tableName, index, args...)

	if err != nil {
		return nil, err
	}

	var res []any

	fields, ok := FieldsByModelTableName[tableName]

	if !ok {
		return nil, errors.New("model not found")
	}

	tableModelField, ok := fields[fieldName]

	if !ok {
		return nil, errors.New("field not found")
	}

	tableModelFieldIsSlice := tableModelField.Type.Kind() == reflect.Slice

	for obj := it.Next(); obj != nil; obj = it.Next() {
		objField := reflect.ValueOf(obj).Elem().FieldByName(tableModelField.Name)

		if tableModelFieldIsSlice {
			for i := 0; i < objField.Len(); i++ {
				res = append(res, objField.Index(i).Interface())
			}
		} else {
			res = append(res, objField.Interface())
			break
		}
	}

	if !tableModelFieldIsSlice {
		if len(res) == 0 {
			return nil, nil
		}

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
