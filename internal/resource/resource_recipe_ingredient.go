package resource

import (
	"github.com/nitwhiz/api2go/v2"
	"github.com/nitwhiz/svapi/pkg/model"
)

type RecipeIngredientResource struct {
}

func (c RecipeIngredientResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	return Search(model.TypeRecipeIngredient, r.QueryParams)
}

func (c RecipeIngredientResource) FindOne(id string, _ api2go.Request) (api2go.Responder, error) {
	return FirstById[model.RecipeIngredient](id)
}
