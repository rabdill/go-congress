package congress

// Client holds the connection information for sending data to the API.
type Client struct {
	// Endpoint is the URL of the ProPublica Congress API
	Endpoint string

	// Key is the user's ProPublica API key
	Key string
}
