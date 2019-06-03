package wgo

import (
	"testing"
	"time"
)

func TestSimpleCase(t *testing.T) {
	done := make(chan struct{})
	Launch(func() {
		defer close(done)
		time.Sleep(200 * time.Millisecond)

	})
	if ok := Wait(time.Duration(time.Second)); !ok {
		t.Fatal("failed to wait")
	}
	select {
	case <-done:
	default:
		t.Fatal("Wait() specified but not working")
	}
}

func TestWaitTimeout(t *testing.T) {
	done := make(chan struct{})
	Launch(func() {
		defer close(done)
		time.Sleep(time.Second)
	})
	if ok := Wait(time.Duration(time.Millisecond)); ok {
		t.Fatal("expected to exceed timeout, but not happening")
	}
	select {
	case <-done:
		t.Fatal("expected not done, but done")
	default:
	}
}

func TestEmbeddedCase(t *testing.T) {
	done := make(chan struct{})
	Launch(func() {
		Launch(func() {
			Launch(func() {
				defer close(done)
				time.Sleep(200 * time.Millisecond)
			})
		})
	})
	if ok := Wait(time.Duration(5 * time.Second)); !ok {
		t.Fatal("failed to wait")
	}
	select {
	case <-done:
	default:
		t.Fatal("Wait() specified but not working")
	}
}
