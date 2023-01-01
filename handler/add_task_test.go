package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"todo_app/entity"
	"todo_app/store"
	"todo_app/testutil"
)

func TestAddTask(t *testing.T) {
	type want struct {
		status  int
		rspFile string
	}
	tests := map[string]struct {
		reqFile string
		want    want
	}{
		"ok": {
			reqFile: "testdata/add_task/ok_req.json.golden",
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/add_task/ok_rsp.json.golden",
			},
		},
		"badRequest": {
			reqFile: "testdata/add_task/bad_req.json.golden",
			want: want{
				status:  http.StatusBadRequest,
				rspFile: "testdata/add_task/bad_rsp.json.golden",
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
				"/tasks",
				bytes.NewReader(testutil.LoadFile(t, tt.reqFile)),
			)

			moq := &store.MockStore{}
			moq.AddTaskFunc = func(
				ctx entity.Context, title string
				) (entity.Task, error) {
				if tt.want.status == http.StatusOK {
					return &entity.Task{ID: 1}, nil
				}
				return nil, errors.New("error from mock")
			}
			sut := AddTask{
				Service: moq,
				Validator: validator.New(),
			}
			sut.ServeHTTP(w, r)

			resp := w.Result()
			testutil.AssertResponse(t,
				resp, tt.want.status, testutil.LoadFile(t, tt.want.rspFile),
			)
		})
	}
}
