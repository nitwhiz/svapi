package resource

import (
	"github.com/nitwhiz/api2go/v2"
	"github.com/nitwhiz/svapi/pkg/model"
)

type CategoryResource struct{}

func (c CategoryResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	return Search(model.TypeCategory, r.QueryParams)
}

func (c CategoryResource) FindOne(id string, _ api2go.Request) (api2go.Responder, error) {
	return FirstById[model.Category](id)
}
