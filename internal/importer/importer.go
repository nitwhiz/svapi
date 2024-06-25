package importer

import (
	"github.com/google/uuid"
	"github.com/nitwhiz/stardew-valley-guide-api/internal/data"
	"github.com/nitwhiz/stardew-valley-guide-api/pkg/model"
	"gorm.io/gorm"
	"log"
)

type Importable interface {
	Items | Npcs | GiftTastes
}

type Model interface {
	model.Item | model.Npc | model.GiftTaste
}

func needsUpdate(db *gorm.DB, identifier string, internalVersion string) bool {
	v := &model.Version{
		ID: identifier,
	}

	db.First(&v)

	return v.Version == "" || v.Version != internalVersion
}

func ImportItems(db *gorm.DB) error {
	internalData, err := data.Parse[Items]("items.json")

	if err != nil {
		return err
	}

	if !needsUpdate(db, "items", internalData.Version) {
		return nil
	}

	log.Println("updating items ...")

	db.Where("1 = 1").Delete(&model.Item{})

	for _, d := range internalData.Objects {
		itemModel := model.Item{
			ID:           d.ID,
			Category:     d.Category,
			Type:         d.Type,
			DisplayNames: []model.ItemName{},
		}

		for langCode, v := range d.DisplayNames {
			itemNameModel := model.ItemName{
				ID:       uuid.New().String(),
				Language: langCode,
				Name:     v,
			}

			itemModel.DisplayNames = append(itemModel.DisplayNames, itemNameModel)
		}

		db.Create(&itemModel)
	}

	db.Save(&model.Version{
		ID:      "items",
		Version: internalData.Version,
	})

	return nil
}

func ImportNpcs(db *gorm.DB) error {
	internalData, err := data.Parse[Npcs]("npcs.json")

	if err != nil {
		return err
	}

	if !needsUpdate(db, "npcs", internalData.Version) {
		return nil
	}

	log.Println("updating npcs ...")

	db.Where("1 = 1").Delete(&model.Npc{})

	for _, d := range internalData.Npcs {
		npcModel := model.Npc{
			ID:             d.ID,
			BirthdaySeason: d.BirthdaySeason,
			BirthdayDay:    d.BirthdayDay,
			DisplayNames:   []model.NpcName{},
		}

		for langCode, v := range d.DisplayNames {
			npcNameModel := model.NpcName{
				ID:       uuid.New().String(),
				Language: langCode,
				Name:     v,
			}

			npcModel.DisplayNames = append(npcModel.DisplayNames, npcNameModel)
		}

		db.Create(&npcModel)
	}

	db.Save(&model.Version{
		ID:      "npcs",
		Version: internalData.Version,
	})

	return nil
}

func ImportGiftTastes(db *gorm.DB) error {
	internalData, err := data.Parse[GiftTastes]("gift-tastes-by-npc.json")

	if err != nil {
		return err
	}

	if !needsUpdate(db, "giftTastes", internalData.Version) {
		return nil
	}

	log.Println("updating gift tastes ...")

	db.Where("1 = 1").Delete(&model.GiftTaste{})

	for _, d := range internalData.TastesByNpc {
		itemTasteMap := map[string][]string{
			model.TasteDislike: d.DislikeItems,
			model.TasteHate:    d.HateItems,
			model.TasteLike:    d.LikeItems,
			model.TasteLove:    d.LoveItems,
			model.TasteNeutral: d.NeutralItems,
		}

		for itemTaste, itemIds := range itemTasteMap {
			for _, itemId := range itemIds {
				giftTasteModel := model.GiftTaste{
					ID:     uuid.New().String(),
					ItemID: itemId,
					NpcID:  d.NpcID,
					Taste:  itemTaste,
				}

				db.Create(&giftTasteModel)
			}
		}

	}

	db.Save(&model.Version{
		ID:      "giftTastes",
		Version: internalData.Version,
	})

	return nil
}
