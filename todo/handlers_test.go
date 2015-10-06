package todo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

const (
	cType      = "application/json"
	entryPoint = "/todo"
)

var todoItem = &item{"Buy strawberries."}

// Test using httptest's Recorder.s
func TestAddHandlerUsingRecord(t *testing.T) {
	// Configuring the http router.
	e := echo.New()
	e.Post(entryPoint, AddHandler(InMemoryStore()))

	// Issuing request.
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", entryPoint, jsonBuffer(todoItem))
	r.Header.Set("Content-Type", cType)
	e.ServeHTTP(w, r)

	// Asserting on results.
	assert.Equal(t, w.Code, http.StatusCreated)
	assert.Contains(t, w.Body.String(), jsonString(todoItem))
}

// Endo to end test of the API server. This tests brings up a httptest
// server as a blackbox component and invokes API methods using standard
// HTTP calls.
func TestEnd2End(t *testing.T) {
	s := InMemoryStore()
	e := echo.New()
	e.Post(entryPoint, AddHandler(s))
	e.Get(entryPoint, GetHandler(s))

	srv := httptest.NewServer(e)
	defer srv.Close()
	addr := fmt.Sprintf("%s%s", srv.URL, entryPoint)

	// Issuing requests many times concurrently TSAND and VSAND we can detect
	// race conditions.
	// Care to check, please run go test -race :)
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Add.
			func() {
				r, _ := http.Post(addr, cType, jsonBuffer(todoItem))
				defer r.Body.Close()
				assert.Equal(t, r.StatusCode, http.StatusCreated)
			}()
			// Add with failed precondition.
			func() {
				r, _ := http.Post(addr, cType, bytes.NewBufferString(""))
				defer r.Body.Close()
				assert.Equal(t, r.StatusCode, http.StatusPreconditionFailed)
			}()
			// Get.
			func() {
				r, _ := http.Get(addr)
				defer r.Body.Close()
				if assert.Equal(t, r.StatusCode, http.StatusOK) {
					b, _ := ioutil.ReadAll(r.Body)
					assert.Contains(t, string(b), jsonString(todoItem))
				}
			}()
		}()
	}
	wg.Wait()
}

func jsonBuffer(i *item) *bytes.Buffer {
	b, _ := json.Marshal(i)
	return bytes.NewBuffer(b)
}

func jsonString(i *item) string {
	b, _ := json.Marshal(i)
	return string(b)
}
