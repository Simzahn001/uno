package uno

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a        Value
		b        Value
		expected Value
		wantErr  bool
	}{
		{
			name:     "same powerOfTen and quantity",
			a:        Value{value: 100, powerOfTen: 2, physicalQuantity: 'N'},
			b:        Value{value: 200, powerOfTen: 2, physicalQuantity: 'N'},
			expected: Value{value: 300, powerOfTen: 2, physicalQuantity: 'N'},
			wantErr:  false,
		},
		{
			name:     "different powerOfTen, same quantity",
			a:        Value{value: 1, powerOfTen: 3, physicalQuantity: 'N'},
			b:        Value{value: 1, powerOfTen: 0, physicalQuantity: 'N'},
			expected: Value{value: 1001, powerOfTen: 0, physicalQuantity: 'N'},
			wantErr:  false,
		},
		{
			name:     "different physicalQuantity",
			a:        Value{value: 50, powerOfTen: 1, physicalQuantity: 'N'},
			b:        Value{value: 50, powerOfTen: 1, physicalQuantity: 'J'},
			expected: Value{value: 100, powerOfTen: 1, physicalQuantity: '#'},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Add(tt.a, tt.b)

			if tt.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if result.value != tt.expected.value ||
				result.powerOfTen != tt.expected.powerOfTen ||
				result.physicalQuantity != tt.expected.physicalQuantity {
				t.Errorf("Add() = %+v; want %+v", result, tt.expected)
			}
		})
	}
}

func BenchmarkAdd_SamePower(b *testing.B) {
	a := Value{value: 1000, powerOfTen: 0, physicalQuantity: 'N'}
	c := Value{value: 2000, powerOfTen: 0, physicalQuantity: 'N'}
	for i := 0; i < b.N; i++ {
		_, _ = Add(a, c)
	}
}

func BenchmarkAdd_DifferentPower(b *testing.B) {
	a := Value{value: 1, powerOfTen: 3, physicalQuantity: 'N'}
	c := Value{value: 1000, powerOfTen: 0, physicalQuantity: 'N'}
	for i := 0; i < b.N; i++ {
		_, _ = Add(a, c)
	}
}

func BenchmarkAddManyWideRange(b *testing.B) {
	values := make([]Value, 1000)
	for i := range values {
		values[i] = Value{
			value:            (i + 1) * intPowTest(7, i%5), // große Schwankungen
			powerOfTen:       -5 + (i % 11),                // von -5 bis +5
			physicalQuantity: 'N',
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sum := values[0]
		for j := 1; j < len(values); j++ {
			sum, _ = Add(sum, values[j])
		}
	}
}

func BenchmarkMultiplyManyWideRange(b *testing.B) {
	values := make([]Value, 1000)
	for i := range values {
		values[i] = Value{
			value:            (i + 1) * intPowTest(7, i%5), // große Schwankungen
			powerOfTen:       -5 + (i % 11),                // von -5 bis +5
			physicalQuantity: 'm',
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		product := values[0]
		for j := 1; j < len(values); j++ {
			product, _ = Multiply(product, values[j])
		}
	}
}

func BenchmarkDivideManyWideRange(b *testing.B) {
	values := make([]Value, 1000)
	for i := range values {
		values[i] = Value{
			value:            (i + 1) * intPowTest(7, i%5), // große Schwankungen
			powerOfTen:       -5 + (i % 11),                // von -5 bis +5
			physicalQuantity: 'F',
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		quotient := values[0]
		for j := 1; j < len(values); j++ {
			quotient, _ = Divide(quotient, values[j])
		}
	}
}

// Hilfsfunktion für ganzzahlige Potenzen
func intPowTest(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

func TestGetProductQuantity(t *testing.T) {
	tests := []struct {
		name     string
		a        rune
		b        rune
		expected rune
		wantErr  bool
	}{
		{
			name:     "standard case",
			a:        'F',
			b:        'l',
			expected: 'M',
			wantErr:  false,
		},
		{
			name:     "switched first and second rune",
			a:        'l',
			b:        'F',
			expected: 'M',
			wantErr:  false,
		},
		{
			name:     "same rune",
			a:        'l',
			b:        'l',
			expected: 'A',
			wantErr:  false,
		},
		{
			name:     "invalid combination",
			a:        'F',
			b:        'F',
			expected: '#',
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetProductQuantity(tt.a, tt.b)

			if tt.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Add() = %+v; want %+v", result, tt.expected)
			}
		})
	}
}

func BenchmarkGetProductQuantity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GetProductQuantity('t', 'l')
		_, _ = GetProductQuantity('t', 'v')
		_, _ = GetProductQuantity('m', 'a')
		_, _ = GetProductQuantity('F', 'l')
		_, _ = GetProductQuantity('l', 'l')
		_, _ = GetProductQuantity('A', 'l')
		_, _ = GetProductQuantity('x', 'y') // nicht vorhanden
	}
}

func TestGetDivisionQuantity(t *testing.T) {
	tests := []struct {
		name        string
		numerator   rune
		denominator rune
		expected    rune
		wantErr     bool
	}{
		{
			name:        "standard case",
			numerator:   'l',
			denominator: 't',
			expected:    'v',
			wantErr:     false,
		},
		{
			name:        "switched numerator and denominator",
			numerator:   't',
			denominator: 'l',
			expected:    'v',
			wantErr:     false,
		},
		{
			name:        "same rune",
			numerator:   'l',
			denominator: 'l',
			expected:    'A',
			wantErr:     false,
		},
		{
			name:        "invalid combination",
			numerator:   'F',
			denominator: 'F',
			expected:    '#',
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetDivisionQuantity(tt.numerator, tt.denominator)

			if tt.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("GetDivisionQuantity() = %+v; want %+v", result, tt.expected)
			}
		})
	}
}

func BenchmarkGetDivisionQuantity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GetDivisionQuantity('l', 't') // → 'v'
		_, _ = GetDivisionQuantity('t', 'l') // → 'v'
		_, _ = GetDivisionQuantity('l', 'l') // → 'A'
		_, _ = GetDivisionQuantity('F', 'F') // → Fehlerfall
	}
}
