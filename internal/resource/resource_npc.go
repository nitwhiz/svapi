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

func (n NpcResource) getQueryOptions(r api2go.Request) *storage.QueryOptions {
	opts := &storage.QueryOptions{}

	ApplyFilters(
		r.QueryParams,
		map[string]string{
			"filter[npc]":            "npc_names.id",
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

	TryApplyJoinFilter(
		r.QueryParams,
		&JoinFilterOptions{
			QueryParamKey: "npcNamesID",
			TableName:     "npc_names",
			FilterColumn:  "npc_names.id",
			LeftOn:        "npcs.id",
			RightOn:       "npc_names.npc_id",
		},
		opts,
	)

	ApplySorting(r.QueryParams, opts)
	ApplyPagination(r.QueryParams, opts)

	return opts
}

func (n NpcResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	opts := n.getQueryOptions(r)

	totalCount := storage.QueryTotalCount[model.Npc](n.DB, opts)

	return totalCount, &Response{Res: storage.QueryAll[model.Npc](n.DB, opts)}, nil
}

func (n NpcResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	return &Response{Res: storage.QueryAll[model.Npc](n.DB, n.getQueryOptions(r))}, nil
}

func (n NpcResource) FindOne(id string, r api2go.Request) (api2go.Responder, error) {
	res, err := storage.QueryOne[model.Npc](n.DB, id, &storage.QueryOptions{})

	return &Response{Res: res}, err
}
