package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_If_It_Gets_And_Error_If_ID_Is_Blank(t *testing.T) {
	order := Order{}

	assert.Error(t, order.Validate(), "ID is required")
}

func Test_If_It_Gets_And_Error_If_Price_Is_Less_Than_Zero(t *testing.T) {
	order := Order{
		ID:    "1",
		Price: -1,
	}

	assert.Error(t, order.Validate(), "Price must be greater than zero")
}

func Test_If_It_Gets_And_Error_If_Tax_Is_Less_Than_Zero(t *testing.T) {
	order := Order{
		ID:    "1",
		Price: 10,
		Tax:   -1,
	}

	assert.Error(t, order.Validate(), "Tax must be greater than zero")
}

func Test_Final_Price(t *testing.T) {
	order := Order{
		ID:    "1",
		Price: 10,
		Tax:   1,
	}

	assert.NoError(t, order.Validate())
	assert.Equal(t, "1", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 1.0, order.Tax)
	order.CalculateFinalPrice()
	assert.Equal(t, 11.0, order.FinalPrice)
}
