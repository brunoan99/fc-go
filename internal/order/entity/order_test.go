package entity_test

import (
	"testing"

	"github.com/brunoan99/fullcycle-go/internal/order/entity"
	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyId_WhenCreateANewOrder_ThenShouldReceiveAndError(t *testing.T) {
	order := entity.Order{}
	assert.Error(t, order.IsValid(), "invalid id")
}

func TestGivenAnEmptyPrice_WhenCreateANewOrder_ThenShouldReceiveAndError(t *testing.T) {
	order := entity.Order{
		ID: "abcdefg",
	}
	assert.Error(t, order.IsValid(), "invalid price")
}

func TestGivenAnEmptyTax_WhenCreateANewOrder_ThenShouldReceiveAndError(t *testing.T) {
	order := entity.Order{
		ID:    "abcdefg",
		Price: 10,
	}
	assert.Error(t, order.IsValid(), "invalid tax")
}

func TestGivenAValidParams_WhenCreateANewOrder_ThenShouldRecieveCreateORderWithAllParams(t *testing.T) {
	order, err := entity.NewOrder("123", 10, 2)

	assert.NoError(t, err)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
}

func TestGivenAValidParams_WhenCallCalculateFinalPrice_ThenShouldCalculateFinaPriceAndSetItOnFinalPriceProrperty(t *testing.T) {
	order, err := entity.NewOrder("123", 10, 2)
	assert.NoError(t, err)
	err = order.CalculateFinalPrice()
	assert.NoError(t, err)
	assert.Equal(t, 12.0, order.FinalPrice)
}
