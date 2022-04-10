package resource

import (
	"github.com/manyminds/api2go"
	"github.com/nitwhiz/stardew-valley-guide-api/internal/storage"
	"github.com/nitwhiz/stardew-valley-guide-api/pkg/model"
	"gorm.io/gorm"
)

type NpcNameResource struct {
	DB *gorm.DB
}

func (n NpcNameResource) getQueryOptions(r api2go.Request) *storage.QueryOptions {
	opts := &storage.QueryOptions{}

	ApplyFilters(
		r.QueryParams,
		map[string]string{
			"npcsID":           "npc_id",
			"filter[npc]":      "npc_id",
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

func (n NpcNameResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	opts := n.getQueryOptions(r)

	totalCount := storage.QueryTotalCount[model.NpcName](n.DB, opts)

	return totalCount, &Response{Res: storage.QueryAll[model.NpcName](n.DB, opts)}, nil
}

func (n NpcNameResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	return &Response{Res: storage.QueryAll[model.NpcName](n.DB, n.getQueryOptions(r))}, nil
}

func (n NpcNameResource) FindOne(id string, r api2go.Request) (api2go.Responder, error) {
	res, err := storage.QueryOne[model.NpcName](n.DB, id, &storage.QueryOptions{})

	return &Response{Res: res}, err
}
