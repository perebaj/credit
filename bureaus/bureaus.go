package bureaus

import "net/http"

// Client is a client for any bureau.
type Client struct {
	Client *http.Client
	URL    string
}
