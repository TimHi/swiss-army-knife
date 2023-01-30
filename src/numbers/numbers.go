package numbers

import (
	"errors"
	"math/big"
)

// Result of this call is plattform dependend.
// Check wether it is a 64 bit or 32 bit OS is done using ^uint(0)
func MaxInt() int {
	const UintSize = 32 << (^uint(0) >> 32 & 1)
	const MaxInt = 1<<(UintSize-1) - 1
	return MaxInt
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Mod(x, y int64) (int64, error) {
	if y == 0 {
		return 0, errors.New("attempted to divide with 0")
	}
	bx, by := big.NewInt(x), big.NewInt(y)
	return new(big.Int).Mod(bx, by).Int64(), nil
}
