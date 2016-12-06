package user

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/shirou/swagger_sample/models"
)

func TestSearch(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	params := GetSearchParams{
		HTTPRequest: req,
		UserID:      swag.Int64(10),
	}
	r := Search(params)
	w := httptest.NewRecorder()
	r.WriteResponse(w, runtime.JSONProducer())
	if 200 != w.Code {
		t.Error("status code")
	}
	var a models.User
	err = json.Unmarshal(w.Body.Bytes(), &a)
	if err != nil {
		t.Error("unmarshal")
	}
	if 10 == swag.Int64Value(a.UserID) {
		t.Error("wrong user id")
	}
}
