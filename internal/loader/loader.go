package loader

import (
	"github.com/gofrs/uuid/v5"
	"github.com/hashicorp/go-memdb"
	"github.com/nitwhiz/svapi/internal/storage"
	"github.com/nitwhiz/svapi/pkg/model"
)

var namespaceId = uuid.FromStringOrNil("798733e8-df87-41be-911c-96ab8853b8a2")

func newUUID(v string) string {
	return uuid.NewV5(namespaceId, v).String()
}

func first[ModelType storage.Model](txn *memdb.Txn, index string, args ...interface{}) (*ModelType, error) {
	m := new(ModelType)

	raw, err := txn.First((*m).TableName(), index, args...)

	if err != nil {
		return nil, err
	}

	if raw == nil {
		return nil, nil
	}

	return raw.(*ModelType), nil
}

func findOrCreateLanguageByCode(txn *memdb.Txn, code string) (*model.Language, error) {
	l, err := storage.First(txn, model.TypeLanguage, "code", code)

	if err != nil {
		return nil, err
	}

	if l == nil {
		l = &model.Language{
			ID:   newUUID(code),
			Code: code,
		}

		if err := storage.Insert(txn, l); err != nil {
			return nil, err
		}
	}

	return l.(*model.Language), nil
}

func Load() error {
	txn := storage.Database.Txn(true)

	if err := loadCategories(txn); err != nil {
		return err
	}

	if err := loadItems(txn); err != nil {
		return err
	}

	if err := loadNpcs(txn); err != nil {
		return err
	}

	if err := loadGiftTastes(txn); err != nil {
		return err
	}

	if err := loadRecipes(txn); err != nil {
		return err
	}

	txn.Commit()

	return nil
}
