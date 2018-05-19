package cacheline

import "sync/atomic"

// SharedContext Common interface for update/load of SharedContext variables
type SharedContext interface {
	IncrementAllCounters()
	FetchCounterA() uint64
	IncrementCounterA()
	IncrementCounterB()
	FetchCounterB() uint64
}

// SharedContextNoPadding vanilla container with no CPU cache-line padding
type SharedContextNoPadding struct {
	counterA uint64
	counterB uint64
}

// IncrementAllCounters Increments all counters in shared context
func (scnp *SharedContextNoPadding) IncrementAllCounters() {
	atomic.AddUint64(&scnp.counterA, 1)
	atomic.AddUint64(&scnp.counterB, 1)
}

// IncrementCounterA Increments first counter indexed by B in SharedContext via atomic primitive
func (scnp *SharedContextNoPadding) IncrementCounterA() {
	atomic.AddUint64(&scnp.counterA, 1)
}

// IncrementCounterB Increments counter indexed by A in SharedContext via atomic primitive
func (scnp *SharedContextNoPadding) IncrementCounterB() {
	atomic.AddUint64(&scnp.counterB, 1)
}

// FetchCounterA return counter indexed by A
func (scnp *SharedContextNoPadding) FetchCounterA() uint64 {
	return atomic.LoadUint64(&scnp.counterA)
}

// FetchCounterB return counter indexed by B
func (scnp *SharedContextNoPadding) FetchCounterB() uint64 {
	return atomic.LoadUint64(&scnp.counterB)
}

// SharedContextWithPadding container with CPU cache-line padding
type SharedContextWithPadding struct {
	counterA uint64
	_p1      [8]uint64
	counterB uint64
}

// IncrementAllCounters Increments all counters in shared context
func (scwp *SharedContextWithPadding) IncrementAllCounters() {
	atomic.AddUint64(&scwp.counterA, 1)
	atomic.AddUint64(&scwp.counterB, 1)
}

// FetchCounterA return counter indexed by A
func (scwp *SharedContextWithPadding) FetchCounterA() uint64 {
	return atomic.LoadUint64(&scwp.counterA)
}

// FetchCounterB return counter indexed by B
func (scwp *SharedContextWithPadding) FetchCounterB() uint64 {
	return atomic.LoadUint64(&scwp.counterB)
}

// IncrementCounterA Increments first counter indexed by B in SharedContext via atomic primitive
func (scwp *SharedContextWithPadding) IncrementCounterA() {
	atomic.AddUint64(&scwp.counterA, 1)
}

// IncrementCounterB Increments counter indexed by A in SharedContext via atomic primitive
func (scwp *SharedContextWithPadding) IncrementCounterB() {
	atomic.AddUint64(&scwp.counterB, 1)
}
