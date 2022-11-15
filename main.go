package main

import (
	"fmt"

	"github.com/7ail/toy/internal/bootstrap"
)

func main() {
	ids := []int{1, 2, 3}

	for _, id := range ids {
		fmt.Printf("CASE %v==========\n", id)
		user := bootstrap.NewUserManager(id)
		_, err := user.Name()
		if user.IsNotFoundError(err) {
			fmt.Println(fmt.Errorf("user.Name(): %v", err))
			fmt.Println("Create user")
		}

		if user.IsInternalServerError(err) {
			fmt.Println(fmt.Errorf("user.Name(): %v", err))
			fmt.Println("Log in datadog")
			fmt.Println("Increment StatsD for PagerDuty")
		}

		if retryIn, e := user.IsRateLimitError(err); e == nil {
			fmt.Println(fmt.Errorf("user.Name(): %v", err))
			fmt.Printf("Retry in %v seconds\n", retryIn)
		}
		fmt.Println("CASE ===========")
	}
}
