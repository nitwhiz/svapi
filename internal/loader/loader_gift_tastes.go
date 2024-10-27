package loader

import (
	"github.com/hashicorp/go-memdb"
	"github.com/nitwhiz/svapi/internal/data"
	"github.com/nitwhiz/svapi/pkg/model"
)

func loadGiftTastes(txn *memdb.Txn) error {
	tastes, err := data.Parse[GiftTastes]("gift-tastes-by-npc.json")

	if err != nil {
		return err
	}

	for _, taste := range tastes.TastesByNpc {
		if taste.NpcID == "" {
			continue
		}

		npc, err := first[model.Npc](txn, "internalId", taste.NpcID)

		if err != nil {
			return err
		}

		if npc == nil {
			continue
		}

		itemTasteMap := map[string][]string{
			model.TasteDislike: taste.DislikeItems,
			model.TasteHate:    taste.HateItems,
			model.TasteLike:    taste.LikeItems,
			model.TasteLove:    taste.LoveItems,
			model.TasteNeutral: taste.NeutralItems,
		}

		for itemTaste, itemIds := range itemTasteMap {
			for _, itemId := range itemIds {
				if itemId == "" {
					continue
				}

				item, err := first[model.Item](txn, "internalId", itemId)

				if err != nil {
					return err
				}

				if item == nil {
					continue
				}

				giftTaste := &model.GiftTaste{
					ID:    newUUID(npc.ID + "_" + item.ID + "_" + itemTaste),
					Npc:   npc,
					Item:  item,
					Taste: itemTaste,
				}

				npc.GiftTastes = append(npc.GiftTastes, giftTaste)
				item.GiftTastes = append(item.GiftTastes, giftTaste)

				if err := txn.Insert(giftTaste.TableName(), giftTaste); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
