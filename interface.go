package main

import (
	"errors"
)

type Interest interface {
	simpleInterest(p float64, r float64, t float64) (float64, error)
}

func simpleInterest(p float64, r float64, t float64) (float64, error) {
	i := p * r * t / 100
	if i == 0 {
		return 0, errors.New("Error calculating")
	}
	return i, nil
}
