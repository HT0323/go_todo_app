package handler

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HT0323/go_todo_app/entity"
	"github.com/HT0323/go_todo_app/testutil"
	"github.com/go-playground/validator/v10"
)

func TestRegisterUser(t *testing.T) {
	t.Parallel()
	type want struct {
		status  int
		rspFile string
	}
	tests := map[string]struct {
		reqFile string
		want    want
	}{
		"ok": {
			reqFile: "testdata/register_user/ok_req.json.golden",
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/register_user/ok_rsp.json.golden",
			},
		},
		"name_duplicate_ng": {
			reqFile: "testdata/register_user/name_duplicate_bad_req.json.golden",
			want: want{
				status:  http.StatusInternalServerError,
				rspFile: "testdata/register_user/name_duplicate_bad_rsp.json.golden",
			},
		},
		"para_ng": {
			reqFile: "testdata/register_user/ng_para_bad_req.json.golden",
			want: want{
				status:  http.StatusBadRequest,
				rspFile: "testdata/register_user/ng_para_bad_rsp.json.golden",
			},
		},
	}
	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodPost,
				"/register",
				bytes.NewReader(testutil.LoadFile(t, tt.reqFile)),
			)

			moq := &RegisterUserServiceMock{}
			moq.RegisterUserFunc = func(
				ctx context.Context, name string, password string, role string,
			) (*entity.User, error) {
				if tt.want.status == http.StatusOK {
					return &entity.User{ID: 1}, nil
				}
				if tt.want.status == http.StatusInternalServerError {
					return nil, fmt.Errorf("failed to register: cannot create same name user: duplicate entry")
				}

				return nil, errors.New("error from mock")
			}

			sut := RegisterUser{
				Service:   moq,
				Validator: validator.New(),
			}
			sut.ServeHTTP(w, r)

			resp := w.Result()
			testutil.AssertResponse(t, resp, tt.want.status, testutil.LoadFile(t, tt.want.rspFile))
		})
	}

}
