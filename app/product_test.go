// When you want to test the internal functions of the package, you should use the same name of the package you'll test
// so you can have access to the private methods from that package.
package app_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/gscarv13/fc-hex-architecture/app"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := app.Product{}
	product.Name = "Product1"
	product.Status = app.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()

	// This is a 'console.log' equivalent in golang
	// to see the output in the console run go test -v ./...
	// t.Log(err.Error())

	require.Equal(t, "the price must be grater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := app.Product{}
	product.Name = "Product1"
	product.Status = app.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be zero in order to disable the product", err.Error())

}

func TestProduct_IsValid(t *testing.T) {
	product := app.Product{}
	product.ID = uuid.New().String()
	product.Name = "Hey"
	product.Status = app.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = app.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal to zero", err.Error())
}
