package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		return
	}
	p := m + n - 1
	pm := m - 1
	pn := n - 1
	for pm >= 0 && pn >= 0 {
		if nums2[pn] >= nums1[pm] {
			nums1[p] = nums2[pn]
			pn--
		} else {
			nums1[p] = nums1[pm]
			pm--
		}
		p--
	}
}
