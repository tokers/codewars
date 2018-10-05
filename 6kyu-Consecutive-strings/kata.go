package kata

import "strings"

func LongestConsec(strarr []string, k int) string {
    n := len(strarr)
    if n == 0 || k > n || k <= 0 {
        return ""
    }
    
    sum := 0
    for i := 0; i < k; i++ {
        sum += len(strarr[i])
    }
    
    index := 0
    lengest := sum
    for i := k; i < n; i++ {
        sum = sum - len(strarr[i - k]) + len(strarr[i])
        if lengest < sum {
            index = i - k + 1
            lengest = sum
        }
    }
    
    var b strings.Builder
    for i := index; i < index + k; i++ {
        b.WriteString(strarr[i])
    }
    return b.String()
}
