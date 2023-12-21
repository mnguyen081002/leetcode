/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

package main

func isBadVersion(version int) bool {
	return version == 4
}

func firstBadVersion(n int) int {
	if n == 1 {
		return n
	}
	return binarySearch(n, 0, n)
}

func binarySearch(n, left, right int) int {
	if right >= left {
		mid := (left + right) / 2
		if isBadVersion(mid) {
			return binarySearch(n, left, mid-1)
		}
		return binarySearch(n, mid+1, right)
	}
	return left
}
