package client

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gamezop/razorpay-go/pkg/requestor"
	"github.com/gamezop/razorpay-go/pkg/resource"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

const (
	testSecret = "your-secret-key"
	testAPI    = "your-api-key"
)

func TestContactAPIs(t *testing.T) {
	t.Run("create-funding-account", func(t *testing.T) {
		defer gock.Off()
		httpClient := &http.Client{}

		expected := map[string]interface{}{
			"id":           "cont_00000000000001",
			"entity":       "contact",
			"name":         "Gaurav Kumar",
			"contact":      "9123456789",
			"email":        "gaurav.kumar@example.com",
			"type":         "customer",
			"reference_id": "Acme Contact ID 12345",
			"batch_id":     nil,
			"active":       true,
			"created_at":   1545320320,
		}

		gock.InterceptClient(httpClient)
		gock.New(requestor.BASE_URL).
			Post("/contacts").
			BasicAuth(testAPI, testSecret).
			Reply(200).
			JSON(expected)
		rzpSvc := NewClient(httpClient, testAPI, testSecret)
		contact, err := rzpSvc.CreateContact(context.Background(), resource.RequestCreateContact{
			Name:        "Gaurav Kumar",
			Email:       "gaurav.kumar@example.com",
			Contact:     "9123456789",
			Type:        "customer",
			ReferenceID: "Acme Contact ID 12345",
		})
		require.Nil(t, err, "expected error to be nil")

		expectedBarr, err := json.Marshal(expected)
		require.Nil(t, err, "expected error to be nil")
		contactBarr, err := json.Marshal(contact)
		require.Nil(t, err, "expected error to be nil")
		require.JSONEq(t, string(expectedBarr), string(contactBarr))
	})
}
