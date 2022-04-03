package resource

import (
	"github.com/manyminds/api2go"
	"github.com/nitwhiz/stardew-valley-guide-api/internal/storage"
	"github.com/nitwhiz/stardew-valley-guide-api/pkg/model"
	"gorm.io/gorm"
)

type ItemNameResource struct {
	DB *gorm.DB
}

func (i ItemNameResource) GetAll(r api2go.Request, opts *storage.QueryOptions) []model.ItemName {
	ApplyFilters(
		r.QueryParams,
		map[string]string{
			"itemsID":          "item_id",
			"filter[item]":     "item_id",
			"filter[language]": "language",
			"filter[name]":     "name",
			"filter[query]":    "name",
		},
		opts,
	)

	ApplySorting(r.QueryParams, opts)

	return storage.QueryAll[model.ItemName](i.DB, opts)
}

func (i ItemNameResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	queryOpts := &storage.QueryOptions{}

	ApplyPagination(r.QueryParams, queryOpts)

	totalCount := storage.QueryTotalCount[model.ItemName](i.DB, queryOpts)

	return totalCount, &Response{Res: i.GetAll(r, queryOpts)}, nil
}

func (i ItemNameResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	return &Response{Res: i.GetAll(r, &storage.QueryOptions{})}, nil
}

func (i ItemNameResource) FindOne(id string, r api2go.Request) (api2go.Responder, error) {
	res, err := storage.QueryOne[model.ItemName](i.DB, id, &storage.QueryOptions{})

	return &Response{Res: res}, err
}
