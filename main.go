package goarea

import "math"

const Pi = 3.1416

func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

func Rect(base, altrua float64) float64 {
	return base * altrua
}

func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
