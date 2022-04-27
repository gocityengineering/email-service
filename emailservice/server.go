package emailservice

import (
	"fmt"
	"io/ioutil"
	"log"
  "net"
	"net/http"
)

func Serve(server string, port int, dryrun bool) error {
  // port valid and available?
  ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
  if err != nil {
    return err
  }
  ln.Close()

  if dryrun {
    return nil
  }

	http.Handle("/", handler())
	http.Handle("/health", healthHandler())
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", server, port), nil))

  return nil
}

func handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			fmt.Fprintf(w, `POST JSON data to send mail
{
  "recipients": [ "USER1@DOMAIN", "USER2@DOMAIN" ],
  "subject": "SUBJECT"
  "markdownBody": "MARKDOWN"
}`)
		case "POST":
			bytes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "500: error in ioutil.ReadAll: %v", err)
				return
			}
			err = processData(bytes, false)
			if err != nil {
				// set status code
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "500: error in processData: %v", err)
				return
			}
			fmt.Fprintf(w, `OK`)
		}
	})
}

func healthHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `OK`)
	})
}
