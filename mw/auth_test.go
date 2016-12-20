package mw

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/praelatus/backend/models"
)

func TestGetToken(t *testing.T) {
	bearer, _ := http.NewRequest("GET", "/", nil)
	bearer.Header.Set("Authorization", "Bearer TESTTOKEN")

	token, _ := http.NewRequest("GET", "/", nil)
	token.Header.Set("Authorization", "Token TESTTOKEN")

	tk := getToken(bearer)
	if tk != "TESTTOKEN" {
		t.Errorf("Expected TESTTOKEN Got %s", tk)
	}

	tk = ""
	tk = getToken(token)
	if tk != "TESTTOKEN" {
		t.Errorf("Expected TESTTOKEN Got %s", tk)
	}

	fail, _ := http.NewRequest("", "/", nil)

	tk = ""
	tk = getToken(fail)
	if tk != "" {
		t.Errorf("Expected \"\" Got %s", tk)
	}
}

func TestValidateToken(t *testing.T) {
	u, e := models.NewUser("testuser", "test", "Test Testerson",
		"test@example.com", false)
	if e != nil {
		t.Error(e)
	}

	token, e := JWTSignUser(*u)
	if e != nil {
		t.Error(e)
	}

	tokenUser := validateToken(token)
	if tokenUser == nil {
		t.Error("Expected a user got nil instead")
		t.FailNow()
	}

	if tokenUser.Username != u.Username {
		t.Errorf("Expected %s Got %s; %s", u.Username, tokenUser.Username, tokenUser)
	}
}

type mockAuthHandler struct{}

func (m mockAuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u := GetUser(r.Context())
	w.Write([]byte(u.String()))
}

func TestAuth(t *testing.T) {
	u, e := models.NewUser("testuser", "test", "Test Testerson",
		"test@example.com", false)
	if e != nil {
		t.Error(e)
	}

	token, e := JWTSignUser(*u)
	if e != nil {
		t.Error(e)
	}

	mock := mockAuthHandler{}
	auth := Auth(mock)

	r, e := http.NewRequest("GET", "/", nil)
	if e != nil {
		t.Fatal(e)
	}

	r.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()

	auth.ServeHTTP(w, r)

	var user models.User

	e = json.Unmarshal(w.Body.Bytes(), &user)
	if e != nil {
		t.Error(e)
	}

	if user.Username != u.Username {
		t.Errorf("Expected %s Got %s", u.Username, user.Username)
	}
}

func TestGetUser(t *testing.T) {
	u := models.User{Username: "testuser"}
	ctx := context.WithValue(context.Background(), currentUser, &u)

	tu := GetUser(ctx)
	if tu.Username != u.Username {
		t.Errorf("Expected %s Got %s", u.Username, tu.Username)
	}
}
