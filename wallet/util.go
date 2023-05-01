package wallet

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strconv"
)

// NewPaymentID64 generates a 64 bit payment ID (hex encoded).
// With 64 bit IDs, there is a non-negligible chance of a collision
// if they are randomly generated. It is up to recipients generating
// them to sanity check for uniqueness.
//
// 1 million IDs at 64-bit (simplified): 1,000,000^2 / (2^64 * 2) = ~1/36,893,488 so
// there is a 50% chance a collision happens around 5.06 billion IDs generated.
func NewPaymentID64() string {
	buf := make([]byte, 8)
	rand.Read(buf)
	return hex.EncodeToString(buf)
}

// NewPaymentID256 generates a 256 bit payment ID (hex encoded).
func NewPaymentID256() string {
	buf := make([]byte, 32)
	rand.Read(buf)
	return hex.EncodeToString(buf)
}

// OXENToDecimal converts a raw atomic OXEN balance to a more
// human readable format.
func OXENToDecimal(OXEN uint64) string {
	str0 := fmt.Sprintf("%013d", OXEN)
	l := len(str0)
	return str0[:l-12] + "." + str0[l-12:]
}

// OXENToFloat64 converts raw atomic OXEN to a float64
func OXENToFloat64(OXEN uint64) float64 {
	return float64(OXEN) / 1e12
}

// Float64ToOXEN converts a float64 to a raw atomic OXEN
func Float64ToOXEN(OXEN float64) uint64 {
	return uint64(OXEN * 1e12)
}

// StringToOXEN converts a string to a raw atomic OXEN
func StringToOXEN(OXEN string) (uint64, error) {
	f, err := strconv.ParseFloat(OXEN, 64)
	if err != nil {
		return 0, err
	}
	return uint64(f * 1e12), nil
}
