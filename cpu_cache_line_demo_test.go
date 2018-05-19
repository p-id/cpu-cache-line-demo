package cacheline

import (
	"runtime"
	"sync"
	"testing"
)

func testAtomicIncrement(sc SharedContext) {
	runtime.GOMAXPROCS(4)
	paraNum := 10000
	addTimes := 10000
	var wg sync.WaitGroup
	wg.Add(paraNum * 3)
	for i := 0; i < paraNum; i++ {
		go func() {
			for j := 0; j < addTimes; j++ {
				sc.IncrementAllCounters()
			}
			wg.Done()
		}()
	}
	for i := 0; i < paraNum; i++ {
		go func() {
			for j := 0; j < addTimes; j++ {
				sc.FetchCounterA()
			}
			wg.Done()
		}()
	}
	for i := 0; i < paraNum; i++ {
		go func() {
			for j := 0; j < addTimes; j++ {
				sc.FetchCounterB()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func testSharedContextImpl(sc SharedContext, t *testing.T) {
	crrCounterA := sc.FetchCounterA()
	sc.IncrementCounterA()
	if (crrCounterA + 1) != sc.FetchCounterA() {
		t.Error("IncrementCounterA() failed to increment from previous value")
	}

	crrCounterB := sc.FetchCounterB()
	sc.IncrementCounterB()
	if (crrCounterB + 1) != sc.FetchCounterB() {
		t.Error("IncrementCounterB() failed to increment from previous value")
	}

	crrCounterA = sc.FetchCounterA()
	crrCounterB = sc.FetchCounterB()
	sc.IncrementAllCounters()

	if (crrCounterA+1) != sc.FetchCounterA() ||
		(crrCounterB+1) != sc.FetchCounterB() {
		t.Error("IncrementAllCounters() failed to increment from previous value")
	}
}

func BenchmarkSharedContextIncrementWithPadding(b *testing.B) {
	scwp := &SharedContextWithPadding{}
	b.ResetTimer()
	testAtomicIncrement(scwp)
}

func BenchmarkSharedContextIncrementNoPadding(b *testing.B) {
	scnp := &SharedContextNoPadding{}
	b.ResetTimer()
	testAtomicIncrement(scnp)
}

func TestSharedContextWithPadding(t *testing.T) {
	scwp := &SharedContextWithPadding{}
	testSharedContextImpl(scwp, t)
}

func TestSharedContextNoPadding(t *testing.T) {
	scnp := &SharedContextNoPadding{}
	testSharedContextImpl(scnp, t)
}
