package auth

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/fatih/color"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type OAuth struct {
	config oauth2.Config
	token  oauth2.Token
	ctx    context.Context
	client *http.Client
}

func New(config oauth2.Config) *OAuth {
	return &OAuth{
		config: config,
	}
}

func (o *OAuth) Authenticate() {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	srv, wait := o.startCallbackServer(ctx)
	o.openAuthPage(ctx)
	wait.Wait()
	srv.Shutdown(ctx)
}

func (o *OAuth) Token() oauth2.Token {
	return o.token
}

// Client - Gets an HTTP client
func (o *OAuth) Client(ctx context.Context) *http.Client {
	return o.config.Client(ctx, &o.token)
}

func (o *OAuth) startCallbackServer(ctx context.Context) (*http.Server, *sync.WaitGroup) {
	srv := &http.Server{Addr: ":9999"}
	waiter := sync.WaitGroup{}
	http.HandleFunc("/oauth/callback", o.getOauthCallback(ctx, &waiter))
	go func() {
		// always returns error. ErrServerClosed on graceful close
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			// unexpected error. port in use?
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()
	// returning reference so caller can call Shutdown()
	return srv, &waiter
}

func (o *OAuth) openAuthPage(ctx context.Context) {
	// add transport for self-signed certificate to context
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	sslcli := &http.Client{Transport: tr}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, sslcli)

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := o.config.AuthCodeURL("state", oauth2.AccessTypeOffline)

	log.Println(color.CyanString("You will now be taken to your browser for authentication"))
	time.Sleep(1 * time.Second)
	open.Run(url)
	time.Sleep(1 * time.Second)
	log.Printf("Authentication URL: %s\n", url)
}

func (o *OAuth) getOauthCallback(ctx context.Context, waiter *sync.WaitGroup) func(w http.ResponseWriter, r *http.Request) {
	waiter.Add(1)
	return func(w http.ResponseWriter, r *http.Request) {
		queryParts, _ := url.ParseQuery(r.URL.RawQuery)

		log.Printf("Starting Callback Handler")
		// Use the authorization code that is pushed to the redirect
		// URL.
		code := queryParts["code"][0]
		log.Printf("code: %s\n", code)

		// Exchange will do the handshake to retrieve the initial access token.
		tok, err := o.config.Exchange(ctx, code)
		if err != nil {
			log.Fatal(err)
		}
		o.token = *tok

		// show succes page
		msg := "<p><strong>Success!</strong></p>"
		msg = msg + "<p>You are authenticated and can now return to the CLI.</p>"
		fmt.Fprintf(w, msg)

		// Done Processing
		waiter.Done()
	}
}
