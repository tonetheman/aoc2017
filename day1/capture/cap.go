package capture

import "strconv"

// Captcha computes the captch from day 1
func Captcha(s string) int {
	total := 0
	var v1, v2 int
	for i := 0; i < len(s); i++ {
		v1, _ = strconv.Atoi(string(s[i]))
		index := (i + 1) % len(s)
		v2, _ = strconv.Atoi(string(s[index]))
		if v1 == v2 {
			//fmt.Println(i, "adding", v1)
			total += v1
		}
		//fmt.Println(v1, v2)
	}
	//fmt.Println(v1, v2)
	return total
}
