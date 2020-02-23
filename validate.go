package cm

import "fmt"

type IntValidator func(int) error

func ValidateIntBetween(lower, upper int) IntValidator {
	return IntValidator(func(v int) error {
		switch {
		case v < lower:
			return fmt.Errorf("int %d is below lower limit %d", v, lower)
		case v > upper:
			return fmt.Errorf("int %d is above upper limit %d", v, upper)
		default:
			return nil
		}
	})
}

func ValidateIntInSet(values ...int) IntValidator {
	set := make(map[int]struct{}, len(values))
	for _, v := range values {
		set[v] = struct{}{}
	}

	return IntValidator(func(v int) error {
		if _, ok := set[v]; !ok {
			keys := make([]int, 0, len(set))
			for k := range set {
				keys = append(keys, k)
			}

			return fmt.Errorf("int %d not in set %v", v, keys)
		}

		return nil
	})
}

type StringValidator func(string) error

func ValidateStringInSet(values ...string) StringValidator {
	set := make(map[string]struct{}, len(values))
	for _, v := range values {
		set[v] = struct{}{}
	}

	return StringValidator(func(v string) error {
		if _, ok := set[v]; !ok {
			keys := make([]string, 0, len(set))
			for k := range set {
				keys = append(keys, k)
			}

			return fmt.Errorf("string '%s' not in set %v", v, keys)
		}

		return nil
	})
}
