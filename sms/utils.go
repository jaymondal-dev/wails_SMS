package sms

import (
    "math/rand"
    "time"
)

func Generate16CharAlphaString() string {
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    const length = 16

    var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

    result := make([]byte, length)
    for i := range result {
        result[i] = charset[seededRand.Intn(len(charset))]
    }

    return string(result)
}

add