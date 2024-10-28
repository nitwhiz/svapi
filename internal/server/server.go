package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/manyminds/api2go"
	"github.com/manyminds/api2go/routing"
	"github.com/nitwhiz/svapi/internal/data"
	"github.com/nitwhiz/svapi/internal/loader"
	"github.com/nitwhiz/svapi/internal/resource"
	"github.com/nitwhiz/svapi/internal/storage"
	"github.com/nitwhiz/svapi/pkg/model"
	"net/http"
)

var Router *gin.Engine

func RegisterModels() {
	storage.RegisterModelAndResource(&model.Language{}, resource.LanguageResource{})
	storage.RegisterModelAndResource(&model.Category{}, resource.CategoryResource{})
	storage.RegisterModelAndResource(&model.CategoryName{}, resource.CategoryNameResource{})
	storage.RegisterModelAndResource(&model.Npc{}, resource.NpcResource{})
	storage.RegisterModelAndResource(&model.NpcName{}, resource.NpcNameResource{})
	storage.RegisterModelAndResource(&model.Item{}, resource.ItemResource{})
	storage.RegisterModelAndResource(&model.ItemName{}, resource.ItemNameResource{})
	storage.RegisterModelAndResource(&model.GiftTaste{}, resource.GiftTasteResource{})
	storage.RegisterModelAndResource(&model.Recipe{}, resource.RecipeResource{})
	storage.RegisterModelAndResource(&model.RecipeIngredient{}, resource.RecipeIngredientResource{})
	storage.RegisterModelAndResource(&model.RecipeIngredientGroup{}, resource.RecipeIngredientGroupResource{})
}

func InitRouter() error {
	Router = gin.Default()
	Router.Use(cors.Default())

	api := api2go.NewAPIWithRouting(
		"v1",
		api2go.NewStaticResolver("/"),
		routing.Gin(Router),
	)

	for m, r := range storage.ResourceByModel {
		api.AddResource(m, r)
	}

	texturesFS, err := data.GetTexturesFS()

	if err != nil {
		return err
	}

	Router.StaticFS("/v1/textures", http.FS(texturesFS))

	return nil
}

func Init() error {
	RegisterModels()

	if err := InitRouter(); err != nil {
		return err
	}

	if err := storage.Init(); err != nil {
		return err
	}

	if err := loader.LoadAll(); err != nil {
		return err
	}

	return nil
}

func Start() error {
	if err := Init(); err != nil {
		return err
	}

	return Router.Run("0.0.0.0:4200")
}
