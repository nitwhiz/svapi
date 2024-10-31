package resource

import (
	"github.com/manyminds/api2go"
	"github.com/nitwhiz/svapi/pkg/model"
)

type CategoryNameResource struct{}

func (c CategoryNameResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	return Search(model.TypeCategoryName, r.QueryParams)
}

func (c CategoryNameResource) FindOne(id string, _ api2go.Request) (api2go.Responder, error) {
	return First[model.CategoryName](id)
}
