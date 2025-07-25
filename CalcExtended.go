package CalcExtended

import (
	"strconv"
	"strings"
)

func numsConvF(n1, n2 string) (float64, float64, error) {
	v1, err := strconv.ParseFloat(strings.TrimRight(n1, "+-*/"), 64)
	if err != nil {
		return 0, 0, err
	}
	v2, err := strconv.ParseFloat(strings.TrimRight(n2, "+-*/"), 64)
	if err != nil {
		return 0, 0, err
	}
	return v1, v2, nil
}

func zeroRemove(n string) string {
	f, b, _ := strings.Cut(n, ".")
	sl := strings.Split(b, "")
	for i := len(sl) - 1; i > 0; i-- {
		if sl[i] == "0" {
			sl = sl[:i]
		} else {
			break
		}
	}
	b = strings.Join(sl, "")
	sl = []string{f, b}
	if b != "0" {
		n = strings.Join(sl, ".")
	} else {
		n = f
	}
	return n
}

func CalcExpr(expression string) (string, error) {
	var vars []string
	var sum string
	i := 0
	exp := []byte(expression)
	for i < len(expression) {
		i = strings.IndexAny(string(exp), "+-*/")
		if i == -1 {
			vars = append(vars, string(exp))
			break
		}
		vars = append(vars, string(exp[:i+1]))
		exp = exp[i+1:]
	}
	if strings.ContainsAny(expression, "*/") {
		for idx, val := range vars {
			switch {
			case strings.HasSuffix(val, "*"):
				v1, v2, err := numsConvF(vars[idx], vars[idx+1])
				if err != nil {
					return "", err
				}
				sign, _ := strings.CutPrefix(vars[idx+1], zeroRemove(strconv.FormatFloat(v2, 'f', 10, 64)))
				sl := []string{zeroRemove(strconv.FormatFloat(v1*v2, 'f', 10, 64)), sign}
				vars[idx+1] = strings.Join(sl, "")
				sl[0], sl[1] = "0", "+"
				vars[idx] = strings.Join(sl, "")
				sum = zeroRemove(strconv.FormatFloat(v1*v2, 'f', 4, 64))
			case strings.HasSuffix(val, "/"):
				v1, v2, err := numsConvF(vars[idx], vars[idx+1])
				if err != nil {
					return "", err
				}
				sign, _ := strings.CutPrefix(vars[idx+1], zeroRemove(strconv.FormatFloat(v2, 'f', 10, 64)))
				sl := []string{zeroRemove(strconv.FormatFloat(v1/v2, 'f', 10, 64)), sign}
				vars[idx+1] = strings.Join(sl, "")
				sl[0], sl[1] = "0", "+"
				vars[idx] = strings.Join(sl, "")
				sum = zeroRemove(strconv.FormatFloat(v1/v2, 'f', 4, 64))
			default:
				continue
			}
		}
	}
	var vars2 []string
	for _, val := range vars {
		if val != "0+" {
			vars2 = append(vars2, val)
		}
	}
	vars = vars2
	for idx, val := range vars {
		switch {
		case strings.HasSuffix(val, "+"):
			v1, v2, err := numsConvF(vars[idx], vars[idx+1])
			if err != nil {
				return "", err
			}
			sign, _ := strings.CutPrefix(vars[idx+1], zeroRemove(strconv.FormatFloat(v2, 'f', 10, 64)))
			sl := []string{zeroRemove(strconv.FormatFloat(v1+v2, 'f', 10, 64)), sign}
			vars[idx+1] = strings.Join(sl, "")
			sum = zeroRemove(strconv.FormatFloat(v1+v2, 'f', 4, 64))
		case strings.HasSuffix(val, "-"):
			v1, v2, err := numsConvF(vars[idx], vars[idx+1])
			if err != nil {
				return "", err
			}
			sign, _ := strings.CutPrefix(vars[idx+1], zeroRemove(strconv.FormatFloat(v2, 'f', 10, 64)))
			sl := []string{zeroRemove(strconv.FormatFloat(v1-v2, 'f', 10, 64)), sign}
			vars[idx+1] = strings.Join(sl, "")
			sum = zeroRemove(strconv.FormatFloat(v1-v2, 'f', 4, 64))
		}
	}
	return sum, nil
}
