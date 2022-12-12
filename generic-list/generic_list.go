package genericlist

import (
	"errors"
)

type GenericList[T comparable] struct {
	data []T
}

func New[T comparable]() *GenericList[T] {
	return &GenericList[T]{
		data: []T{},
	}
}

func (l *GenericList[T]) Insert(value T) {
	l.data = append(l.data, value)
}

func (l *GenericList[T]) InsertAll(values ...T) {
	l.data = append(l.data, values...)
}

func (l *GenericList[T]) Get(i int) (T, error) {
	var value T
	if err := indexIsValid(i, len(l.data)); err != nil {
		return value, err
	}
	for it := 0; it < len(l.data); it++ {
		if i == it {
			value = l.data[it]
		}
	}
	return value, nil
}

func (l *GenericList[T]) Remove(index int) (T, error) {
	var value T
	if err := indexIsValid(index, len(l.data)); err != nil {
		return value, err
	}
	for it := 0; it < len(l.data); it++ {
		if index == it {
			value = l.data[it]
			l.data = append(l.data[:index], l.data[index+1:]...)
		}
	}
	return value, nil
}

func (l *GenericList[T]) RemoveByValue(value T) error {
	index, err := l.GetIndex(value)
	if err != nil {
		return err
	}
	l.data = append(l.data[:index], l.data[index+1:]...)
	return nil
}

func (l *GenericList[T]) GetIndex(value T) (int, error) {
	for index, val := range l.data {
		if val == value {
			return index, nil
		}
	}
	return 0, errors.New("value not found")
}

func (l *GenericList[T]) UpdateValue(oldValue, newValue T) error {
	index, err := l.GetIndex(oldValue)
	if err != nil {
		return err
	}
	l.data[index] = newValue
	return nil
}

func indexIsValid(index, length int) error {
	if index < 0 {
		return errors.New("index can't be negative")
	} else if index > length-1 {
		return errors.New("index is too high")
	} else {
		return nil
	}
}
