package delete

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/Prrromanssss/URLShortener/internal/lib/logger/sl"
	"github.com/Prrromanssss/URLShortener/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=URLDeletter
type URLDeletter interface {
	DeleteURL(alias string) error
}

func New(log *slog.Logger, urlDeletter URLDeletter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.delete.New"

		log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info("alias is empty")
			render.JSON(w, r, "invalid request")

			return
		}

		err := urlDeletter.DeleteURL(alias)
		if errors.Is(err, storage.ErrURLNotFound) {
			log.Info("url not found", "alias", alias)
			render.JSON(w, r, "not found")

			return
		}
		if err != nil {
			log.Info("failed to delete url", sl.Err(err))
			render.JSON(w, r, "internal error")

			return
		}

		log.Info("delete url")

		render.Status(r, http.StatusNoContent)
		render.JSON(w, r, nil)
	}
}
