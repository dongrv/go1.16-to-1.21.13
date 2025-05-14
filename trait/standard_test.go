package trait

import "testing"

func TestNewTrait(t *testing.T) {
	t.Run("slices", func(t *testing.T) { Slices() })
	t.Run("maps", func(t *testing.T) { Maps() })
	t.Run("context", func(t *testing.T) { Context() })
	t.Run("panic-recover", func(t *testing.T) { PanicRecover() })
}
