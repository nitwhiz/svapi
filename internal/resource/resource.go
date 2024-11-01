package resource

import (
	"errors"
	"github.com/manyminds/api2go"
	"github.com/nitwhiz/svapi/internal/storage"
	"github.com/nitwhiz/svapi/pkg/flags"
	"github.com/nitwhiz/svapi/pkg/responder"
	"reflect"
	"strings"
)

type search struct {
	model     storage.Model
	table     string
	id        string
	fieldName string
	offset    int
	limit     int
}

func FirstById[ModelType storage.Model](id string) (api2go.Responder, error) {
	txn := storage.Database.Txn(false)
	defer txn.Commit()

	m := new(ModelType)

	res, err := storage.First(txn, (*m).TableName(), "id", id)

	if err != nil {
		return nil, err
	}

	return &responder.Response{Res: res}, nil
}

func buildSearch(tableName string, queryParams map[string][]string) (*search, error) {
	s := search{
		model:     nil,
		table:     "",
		id:        "",
		fieldName: "",
	}

	srcModel, ok := storage.ModelByTable[tableName]

	if !ok {
		return nil, errors.New("model not found")
	}

	searchModel := reflect.New(reflect.TypeOf(srcModel).Elem()).Interface()

	searchValue := reflect.ValueOf(searchModel).Elem()
	searchType := searchValue.Type()

	s.table = srcModel.TableName()

	isFilterSearch := false

	for param, values := range queryParams {
		if len(values) == 0 {
			continue
		}

		if strings.HasSuffix(param, "ID") {
			// categoriesID, itemsID, etc.
			// only support one of these filters

			typeName := strings.TrimSuffix(param, "ID")
			typeModel, ok := storage.ModelByTable[typeName]

			if !ok {
				continue
			}

			s.table = typeModel.TableName()
			s.id = values[0]

			nameParams, ok := queryParams[typeName+"Name"]

			if ok {
				s.fieldName = nameParams[0]
			}

			break
		} else if param, ok := strings.CutPrefix(param, "filter"); ok {

			// filter[someField] = x -> searchStruct.SomeField.SetABC(x)

			filterField := strings.ToLower(param[1:strings.Index(param, "]")])

			for i := 0; i < searchType.NumField(); i++ {
				if strings.ToLower(searchType.Field(i).Name) == filterField {
					f := searchValue.Field(i)

					switch f.Type().Kind() {
					case reflect.String:
						f.SetString(values[0])
						isFilterSearch = true
						break
					case reflect.Pointer:
						// resolve as another model with only it's id being set

						v := reflect.New(f.Type().Elem()).Elem()
						idField := v.FieldByName("ID")

						if idField.IsValid() {
							idField.SetString(values[0])
							f.Set(v.Addr())
							isFilterSearch = true
						}
						break
					case reflect.Slice:
						slicePointerType := f.Type().Elem().Elem()

						if slicePointerType == reflect.TypeOf(flags.Flag{}) {
							// handle flags

							fs := make([]*flags.Flag, len(values))

							for idx, flagSegment := range values {
								fs[idx] = flags.Get(flagSegment)
							}

							f.Set(reflect.ValueOf(fs))
							isFilterSearch = true
						} else if slicePointerType.Implements(reflect.TypeFor[storage.Model]()) {
							// create instances, set ID

							v := reflect.New(slicePointerType).Elem()
							idField := v.FieldByName("ID")

							if idField.IsValid() {
								idField.SetString(values[0])

								s := reflect.MakeSlice(reflect.SliceOf(reflect.PointerTo(slicePointerType)), 1, 1)

								s.Index(0).Set(v.Addr())

								f.Set(s)
								isFilterSearch = true
							}
						}

						break
					default:
						break
					}

					break
				}
			}
		}
	}

	if isFilterSearch {
		s.model = searchModel.(storage.Model)
	}

	return &s, nil
}

func Search(tableName string, queryParams map[string][]string) (response api2go.Responder, err error) {
	s, err := buildSearch(tableName, queryParams)

	if err != nil {
		return nil, err
	}

	txn := storage.Database.Txn(false)
	defer txn.Commit()

	var res interface{}

	if s.model != nil {
		res, err = storage.SearchAll(txn, s.model)
	} else {
		if s.id == "" {
			res, err = storage.FindAll(txn, s.table, "id")
		} else {
			res, err = storage.FindAllInField(txn, s.table, "id", strings.ToLower(s.fieldName), s.id)
		}
	}

	if err != nil {
		return nil, err
	}

	return &responder.Response{Res: res}, nil
}
