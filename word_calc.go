package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	errInvalidType   = errors.New("Данный оператор не поддреживается между этими типами")
	errStringInvalid = errors.New("Строка имела неверный формат")
	errMaxValue      = errors.New("Число не должно превышать 10")
	errStringMax     = errors.New("Строка имела более 10 символов")
)

func process_string(str string) (string, string, string, bool, error) {
	var l, r, s string
	if string(str[0]) == "\"" {
		l = str[1:]
		l = l[:strings.Index(l, "\"")]
		_, cutted, _ := strings.Cut(str, l)
		s = string(cutted[2])
		cutted = cutted[1:]
		if strings.Index(cutted, "\"") == -1 {
			splitted := strings.Split(cutted, " ")
			s = splitted[1]
			r = splitted[2]
			r = r[:len(r)-1]
			return l, r, s, true, nil
		} else {
			cutted = cutted[1:]
			if strings.Count(cutted, "\"") <= 1 {
				panic(errStringInvalid)
			} else {
				r = cutted[strings.Index(cutted, "\"")+1 : len(cutted)-2]
				return l, r, s, false, nil
			}
		}

	} else {
		panic(errStringInvalid)
	}
}

func subString(s, n string) string {
	res, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	s = s[:len(s)/res]
	return s
}
func multString(l, n string) string {
	m, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	res := ""
	for i := 0; i < m; i++ {
		res += l
	}
	if len(res) > 40 {
		return res[:40] + "..."
	} else {
		return res
	}
}

func cut(l, s string) (string, bool) {
	result := ""
	var index int
	flag := false
	for k, v := range l {
		if v == rune(s[0]) {
			if len(s) > len(l[k:]) {
				continue
			} else {
				if l[k:k+len(s)] == s {
					index = k + len(s)
					flag = true
					break
				}
			}
		}
		result += string(v)
	}
	result += l[index:]
	if flag {
		return result, true
	} else {
		return l, false
	}
}

func calculate(l, r, m string, isConverted bool) (string, error) {
	switch m {
	case "+":
		if isConverted {
			return "", errInvalidType
		}
		return l + r, nil

	case "-":
		if isConverted {
			return "", errInvalidType
		}
		//fmt.Println(l, r)
		after, found := cut(l, r)
		//fmt.Println(after)
		if found {
			return string(after), nil
		} else {
			return l, nil
		}

	case "*":
		if !isConverted {
			return "", errInvalidType
		}
		return multString(l, r), nil

	case "/":
		if !isConverted {
			return "", errInvalidType
		}
		return subString(l, r), nil

	default:
		return "", errors.New("Такой операции нет")
	}
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	for true {
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-1]
		left, right, middle, isConverted, err := process_string(text)
		if err != nil {
			panic(err)
		}
		val, _ := strconv.Atoi(right)
		if isConverted && val > 10 {
			panic(errMaxValue)
		}
		if len(left) > 10 {
			panic(errStringMax)
		}
		res, err := calculate(left, right, middle, isConverted)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(res))

	}
}
