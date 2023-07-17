package handler

import (
	"github.com/golang/mock/gomock"
	"testing"
	"zangetsu/internal/domain/entity"
	mock_service "zangetsu/internal/domain/service/mocks"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(s *mock_service.MockIUserService, user entity.UserViewModel)

	testTable := []struct {
		name                string
		inputBody           string
		inputUser           entity.UserViewModel
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"firstName":"Test","secondName":"Test","email":"test@gmail.com","password":"test"}`,
			inputUser: entity.UserViewModel{
				FirstName:  "Test",
				SecondName: "Test",
				Email:      "test@gmail.com",
				Password:   "test",
			},
			mockBehavior: func(s *mock_service.MockIUserService, user entity.UserViewModel) {
				s.EXPECT().SignUp(user).Return(entity.UserViewModel{FirstName: "Test", SecondName: "Test", Email: "test@gmail.com", Password: "test"}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"firstName":"Test","secondName":"Test","email":"test@gmail.com","password":"test"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			iUserService := mock_service.NewMockIUserService(c)
			testCase.mockBehavior(iUserService, testCase.inputUser)

			//services := &service.Service{service: iUserService}
		})
	}
}
