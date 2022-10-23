package handler

import (
	"net/http"
)

type GetTask struct {
	Service GetTaskService
}

func (gt *GetTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	taskID := r.Context().Value("taskID").(int)
	task, err := gt.Service.GetTask(ctx, taskID)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	RespondJSON(ctx, w, task, http.StatusOK)
}
