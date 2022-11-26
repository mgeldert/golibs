package set

type Set[T comparable] map[T]struct{}

func Add[T comparable](set map[T]struct{}, value T) bool {
	setLength := len(set)
	set[value] = struct{}{}
	return len(set) != setLength
}

func Remove[T comparable](set map[T]struct{}, value T) bool {
	setLength := len(set)
	delete(set, value)
	return len(set) != setLength
}

func Contains[T comparable](set map[T]struct{}, value T) bool {
	_, exists := set[value]
	return exists
}

func Union[T comparable](firstSet map[T]struct{}, otherSets ...map[T]struct{}) map[T]struct{} {
	union := map[T]struct{}{}
	for k := range firstSet {
		union[k] = struct{}{}
	}
	for _, s := range otherSets {
		for k := range s {
			union[k] = struct{}{}
		}
	}
	return union
}

func Intersection[T comparable](firstSet map[T]struct{}, otherSets ...map[T]struct{}) map[T]struct{} {
	intersection := map[T]struct{}{}
	for key := range firstSet {
		intersects := true
		for _, variadicSet := range otherSets {
			if _, exists := variadicSet[key]; !exists {
				intersects = false
				break
			}
		}
		if intersects {
			intersection[key] = struct{}{}
		}
	}
	return intersection
}

func ToArray[T comparable](set map[T]struct{}) []T {
	values := []T{}
	for k := range set {
		values = append(values, k)
	}
	return values
}
