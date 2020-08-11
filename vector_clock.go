package vclock

// Vector clock is logically a list of tuple [(pid1, tick1), (pid2, tick2), ...]
// where pid is the process id and tick is the clock value.
type VClock map[string]uint64

type Relation uint8

const (
  Equal = iota
  Ancestor
  Descendant
  Conflict
)

// Advance tick for the given process id by 1
func (v VClock) Tick(pid string) {
  v[pid] = v[pid] + 1
}

// Get the tick for the given process id, return 0 if pid doesn't exist
func (v VClock) Get(pid string) uint64 {
  return v[pid]
}

// Set the tick for the given process id
func (v VClock) Set(pid string, tick uint64) {
  if (tick == 0) {
    panic("Tick must be positive value")
  }
  v[pid] = tick
}

// Copy this vector clock
func (v VClock) Copy() VClock {
  res := make(VClock, len(v))
  for pid := range v {
    res[pid] = v[pid]
  }
  return res
}

// Merge this vector clock with other vector clock, return the resulting vector
// clock
func (v VClock) Merge(other VClock) VClock {
  res := v.Copy()
  for pid := range other {
    if res[pid] < other[pid] {
      res[pid] = other[pid]
    }
  }
  return res
}

// Whether this vector clock is equal to the given vector clock
func (v VClock) Equal(other VClock) bool {
  if len(other) != len(v) {
    return false
  }

  for pid := range other {
    if v[pid] != other[pid] {
      return false
    }
  }
  return true
}

// Whether this vector clock is descendant of the given vector clock
func (v VClock) Descendant(other VClock) bool {
  // v is descendant of other iff:
  // - all elements in other is less than or equal than v, and
  // - v != other

  isEqual := (len(other) == len(v))
  for pid := range other {
    if other[pid] > v[pid] {
      return false
    } else if isEqual && other[pid] < v[pid] {
      isEqual = false
    }
  }

  return !isEqual
}

// Relation of this vector clock to the given vector clock
func (v VClock) Relation(other VClock) Relation {
  if (v.Equal(other)) {
    return Equal
  } else if (other.Descendant(v)) {
    return Ancestor
  } else if (v.Descendant(other)) {
    return Descendant
  } else {
    return Conflict
  }
}
