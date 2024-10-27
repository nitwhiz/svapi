package resource

import (
	"github.com/manyminds/api2go"
	"github.com/nitwhiz/svapi/pkg/model"
)

type GiftTasteResource struct{}

func (g GiftTasteResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	return Search(model.TypeGiftTaste, r.QueryParams)
}

func (g GiftTasteResource) FindOne(id string, r api2go.Request) (api2go.Responder, error) {
	return First[model.GiftTaste](id)
}
