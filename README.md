# Razorpay Go Client

Golang bindings for interacting with the Razorpay API

## Usage

You need to setup your key and secret using the following:
You can find your API keys at <https://dashboard.razorpay.com/#/app/keys>.

```go
package main

import (
	"context"
	"net/http"
	"time"

	razorpay "github.com/gamezop/razorpay-go/pkg/client"
	"github.com/gamezop/razorpay-go/pkg/resource"
)

func main() {
	httpClient := &http.Client{Timeout: time.Second * 30}
	client := razorpay.NewClient(httpClient, "<YOUR_API_KEY>", "<YOUR_API_SECRET>")

	// example api:
	contact, err = client.CreateContact(context.Background(), resource.RequestCreateContact{
		Name:        "Gaurav Kumar",
		Email:       "gaurav.kumar@example.com",
		Contact:     "9123456789",
		Type:        "customer",
		ReferenceID: "Acme Contact ID 12345",
	})
}
```