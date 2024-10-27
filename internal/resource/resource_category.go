package resource

import (
	"github.com/manyminds/api2go"
	"github.com/nitwhiz/svapi/pkg/model"
)

type CategoryResource struct{}

func (c CategoryResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	return Search(model.TypeCategory, r.QueryParams)
}

func (c CategoryResource) FindOne(id string, r api2go.Request) (api2go.Responder, error) {
	return First[model.Category](id)
}