package containers_test

import (
	"testing"

	"github.com/SharkByteSoftware/go-snk/containers"
	"github.com/SharkByteSoftware/go-snk/containers/lists"
	"github.com/SharkByteSoftware/go-snk/containers/sets"
	"github.com/SharkByteSoftware/go-snk/containers/stacks"
	"github.com/stretchr/testify/assert"
)

func TestContainers_Interface(t *testing.T) {
	list := lists.New(1, 2, 3)
	runAssertions[int](t, list)

	set := sets.New(1, 2, 3)
	runAssertions[int](t, set)

	stack := stacks.New[int](1, 2, 3)
	runAssertions[int](t, stack)
}

func runAssertions[T any](t *testing.T, container containers.Container[T]) {
	assert.False(t, container.IsEmpty())
	assert.Equal(t, 3, container.Size())
	assert.Len(t, container.Values(), 3)

	container.Clear()
	assert.True(t, container.IsEmpty())
}
