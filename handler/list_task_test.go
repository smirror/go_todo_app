package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"todo_app/entity"
	"todo_app/store"
	"todo_app/testutil"
)

func TestListTask(t *testing.T) {
	type want struct {
		status  int
		rspFile string
	}
	tests := map[string]struct {
		tasks *[]*entity.Task
		want  want
	}{
		"ok": {
			tasks: []*entity.Task{
				{
					ID:     1,
					Title:  "test1",
					Status: entity.TaskStatusTodo,
				},
				{
					ID:     2,
					Title:  "test2",
					Status: entity.TaskStatusDone,
				},
			},
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/list_task/ok_rsp.json.golden",
			},
		},
		"empty": {
			tasks: []*entity.Task{},
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/list_task/empty_rsp.json.golden",
			},
		},
	}
	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/tasks", nil)

			sut := ListTask{Store: &store.TaskStore{Tasks: tt.tasks}}
			sut.ServeHTTP(w, r)

			resp := w.Result()
			testutil.AssertResponse(t,
				resp, tt.want.status, testutil.LoadFile(t, tt.want.rspFile),
			)
		})
	}
}
