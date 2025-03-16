package resource

import (
	"github.com/nitwhiz/api2go/v2"
	"github.com/nitwhiz/svapi/pkg/model"
)

type RecipeIngredientGroupResource struct {
}

func (c RecipeIngredientGroupResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	return Search(model.TypeRecipeIngredientGroup, r.QueryParams)
}

func (c RecipeIngredientGroupResource) FindOne(id string, _ api2go.Request) (api2go.Responder, error) {
	return FirstById[model.RecipeIngredientGroup](id)
}
