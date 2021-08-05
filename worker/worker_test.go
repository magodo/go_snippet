package worker_test

import (
	"context"
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
		expect = append(expect, i)
		if !wp.AddTask(worker.NewTask(
			context.TODO(),
			i,
			func(_ context.Context, v interface{}) (interface{}, error) { return v, nil },
		)) {
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
		if !wp.AddTask(worker.NewTask(
			context.TODO(),
			i,
			func(_ context.Context, v interface{}) (interface{}, error) {
				v = v.(int)
				if v == 2 {
					return v, err
				}
				return v, nil
			},
		)) {
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
		if !wp.AddTask(worker.NewTask(
			context.TODO(),
			i,
			func(_ context.Context, v interface{}) (interface{}, error) {
				return v, nil
			},
		)) {
			break
		}
	}

	require.ErrorIs(t, wp.Done(), err)
}
