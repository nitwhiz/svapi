package resource

import (
	"github.com/nitwhiz/api2go/v2"
	"github.com/nitwhiz/svapi/pkg/model"
)

type LanguageResource struct{}

func (l LanguageResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	return Search(model.TypeLanguage, r.QueryParams)
}

func (l LanguageResource) FindOne(id string, _ api2go.Request) (api2go.Responder, error) {
	return FirstById[model.Language](id)
}
