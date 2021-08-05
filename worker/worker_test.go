package worker_test

import (
	"errors"
	"sort"
	"testing"
	"worker"

	"github.com/stretchr/testify/require"
)

func TestSuccessfulRun(t *testing.T) {
	wp := worker.NewWorkPool(8)
	expect := []int{}
	results := []int{}
	wp.Run(func(v interface{}) error {
		results = append(results, v.(int))
		return nil
	})

	for i := 0; i < 10; i++ {
		num := i
		expect = append(expect, num)
		if !wp.AddTask(func() (interface{}, error) { return num, nil }) {
			break
		}
	}

	require.NoError(t, wp.Done())

	// The order in results might diff due to the parallelism nature of the worker pool
	sort.Ints(results)
	require.Equal(t, expect, results)
}

func TestTaskError(t *testing.T) {
	wp := worker.NewWorkPool(8)

	results := []int{}
	wp.Run(func(v interface{}) error {
		results = append(results, v.(int))
		return nil
	})

	err := errors.New("Error")
	for i := 0; i < 10; i++ {
		num := i
		if !wp.AddTask(func() (interface{}, error) {
			if num == 2 {
				return num, err
			}
			return num, nil
		}) {
			break
		}
	}

	require.ErrorIs(t, wp.Done(), err)
}

func TestResultHandlerError(t *testing.T) {
	wp := worker.NewWorkPool(8)

	err := errors.New("Error")
	results := []int{}
	wp.Run(func(v interface{}) error {
		v = v.(int)
		if v == 2 {
			return err
		}
		results = append(results, v.(int))
		return nil
	})
	for i := 0; i < 10; i++ {
		num := i
		if !wp.AddTask(func() (interface{}, error) {
			return num, nil
		}) {
			break
		}
	}

	require.ErrorIs(t, wp.Done(), err)
}

func TestSkipTask(t *testing.T) {
	wp := worker.NewWorkPool(8)

	results := []int{}
	wp.Run(func(v interface{}) error {
		v = v.(int)
		results = append(results, v.(int))
		return nil
	})

	expect := []int{}
	for i := 0; i < 10; i++ {
		if i != 2 {
			expect = append(expect, i)
		}
		num := i
		if !wp.AddTask(func() (interface{}, error) {
			if num == 2 {
				return num, worker.ErrSkipTask
			}
			return num, nil
		}) {
			break
		}
	}

	require.NoError(t, wp.Done())

	// The order in results might diff due to the parallelism nature of the worker pool
	sort.Ints(results)
	require.Equal(t, expect, results)
}
