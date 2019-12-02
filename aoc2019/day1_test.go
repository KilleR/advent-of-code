package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getFuel(t *testing.T) {
	tests := []struct {
		name     string
		mass     float64
		wantFuel float64
	}{
		{"zero value", 2, 0},
		{"basic small", 12, 2},
		{"rounding small", 14, 2},
		{"medium", 1969, 654},
		{"large", 100756, 33583},
	}

	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(getFuel(test.mass), test.wantFuel, test.name)
	}
}

func Test_getExtraFuel(t *testing.T) {
	type args struct {
		mass float64
	}
	tests := []struct {
		name     string
		args     args
		wantFuel float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFuel := getExtraFuel(tt.args.mass); gotFuel != tt.wantFuel {
				t.Errorf("getExtraFuel() = %v, want %v", gotFuel, tt.wantFuel)
			}
		})
	}
}
