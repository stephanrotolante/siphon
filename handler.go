package main

import (
	"io"
	"math/rand"
	"net/http"

	"github.com/stephanrotolante/siphon/config"
)

func Handler(siphonConfig *config.SiphonConfig) http.Handler {

	var client = &http.Client{}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if "GET" == r.Method || "DELETE" == r.Method {

			randomInt := rand.Intn(100)

			if randomInt < siphonConfig.Target.Probability {

				req, err := http.NewRequest(r.Method, siphonConfig.Target.TargetHost, r.Body)

				req.Header = r.Header
				if err != nil {
					panic(err)
				}

				resp, err := client.Do(req)

				if err != nil {
					panic(err)
				}

				defer resp.Body.Close()

				for key, values := range resp.Header {
					for _, v := range values {
						w.Header().Add(key, v)
					}

				}

				w.WriteHeader(resp.StatusCode)

				io.Copy(w, resp.Body)

				return
			}
		}

		req, err := http.NewRequest(r.Method, siphonConfig.Target.DefaultHost, r.Body)

		req.Header = r.Header
		if err != nil {
			panic(err)
		}

		resp, err := client.Do(req)

		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()

		for key, values := range resp.Header {
			for _, v := range values {
				w.Header().Add(key, v)
			}

		}

		w.WriteHeader(resp.StatusCode)

		io.Copy(w, resp.Body)

	})
}
