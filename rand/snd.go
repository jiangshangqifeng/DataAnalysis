package rand

import (
	"math/rand"
	"sort"
	"time"
)

//自定义一个类型
type i64s []int64

func (s i64s) Len() int           { return len(s) }
func (s i64s) Less(i, j int) bool { return s[i] < s[j] }
func (s i64s) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// random
func Rand(total int64, capacity int) i64s {
	ret := []int64{}
	weights :=[]int64{}

	rand.Seed(time.Now().UnixNano())
	sum := int64(0)
	for {
		if len(weights) == capacity {
			break
		}
		i := RandInt64(1, 10000)
		weights = append(weights, i)
		sum += i
	}

	for _, weight := range weights {
		item := total * weight /sum
		ret = append(ret, item)
	}
	return ret
}

func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

//Pareto's law
func ParetosLaw(cardinal int64, capacity int) i64s {
	one5 := cardinal / 5
	four5 := cardinal - one5
	niche := capacity / 5
	mass := capacity - niche

	result := Rand(four5, niche)
	result = append(result, Rand(one5, mass)...)

	sort.Sort(result)
	return result
}

// standard normal distribution
func SND(cardinal int64, capacity int) i64s {
	var ret i64s
	weights :=[]int64{}
	sum := int64(0)

	rand.Seed(time.Now().UnixNano())
	for {
		if len(weights) == capacity {
			break
		}

		weight := rand.NormFloat64()*100 + 300
		if weight > 0 {
			weights = append(weights, int64(weight))
			sum += int64(weight)
		}
	}

	for _, weight := range weights {
		item := cardinal * weight /sum
		ret = append(ret, item)
	}
	sort.Sort(ret)
	return ret
}
