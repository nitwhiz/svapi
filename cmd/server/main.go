package main

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

var isRelease = false

func main() {
	if isRelease {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(cors.Default())

	api := api2go.NewAPIWithRouting(
		"v1",
		api2go.NewStaticResolver("/"),
		routing.Gin(router),
	)

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

	err := storage.InitDB()

	if err != nil {
		panic(err)
	}

	if err := loader.Load(); err != nil {
		panic(err)
	}

	for m, r := range storage.ResourceByModel {
		api.AddResource(m, r)
	}

	texturesFS, err := data.GetTexturesFS()

	if err != nil {
		panic(err)
	}

	router.StaticFS("/v1/textures", http.FS(texturesFS))

	err = router.Run("0.0.0.0:4200")

	if err != nil {
		panic(err)
	}
}
