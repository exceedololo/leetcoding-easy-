func climbStairs(n int) int {
/*if n == 0{
      return 0
  }else if n ==1 {
      return 1
  }else if n == 2{
      return 2
  }else{
      return climbStairs(n-1) + climbStairs(n-2)
  }*/
arr := make([]int, 0, n)
for i :=0; i<n;i++{
if i == 0{
arr = append(arr, 1)
}else if i == 1{
arr = append(arr, 2)
}else{
arr = append(arr, arr[i-1]+arr[i-2])
}
}
return arr[n-1]
}