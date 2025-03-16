package resource

import (
	"github.com/nitwhiz/api2go/v2"
	"github.com/nitwhiz/svapi/pkg/model"
)

type NpcNameResource struct{}

func (n NpcNameResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	return Search(model.TypeNpcName, r.QueryParams)
}

func (n NpcNameResource) FindOne(id string, _ api2go.Request) (api2go.Responder, error) {
	return FirstById[model.NpcName](id)
}
