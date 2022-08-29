package generics

type Stack[T any] struct {
	values []T
}

func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false

		// ** alternative
		// zero := reflect.Zero(reflect.TypeOf(s.values).Elem())
		// return zero.Interface().(T), false
	}

	last := len(s.values) - 1
	el := s.values[last]
	s.values = s.values[:last]

	return el, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}
