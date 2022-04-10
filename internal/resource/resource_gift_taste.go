package resource

import (
	"github.com/manyminds/api2go"
	"github.com/nitwhiz/stardew-valley-guide-api/internal/storage"
	"github.com/nitwhiz/stardew-valley-guide-api/pkg/model"
	"gorm.io/gorm"
)

type GiftTasteResource struct {
	DB *gorm.DB
}

func (g GiftTasteResource) getQueryOptions(r api2go.Request) *storage.QueryOptions {
	opts := &storage.QueryOptions{}

	if opts != nil {
		opts.Preload = []string{
			"Npc.DisplayNames",
			"Item.DisplayNames",
		}
	}

	ApplyFilters(
		r.QueryParams,
		map[string]string{
			"filter[item]":  "item_id",
			"filter[npc]":   "npc_id",
			"filter[taste]": "taste",
		},
		opts,
	)

	ApplySorting(r.QueryParams, opts)
	ApplyPagination(r.QueryParams, opts)

	return opts
}

func (g GiftTasteResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	queryOpts := g.getQueryOptions(r)

	totalCount := storage.QueryTotalCount[model.GiftTaste](g.DB, queryOpts)

	return totalCount, &Response{Res: storage.QueryAll[model.GiftTaste](g.DB, queryOpts)}, nil
}

func (g GiftTasteResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	return &Response{Res: storage.QueryAll[model.GiftTaste](g.DB, g.getQueryOptions(r))}, nil
}

func (g GiftTasteResource) FindOne(id string, r api2go.Request) (api2go.Responder, error) {
	queryOpts := &storage.QueryOptions{
		Preload: []string{
			"Npc.DisplayNames",
			"Item.DisplayNames",
		},
	}

	res, err := storage.QueryOne[model.GiftTaste](g.DB, id, queryOpts)

	return &Response{Res: res}, err
}
