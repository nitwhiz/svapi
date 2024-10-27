package storage

import (
	"github.com/hashicorp/go-memdb"
	"github.com/manyminds/api2go/jsonapi"
)

type Model interface {
	jsonapi.MarshalIdentifier
	TableName() string
	Indexes() map[string]*memdb.IndexSchema
	SearchIndexContents() [][]string
}
