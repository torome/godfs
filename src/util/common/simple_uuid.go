package common

import (
    "math/rand"
)

var seeds = []rune("abcdefghijklmnopqrstuvwxyz1234567890")

// create simple uuid from rand seed
func UUID() string {
    var buffer = make([]rune, 30)
    for i := 0; i < 30; i++ {
        index := rand.Int31n(30)
        buffer[i] = seeds[index]
    }
    return string(buffer)
}
