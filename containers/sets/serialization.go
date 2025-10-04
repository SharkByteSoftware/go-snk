package sets

import (
	"encoding/json"
	"fmt"
)

// ToJSON outputs the JSON representation of the set.
func (s *Set[T]) ToJSON() ([]byte, error) {
	result, err := json.Marshal(s.Values())
	if err != nil {
		return nil, fmt.Errorf("failed to marshal set: %w", err)
	}

	return result, nil
}

// FromJSON populates the set from the input JSON representation.
func (s *Set[T]) FromJSON(data []byte) error {
	var items []T

	err := json.Unmarshal(data, &items)
	if err != nil {
		return fmt.Errorf("failed to unmarshal set: %w", err)
	}

	s.Clear()
	s.Add(items...)

	return nil
}

// UnmarshalJSON @implements json.Unmarshaler.
func (s *Set[T]) UnmarshalJSON(bytes []byte) error {
	return s.FromJSON(bytes)
}

// MarshalJSON @implements json.Marshaler.
func (s *Set[T]) MarshalJSON() ([]byte, error) {
	return s.ToJSON()
}
