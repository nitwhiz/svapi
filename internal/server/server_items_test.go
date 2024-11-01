package server

import (
	"testing"
)

var ItemTestRequests = []testRequest{
	{
		name:        "Item By Internal ID",
		url:         "/v2/items",
		queryParams: map[string]string{"filter[internalId]": "(O)874"},
		fixture:     "item_internal_id_O874.json",
	},
	{
		name:        "Items By 'giftable' Flag",
		url:         "/v2/items",
		queryParams: map[string]string{"filter[flags]": "giftable"},
		fixture:     "items_flags_giftable.json",
	},
	{
		name:        "Items By 'cooking' Type",
		url:         "/v2/items",
		queryParams: map[string]string{"filter[type]": "cooking"},
		fixture:     "items_type_cooking.json",
	},
	{
		name:        "Items By 'giftable' Flag And 'ring' Type",
		url:         "/v2/items",
		queryParams: map[string]string{"filter[type]": "ring", "filter[flags]": "giftable"},
		fixture:     "items_flags_giftable_type_ring.json",
	},
	{
		name:        "Category for Item 003c63d2-c511-5051-b8d4-28cdbc08ad6c",
		url:         "/v2/items/003c63d2-c511-5051-b8d4-28cdbc08ad6c/category",
		queryParams: map[string]string{},
		fixture:     "item_relation_category.json",
	},
}

func testItems(t *testing.T) {
	t.Parallel()

	for _, tr := range ItemTestRequests {
		runTest(t, tr)
	}
}

func benchmarkItems(b *testing.B) {
	for _, tr := range ItemTestRequests {
		runBenchmark(b, tr)
	}
}
