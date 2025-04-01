package random

import (
	cr "crypto/rand"
	"log"
	"math"
	"math/big"
	mr "math/rand"
	"reflect"
)

var seed *big.Int

// RandomValueType 隨機值型別 interface
type RandomValueType interface {
	int | int32 | int64
}

func init() {
	var err error
	seed, err = cr.Int(cr.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		log.Fatalln("Random seed error:", err.Error())
	}
	mr.NewSource(seed.Int64())
}

// GetRandomInt 取得 0~n-1 隨機整數(接受型別[int | int32 | int64])
func GetRandomInt[T RandomValueType](n T) T {
	if n <= 0 {
		return 0
	}
	switch reflect.TypeOf(n).Kind() {
	case reflect.Int:
		return T(mr.Intn(int(n)))
	case reflect.Int32:
		return T(mr.Int31n(int32(n)))
	case reflect.Int64:
		return T(mr.Int63n(int64(n)))
	}
	return T(mr.Int63n((int64(n))))
}
