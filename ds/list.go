package ds

type List[T any] struct {
	next *List[T]
	list *List[T]
	prev *List[T]
}

func NewList[T any]() *List[T] {
	return &List[T]{}
}

func (l *List[T]) Back() *T {
	return nil
}

func (l *List[T]) Front() *T {
	return nil
}

func (l *List[T]) Init() *List[T] {
	return nil
}

func (l *List[T]) InsertAfter(v any, mark *T) *T {
	return nil
}

func (l *List[T]) InsertBefore(v any, mark *T) *T {
	return nil
}

func (l *List[T]) Len() int {
}

func (l *List[T]) MoveAfter(e, mark *T) {
}

func (l *List[T]) MoveBefore(e, mark *T) {
}

func (l *List[T]) MoveToBack(e *T) {
}

func (l *List[T]) MoveToFront(e *T) {
}

func (l *List[T]) PushBack(v any) *T {
	return nil
}

func (l *List[T]) PushBackList(other *List[T]) {
}

func (l *List[T]) PushFront(v any) *T {
	return nil
}

func (l *List[T]) PushFrontList(other *List[T]) {
}

func (l *List[T]) Remove(e *T) any {
	return nil
}
