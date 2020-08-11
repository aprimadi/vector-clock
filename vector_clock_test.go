package vclock

import (
  "reflect"
  "testing"
)

func TestGetSet(t *testing.T) {
  vc := VClock{}
  vc.Set("a", 1)
  if vc.Get("a") != 1 {
    t.Errorf("Got: %v, expect: %v", vc.Get("a"), 1)
  }
}

func TestTick(t *testing.T) {
  vc := VClock{}
  var i uint64
  for i = 1; i <= 10; i++ {
    vc.Tick("a")
    if vc.Get("a") != i {
      t.Errorf("Got: %v, expect: %v", vc.Get("a"), i)
    }
  }
}

func TestCopy(t *testing.T) {
  vc1 := VClock{"a": 1}
  vc2 := vc1.Copy()
  vc2.Tick("a")
  if vc1.Get("a") != 1 {
    t.Errorf("Got: %v, expect: %v", vc1.Get("a"), 1)
  }
  if vc2.Get("a") != 2 {
    t.Errorf("Got: %v, expect: %v", vc2.Get("a"), 2)
  }
}

func TestMerge(t *testing.T) {
  vc1 := VClock{"a": 1, "b": 1}
  vc2 := VClock{"b": 2, "c": 1}
  vc := vc1.Merge(vc2)
  expected := VClock{"a": 1, "b": 2, "c": 1}
  if !reflect.DeepEqual(vc, expected) {
    t.Errorf("Got: %v, expect: %v", vc, expected)
  }
}

func TestEqual(t *testing.T) {
  vc1 := VClock{"a": 1, "b": 2}
  vc2 := VClock{"a": 1, "b": 2}
  vc3 := VClock{"a": 1, "b": 2, "c": 3}
  if !vc1.Equal(vc2) {
    t.Errorf("%v, %v should be equal", vc1, vc2)
  }
  if vc1.Equal(vc3) {
    t.Errorf("%v, %v should not be equal", vc1, vc3)
  }
}

func TestDescendant(t *testing.T) {
  vc := VClock{"a": 1, "b": 1}
  vc1 := VClock{"a": 1, "b": 2}
  vc2 := VClock{"a": 1, "c": 1}
  vc3 := VClock{"a": 1}
  vc4 := VClock{"a": 1, "b": 1, "c": 1}
  if !vc1.Descendant(vc) {
    t.Errorf("%v should be descendant of %v", vc1, vc)
  }
  if vc2.Descendant(vc) {
    t.Errorf("%v should not be descendant of %v", vc2, vc)
  }
  if vc3.Descendant(vc) {
    t.Errorf("%v should not be descendant of %v", vc3, vc)
  }
  if !vc4.Descendant(vc) {
    t.Errorf("%v should be descendant of %v", vc4, vc)
  }
}

func TestRelation(t *testing.T) {
  vc := VClock{"a": 1, "b": 1}
  vc0 := VClock{"a": 1, "b": 1}
  vc1 := VClock{"a": 1, "b": 1, "c": 1}
  vc2 := VClock{"b": 1}
  vc3 := VClock{"a": 1, "c": 1}
  if vc.Relation(vc0) != Equal {
    t.Errorf("%v, %v should be equal", vc, vc0)
  }
  if vc.Relation(vc1) != Ancestor {
    t.Errorf("%v should be ancestor of %v", vc, vc1)
  }
  if vc.Relation(vc2) != Descendant {
    t.Errorf("%v should be descendant of %v", vc, vc2)
  }
  if vc.Relation(vc3) != Conflict {
    t.Errorf("%v should conflict with %v", vc, vc3)
  }
}
