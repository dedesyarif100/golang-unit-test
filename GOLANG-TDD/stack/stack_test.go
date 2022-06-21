package stack_test

import (
	"testing"
	"GOLANG-TDD/stack"
	"github.com/stretchr/testify/assert"
)

// Requirements :
// * A STACK IS EMPTY ON CONSTRUCTION
// * A STACK HAS SIZE 0 ON CONSTRUCTION
// * AFTER & PUSH TO AN EMPTY STACK (n > 0), THE STACK IS NON-EMPTY AND ITS SIZE EQUALS N


func TestNewSet(t *testing.T) {
	s := stack.New()
	t.Run("A STACK IS EMPTY ON CONSTRUCTION", func(t *testing.T) {
		assert.True(t, s.IsEmpty())
	})

	t.Run("A STACK HAS SIZE 0 ON CONSTRUCTION", func(t *testing.T) {
		assert.Equal(t, 0, s.Size())
	})
}

func TestInsert(t *testing.T) {
	t.Run("AFTER & PUSH TO AN EMPTY STACK (n > 0), THE STACK IS NON-EMPTY AND ITS SIZE EQUALS N", func(t *testing.T) {
		s := stack.New()
		s.Push(1)
		s.Push(2)
		s.Push(3)
		assert.False(t, s.IsEmpty())
		assert.Equal(t, 3, s.Size())
	})

	t.Run("IF ONE PUSH X THEN POPS, THE VALUE POPPED IS X, AND THE SIZE IS ONE LESS THAN OLD SIZE", func(t *testing.T) {
		s := stack.New()
		s.Push(1)
		s.Push(2)
		s.Push(6)
		assert.Equal(t, 3, s.Size())
		val, _ := s.Pop()
		assert.Equal(t, 6, val)
		assert.Equal(t, 2, s.Size())
	})

	t.Run("IF ONE PUSHED X THEN PEEKS, THE VALUE RETURNED IS X, BUT THE SIZE STAYS THE SAME", func(t *testing.T) {
		s := stack.New()
		s.Push(1)
		s.Push(2)
		s.Push(6)
		assert.Equal(t, 3, s.Size())
		val, _ := s.Peek()
		assert.Equal(t, 6, val)
		assert.Equal(t, 3, s.Size())
	})
}


func TestError(t *testing.T) {
	t.Run("POPPING FROM AN EMPTY STACK RETURN AN ERROR: ErrorNoSuchElement", func(t *testing.T) {
		s := stack.New()
		_, err := s.Pop()
		if err == nil {
			t.Fail()
			t.Logf("EXPECT ERROR IS NOT NIL, BUT GOT %v", err)
		}
		assert.Equal(t, stack.ErrNoSuchElement, err)
	})

	t.Run("PEEKING INTO AN EMPTY STACK RETURN AN ERROR: ErrorNoSuchElement", func(t *testing.T) {
		s := stack.New()
		_, err := s.Peek()
		if err == nil {
			t.Fail()
			t.Logf("EXPECT ERROR IS NOT NIL, BUT GOT %v", err)
		}
		assert.Equal(t, stack.ErrNoSuchElement, err)
	})
}
