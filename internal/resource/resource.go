package resource

import (
	"fmt"
	"github.com/nitwhiz/stardew-valley-guide-api/internal/storage"
	"strconv"
	"strings"
)

type JoinFilterOptions struct {
	QueryParamKey string
	TableName     string
	FilterColumn  string
	LeftOn        string
	RightOn       string
}

// TryApplyJoinFilter only applies the join filter if JoinFilterOptions.QueryParamKey is found in queryParams
func TryApplyJoinFilter(queryParams map[string][]string, joinFilterOpts *JoinFilterOptions, queryOpts *storage.QueryOptions) {
	filter, ok := queryParams[joinFilterOpts.QueryParamKey]

	if ok {
		if queryOpts.WhereColumns == nil {
			queryOpts.WhereColumns = map[string]any{
				joinFilterOpts.FilterColumn: filter,
			}
		} else {
			queryOpts.WhereColumns[joinFilterOpts.FilterColumn] = filter
		}

		joinStr := "JOIN " + joinFilterOpts.TableName + " ON " + joinFilterOpts.LeftOn + " = " + joinFilterOpts.RightOn

		if queryOpts.Join == nil {
			queryOpts.Join = append(queryOpts.Join, joinStr)
		} else {
			queryOpts.Join = []string{
				joinStr,
			}
		}
	}
}

func ApplyPagination(queryParams map[string][]string, queryOpts *storage.QueryOptions) {
	if offsetQuery, ok := queryParams["page[offset]"]; ok && len(offsetQuery) > 0 {
		offset, _ := strconv.ParseUint(offsetQuery[0], 10, 64)

		queryOpts.Offset = offset
	}

	if limitQuery, ok := queryParams["page[limit]"]; ok && len(limitQuery) > 0 {
		limit, _ := strconv.ParseUint(limitQuery[0], 10, 64)

		queryOpts.Limit = limit
	}
}

func ApplySorting(queryParams map[string][]string, queryOpts *storage.QueryOptions) {
	fmt.Println("%+v\n", queryOpts)

	if queryOpts == nil {
		return
	}

	if queryOpts.Orders == nil {
		queryOpts.Orders = []storage.Order{}
	}

	if sortingQuery, ok := queryParams["sort"]; ok && len(sortingQuery) > 0 {
		for _, sortField := range sortingQuery {
			desc := strings.HasPrefix(sortField, "-")

			if desc && len(sortField) <= 1 {
				continue
			}

			order := storage.Order{}

			if desc {
				order.Direction = storage.OrderDirectionDesc
			} else {
				order.Direction = storage.OrderDirectionAsc
			}

			order.Field = strings.TrimPrefix(sortField, "-")

			queryOpts.Orders = append(queryOpts.Orders, order)
		}
	}
}

func ApplyFilters(queryParams map[string][]string, mappings map[string]string, queryOpts *storage.QueryOptions) {
	if queryOpts == nil {
		return
	}

	for queryParamKey, queryOptsColumn := range mappings {
		filterValue, ok := queryParams[queryParamKey]

		if ok {
			if queryOpts.WhereColumns == nil {
				queryOpts.WhereColumns = map[string]any{
					queryOptsColumn: filterValue,
				}
			} else {
				queryOpts.WhereColumns[queryOptsColumn] = filterValue
			}
		}
	}
}
