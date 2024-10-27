package loader

import (
	"github.com/hashicorp/go-memdb"
	"github.com/nitwhiz/svapi/internal/data"
	"github.com/nitwhiz/svapi/internal/storage"
	"github.com/nitwhiz/svapi/pkg/model"
)

func loadNpcs(txn *memdb.Txn) error {
	npcs, err := data.Parse[Npcs]("npcs.json")

	if err != nil {
		return err
	}

	for _, npc := range npcs.Npcs {
		npcModel := &model.Npc{
			ID:             newUUID(npc.ID),
			InternalID:     npc.ID,
			TextureName:    npc.TextureName,
			BirthdaySeason: npc.BirthdaySeason,
			BirthdayDay:    npc.BirthdayDay,
			Names:          []*model.NpcName{},
		}

		for langCode, name := range npc.DisplayNames {
			lang, err := findOrCreateLanguageByCode(txn, langCode)

			if err != nil {
				return err
			}

			npcNameModel := &model.NpcName{
				ID:       newUUID(npc.ID + "_" + lang.ID),
				Npc:      npcModel,
				Language: lang,
				Name:     name,
			}

			npcModel.Names = append(npcModel.Names, npcNameModel)
			lang.NpcNames = append(lang.NpcNames, npcNameModel)

			if err := storage.Insert(txn, npcNameModel); err != nil {
				return err
			}
		}

		if err := storage.Insert(txn, npcModel); err != nil {
			return err
		}
	}

	return nil
}
