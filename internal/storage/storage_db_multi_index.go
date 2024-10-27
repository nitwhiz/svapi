package storage

import (
	"fmt"
	"github.com/nitwhiz/svapi/pkg/util"
	"strings"
)

type SearchIndex struct {
}

func (SearchIndex) FromArgs(args ...interface{}) ([]byte, error) {
	// this only searches for the very first value of SearchIndexContents()!

	// args is the output of SearchIndexContents. how to use this if we're using [][]byte?

	if len(args) == 0 {
		return []byte{}, nil
	}

	if len(args) != 1 {
		return nil, fmt.Errorf("wrong number of args %d, expected 1", len(args))
	}

	indexValues, ok := args[0].([][]string)

	if !ok {
		return nil, fmt.Errorf("wrong type for arg %T, expected [][]string", args[0])
	}

	if len(indexValues) == 0 {
		return []byte{}, nil
	}

	return append([]byte(strings.Join(indexValues[0], "_")), 0), nil
}

func (SearchIndex) FromObject(raw interface{}) (bool, [][]byte, error) {
	m, ok := raw.(Model)

	if !ok {
		return false, nil, fmt.Errorf("wrong type for arg %T, expected Model", raw)
	}

	var result [][]byte

	for _, idxContents := range m.SearchIndexContents() {
		searchIndexCombinations := util.Combinations(idxContents...)

		for i := range searchIndexCombinations {
			result = append(result, append([]byte(searchIndexCombinations[i]), 0))
		}
	}

	return true, result, nil
}
