package modules

import (
	"net/http/httptest"
	"testing"

	"github.com/adiet95/go-cart/src/database"
	"github.com/adiet95/go-cart/src/libs"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repo = RepoMock{mock.Mock{}}
var service = NewService(&repo)
var ctrl = NewCtrl(service)

func TestCtrlGetAll(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()

	var dataMocks = database.Carts{
		{ProductName: "Mobil Langkah", Quantity: 1},
		{ProductName: "Motor", Quantity: 3},
	}

	repo.mock.On("GetCart").Return(&dataMocks, nil)

	req := httptest.NewRequest("GET", "/test/get", nil)

	r.GET("/test/user", ctrl.GetCart)

	r.ServeHTTP(w, req)

	var user *database.Carts
	respon := libs.Response{
		Data: &user,
	}
	assert.False(t, respon.IsError, "error must be false")
}

func TestCtrlAddCart(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	var dataMock = database.Cart{ProductName: "Motor", Quantity: 3}

	repo.mock.On("FindCart", "Motor").Return(&dataMock, nil)
	repo.mock.On("UpdateCart", &dataMock, "Motor").Return(&dataMock, nil)

	req := httptest.NewRequest("POST", "/test/add", nil)

	r.POST("/test/user", ctrl.AddCart)

	r.ServeHTTP(w, req)

	var user *database.Carts
	respon := libs.Response{
		Data: &user,
	}
	assert.False(t, respon.IsError, "error must be false")
	assert.Equal(t, "", respon.Status)

}

func TestCtrlDeleteCart(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	var dataMock = database.Cart{ProductName: "Motor", Quantity: 3, CartId: "12345"}
	repo.mock.On("DeleteUser", "12345").Return(&dataMock, nil)

	req := httptest.NewRequest("DELETE", "/test/delete", nil)

	r.DELETE("/test/delete", ctrl.DeleteCart)

	r.ServeHTTP(w, req)

	var user *database.Carts
	respon := libs.Response{
		Data: &user,
	}
	assert.False(t, respon.IsError, "error must be false")
	assert.Equal(t, "", respon.Status)

}
