func reverseBits(n int) int {
    var res int32 = 0
    var num int32 = int32(n)
    for i := 0; i < 32; i++ {
        res = (res << 1) | (num & 1)
        num>>=1
    }
    return int(res)
}