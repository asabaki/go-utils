package go_utils

var void any

// The Set type lets you store unique values of any type.
type Set[T comparable] map[T]any

// The Has() method returns a boolean indicating whether an element with the specified value exists in a Set object or not.
func (s Set[T]) Has(val T) bool {
	_, exists := s[val]
	return exists
}

// The Clear() method removes all elements from a Set object.
func (s Set[T]) Clear() {
	for K := range s {
		delete(s, K)
	}
}

// The Add() method inserts a new element with a specified value in to a Set object, if there isn't an element with the same value already in the Set.
func (s Set[T]) Add(val T) {
	s[val] = void
}

// The Delete() method removes a specified value from a Set object, if it is in the set.
// Returns `true` if value was already in Set; otherwise `false`.
func (s Set[T]) Delete(val T) bool {
	if _, ok := s[val]; ok {
		delete(s, val)
		return true
	}
	return false
}

// The Values() method returns a slice containing the values in the Set
func (s Set[T]) Values() []T {
	vals := make([]T, 0)
	for k := range s {
		vals = append(vals, k)
	}
	return vals
}

// The Size() method returns the number of (unique) elements in a Set object.
func (s Set[T]) Size() int {
	return len(s)
}
