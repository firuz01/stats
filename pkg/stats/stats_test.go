package stats

import (
	"github.com/firuz01/bank/v2/pkg/types"
	"reflect"
	"testing"
	
)

func TestCategoriesAvg_nil(t *testing.T) {
	var payments []types.Payment
	result := CategoriesAvg(payments)

	if len(result) != 0 {
		t.Error("result len != 0 ")
	}
}

func TestCategoriesAvg_empty(t *testing.T) {
	payments := []types.Payment {}
	result := CategoriesAvg(payments)

	if len(result) != 0 {
		t.Error("result len != 0 ")
	}
}

func TestCategoriesAvg_FoundONEResult(t *testing.T) {
	payments := []types.Payment {
		{Amount: 1000, Category: "auto", Status: "FAIL"},
		{Amount: 2000, Category: "pharmacy", Status: "FAIL"},
		{Amount: 3000, Category: "shop", Status: "OK"},
	}
	expected := map [types.Category] types.Money {
		"shop": 3000,
	} 
	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}
func TestCategoriesAvg_FoundMoreResult(t *testing.T) {
	payments := []types.Payment {
		{Amount: 2000, Category: "shop", Status: "OK"},
		{Amount: 2000, Category: "pharmacy", Status: "OK"},
		{Amount: 3000, Category: "shop", Status: "OK"},
		{Amount: 4000, Category: "shop", Status: "OK"},
	}
	expected := map [types.Category] types.Money {
		"shop": 3000,
		"pharmacy": 2000,
	} 
	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestPeriodsDynamic_empty(t *testing.T) {
	first := map [types.Category] types.Money {}
	second := map [types.Category] types.Money {}
	expected := map [types.Category] types.Money {}
	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestPeriodsDynamic_firstempty(t *testing.T) {
	first := map [types.Category] types.Money {}
	second := map [types.Category] types.Money {
		"auto": 1000,
	}
	expected := map [types.Category] types.Money {
		"auto": 1000,
	}
	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestPeriodsDynamic_secondempty(t *testing.T) {
	first := map [types.Category] types.Money {
		"auto": 1000,
	}
	second := map [types.Category] types.Money {}
	expected := map [types.Category] types.Money {
		"auto": -1000,
	}
	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestPeriodsDynamic_full(t *testing.T) {
	first := map [types.Category] types.Money {
		"pharm": 1000,
	}
	second := map [types.Category] types.Money {
		"pharm": 2000,
	}
	expected := map [types.Category] types.Money {
		"pharm": 1000,
	}
	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}