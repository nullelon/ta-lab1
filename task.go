package lab1

import "errors"

func DecimalToRomanNumeral(num int) (RomanNumeral, error) {
	if num < 1 && num > 3000 {
		return nil, errors.New("wrong argument")
	}
	result := NewRomanNumeral()

	if i := num / 1000; i >= 1 {
		switch i {
		case 1:
			result = append(result, M)
		case 2:
			result = append(result, M, M)
		case 3:
			result = append(result, M, M, M)
		}
	}
	if i := (num / 100) % 10; i >= 1 {
		switch i {
		case 1:
			result = append(result, C)
		case 2:
			result = append(result, C, C)
		case 3:
			result = append(result, C, C, C)
		case 4:
			result = append(result, C, D)
		case 5:
			result = append(result, D)
		case 6:
			result = append(result, D, C)
		case 7:
			result = append(result, D, C, C)
		case 8:
			result = append(result, D, C, C, C)
		case 9:
			result = append(result, C, M)
		}
	}
	if i := (num / 10) % 10; i >= 1 {
		switch i {
		case 1:
			result = append(result, X)
		case 2:
			result = append(result, X, X)
		case 3:
			result = append(result, X, X, X)
		case 4:
			result = append(result, X, L)
		case 5:
			result = append(result, L)
		case 6:
			result = append(result, L, X)
		case 7:
			result = append(result, L, X, X)
		case 8:
			result = append(result, L, X, X, X)
		case 9:
			result = append(result, X, C)
		}
	}
	if i := num % 10; i >= 1 {
		switch i {
		case 1:
			result = append(result, I)
		case 2:
			result = append(result, I, I)
		case 3:
			result = append(result, I, I, I)
		case 4:
			result = append(result, I, V)
		case 5:
			result = append(result, V)
		case 6:
			result = append(result, V, I)
		case 7:
			result = append(result, V, I, I)
		case 8:
			result = append(result, V, I, I, I)
		case 9:
			result = append(result, I, X)
		}
	}

	return result, nil
}
