package ofctl

import "golang.org/x/exp/constraints"

func roundUp[T constraints.Integer](val T, roundTo T) T {
	return val + (roundTo-val)%roundTo
}
