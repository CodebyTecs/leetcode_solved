func rotate(nums []int, k int)  {
   n := len(nums)
   if n == 0 {
        return
   }
   
   k %= n
   if k == 0 {
        return
   }

   reverse := func(a []int, l, r int) {
        for l < r {
            a[l], a[r] = a[r], a[l]
            l++
            r--
        }
   }
   reverse(nums, 0, n-1)
   reverse(nums, 0, k-1)
   reverse(nums, k, n-1)
}