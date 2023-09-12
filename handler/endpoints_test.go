package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
)

func TestHello(t *testing.T) {

}
func TestGetProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepository := repository.NewMockRepositoryInterface(ctrl)
	mockRepository.EXPECT().GetProfile(gomock.Any(), gomock.Any()).Return(repository.User{}, nil)
	s := Server{
		Repository: mockRepository,
	}
	token := "xxx"
	recorder := httptest.NewRecorder()
	err := s.GetProfile(echo.New().NewContext(httptest.NewRequest(echo.GET, "/profile?token=xxx", nil), recorder), generated.GetProfileParams{
		Token: &token,
	})
	if err != nil {
		t.Fatal("should not return err")
	}
	if recorder.Result().StatusCode != http.StatusOK {
		t.Fatal("unexpectetd status response")
	}
}

func TestGetProfile_err(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepository := repository.NewMockRepositoryInterface(ctrl)
	mockRepository.EXPECT().GetProfile(gomock.Any(), gomock.Any()).Return(repository.User{}, errors.New("test"))
	s := Server{
		Repository: mockRepository,
	}
	token := "xxx"
	recorder := httptest.NewRecorder()
	err := s.GetProfile(echo.New().NewContext(httptest.NewRequest(echo.GET, "/profile?token=xxx", nil), recorder), generated.GetProfileParams{
		Token: &token,
	})
	if err != nil {
		t.Fatal("should not return err")
	}
	if recorder.Result().StatusCode != http.StatusForbidden {
		t.Fatal("unexpectetd status response")
	}
}

func TestPutProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepository := repository.NewMockRepositoryInterface(ctrl)
	mockRepository.EXPECT().UpdateProfile(gomock.Any(), gomock.Any(), gomock.Any()).Return(repository.User{}, nil)
	s := Server{
		Repository: mockRepository,
	}
	token := "xxx"
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(echo.PUT, "/profile?token=xxx", bytes.NewBufferString(`{
		"fullName": "test1234",
		"id": 3,
		"phoneNumbers": "+6285864632842"
	}`))
	req.Header.Add("Content-Type", "application/json")
	err := s.PutProfile(echo.New().NewContext(req, recorder), generated.PutProfileParams{
		Token: &token,
	})
	if err != nil {
		t.Fatal("should not return err")
	}
	if recorder.Result().StatusCode != http.StatusOK {
		t.Fatal("unexpectetd status response")
	}
}

func TestPutProfile_err(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepository := repository.NewMockRepositoryInterface(ctrl)
	mockRepository.EXPECT().UpdateProfile(gomock.Any(), gomock.Any(), gomock.Any()).Return(repository.User{}, errors.New("test"))
	s := Server{
		Repository: mockRepository,
	}
	token := "xxx"
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(echo.PUT, "/profile?token=xxx", bytes.NewBufferString(`{
		"fullName": "test1234",
		"id": 3,
		"phoneNumbers": "+6285864632842"
	}`))
	req.Header.Add("Content-Type", "application/json")
	err := s.PutProfile(echo.New().NewContext(req, recorder), generated.PutProfileParams{
		Token: &token,
	})
	if err != nil {
		t.Fatal("should not return err")
	}
	if recorder.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("unexpectetd status response")
	}
}

func TestPutProfile_err_bind(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepository := repository.NewMockRepositoryInterface(ctrl)
	s := Server{
		Repository: mockRepository,
	}
	token := "xxx"
	recorder := httptest.NewRecorder()
	err := s.PutProfile(echo.New().NewContext(httptest.NewRequest(echo.PUT, "/profile?token=xxx", bytes.NewBufferString(`xxx`)), recorder), generated.PutProfileParams{
		Token: &token,
	})
	if err != nil {
		t.Fatal("should not return err")
	}
	if recorder.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("unexpectetd status response")
	}
}

func TestRegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepository := repository.NewMockRepositoryInterface(ctrl)
	mockRepository.EXPECT().RegisterUser(gomock.Any(), gomock.Any()).Return(repository.User{}, nil)
	s := Server{
		Repository: mockRepository,
	}
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(echo.POST, "/register", bytes.NewBufferString(`{
		"fullName": "test1234",
		"id": 3,
		"phoneNumbers": "+6285864632842"
	}`))
	req.Header.Add("Content-Type", "application/json")
	err := s.RegisterUser(echo.New().NewContext(req, recorder))
	if err != nil {
		t.Fatal("should not return err")
	}
	if recorder.Result().StatusCode != http.StatusOK {
		t.Fatal("unexpectetd status response")
	}
}

func TestRegisterUser_err(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepository := repository.NewMockRepositoryInterface(ctrl)
	mockRepository.EXPECT().RegisterUser(gomock.Any(), gomock.Any()).Return(repository.User{}, errors.New("test"))
	s := Server{
		Repository: mockRepository,
	}
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(echo.POST, "/register", bytes.NewBufferString(`{
		"fullName": "test1234",
		"id": 3,
		"phoneNumbers": "+6285864632842"
	}`))
	req.Header.Add("Content-Type", "application/json")
	err := s.RegisterUser(echo.New().NewContext(req, recorder))
	if err != nil {
		t.Fatal("should not return err")
	}
	if recorder.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("unexpectetd status response")
	}
}

func TestRegisterUser_err_bind(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepository := repository.NewMockRepositoryInterface(ctrl)
	s := Server{
		Repository: mockRepository,
	}
	recorder := httptest.NewRecorder()
	err := s.RegisterUser(echo.New().NewContext(httptest.NewRequest(echo.POST, "/register", bytes.NewBufferString(`xxx`)), recorder))
	if err != nil {
		t.Fatal("should not return err")
	}
	if recorder.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("unexpectetd status response")
	}
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepository := repository.NewMockRepositoryInterface(ctrl)
	mockRepository.EXPECT().Login(gomock.Any(), gomock.Any()).Return(repository.User{}, nil)
	s := Server{
		Repository: mockRepository,
	}
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(echo.POST, "/login", bytes.NewBufferString(`{
		"phoneNumbers": "+6285864632842",
		"password": "test123"
	}`))
	req.Header.Add("Content-Type", "application/json")
	err := s.LoginUser(echo.New().NewContext(req, recorder))
	if err != nil {
		t.Fatal("should not return err")
	}
	if recorder.Result().StatusCode != http.StatusOK {
		t.Fatal("unexpectetd status response")
	}
}

func TestLogin_err(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepository := repository.NewMockRepositoryInterface(ctrl)
	mockRepository.EXPECT().Login(gomock.Any(), gomock.Any()).Return(repository.User{}, errors.New("test"))
	s := Server{
		Repository: mockRepository,
	}
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(echo.POST, "/login", bytes.NewBufferString(`{
		"phoneNumbers": "+6285864632842",
		"password": "test123"
	}`))
	req.Header.Add("Content-Type", "application/json")
	err := s.LoginUser(echo.New().NewContext(req, recorder))
	if err != nil {
		t.Fatal("should not return err")
	}
	if recorder.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("unexpectetd status response")
	}
}

func TestLogin_err_bind(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepository := repository.NewMockRepositoryInterface(ctrl)
	s := Server{
		Repository: mockRepository,
	}
	recorder := httptest.NewRecorder()
	err := s.LoginUser(echo.New().NewContext(httptest.NewRequest(echo.POST, "/login", bytes.NewBufferString(`xxx`)), recorder))
	if err != nil {
		t.Fatal("should not return err")
	}
	if recorder.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("unexpectetd status response")
	}
}
