package delete_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	delUrl "github.com/Prrromanssss/URLShortener/internal/http-server/handlers/url/delete"
	"github.com/Prrromanssss/URLShortener/internal/http-server/handlers/url/delete/mocks"
	"github.com/Prrromanssss/URLShortener/internal/lib/logger/handlers/slogdiscard"
)

func TestDeleteHandler(t *testing.T) {
	cases := []struct {
		name      string
		alias     string
		respError string
		mockError error
	}{
		{
			name:  "Success",
			alias: "test_alias",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			urlDeletterMock := mocks.NewURLDeletter(t)

			if tc.respError == "" || tc.mockError != nil {
				urlDeletterMock.On("DeleteURL", tc.alias).
					Return(nil).
					Once()
			}

			r := chi.NewRouter()
			r.Delete("/{alias}", delUrl.New(slogdiscard.NewDiscardLogger(), urlDeletterMock))

			ts := httptest.NewServer(r)
			defer ts.Close()

			req, err := http.NewRequest("DELETE", ts.URL+"/"+tc.alias, nil)
			require.NoError(t, err)

			resp, err := http.DefaultClient.Do(req)
			require.NoError(t, err)
			defer resp.Body.Close()

			assert.Equal(t, http.StatusNoContent, resp.StatusCode)

		})
	}
}
