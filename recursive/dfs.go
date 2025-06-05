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
    visited := make(map[rune]bool)
    minChar := char

    var dfs func(rune)
    dfs = func(c rune) {
        if visited[c] {
            return
        }
        visited[c] = true
        if c < minChar {
            minChar = c
        }
        for _, neighbor := range mid[c] {
            dfs(neighbor)
        }

    }
    dfs(char)
    return minChar
}

package main
import "fmt"

func main() {
    s1 := "parker"
    s2 := "morris"
    baseStr := "parser"
  fmt.Println(smallestEquivalentString(s1, s2, baseStr))
}