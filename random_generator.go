package tcl

import (
	"math/rand"
	"time"
)

const TimestampEffectiveBits = 10
const TimestampMask = (1 << TimestampEffectiveBits) - 1

var globalRand *rand.Rand
var globalRandSource rand.Source

func UpdateRandomSeedWithCurrentTimestamp() {
	timestamp := time.Now().Unix()
	effectiveValue := timestamp & TimestampMask
	offset := Intn(64 - TimestampEffectiveBits + 1)
	newSeed := globalRandSource.Int63() ^ (effectiveValue << uint(offset))
	globalRandSource.Seed(newSeed)
}

func Intn(n int) int {
	return globalRand.Intn(n)
}

func init() {
	globalRandSource = rand.NewSource(time.Now().UnixNano())
	globalRand = rand.New(globalRandSource)
}
