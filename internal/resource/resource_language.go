package resource

import (
	"github.com/manyminds/api2go"
	"github.com/nitwhiz/svapi/pkg/model"
)

type LanguageResource struct{}

func (l LanguageResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	return Search(model.TypeLanguage, r.QueryParams)
}

func (l LanguageResource) FindOne(id string, _ api2go.Request) (api2go.Responder, error) {
	return First[model.Language](id)
}
