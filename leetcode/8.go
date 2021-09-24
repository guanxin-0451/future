package leetcode

import "strconv"

const IntMax = 2 * 31

const (
	Start  = "start"
	End    = "end"
	Signed = "signed"
	Number = "number"
)

type Atoi struct {
	Mod     map[string][]string
	S       string
	I       int64
	Status  string
	IsMinus bool
}

func (a *Atoi) isBlank(s string) bool {
	if s == " " {
		return true
	}
	return false
}

func (a *Atoi) isNumber(s string) bool {
	if s >= "0" && s <= "9" {
		return true
	}
	return false
}

func (a *Atoi) isSign(s string) bool {
	if s == "-" || s == "+" {
		return true
	}
	return false
}

func (a *Atoi) isOther(s string) bool {
	if a.isBlank(s) || a.isNumber(s) || a.isSign(s) {
		return false
	}
	return true
}

func (a *Atoi) toInt() int64 {
	for _, c := range a.S {
		ch := string(c)
		if a.isBlank(ch) {
			a.Status = a.Mod[a.Status][0]
		} else if a.isSign(ch) {
			a.Status = a.Mod[a.Status][1]
		} else if a.isNumber(ch) {
			a.Status = a.Mod[a.Status][2]
		} else {
			a.Status = a.Mod[a.Status][3]
		}

		if a.Status == Signed {
			if ch == "-" {
				a.IsMinus = true
			}
		}

		if a.Status == Number {
			i64, _ := strconv.ParseInt(ch, 10, 64)
			if IntMax/10 < a.I {
				a.I = IntMax
				a.Status = End
			} else {
				a.I = a.I*10 + i64
			}
		}

	}

	if a.IsMinus {
		a.I = 0 - a.I
	}
	return a.I
}

func NewAtoi(s string) *Atoi {
	atoi := Atoi{
		Mod: map[string][]string{
			//      " "    "+/-"   "0-9"   "other"
			Start:  {Start, Signed, Number, End},
			Signed: {End, End, Number, End},
			Number: {End, End, Number, End},
			End:    {End, End, End, End},
		},
		S:       s,
		IsMinus: false,
		Status:  Start,
	}

	return &atoi
}

func MyAtoi(s string) int {
	atoi := NewAtoi(s)
	return int(atoi.toInt())
}
