package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/manyminds/api2go"
	"github.com/manyminds/api2go/routing"
	"github.com/nitwhiz/stardew-valley-guide-api/internal/resource"
	"github.com/nitwhiz/stardew-valley-guide-api/internal/storage"
	"github.com/nitwhiz/stardew-valley-guide-api/pkg/model"
	"os"
)

var isRelease = false

func main() {
	_ = godotenv.Load(".env.local")
	_ = godotenv.Load()

	if isRelease {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(cors.Default())

	router.Static("/textures", os.Getenv("API_TEXTURES_DIR"))

	api := api2go.NewAPIWithRouting(
		"v1",
		api2go.NewStaticResolver("/"),
		routing.Gin(router),
	)

	db, err := storage.InitDB(isRelease)

	if err != nil {
		panic(err)
	}

	api.AddResource(model.ItemName{}, resource.ItemNameResource{DB: db})
	api.AddResource(model.Item{}, resource.ItemResource{DB: db})

	api.AddResource(model.NpcName{}, resource.NpcNameResource{DB: db})
	api.AddResource(model.Npc{}, resource.NpcResource{DB: db})

	api.AddResource(model.GiftTaste{}, resource.GiftTasteResource{DB: db})

	err = router.Run("0.0.0.0:4200")

	if err != nil {
		panic(err)
	}
}
