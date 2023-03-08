package validate

import "strconv"

func IsUint(input string) (uint, error) {
	i, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(i), nil
}
