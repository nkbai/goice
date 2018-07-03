package turn

import "github.com/nkbai/goice/stun"

// DontFragmentAttr represents DONT-FRAGMENT attribute.
type DontFragmentAttr bool

// AddTo adds DONT-FRAGMENT attribute to message.
func (DontFragmentAttr) AddTo(m *stun.Message) error {
	m.Add(stun.AttrDontFragment, nil)
	return nil
}

// IsSet returns true if DONT-FRAGMENT attribute is set.
func (DontFragmentAttr) IsSet(m *stun.Message) bool {
	_, err := m.Get(stun.AttrDontFragment)
	return err == nil
}

// DontFragment is shorthand for DontFragmentAttr.
var DontFragment DontFragmentAttr = true