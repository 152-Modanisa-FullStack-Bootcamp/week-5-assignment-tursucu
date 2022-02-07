package assignment

import (
	"math"
	"reflect"
	"sort"
	"strings"
)

func AddUint32(x, y uint32) (uint32, bool) {
	sum := x + y
	if sum > x || sum > y {
		return sum, false
	}
	return sum, true
}

func CeilNumber(f float64) float64 {
	const inc float64 = 0.25
	return math.Ceil(f/inc) * inc
}

func AlphabetSoup(s string) string {
	charArr := []rune(s)
	sort.Slice(charArr, func(i, j int) bool {
		return charArr[i] < charArr[j]
	})
	return string(charArr)
}

func StringMask(s string, n uint) string {
	var value string
	change := "*"
	charArr := []rune(s)
	l := uint(len(charArr))

	if l == 0 {
		return change
	}

	for i := range charArr {
		char := string(charArr[i])
		if n >= l {
			value += change
		} else if uint(i) >= n {
			value += change
		} else {
			value += char
		}
	}

	return value
}

func WordSplit(arr [2]string) string {
	if len(arr) != 2 || (len(arr[0]) == 0 || len(arr[1]) == 0) {
		return "not possible"
	}

	var generator func([]string, int, []string, *[][]string, int)
	generator = func(set []string, breakPoint int, selected []string, results *[][]string, maxLength int) {
		if len(set) == breakPoint {
			if len(selected) == maxLength {
				arr := make([]string, maxLength)
				copy(arr, selected)
				*results = append(*results, arr)
			}
			return
		}
		selected = append(selected, set[breakPoint])
		generator(set, breakPoint+1, selected, results, maxLength)
		selected = selected[:len(selected)-1]
		generator(set, breakPoint+1, selected, results, maxLength)
	}

	results := make([][]string, 0)
	selected := make([]string, 0)
	items := strings.Split(arr[1], ",")
	generator(items, 0, selected, &results, 2)
	for _, value := range results {
		tempVal := ""
		var reverse []string
		for _, val := range value {
			tempVal += val
			reverse = append(reverse, val)
		}
		rLen := len(reverse) / 2
		for u := 0; u < rLen; u++ {
			temp := reverse[u]
			reverse[u] = reverse[len(reverse)-1-u]
			reverse[len(reverse)-1-u] = temp
		}
		tempVal = ""
		for _, val := range reverse {
			tempVal += val
		}
		if arr[0] == tempVal {
			return strings.Join(reverse, ",")
		}
	}
	return "not possible"
}

func VariadicSet(arg ...interface{}) []interface{} {
	args := reflect.ValueOf(arg[0])
	l := args.Len()
	var set []interface{}

	for i := 0; i < l; i++ {
		argValue := args.Index(i).Interface()
		argsSet := reflect.ValueOf(set)
		lSet := argsSet.Len()
		nilControl := argsSet.IsNil()
		var setValueControl bool

		if !nilControl {
			for b := 0; b < lSet; b++ {
				argsSetValue := argsSet.Index(b).Interface()
				if argsSetValue == argValue {
					setValueControl = true
				}
			}
		}

		if !setValueControl {
			set = append(set, argValue)
		}
	}
	return set
}
