package resource

import (
	"github.com/manyminds/api2go"
	"github.com/nitwhiz/svapi/pkg/model"
)

type ItemNameResource struct{}

func (n ItemNameResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	return Search(model.TypeItemName, r.QueryParams)
}

func (n ItemNameResource) FindOne(id string, _ api2go.Request) (api2go.Responder, error) {
	return First[model.ItemName](id)
}
