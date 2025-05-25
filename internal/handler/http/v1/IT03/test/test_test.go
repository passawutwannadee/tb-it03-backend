package it03apiv1_test

import (
	"testing"

	it03apiv1 "github.com/passawutwannadee/tb-it03/internal/handler/http/v1/IT03"
	it03_mock "github.com/passawutwannadee/tb-it03/internal/usecase/it03/mock"
	"github.com/stretchr/testify/suite"
)

type IT03HandlerTestSuite struct {
	suite.Suite
	mockIT03 *it03_mock.MockUseCase
	handler  *it03apiv1.Handler
}

func (t *IT03HandlerTestSuite) SetupTest() {
	// Setup code here
	t.mockIT03 = &it03_mock.MockUseCase{}
	t.handler = it03apiv1.New(it03apiv1.Dependencies{
		IT03: t.mockIT03,
	})
}

func TestIT03HandlerTestSuite(t *testing.T) {
	suite.Run(t, new(IT03HandlerTestSuite))
}
