package cors

import (
	"net/http"
	"net/http/httptest"
	"strings"
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
			return "otherorigin.com" // allow all origins
		},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin", "X-Test", "Accept", "Accept-Language", "Content-Language"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"X-Test"},
		CacheMaxAge:      120,
	}))

	// send preflight request
	req := httptest.NewRequest(http.MethodOptions, "http://localhost:8080/", nil)
	req.Header.Set(KeyOrigin, "http://unauthorized.com")
	req.Header.Set(KeyRequestMethods, http.MethodDelete)
	req.Header.Set(KeyRequestHeaders, strings.Join([]string{
		"X-Test: value for x-test header",
		"Unauthorized: hehe",
	}, ","))
	resrec := httptest.NewRecorder()

	router.ServeHTTP(resrec, req)

	if resrec.Code != http.StatusBadRequest {
		t.Error("should not be authorized")
		return
	}

	// todo: send normal request

	// process result
	// body, err := ioutil.ReadAll(resrec.Body)
	// if err != nil {
	// 	t.Error(err)
	// }
	// fmt.Println(string(body))
}
