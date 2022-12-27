package ethernet

import "testing"

func FuzzFrameMarshalBinary(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		var f Frame
		if err := f.UnmarshalBinary(b); err != nil {
			return
		}
		if _, err := f.MarshalBinary(); err != nil {
			t.Fatal(err)
		}
	})
}

func FuzzFrameMarshalFCS(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		var f Frame
		if err := f.UnmarshalBinary(b); err != nil {
			return
		}
		if _, err := f.MarshalFCS(); err != nil {
			t.Fatal(err)
		}
	})
}

func FuzzVLANMarshalBinary(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		var v VLAN
		if err := v.UnmarshalBinary(b); err != nil {
			return
		}
		if _, err := v.MarshalBinary(); err != nil {
			t.Fatal(err)
		}
	})
}
