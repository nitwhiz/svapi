package resource

import (
	"github.com/nitwhiz/api2go/v2"
	"github.com/nitwhiz/svapi/pkg/model"
)

type ItemNameResource struct{}

func (n ItemNameResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	return Search(model.TypeItemName, r.QueryParams)
}

func (n ItemNameResource) FindOne(id string, _ api2go.Request) (api2go.Responder, error) {
	return FirstById[model.ItemName](id)
}
