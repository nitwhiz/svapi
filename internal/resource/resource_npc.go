package resource

import (
	"github.com/manyminds/api2go"
	"github.com/nitwhiz/stardew-valley-guide-api/internal/storage"
	"github.com/nitwhiz/stardew-valley-guide-api/pkg/model"
	"gorm.io/gorm"
)

type NpcResource struct {
	DB *gorm.DB
}

func (n NpcResource) GetAll(r api2go.Request, opts *storage.QueryOptions) []model.Npc {
	ApplyFilters(
		r.QueryParams,
		map[string]string{
			"filter[birthdaySeason]": "birthdaySeason",
			"filter[birthdayDay]":    "birthdayDay",
		},
		opts,
	)

	TryApplyJoinFilter(
		r.QueryParams,
		&JoinFilterOptions{
			QueryParamKey: "giftTastesID",
			TableName:     "gift_tastes",
			FilterColumn:  "gift_tastes.id",
			LeftOn:        "npcs.id",
			RightOn:       "gift_tastes.npc_id",
		},
		opts,
	)

	ApplySorting(r.QueryParams, opts)

	return storage.QueryAll[model.Npc](n.DB, opts)
}

func (n NpcResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	queryOpts := &storage.QueryOptions{}

	ApplyPagination(r.QueryParams, queryOpts)

	totalCount := storage.QueryTotalCount[model.Npc](n.DB, queryOpts)

	return totalCount, &Response{Res: n.GetAll(r, queryOpts)}, nil
}

func (n NpcResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	return &Response{Res: n.GetAll(r, &storage.QueryOptions{})}, nil
}

func (n NpcResource) FindOne(id string, r api2go.Request) (api2go.Responder, error) {
	res, err := storage.QueryOne[model.Npc](n.DB, id, &storage.QueryOptions{})

	return &Response{Res: res}, err
}
