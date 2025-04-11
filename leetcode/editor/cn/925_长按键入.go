package main

import "fmt"

// leetcode submit region begin(Prohibit modification and deletion)
func isLongPressedName(name string, typed string) bool {
	// todo 同一个字母，下面的个数大于上面个数就行
	// alex
	// aaleex
	if len(name) > len(typed) {
		return false
	}
	if len(name) == len(typed) && name != typed {
		return false
	}

	pre := name[0]
	cnt := 1
	j := 0
	Tcnt := 0
	for i := 1; i < len(name); i++ {
		if name[i] == pre {
			cnt++
		} else {
			for ; j < len(typed); j++ {
				if typed[j] == pre {
					Tcnt++
				} else {
					break
				}
			}
			if cnt > Tcnt {
				return false
			}

			// 换字母
			cnt = 1
			pre = name[i]
			Tcnt = 0
		}
	}

	// todo 处理最后一个字母
	for ; j < len(typed); j++ {
		if typed[j] == pre {
			Tcnt++
		} else {
			break
		}
	}
	if cnt > Tcnt || j != len(typed) {
		return false
	}

	return true

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	name := isLongPressedName("leelee", "lleeelee")
	fmt.Print(name)
}
