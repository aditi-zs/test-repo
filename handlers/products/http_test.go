package products

import (
	"bytes"
	"context"
	"github.com/aditi-zs/test-repo/handlers"
	"github.com/aditi-zs/test-repo/models"
	"github.com/aditi-zs/test-repo/services"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gofr.dev/pkg/gofr"
	gofrHTTP "gofr.dev/pkg/gofr/http"
	"gofr.dev/pkg/gofr/http/response"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func getContext(method string, pathParams map[string]string, qParam url.Values, body []byte) *gofr.Context {
	req := httptest.NewRequest(method, "/product", bytes.NewReader(body))
	req.Header.Set("content-type", "application/json")

	//setURLVars are used to set path param
	req = mux.SetURLVars(req, pathParams)

	// to set path param use this
	req.URL.RawQuery = qParam.Encode()

	r := gofrHTTP.NewRequest(req)

	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   r,
		Container: nil,
	}

	ctx.Context = context.Background()

	return ctx
}

func initializeMocks(t *testing.T) (*services.MockProduct, handlers.Product) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProduct := services.NewMockProduct(ctrl)
	product := New(mockProduct)

	return mockProduct, product
}

func TestApplicationHandler_Index(t *testing.T) {
	mockProduct, product := initializeMocks(t)

	qParam := url.Values{"name": {"book"}}

	tests := []struct {
		desc   string
		expOut interface{}
		expErr error
	}{
		{"success", response.Raw{Data: []*models.Product{}}, nil},
	}

	for i, tc := range tests {
		ctx := getContext(http.MethodGet, map[string]string{"name": "book"}, qParam, nil)

		mockProduct.EXPECT().GetAll(ctx, "book").Return(nil, nil)

		output, err := product.Get(ctx)

		assert.Equal(t, tc.expErr, err, "TEST[%d], failed.\n%s", i, tc.desc)

		assert.Equal(t, tc.expOut, output, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}
