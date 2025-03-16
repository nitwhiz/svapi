package resource

import (
	"github.com/nitwhiz/api2go/v2"
	"github.com/nitwhiz/svapi/pkg/model"
)

type ItemResource struct{}

func (i ItemResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	return Search(model.TypeItem, r.QueryParams)
}

func (i ItemResource) FindOne(id string, _ api2go.Request) (api2go.Responder, error) {
	return FirstById[model.Item](id)
}
