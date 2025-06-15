package utils

import "strconv"

func ParseUint(s string) uint64 {
    id, _ := strconv.ParseUint(s, 10, 64)
    return id
}
