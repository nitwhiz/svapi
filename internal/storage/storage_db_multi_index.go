package storage

import (
	"fmt"
	"github.com/nitwhiz/svapi/pkg/util"
)

type SearchIndex struct {
}

func (SearchIndex) FromArgs(args ...interface{}) ([]byte, error) {
	if len(args) == 0 {
		return []byte{}, nil
	}

	if len(args) != 1 {
		return nil, fmt.Errorf("wrong number of args %d, expected 1", len(args))
	}

	indexValue, ok := args[0].(string)

	if !ok {
		return nil, fmt.Errorf("wrong type for arg %T, expected string", args[0])
	}

	return append([]byte(indexValue), 0), nil
}

func (SearchIndex) FromObject(raw interface{}) (bool, [][]byte, error) {
	m, ok := raw.(Model)

	if !ok {
		return false, nil, fmt.Errorf("wrong type for arg %T, expected Model", raw)
	}

	var result [][]byte

	searchIndexCombinations := util.Combinations(m.SearchIndexContents()...)

	for i := range searchIndexCombinations {
		result = append(result, append([]byte(searchIndexCombinations[i]), 0))
	}

	return true, result, nil
}
