package importer

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/nitwhiz/stardew-valley-guide-api/pkg/model"
	"gorm.io/gorm"
	"os"
)

type Importable interface {
	Item | Npc | GiftTasteByNpc
}

type Model interface {
	model.Item | model.Npc | model.GiftTaste
}

func ReadJSON[I Importable](jsonPath string) ([]I, error) {
	var data map[string][]I

	bs, err := os.ReadFile(jsonPath)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bs, &data); err != nil {
		return nil, err
	}

	for _, v := range data {
		return v, nil
	}

	return nil, nil
}

func ImportItems(jsonPath string, db *gorm.DB) error {
	data, err := ReadJSON[Item](jsonPath)

	if err != nil {
		return err
	}

	if data == nil {
		return errors.New("no item data")
	}

	for _, d := range data {
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

func ImportNpcs(jsonPath string, db *gorm.DB) error {
	data, err := ReadJSON[Npc](jsonPath)

	if err != nil {
		return err
	}

	if data == nil {
		return errors.New("no npc data")
	}

	for _, d := range data {
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

func ImportGiftTastes(jsonPath string, db *gorm.DB) error {
	data, err := ReadJSON[GiftTasteByNpc](jsonPath)

	if err != nil {
		return err
	}

	if data == nil {
		return errors.New("no gift taste data")
	}

	for _, d := range data {
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
