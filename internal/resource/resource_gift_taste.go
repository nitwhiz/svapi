package resource

import (
	"github.com/nitwhiz/api2go/v2"
	"github.com/nitwhiz/svapi/pkg/model"
)

type GiftTasteResource struct{}

func (g GiftTasteResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	return Search(model.TypeGiftTaste, r.QueryParams)
}

func (g GiftTasteResource) FindOne(id string, _ api2go.Request) (api2go.Responder, error) {
	return FirstById[model.GiftTaste](id)
}
