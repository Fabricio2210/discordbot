package random

import (
    "math/rand"
    "time"
)

// ChooseRandomItem selects a random item from a slice of strings
func ChooseRandomItem(items []string) string {
    rand.Seed(time.Now().UnixNano())
    index := rand.Intn(len(items))
    return items[index]
}
