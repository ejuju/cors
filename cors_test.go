package cors

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestSet(t *testing.T) {
	// init test server
	router := mux.NewRouter()

	// register test endpoint
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("CORS protected resource"))
	})

	// register cors middleware
	router.Use(Set(Policy{
		AllowOrigin: func(origin string) string {
			return origin // allow all origins
		},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin", "X-Test", "Accept", "Accept-Language", "Content-Language"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"X-Test"},
		CacheMaxAge:      120,
	}))

	// send test request
	req := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/", nil)
	req.Header.Set("Origin", "http://example.com/page1")
	resrec := httptest.NewRecorder()
	router.ServeHTTP(resrec, req)

	// process result
	fmt.Println(resrec.HeaderMap, resrec.Code)
	body, err := ioutil.ReadAll(resrec.Body)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(body))
}
