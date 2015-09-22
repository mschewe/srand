// Package srand contains utility functions for working with secure randoms.
package srand

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"math/big"
)

const (
	maxIntRange = 999999
)

var MinMaxError = errors.New("MIN value cannot be greater than MAX value.")

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

// GenerateRandomInt returns a random number output between 0 and n.
// It will return an error if the system's secure random
// number generator fails to function correctly.
func GenerateRandomInt(n int64) (int64, error) {
	max := *big.NewInt(n)
	r, err := rand.Int(rand.Reader, &max)
	if err != nil {
		return 0, err
	}
	return r.Int64(), nil
}

// GenerateRandomIntRange returns a random number output within the given range [m, n).
func GenerateRandomIntRange(m, n int64) (int64, error) {
	if m > n {
		return 0, MinMaxError
	}
	if n == m {
		return n, nil
	}
	max := n - m
	r, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}
	return (m + r.Int64()), nil
}

// GenerateRandomFloatRange returns a random floating point number output within the range [m, n).
func GenerateRandomFloat() (float64, error) {
	rand1, err := GenerateRandomInt(maxIntRange)
	if err != nil {
		return 0, err
	}
	rand2, err := GenerateRandomInt(maxIntRange)
	if err != nil {
		return 0, err
	}
	return float64(((rand1<<15 + rand2) & ((1 << 24) - 1))) / (1 << 24), nil
}

// GenerateRandomFloatRange returns a random floating point number output within the range [m, n).
func GenerateRandomFloatRange(m, n float64) (float64, error) {
	if m > n {
		return 0, MinMaxError
	}
	if n == m {
		return n, nil
	}
	rand1, err := GenerateRandomInt(maxIntRange)
	if err != nil {
		return 0, err
	}
	rand2, err := GenerateRandomInt(maxIntRange)
	if err != nil {
		return 0, err
	}
	return (m + (float64(((rand1<<15 + rand2) & ((1 << 24) - 1))) / (1 << 24) / (n - m))), nil
}

// RandMax returns the maximum value returned by the rand function.
func RandMax(n int64) int64 {
	return int64((1 << 63) - 1 - (1<<63)%uint64(n))
}

// GenerateUUID returns a securely generated UUID according to RFC 4122.
func GenerateUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random)
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
