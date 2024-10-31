package loader

import (
	"fmt"
	"github.com/hashicorp/go-memdb"
	"github.com/nitwhiz/svapi/internal/data"
	"github.com/nitwhiz/svapi/internal/storage"
	"github.com/nitwhiz/svapi/pkg/flags"
	"github.com/nitwhiz/svapi/pkg/model"
)

func loadItems(txn *memdb.Txn) error {
	items, err := data.Parse[Items]("items.json")

	if err != nil {
		return err
	}

	for _, obj := range items.Objects {
		catModel, err := first[model.Category](txn, "internalId", fmt.Sprintf("%d", obj.Category))

		if err != nil {
			return err
		}

		if catModel == nil {
			continue
		}

		itemModel := &model.Item{
			ID:          newUUID(obj.ID),
			InternalID:  obj.ID,
			TextureName: obj.TextureName,
			Category:    catModel,
			Type:        obj.Type,
			Flags:       []*flags.Flag{},
			Names:       []*model.ItemName{},
		}

		if obj.IsGiftable {
			itemModel.Flags = append(itemModel.Flags, flags.IsGiftable)
		}

		if obj.IsBigCraftable {
			itemModel.Flags = append(itemModel.Flags, flags.IsBigCraftable)
		}

		catModel.Items = append(catModel.Items, itemModel)

		for langCode, name := range obj.DisplayNames {
			lang, err := findOrCreateLanguageByCode(txn, langCode)

			if err != nil {
				return err
			}

			itemNameModel := &model.ItemName{
				ID:       newUUID(obj.ID + "_" + lang.ID),
				Item:     itemModel,
				Language: lang,
				Name:     name,
			}

			itemModel.Names = append(itemModel.Names, itemNameModel)
			lang.ItemNames = append(lang.ItemNames, itemNameModel)

			if err := storage.Insert(txn, itemNameModel); err != nil {
				return err
			}
		}

		if err := storage.Insert(txn, itemModel); err != nil {
			return err
		}
	}

	return nil
}
