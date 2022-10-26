package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/HT0323/go_todo_app/entity"
	"github.com/HT0323/go_todo_app/testutil"
)

type taskIDKey struct{}

func TestGetTask(t *testing.T) {
	type want struct {
		status  int
		rspFile string
	}
	tests := map[string]struct {
		task entity.Task
		want want
	}{
		"ok": {
			task: entity.Task{
				ID:       1,
				Title:    "test1",
				Status:   entity.TaskStatusTodo,
				Created:  time.Date(2022, 4, 1, 0, 0, 0, 0, time.Local),
				Modified: time.Date(2022, 4, 1, 0, 0, 0, 0, time.Local),
				UserID:   1,
			},
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/get_task/ok_rsp.json.golden",
			},
		},
	}
	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/tasks/1", nil)

			moq := &GetTaskServiceMock{}
			moq.GetTaskFunc =
				func(ctx context.Context, taskID int) (*entity.Task, error) {
					return &tt.task, nil
				}
			ctx := r.Context()
			setCtx := context.WithValue(ctx, "taskID", 1)
			r = r.WithContext(setCtx)
			sut := GetTask{Service: moq}
			sut.ServeHTTP(w, r)

			resp := w.Result()
			testutil.AssertResponse(t,
				resp, tt.want.status, testutil.LoadFile(t, tt.want.rspFile),
			)
		})
	}
}
