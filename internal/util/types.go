package util

func IntToInt32(i int) (int32, error) {
	if i < 0 || i > 2147483647 {
		return 0, ErrInt32OutOfRange
	}

	return int32(i), nil
}

func Int32ToInt64(i int32) (int64, error) {
	if i < 0 {
		return 0, ErrInt64OutOfRange
	}

	return int64(i), nil
}

func ConvertToInt32Slice(input []int) ([]int32, error) {
	result := make([]int32, len(input))
	var err error
	for i, v := range input {
		result[i], err = IntToInt32(v)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}
