package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printGreeting(ctx); err != nil {
			fmt.Printf("%v", err)
			cancel()
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printFarewell(ctx); err != nil {
			fmt.Printf("%v", err)
		}
	}()
	wg.Wait()
}
func printGreeting(ctx context.Context) error {
	greeting, err := genGreeting(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s, world!\n", greeting)
	return nil
}
func printFarewell(ctx context.Context) error {
	farewell, err := genFarewell(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s, world!\n", farewell)
	return nil
}
func genGreeting(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	switch locate, err := locate(ctx); {
	case err != nil:
		return "", err
	case locate == "EN/US":

		return "hello", nil
	}
	return "", fmt.Errorf("unsupported locate")
}
func genFarewell(ctx context.Context) (string, error) {
	switch locate, err := locate(ctx); {
	case err != nil:
		return "", err

	case locate == "EN/US":
		return "goodbye", nil
	}
	return "", fmt.Errorf("unsupported locate")
}
func locate(ctx context.Context) (string, error) {
	if deadline, ok := ctx.Deadline(); ok {
		if deadline.Sub(time.Now().Add(1*time.Minute)) <= 0 {
			return "", context.DeadlineExceeded
		}
	}
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(1 * time.Minute):
		return "EN/US", nil
	}
}
