package it03apiv1_test

import (
	"context"
	"net/http"
	"net/http/httptest"

	webv1 "github.com/passawutwannadee/tb-it03/internal/handler/http/v1/web"
	postgresrepo "github.com/passawutwannadee/tb-it03/internal/repo/postgres"
	"github.com/passawutwannadee/tb-it03/internal/usecase/it03"
	"github.com/passawutwannadee/tb-it03/internal/util"
	"github.com/stretchr/testify/assert"
)

func (t *IT03HandlerTestSuite) TestList() {
	tests := []struct {
		name           string
		expectList     bool
		mockList       *it03.PaginatedList
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:       "success",
			expectList: true,
			mockList: &it03.PaginatedList{
				Lists: []postgresrepo.IT03ListRow{
					{
						ID:       1,
						Name:     "test1",
						Reason:   "test1",
						StatusID: 1,
						Status:   util.StrPtr("รออนุมัติ"),
					},
				},
			},
			mockError:      nil,
			expectedStatus: 200,
			expectedBody: func() string {

				res := it03.PaginatedList{
					Lists: []postgresrepo.IT03ListRow{
						{
							ID:       1,
							Name:     "test1",
							Reason:   "test1",
							StatusID: 1,
							Status:   util.StrPtr("รออนุมัติ"),
						},
					},
				}

				return webv1.CreateExpectedResponse(res)
			}(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func() {
			if test.expectList {
				t.mockIT03.EXPECT().OffsetList(context.Background()).
					Return(test.mockList, test.mockError).
					Once()
			}

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()

			t.handler.List(w, req)
			assert.Equal(t.T(), test.expectedStatus, w.Code)
			assert.Equal(t.T(), test.expectedBody, w.Body.String())

			t.mockIT03.AssertExpectations(t.T())
		})
	}
}
