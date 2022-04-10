package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/nitwhiz/stardew-valley-guide-api/internal/importer"
	"github.com/nitwhiz/stardew-valley-guide-api/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	_ = godotenv.Load(".env.local")
	_ = godotenv.Load()

	db, err := gorm.Open(
		postgres.Open(
			strings.Trim(fmt.Sprintf(
				"host=%s port=%s user=%s password=%s dbname=%s %s",
				os.Getenv("IMPORTER_DB_HOST"),
				os.Getenv("IMPORTER_DB_PORT"),
				os.Getenv("IMPORTER_DB_USER"),
				os.Getenv("IMPORTER_DB_PASSWORD"),
				os.Getenv("IMPORTER_DB_DATABASE"),
				os.Getenv("IMPORTER_DB_DSN_OPTS"),
			), " "),
		),
		&gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\n", log.LstdFlags),
				logger.Config{
					SlowThreshold:             time.Second,
					Colorful:                  true,
					IgnoreRecordNotFoundError: false,
					LogLevel:                  logger.Info,
				},
			),
		},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("deleting existing records ...")

	db.Where("1 = 1").Delete(&model.Item{})
	db.Where("1 = 1").Delete(&model.Npc{})

	db.Where("1 = 1").Delete(&model.ItemName{})
	db.Where("1 = 1").Delete(&model.NpcName{})

	db.Where("1 = 1").Delete(&model.GiftTaste{})

	fmt.Println("importing data ...")

	if err := importer.ImportItems(os.Getenv("IMPORTER_JSON_ITEMS"), db); err != nil {
		panic(err)
	}

	if err := importer.ImportNpcs(os.Getenv("IMPORTER_JSON_NPCS"), db); err != nil {
		panic(err)
	}

	if err := importer.ImportGiftTastes(os.Getenv("IMPORTER_JSON_GIFT_TASTES"), db); err != nil {
		panic(err)
	}
}
