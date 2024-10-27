package resource

import (
	"errors"
	"github.com/manyminds/api2go"
	"github.com/nitwhiz/svapi/internal/storage"
	"github.com/nitwhiz/svapi/pkg/responder"
	"github.com/nitwhiz/svapi/pkg/util"
	"reflect"
	"strings"
)

func First[ModelType storage.Model](id string) (api2go.Responder, error) {
	txn := storage.Database.Txn(false)
	defer txn.Commit()

	m := new(ModelType)

	res, err := storage.First(txn, (*m).TableName(), "id", id)

	if err != nil {
		return nil, err
	}

	return &responder.Response{Res: res}, nil
}

func Search(tableName string, queryParams map[string][]string) (api2go.Responder, error) {
	txn := storage.Database.Txn(false)
	defer txn.Commit()

	srcModel, ok := storage.ModelByTable[tableName]

	if !ok {
		return nil, errors.New("model not found")
	}

	// todo: these allocations are unnecessary in non-filter[] situations

	search := reflect.New(reflect.TypeOf(srcModel).Elem()).Interface()

	searchValue := reflect.ValueOf(search).Elem()
	searchType := searchValue.Type()

	isFilterSearch := false

	searchTableName := srcModel.TableName()
	searchFieldName := ""

	var searchIdFilter string

	for param, values := range queryParams {
		if len(values) == 0 {
			continue
		}

		if strings.HasSuffix(param, "ID") {
			// categoriesID, itemsID, etc.
			// we only support one of these filters

			typeName := strings.TrimSuffix(param, "ID")
			typeModel, ok := storage.ModelByTable[typeName]

			if !ok {
				continue
			}

			searchTableName = typeModel.TableName()

			searchIdFilter = values[0]

			nameParams, ok := queryParams[typeName+"Name"]

			if ok {
				searchFieldName = nameParams[0]
			}

			break
		} else if strings.HasPrefix(param, "filter[") && strings.HasSuffix(param, "]") {
			// filter[someField] = x -> searchStruct.SomeField.SetABC(x)

			filterField := strings.ToLower(strings.TrimPrefix(strings.TrimSuffix(param, "]"), "filter["))

			for i := 0; i < searchType.NumField(); i++ {
				if strings.ToLower(searchType.Field(i).Name) == filterField {
					f := searchValue.Field(i)

					switch f.Type().Kind() {
					case reflect.Bool:
						f.SetBool(util.AsBool(values[0]))
						isFilterSearch = true
						break
					case reflect.String:
						f.SetString(values[0])
						isFilterSearch = true
						break
					case reflect.Pointer:
						v := reflect.New(f.Type().Elem()).Elem()
						idField := v.FieldByName("ID")

						if idField.IsValid() {
							idField.SetString(values[0])
							f.Set(v.Addr())
							isFilterSearch = true
						}
					default:
						break
					}

					break
				}
			}
		}
	}

	var res interface{}
	var err error

	if isFilterSearch {
		res, err = storage.SearchAll(txn, (search).(storage.Model))
	} else {
		if searchIdFilter == "" {
			res, err = storage.FindAll(txn, searchTableName, "id")
		} else {
			res, err = storage.FindAllIn(txn, searchTableName, "id", strings.ToLower(searchFieldName), searchIdFilter)
		}
	}

	if err != nil {
		return nil, err
	}

	return &responder.Response{Res: res}, nil
}
