package middlewares

import (
	"context"
	"net/http"

	"github.com/HT0323/go_todo_app/handler"
	"github.com/go-chi/chi/v5"
)

func TaskCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		// var ctx context.Context
		var taskID int
		ctx := r.Context()

		if pathParaTaskID := chi.URLParam(r, "taskID"); pathParaTaskID != "" {
			taskID, err = searchTask(pathParaTaskID)

			if err != nil {
				handler.RespondJSON(ctx, w, &handler.ErrResponse{
					Message: err.Error(),
				}, http.StatusNotFound)
				return
			}
		} else {
			handler.RespondJSON(ctx, w, &handler.ErrResponse{
				Message: err.Error(),
			}, http.StatusNotFound)
			return
		}
		if err != nil {
			handler.RespondJSON(ctx, w, &handler.ErrResponse{
				Message: err.Error(),
			}, http.StatusNotFound)
			return
		}

		ctx = context.WithValue(r.Context(), "taskID", taskID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
