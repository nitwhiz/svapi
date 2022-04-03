package resource

import (
	"github.com/manyminds/api2go"
	"github.com/nitwhiz/stardew-valley-guide-api/internal/storage"
	"github.com/nitwhiz/stardew-valley-guide-api/pkg/model"
	"gorm.io/gorm"
)

type ItemResource struct {
	DB *gorm.DB
}

func (i ItemResource) GetAll(r api2go.Request, opts *storage.QueryOptions) []model.Item {
	ApplyFilters(
		r.QueryParams,
		map[string]string{
			"filter[item]":       "item_names.id",
			"filter[internalId]": "internalId",
			"filter[category]":   "category",
			"filter[type]":       "type",
		},
		opts,
	)

	TryApplyJoinFilter(
		r.QueryParams,
		&JoinFilterOptions{
			QueryParamKey: "giftTastesID",
			TableName:     "gift_tastes",
			FilterColumn:  "gift_tastes.id",
			LeftOn:        "items.id",
			RightOn:       "gift_tastes.item_id",
		},
		opts,
	)

	TryApplyJoinFilter(
		r.QueryParams,
		&JoinFilterOptions{
			QueryParamKey: "itemNamesID",
			TableName:     "item_names",
			FilterColumn:  "item_names.id",
			LeftOn:        "items.id",
			RightOn:       "item_names.item_id",
		},
		opts,
	)

	ApplySorting(r.QueryParams, opts)

	return storage.QueryAll[model.Item](i.DB, opts)
}

func (i ItemResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	queryOpts := &storage.QueryOptions{}

	ApplyPagination(r.QueryParams, queryOpts)

	totalCount := storage.QueryTotalCount[model.Item](i.DB, queryOpts)

	return totalCount, &Response{Res: i.GetAll(r, queryOpts)}, nil
}

func (i ItemResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	return &Response{Res: i.GetAll(r, &storage.QueryOptions{})}, nil
}

func (i ItemResource) FindOne(id string, r api2go.Request) (api2go.Responder, error) {
	res, err := storage.QueryOne[model.Item](i.DB, id, &storage.QueryOptions{})

	return &Response{Res: res}, err
}
