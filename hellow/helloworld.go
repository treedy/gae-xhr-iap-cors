// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START gae_go111_app]

// Sample helloworld is an App Engine app.
package main

// [START import]
import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// [END import]
// [START main_func]

func main() {
	http.HandleFunc("/", indexHandler)

	// [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
	// [END setting_port]
}

// [END main_func]

// [START indexHandler]

// indexHandler responds to requests with our greeting.
func indexHandler(w http.ResponseWriter, r *http.Request) {

	log.Printf("Host is %s", r.Host)
	myHostname := strings.SplitAfterN(r.Host, "-dot-", 2)[1] // Get the base hostname from our own hostname

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "https://"+myHostname)
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	if r.Method == "OPTIONS" { // If it's an OPTIONS request, don't return any other data
		return
	}
	fmt.Fprint(w, "Hello, World!")
}

// [END indexHandler]
// [END gae_go111_app]
