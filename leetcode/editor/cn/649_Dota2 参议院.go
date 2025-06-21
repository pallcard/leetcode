package main

import (
	"fmt"
	"strings"
)

// leetcode submit region begin(Prohibit modification and deletion)
func predictPartyVictory(senate string) string {
	//todo 优化点，字符串可以转 []byte 来处理
	//senateBytes := []byte(senate)
	dRest := false
	rRest := false
	dCnt := 0
	rCnt := 0
	newStr := strings.Builder{}
	for i := 0; i < len(senate); i++ {
		if senate[i] == 'D' {
			if rCnt > 0 {
				rCnt--
			} else {
				dCnt++
				newStr.Write([]byte("D"))
				dRest = true
			}
		} else if senate[i] == 'R' {
			if dCnt > 0 {
				dCnt--
			} else {
				rCnt++
				rRest = true
				newStr.Write([]byte("R"))
			}
		}
		if i == len(senate)-1 {
			if rRest && !dRest {
				return "Radiant"
			} else if !rRest && dRest {
				return "Dire"
			}
			rRest = false
			dRest = false
			i = -1
			senate = newStr.String()
			newStr.Reset()
		}
	}

	return ""

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	res := predictPartyVictory("DDDDRRDDDRDRDRRDDRDDDRDRRRRDRRRRRDRDDRDDRRDDRRRDDRRRDDDDRRRRRRRDDRRRDDRDDDRRRDRDDRDDDRRDRRDRRRDRDRDR")
	fmt.Print(res)
}
