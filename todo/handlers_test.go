package todo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/labstack/echo"
)

const cType = "application/json"

var todoItem = &item{"Buy strawberries."}

// Test using httptest's Recorder.s
func TestAddHandlerUsingRecord(t *testing.T) {
	// Configuring the http router.
	e := echo.New()
	e.Post("/todo", AddHandler(InMemoryStore()))

	// Issuing request.
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/todo", jsonBuffer(todoItem))
	r.Header.Set("Content-Type", cType)
	e.ServeHTTP(w, r)

	// Asserting on results.
	if w.Code != http.StatusCreated {
		t.Errorf("Got:%v, want:%v", w.Code, http.StatusCreated)
	}
	if !strings.Contains(w.Body.String(), jsonString(todoItem)) {
		t.Errorf("Got:%s, want:%v", w.Body.String(), jsonString(todoItem))
	}
}

// Endo to end test of the API server. This tests brings up a httptest
// server as a blackbox component and invokes API methods using standard
// HTTP calls.
func TestEnd2End(t *testing.T) {
	s := InMemoryStore()
	e := echo.New()
	e.Post("/todo", AddHandler(s))
	e.Get("/todo", GetHandler(s))

	srv := httptest.NewServer(e)
	defer srv.Close()
	addr := fmt.Sprintf("%s/todo", srv.URL)

	// Issuing requests many times concurrently TSAND and VSAND we can detect
	// race conditions.
	// Care to check, please run go test -race :)
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Add.
			func() {
				r, _ := http.Post(addr, cType, jsonBuffer(todoItem))
				defer r.Body.Close()
				if r.StatusCode != http.StatusCreated {
					t.Errorf("Got:%v, want:%v",
						r.StatusCode, http.StatusCreated)
				}
			}()
			// Add with failed precondition.
			func() {
				r, _ := http.Post(addr, cType, bytes.NewBufferString(""))
				if r.StatusCode != http.StatusPreconditionFailed {
					t.Errorf("Got:%v, want:%v",
						r.StatusCode, http.StatusPreconditionFailed)
				}
			}()
			// Get.
			func() {
				r, _ := http.Get(addr)
				defer r.Body.Close()
				b, _ := ioutil.ReadAll(r.Body)
				if r.StatusCode != http.StatusOK {
					t.Errorf("Got:%v, want:%v", r.StatusCode, http.StatusOK)
				}
				str := jsonString(todoItem)
				if !strings.Contains(string(b), str) {
					t.Errorf("Got:%s, want:%v", string(b), str)
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
