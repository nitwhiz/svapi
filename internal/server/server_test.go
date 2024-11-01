package server

import (
	"errors"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"os"
	"path"
	"testing"
)

type testRequest struct {
	name        string
	url         string
	queryParams map[string]string
	fixture     string
}

func runTest(t *testing.T, req testRequest) {
	t.Run(req.name, func(t *testing.T) {
		t.Parallel()

		apitest.
			New().
			Handler(router).
			Get(req.url).
			QueryParams(req.queryParams).
			Expect(t).
			Status(http.StatusOK).
			Assert(func(response *http.Response, _ *http.Request) error {
				body, err := io.ReadAll(response.Body)

				if err != nil {
					return err
				}

				exp, err := os.ReadFile(path.Join("../../testdata/fixtures/items/", req.fixture))

				if err != nil {
					return err
				}

				if !assert.JSONEq(t, string(exp), string(body)) {
					return errors.New("mismatching body")
				}

				return nil
			}).
			End()
	})
}

func runBenchmark(b *testing.B, req testRequest) {
	b.Run(req.name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			apitest.
				New().
				Handler(router).
				Get(req.url).
				QueryParams(req.queryParams).
				Expect(b).
				Assert(func(response *http.Response, _ *http.Request) error {
					b.StopTimer()
					defer b.StartTimer()

					assert.Equal(b, http.StatusOK, response.StatusCode)

					body, err := io.ReadAll(response.Body)

					if err != nil {
						return err
					}

					exp, err := os.ReadFile(path.Join("../../testdata/fixtures/items/", req.fixture))

					if err != nil {
						return err
					}

					assert.JSONEq(b, string(exp), string(body))

					return nil
				}).
				End()
		}
	})
}

func TestServer(t *testing.T) {
	if err := Init(false); err != nil {
		t.Fatal(err)
	}

	t.Run("Items", testItems)
}

func BenchmarkServer(b *testing.B) {
	if err := Init(false); err != nil {
		b.Fatal(err)
	}

	b.Run("Items", benchmarkItems)
}
