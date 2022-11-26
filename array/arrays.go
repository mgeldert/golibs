package array

func Contains[T comparable](haystack []T, needle T) bool {
	for _, value := range haystack {
		if value == needle {
			return true
		}
	}
	return false
}

func IndexOf[T comparable](haystack []T, needle T) int {
	for index, value := range haystack {
		if value == needle {
			return index
		}
	}
	return -1
}

func InsertAt[T comparable](array *[]T, value T, position int) {
	if position == 0 {
		*array = append([]T{value}, *array...)
	} else if position >= len(*array) {
		*array = append(*array, value)
	} else {
		*array = append((*array)[:position+1], (*array)[position:]...)
		(*array)[position] = value
	}
}

func DeleteIndex[T comparable](array []T, index int) ([]T, error) {
	if index > len(array)-1 {
		return array, fmt.Errorf("Index out of bounds")
	}
	return append(array[:index], array[index+1:]...), nil
}

func DeleteFirstOccurrence[T comparable](array []T, value T) []T {
	firstOccurrenceIndex := -1

	for index, element := range array {
		if element == value {
			firstOccurrenceIndex = index
			break
		}
	}

	if firstOccurrenceIndex > -1 {
		return append(array[:firstOccurrenceIndex], array[firstOccurrenceIndex+1:]...)
	}

	return array
}

func Push[T comparable](array *[]T, value T) {
	InsertAt(array, value, 0)
}

func Pop[T comparable](array *[]T) (T, error) {
	var zero T
	var err error

	if len(*array) == 0 {
		return zero, fmt.Errorf("Array is empty")
	}

	value := (*array)[0]
	if *array, err = DeleteIndex(*array, 0); err != nil {
		return zero, err
	}

	return value, nil
}

func Shift[T comparable](array *[]T, value T) {
	InsertAt(array, value, len(*array))
}

func Unshift[T comparable](array *[]T) (T, error) {
	var zero T
	var err error

	if len(*array) == 0 {
		return zero, fmt.Errorf("Array is empty")
	}

	value := (*array)[len(*array)-1]
	if *array, err = DeleteIndex(*array, len(*array)-1); err != nil {
		return zero, err
	}

	return value, nil
}
