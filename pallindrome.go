import (
    "strconv"
    "slices"
)

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }

    s := strconv.Itoa(x)      // int → string
    r := []rune(s)            // string → runes
    slices.Reverse(r)         // reverse runes
    reversed := string(r)     // runes → string

    return s == reversed
}
