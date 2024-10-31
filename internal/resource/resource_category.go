package resource

import (
	"github.com/manyminds/api2go"
	"github.com/nitwhiz/svapi/pkg/model"
)

type CategoryResource struct{}

func (c CategoryResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	return Search(model.TypeCategory, r.QueryParams)
}

func (c CategoryResource) FindOne(id string, _ api2go.Request) (api2go.Responder, error) {
	return First[model.Category](id)
}

//func (c CategoryResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
//}
