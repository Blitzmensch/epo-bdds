package bdds_test

import (
	"fmt"

	bdds "github.com/patent-dev/epo-bdds"
)

func ExampleNewClient() {
	// Every BDDS endpoint requires an EPO account; paid products also need
	// a subscription to that product.
	client, err := bdds.NewClient(&bdds.Config{
		Username: "your-epo-username",
		Password: "your-epo-password",
	})
	if err != nil {
		panic(err)
	}
	// client is ready; call ListProducts, GetProduct, DownloadFile, etc.
	fmt.Println(client != nil)
	// Output: true
}
