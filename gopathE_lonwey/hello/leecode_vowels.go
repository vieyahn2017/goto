package main
 
import (
    "fmt"
    //"strings"
    //"strconv"
    //"github.com/stretchr/testify/assert"
    //"testing"
   )

var vowels = map[rune]bool{
    'a': true, 'A': true,
    'e': true, 'E': true,
    'i': true, 'I': true,
    'o': true, 'O': true,
    'u': true, 'U': true,
}

func reverseVowels(s string) string {
    result := []rune(s)
    for i, j := 0, len(s)-1; i < j; {
        for i < j && !vowels[result[i]] {
            i++
        }
        for i < j && !vowels[result[j]] {
            j--
        }
        if i < j {
            result[i], result[j] = result[j], result[i]
            i++
            j--
        }
    }
    return string(result)
}

func main() {

    fmt.Println(reverseVowels("leetcode"))

}