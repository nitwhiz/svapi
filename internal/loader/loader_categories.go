package loader

import (
	"fmt"
	"github.com/hashicorp/go-memdb"
	"github.com/nitwhiz/svapi/internal/data"
	"github.com/nitwhiz/svapi/internal/storage"
	"github.com/nitwhiz/svapi/pkg/model"
)

func loadCategories(txn *memdb.Txn) error {
	cats, err := data.Parse[Categories]("categories.json")

	if err != nil {
		return err
	}

	cat0Model := &model.Category{
		ID:         newUUID("0"),
		InternalID: "0",
		Items:      []*model.Item{},
		Names:      []*model.CategoryName{},
	}

	if err = storage.Insert(txn, cat0Model); err != nil {
		return err
	}

	for _, cat := range cats.Categories {
		internalId := fmt.Sprintf("%d", cat.ID)

		catModel := &model.Category{
			ID:         newUUID(internalId),
			InternalID: internalId,
			Names:      []*model.CategoryName{},
		}

		for langCode, name := range cat.DisplayNames {
			lang, err := findOrCreateLanguageByCode(txn, langCode)

			if err != nil {
				return err
			}

			catNameModel := &model.CategoryName{
				ID:       newUUID(internalId + "_" + lang.Code),
				Category: catModel,
				Language: lang,
				Name:     name,
			}

			catModel.Names = append(catModel.Names, catNameModel)
			lang.CategoryNames = append(lang.CategoryNames, catNameModel)

			if err = storage.Insert(txn, catNameModel); err != nil {
				return err
			}
		}

		if err = storage.Insert(txn, catModel); err != nil {
			return err
		}
	}

	return nil
}
