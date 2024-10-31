package storage

import (
	"fmt"
	"github.com/hashicorp/go-memdb"
	"reflect"
	"testing"
)

type TestModel struct {
	A string
	C []string
}

func (t TestModel) GetID() string {
	return "1"
}

func (t TestModel) TableName() string {
	return "tests"
}

func (t TestModel) Indexes() map[string]*memdb.IndexSchema {
	return map[string]*memdb.IndexSchema{}
}

func (t TestModel) SearchIndexContents() []string {
	var res []string

	for _, s := range t.C {
		res = append(res, fmt.Sprintf("v:%s", s))
	}

	return res
}

func TestSearchIndex_FromObject(t *testing.T) {
	type args struct {
		raw interface{}
	}

	tests := []struct {
		name    string
		args    args
		want    bool
		want1   [][]byte
		wantErr bool
	}{
		{
			name: "basic",
			args: args{
				raw: &TestModel{
					A: "Something",
					C: []string{"A", "B"},
				},
			},
			want: true,
			want1: [][]byte{
				{'v', ':', 'A', '_', 0},
				{'_', 'v', ':', 'B', 0},
				{'v', ':', 'A', '_', 'v', ':', 'B', 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			se := SearchIndex{}
			got, got1, err := se.FromObject(tt.args.raw)

			if (err != nil) != tt.wantErr {
				t.Errorf("FromObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FromObject() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("FromObject() got1 = %v, want %v", got1, tt.want1)

				for _, got1Element := range got1 {
					t.Log("got: " + string(got1Element))
				}

				for _, want1 := range tt.want1 {
					t.Log("want: " + string(want1))
				}
			}
		})
	}
}
