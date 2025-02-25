package top100

import "math"

// 找两个正序数组中位数， 找第k到到数， k = ( len(nums1)+len(nums2) ) / 2
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i := 0
	j := 0

	// 4 --- 1 2
	// 3 --- 1
	k := (len(nums1) + len(nums2)) / 2

	two := (len(nums1)+len(nums2))%2 == 0
	if two {
		k--
	}

	// 1 4
	// 2 3

	// 3+4 /2
	res := 0.0
	index := 0
	find := false
	for i < len(nums1) && j < len(nums2) {
		if index == k {
			find = true
			if two {
				//nums1[i], nums1[i+1], nums2[j], nums2[j+1]
				res += math.Min(float64(nums1[i]), float64(nums2[j]))
				if nums1[i] < nums2[j] {
					if i+1 < len(nums1) {
						res += math.Min(float64(nums1[i+1]), float64(nums2[j]))
					} else {
						res += float64(nums2[j])
					}

				} else {
					if j+1 < len(nums2) {
						res += math.Min(float64(nums1[i]), float64(nums2[j+1]))
					} else {
						res += float64(nums1[i])
					}

				}

				res /= 2
			} else {
				res = math.Min(float64(nums1[i]), float64(nums2[j]))
			}
		}

		if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
		index++
	}

	if !find {
		if len(nums1) < len(nums2) {
			k -= len(nums1)
			if two {
				res = float64(nums2[k]+nums2[k+1]) / 2
			} else {
				res = float64(nums2[k])
			}
		} else {
			k -= len(nums2)
			if two {
				res = float64(nums1[k]+nums1[k+1]) / 2
			} else {
				res = float64(nums1[k])
			}
		}
	}

	return res
}
