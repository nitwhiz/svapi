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

func (i ItemNameResource) getQueryOptions(r api2go.Request) *storage.QueryOptions {
	opts := &storage.QueryOptions{}

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
	ApplyPagination(r.QueryParams, opts)

	return opts
}

func (i ItemNameResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	opts := i.getQueryOptions(r)

	totalCount := storage.QueryTotalCount[model.ItemName](i.DB, opts)

	return totalCount, &Response{Res: storage.QueryAll[model.ItemName](i.DB, opts)}, nil
}

func (i ItemNameResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	return &Response{Res: storage.QueryAll[model.ItemName](i.DB, i.getQueryOptions(r))}, nil
}

func (i ItemNameResource) FindOne(id string, r api2go.Request) (api2go.Responder, error) {
	res, err := storage.QueryOne[model.ItemName](i.DB, id, &storage.QueryOptions{})

	return &Response{Res: res}, err
}
