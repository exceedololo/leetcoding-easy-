func isPalindrome(x int) bool {
str := strconv.Itoa(x)
pravda:= true
dlina := len(str)
for i := 0; i < dlina/2; i++ {
if str[i] != str[dlina-i-1] {
pravda = false
break
}
}
return pravda
}