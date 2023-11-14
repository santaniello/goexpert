package tax

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestCalculateTaxParameterized(t *testing.T) {
	type taxParameterized struct {
		amount, expected float64
	}

	scenariosTests := []taxParameterized{
		{500, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
	}

	for _, scenario := range scenariosTests {
		var result = CalculateTax(scenario.amount)
		if result != scenario.expected {
			t.Errorf("Expected %f but got %f", scenario.expected, result)
		}
	}
}

// go test -bench=.
// go test -bench=. -count 5 -run=^# -benchtime=5s
func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}

// go test -fuzz=. -fuzztime=5s
func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Reveived %f but expected 0", result)
		}
		if amount > 20000 && result != 20 {
			t.Errorf("Reveived %f but expected 20", result)
		}
	})
}

// Testes usando a lib testify e recursos de assert
func TestCalculateTax3(t *testing.T) {
	tax, err := CalculateTax3(1000.0)
	assert.Nil(t, err)
	assert.Equal(t, 10.0, tax)

	tax, err = CalculateTax3(0)
	assert.Error(t, err, "amount must be greater than 0")
	assert.Equal(t, 0.0, tax)
	assert.Contains(t, err.Error(), "greater than 0")
}

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil).Twice()
	repository.On("SaveTax", 0.0).Return(errors.New("error saving tax"))

	err := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)
	err = CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err, "error saving tax")

	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "SaveTax", 3)
}
