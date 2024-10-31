package resource

import (
	"github.com/manyminds/api2go"
	"github.com/nitwhiz/svapi/pkg/model"
)

type RecipeResource struct {
}

func (c RecipeResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	return Search(model.TypeRecipe, r.QueryParams)
}

func (c RecipeResource) FindOne(id string, _ api2go.Request) (api2go.Responder, error) {
	return First[model.Recipe](id)
}
