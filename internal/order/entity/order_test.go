package entity_test

import (
	"testing"

	"github.com/brunoan99/fullcycle-go/internal/order/entity"
	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyId_WhenCreateANewOrder_ThenShouldReceiveAndError(t *testing.T) {
	order := entity.Order{}
	assert.Error(t, order.IsValid(), "invalid price")
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

func TestGivenAnValidParams_WhenCreateANewOrder_ThenShouldRecieveCreateORderWithAllParams(t *testing.T) {
	order, err := entity.NewOrder("123", 10, 2)

	assert.Equal(t, err, nil)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
	assert.Equal(t, 10.2, order.FinalPrice)
}
