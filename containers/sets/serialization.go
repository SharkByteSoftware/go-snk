package sets

import (
	"encoding/json"
	"fmt"
)

// ToJSON outputs the JSON representation of the set.
func (set *Set[T]) ToJSON() ([]byte, error) {
	result, err := json.Marshal(set.Values())
	if err != nil {
		return nil, fmt.Errorf("failed to marshal set: %w", err)
	}

	return result, nil
}

// FromJSON populates the set from the input JSON representation.
func (set *Set[T]) FromJSON(data []byte) error {
	var items []T

	err := json.Unmarshal(data, &items)
	if err != nil {
		return fmt.Errorf("failed to unmarshal set: %w", err)
	}

	set.Clear()
	set.Add(items...)

	return nil
}

func (s *Set[T]) UnmarshalJSON(bytes []byte) error {
	return s.FromJSON(bytes)
}

func (s Set[T]) MarshalJSON() ([]byte, error) {
	return s.ToJSON()
}
