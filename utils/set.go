package utils

type Set[T comparable] struct {
	mapValues map[T]string
}

func NewSet[T comparable]() Set[T] {
	set := Set[T]{}
	set.mapValues = map[T]string{}

	return set
}

func (set *Set[T]) HasValue(value T) bool {
	_, exists := set.mapValues[value]

	return exists
}

func (set *Set[T]) Add(value T) (T, bool) {
	exists := set.HasValue(value)

	if !exists {
		set.mapValues[value] = ""
	}

	return value, exists
}

func (set *Set[T]) AddAll(values []T) {
	for _, value := range values {
		set.mapValues[value] = ""
	}
}

func (set *Set[T]) Combine(otherSet *Set[T]) {
	for key := range otherSet.mapValues {
		set.mapValues[key] = ""
	}
}

func (set *Set[T]) Values() []T {
	allValues := []T{}
	for value := range set.mapValues {
		allValues = append(allValues, value)
	}

	return allValues
}

func (set *Set[T]) Remove(value T) (T, bool) {
	exists := set.HasValue(value)

	if !exists {
		return value, false
	}

	delete(set.mapValues, value)

	return value, true
}
