package storage

import (
	"github.com/hashicorp/go-memdb"
	"github.com/nitwhiz/api2go/v2/jsonapi"
)

type Model interface {
	jsonapi.MarshalIdentifier
	TableName() string
	Indexes() map[string]*memdb.IndexSchema
	SearchIndexContents() []string
}
