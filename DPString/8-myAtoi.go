package week09

import (
	"math"
	"strings"
)

const (
	START  = '0'
	SIGNED = '1'
	NUMBER = '2'
	END    = '3'
)

// using state machine to do it
var states = map[byte][]byte{
	// preState: " ", +/-, NUMBER, other
	START:  []byte{START, SIGNED, NUMBER, END},
	SIGNED: []byte{END, END, NUMBER, END},
	NUMBER: []byte{END, END, NUMBER, END},
	END:    []byte{END, END, END, END},
}

type AtoiStateMachine struct {
	ans   int
	sign  int // 1 or -1
	state byte
}

func NewAtoiStateMachine() *AtoiStateMachine {
	return &AtoiStateMachine{
		ans:   0,
		sign:  1,
		state: START,
	}
}

func (sm *AtoiStateMachine) GetAns(c rune) (int, bool) {
	sm.state = states[sm.state][getType(c)]
	if sm.state == NUMBER {
		pop := sm.sign * int(c)
		if sm.sign == 1 && (sm.ans > math.MaxInt32/10 || sm.ans == math.MaxInt32/10 && pop > 7) {
			sm.state = END
			sm.ans = math.MaxInt32
			return sm.ans, false
		}
		if sm.sign == -1 && (sm.ans < math.MinInt32/10 || sm.ans == math.MinInt32/10 && pop < -8) {
			sm.state = END
			sm.ans = math.MinInt32
			return sm.ans, false
		}
		sm.ans = 10*sm.ans + pop
	}

	if sm.state == SIGNED && c == '-' {
		sm.sign = -1
	}

	return sm.ans, sm.state != END
}

func getType(c rune) int {
	// " "
	if c == ' ' {
		return 0
	}
	// +/-
	if c == '+' || c == '-' {
		return 1
	}
	// number
	if c >= '0' && c <= '9' {
		return 2
	}
	// other
	return 3
}

func myAtoi(str string) int {
	sm := NewAtoiStateMachine()
	var res int
	var ok bool
	for _, c := range str {
		if res, ok = sm.GetAns(c); !ok {
			return res
		}
	}

	return res
}

// origin method
func myAtoi1(str string) int {
	res, sign, flag := 0, 0, 0
	str = strings.Trim(str, " ")
	for _, ss := range str {
		// + -
		if flag == 0 && (ss == '+' || ss == '-') {
			if sign != 0 {
				return 0
			}
			if ss == '+' {
				sign = 1
			} else {
				sign = -1
			}
			continue
		}

		if ss >= '0' && ss <= '9' {
			flag = 1
			// 0
			if ss == '0' {
				if res == 0 {
					continue
				}
			}
			pop := int(ss - '0')
			if sign >= 0 && (res > math.MaxInt32/10 || res == math.MaxInt32/10 && pop > 7) {
				// panic("> MAXINT32")
				return math.MaxInt32
			}
			if sign == -1 && (-1*res < math.MinInt32/10 || -1*res == math.MinInt32/10 && -1*pop < -8) {
				// panic("< MININT32")
				return math.MinInt32
			}
			res = res*10 + pop
			continue
		}

		break
	}

	if sign == -1 {
		return -1 * res
	}

	return res
}
