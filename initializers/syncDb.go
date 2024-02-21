package initializers

import (
	"fmt"
)

func SyncRedis() {
	result, err := DB.SetNX(ctx, "post-id", "1", 0).Result()

	if err != nil {
		panic(err)
	}

	if result {
		fmt.Println("Key was set")
	} else {
		fmt.Println("Key already exists")
	}
}
