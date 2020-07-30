package selfsimilar

import (
	"math"
	"math/rand"
)

//自定义一个类型
type f64s []float64

func (s f64s) Len() int           { return len(s) }
func (s f64s) Less(i, j int) bool { return s[i] > s[j] }
func (s f64s) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// self-similar
func Selfsimilar(n, h float64) float64 {
	return 1 + (n * math.Pow(rand.Float64(), math.Log(h)/math.Log(1.0-h)))
}