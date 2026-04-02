func addStrings(num1 string, num2 string) string {
    i := len(num1) - 1
    j := len(num2) - 1
    ans := []byte{}
    carry := 0
    
    for i >= 0 || j >= 0 || carry > 0 {
        sum := carry

        if i >= 0 {
            sum += int(num1[i] - '0')
            i--
        }

        if j >= 0 {
            sum += int(num2[j] - '0')
            j--
        }

        ans = append(ans, byte(sum%10)+'0')
        carry = sum / 10
    }

    for i := 0; i < len(ans)/2; i++ {
        ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
    }

    return string(ans)
}
