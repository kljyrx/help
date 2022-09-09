package help

import "fmt"

func ScanlnInt() (int, error) {
	var input string
	var n int
	var err error
	if _, err = fmt.Scanln(&input); err != nil {
		return 0, errors.New("input err,not int")
	}
	if n, err = strconv.Atoi(input); err != nil {
		return 0, errors.New("input err,not int")
	}
	return n, nil

}
