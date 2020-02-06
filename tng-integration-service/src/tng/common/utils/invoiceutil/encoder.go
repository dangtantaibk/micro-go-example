package invoiceutil

import "strconv"

type permut struct {
	src int
	dst int
}

var (
	key     = []int{0, 3, 8, 2, 7, 9, 5, 6, 1, 4}
	permut1 = []permut{
		permut{0, 9},
		permut{1, 3},
		permut{2, 7},
		permut{4, 11},
		permut{5, 6},
		permut{8, 10},
		permut{0, 0},
	}
	permut2 = []permut{
		permut{0, 5},
		permut{1, 4},
		permut{2, 3},
		permut{6, 10},
		permut{7, 8},
		permut{9, 11},
		permut{0, 0},
	}
)

func permute12(str []string, en bool) {
	var swap string
	if en {
		for i := 0; i < len(permut1); i++ {
			if permut1[i].src != 0 || permut1[i].dst != 0 {
				swap = str[permut1[i].src]
				str[permut1[i].src] = str[permut1[i].dst]
				str[permut1[i].dst] = swap
			}
		}
		for i := 0; i < len(permut2); i++ {
			if permut2[i].src != 0 || permut2[i].dst != 0 {
				swap = str[permut2[i].src]
				str[permut2[i].src] = str[permut2[i].dst]
				str[permut2[i].dst] = swap
			}
		}
	} else {
		for i := 0; i < len(permut2); i++ {
			if permut2[i].src != 0 || permut2[i].dst != 0 {
				swap = str[permut2[i].src]
				str[permut2[i].src] = str[permut2[i].dst]
				str[permut2[i].dst] = swap
			}
		}
		for i := 0; i < len(permut1); i++ {
			if permut1[i].src != 0 || permut1[i].dst != 0 {
				swap = str[permut1[i].src]
				str[permut1[i].src] = str[permut1[i].dst]
				str[permut1[i].dst] = swap
			}
		}
	}
}

func Encode(in []string, len int) []string {
	var tmp []string
	var k, a, b int
	k, _ = strconv.Atoi(in[6])
	for i := 0; i < len; i++ {
		if i != 6 {
			k = (k % 9) + 1
			a, _ = strconv.Atoi(in[i])
			b = a + key[k]
			if b >= 10 {
				b = b - 10
			}
			tmp = append(tmp, string(b+'0'))
			k = k + i + 1
		} else {
			tmp = append(tmp, in[6])
		}
	}
	permute12(tmp, true)
	return tmp
}

func Decode(in []string, len int) []string {
	var tmp []string
	var k, a, b int
	permute12(in, false)
	k, _ = strconv.Atoi(in[6])
	for i := 0; i < len; i++ {
		if i != 6 {
			k = (k % 9) + 1
			a, _ = strconv.Atoi(in[i])
			if a >= key[k] {
				b = a - key[k]
			} else {
				b = a + 10 - key[k]
			}
			tmp = append(tmp, string(b+'0'))
			k = k + i + 1
		} else {
			tmp = append(tmp, in[6])
		}
	}
	return tmp
}
