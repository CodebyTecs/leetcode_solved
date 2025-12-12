/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    l := 1
    r := n

    for l <= r {
        mid := (l + r)/2
        anw := guess(mid)

        if anw == 0 {
            return mid
        } else if anw == -1 {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }

    return 0
}