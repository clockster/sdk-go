package clockster

import (
	"errors"
	"net/http"
	"strings"
	"time"
)

// DefaultBaseURL is production. A demo stand answers the same API at another host.
const DefaultBaseURL = "https://api.clockster.com"

// DefaultTimeout applies to each request, and is what the client builds its own http.Client with.
const DefaultTimeout = 30 * time.Second

// Version is this package, as it goes out in the User-Agent.
const Version = "0.1.0"

// ErrNoToken is what New answers without a key.
var ErrNoToken = errors.New("clockster: a company API key is required, issued under Settings, API")

// transport is the half of the client that is the same on every operation: where the calls go, who
// they are from, and what carries them. Held by the Client and by every namespace on it.
type transport struct {
	token      string
	baseURL    string
	userAgent  string
	httpClient *http.Client
	timeout    time.Duration
}

// Option changes the client rather than one call.
type Option func(*transport)

// WithBaseURL points the client at a demo stand instead of production.
func WithBaseURL(baseURL string) Option {
	return func(t *transport) {
		t.baseURL = strings.TrimRight(baseURL, "/")
	}
}

// WithHTTPClient replaces the transport: a proxy, a retrying wrapper, a recording one. Its own
// timeout is used rather than DefaultTimeout, and closing it stays yours.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(t *transport) {
		t.httpClient = httpClient
	}
}

// WithTimeout replaces DefaultTimeout on the client this package builds. It does nothing to one
// supplied by WithHTTPClient, which carries its own and is not ours to change.
func WithTimeout(timeout time.Duration) Option {
	return func(t *transport) {
		t.timeout = timeout
	}
}

// WithUserAgent names your integration in our request log, which is worth doing when several talk
// to the same company. The default is clockster-go and this package's version.
func WithUserAgent(userAgent string) Option {
	return func(t *transport) {
		t.userAgent = userAgent
	}
}

// New is the client. The key is issued in the web application, under Settings, API, and
// authenticates one company: everything a call can see and change belongs to it.
//
//	clockster, err := clockster.New(os.Getenv("CLOCKSTER_TOKEN"))
//	if err != nil {
//		return err
//	}
//
//	users, err := clockster.Users.List(ctx, &clockster.UsersListParams{
//		PerPage: clockster.Set(int64(100)),
//		Include: []string{"location"},
//	})
//
// A method answers the parsed body, so rows are users.Data, and a refusal is an error rather than
// a status to read — see [Error]. A Client is safe for concurrent use.
func New(token string, opts ...Option) (*Client, error) {
	if token == "" {
		return nil, ErrNoToken
	}

	held := &transport{
		token:     token,
		baseURL:   DefaultBaseURL,
		userAgent: "clockster-go/" + Version,
		timeout:   DefaultTimeout,
	}

	for _, opt := range opts {
		opt(held)
	}

	// Built after the options rather than before, so a timeout meant for ours is not written onto
	// somebody else's client.
	if held.httpClient == nil {
		held.httpClient = &http.Client{Timeout: held.timeout}
	}

	return newClient(held), nil
}
