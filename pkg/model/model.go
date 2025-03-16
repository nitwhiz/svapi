package model

import (
	"github.com/nitwhiz/api2go/v2/jsonapi"
	"reflect"
	"strings"
)

func findIncludeValue(model any, includeName string) reflect.Value {
	v := reflect.ValueOf(model)

	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	typ := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fieldIncludeName := typ.Field(i).Tag.Get("include")

		if fieldIncludeName == "" {
			continue
		}

		if fieldIncludeName == includeName {
			return v.Field(i)
		}
	}

	return reflect.Value{}
}

func doBuildIncluded(include []string, model any, results *[]jsonapi.MarshalIdentifier) {
	for _, i := range include {
		if strings.Contains(i, ".") {
			segments := strings.SplitN(i, ".", 2)

			if len(segments) < 2 {
				continue
			}

			if segments[1] == "" {
				continue
			}

			baseFieldValue := findIncludeValue(model, segments[0])

			if !baseFieldValue.IsValid() {
				continue
			}

			if baseFieldValue.Type().Kind() == reflect.Slice {
				for j := 0; j < baseFieldValue.Len(); j++ {
					doBuildIncluded([]string{segments[1]}, baseFieldValue.Index(j).Interface(), results)
				}
			} else {
				doBuildIncluded([]string{segments[1]}, baseFieldValue.Interface(), results)
			}

			continue
		}

		fieldValue := findIncludeValue(model, i)

		if !fieldValue.IsValid() {
			continue
		}

		if fieldValue.Type().Kind() == reflect.Slice {
			for j := 0; j < fieldValue.Len(); j++ {
				*results = append(*results, fieldValue.Index(j).Interface().(jsonapi.MarshalIdentifier))
			}
		} else {
			*results = append(*results, fieldValue.Interface().(jsonapi.MarshalIdentifier))
		}
	}
}

func BuildIncluded(include []string, model any) []jsonapi.MarshalIdentifier {
	var results []jsonapi.MarshalIdentifier

	doBuildIncluded(include, model, &results)

	return results
}

func BuildReferencedIDs(model any) []jsonapi.ReferenceID {
	var results []jsonapi.ReferenceID

	if m, ok := model.(jsonapi.MarshalReferences); ok {
		for _, ref := range m.GetReferences() {
			fieldValue := findIncludeValue(model, ref.Name)

			if !fieldValue.IsValid() {
				continue
			}

			if fieldValue.Type().Kind() == reflect.Slice {
				for j := 0; j < fieldValue.Len(); j++ {
					results = append(results, jsonapi.ReferenceID{
						ID:           fieldValue.Index(j).Interface().(jsonapi.MarshalIdentifier).GetID(),
						Type:         ref.Type,
						Name:         ref.Name,
						Relationship: ref.Relationship,
					})
				}
			} else {
				results = append(results, jsonapi.ReferenceID{
					ID:           fieldValue.Interface().(jsonapi.MarshalIdentifier).GetID(),
					Type:         ref.Type,
					Name:         ref.Name,
					Relationship: ref.Relationship,
				})
			}

		}
	}

	return results
}
