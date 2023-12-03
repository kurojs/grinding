/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    left := 1
    right := n
    
    for left <= right {
        pivot := (left + right) / 2
        
        if !isBadVersion(pivot) {
            left = pivot + 1
            continue
        }

        if isFirstBadVersion(pivot) {
            return pivot
        }
        
        right = pivot - 1
    }
    
    return -1
}

func isFirstBadVersion(n int) bool {
    if n == 1 {
        return true
    }

    if n > 1 && !isBadVersion(n - 1) {
        return true
    }
    
    return false
}


/*
2 cases will happend:
    - not bad version -> check next pivot
    - is bad version:
        - is previous good version -> 1st bad version
        - continue
*/