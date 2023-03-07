package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIDIsEmpty(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "invalid id")
}

func TestPriceIsEmpty(t *testing.T) {
	order := Order{ID: "1"}
	assert.Error(t, order.Validate(), "invalid id")
}

func TestTaxIsEmpty(t *testing.T) {
	order := Order{ID: "1", Price: 20}
	assert.Error(t, order.Validate(), "invalid id")
}

func TestAllParamsIsValid(t *testing.T) {
	order := Order{ID: "1", Price: 20, Tax: 5}
	assert.NoError(t, order.Validate())
	assert.Equal(t, 20.0, order.Price)
	order.CalculateFinalPrice()
	assert.Equal(t, 25.0, order.FinalPrice)
}
