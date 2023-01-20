package example

import (
	"context"
	"net/http"

	"github.com/a-h/pathvars"
	"github.com/aviva-verde/offer-api/api/respond"
	"github.com/joerdav/zapray"
	"go.uber.org/zap"
)

var getOffersMatcher = pathvars.NewExtractor("*/example/{message}")

type Handler struct {
	log *zapray.Logger
}

func NewHandler(log *zapray.Logger) (h Handler, err error) {
	h.log = log
	return
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respond.WithNotFound(w)
		return
	}
	if matches, ok := getOffersMatcher.Extract(r.URL); ok {
		h.GetExampleResponse(r.Context(), w, matches["message"])
		return
	}
	respond.WithNotFound(w)
}

type Response struct {
	Message string `json:"message,omitempty"`
}

func (h Handler) GetExampleResponse(ctx context.Context, w http.ResponseWriter, message string) {
	h.log.Info("generating response", zap.String("message", message))
	response := Response{
		Message: message,
	}
	respond.WithJSON(w, response)
}
