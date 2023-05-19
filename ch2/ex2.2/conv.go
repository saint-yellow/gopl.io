package main

import "fmt"

type meter float64
type foot float64

const coefficient float64 = 3.28084

func (m meter) String() string {
	return fmt.Sprintf("%g m", m)
}

func (f foot) String() string {
	return fmt.Sprintf("%g ft", f)
}

func mTof(m meter) foot {
	return foot(m * meter(coefficient))
}

func fTom(f foot) meter {
	return meter(f / foot(coefficient))
}
