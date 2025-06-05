func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
    mid := make(map[rune][]rune)
    var result []rune
    for i, c := range s1 {
        mid[c] = append(mid[c], rune(s2[i]))
        mid[rune(s2[i])] = append(mid[rune(s2[i])], c)
    }
    for _, c := range baseStr {
         result = append(result, findLowestMatch(mid, c))
    }
    return string(result)
}

func findLowestMatch(mid map[rune][]rune, char rune) rune {
    found := 1
    for found == 1 {
        char, found = findLowestLetter(mid[char], char)
    }
    return char
}

func findLowestLetter(s []rune, char rune) (rune, int) {
    if len(s) == 0 {
        return 0, 0
    }
    lowest := rune(s[0])
    for _, c := range s {
        if c < lowest {
            lowest = c
        }
    }
    if char <= lowest {
        lowest = char
        return lowest, 0
    }
    return lowest, 1
}

package main
import "fmt"

func main() {
    s1 := "parker"
    s2 := "morris"
    baseStr := "parser"
  fmt.Println(smallestEquivalentString(s1, s2, baseStr))
}