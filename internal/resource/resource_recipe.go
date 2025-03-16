package resource

import (
	"github.com/nitwhiz/api2go/v2"
	"github.com/nitwhiz/svapi/pkg/model"
)

type RecipeResource struct {
}

func (c RecipeResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	return Search(model.TypeRecipe, r.QueryParams)
}

func (c RecipeResource) FindOne(id string, _ api2go.Request) (api2go.Responder, error) {
	return FirstById[model.Recipe](id)
}
