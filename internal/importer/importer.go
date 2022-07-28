package importer

import (
	"errors"
	"github.com/google/uuid"
	"github.com/nitwhiz/stardew-valley-guide-api/internal/data"
	"github.com/nitwhiz/stardew-valley-guide-api/pkg/model"
	"gorm.io/gorm"
)

type Importable interface {
	Item | Npc | GiftTasteByNpc
}

type Model interface {
	model.Item | model.Npc | model.GiftTaste
}

func ReadData[I Importable](path string) ([]I, error) {
	d, err := data.Parse[map[string][]I](path)

	if err != nil {
		return nil, err
	}

	for _, v := range d {
		return v, nil
	}

	return nil, nil
}

func ImportItems(db *gorm.DB) error {
	internalData, err := ReadData[Item]("items.json")

	if err != nil {
		return err
	}

	if internalData == nil {
		return errors.New("no item data")
	}

	for _, d := range internalData {
		itemModel := model.Item{
			ID:           d.ID,
			InternalID:   d.InternalID,
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

	return nil
}

func ImportNpcs(db *gorm.DB) error {
	internalData, err := ReadData[Npc]("npcs.json")

	if err != nil {
		return err
	}

	if internalData == nil {
		return errors.New("no npc data")
	}

	for _, d := range internalData {
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

	return nil
}

func ImportGiftTastes(db *gorm.DB) error {
	internalData, err := ReadData[GiftTasteByNpc]("gift-tastes-by-npc.json")

	if err != nil {
		return err
	}

	if internalData == nil {
		return errors.New("no gift taste data")
	}

	for _, d := range internalData {
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

	return nil
}
