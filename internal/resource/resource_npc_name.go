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

func (n NpcNameResource) GetAll(r api2go.Request, opts *storage.QueryOptions) []model.NpcName {
	ApplyFilters(
		r.QueryParams,
		map[string]string{
			"npcsID":           "npc_id",
			"filter[language]": "language",
			"filter[name]":     "name",
		},
		opts,
	)

	ApplySorting(r.QueryParams, opts)

	return storage.QueryAll[model.NpcName](n.DB, opts)
}

func (n NpcNameResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	queryOpts := &storage.QueryOptions{}

	ApplyPagination(r.QueryParams, queryOpts)

	totalCount := storage.QueryTotalCount[model.NpcName](n.DB, queryOpts)

	return totalCount, &Response{Res: n.GetAll(r, queryOpts)}, nil
}

func (n NpcNameResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	return &Response{Res: n.GetAll(r, &storage.QueryOptions{})}, nil
}

func (n NpcNameResource) FindOne(id string, r api2go.Request) (api2go.Responder, error) {
	res, err := storage.QueryOne[model.NpcName](n.DB, id, &storage.QueryOptions{})

	return &Response{Res: res}, err
}
