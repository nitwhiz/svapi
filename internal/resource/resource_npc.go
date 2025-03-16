package resource

import (
	"github.com/nitwhiz/api2go/v2"
	"github.com/nitwhiz/svapi/pkg/model"
)

type NpcResource struct{}

func (n NpcResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	return Search(model.TypeNpc, r.QueryParams)
}

func (n NpcResource) FindOne(id string, _ api2go.Request) (api2go.Responder, error) {
	return FirstById[model.Npc](id)
}
