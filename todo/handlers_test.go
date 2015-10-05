package todo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
)

const contentType = "application/json"

var todoItem = &item{"Buy strawberries."}

// Test using httptest's Recorder.s
func TestAddHandlerUsingRecord(t *testing.T) {
	// Configuring the http router.
	e := echo.New()
	e.Post("/todo", NewAddHandler(NewStore()))

	// Issuing request.
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/todo", jsonBuffer(todoItem))
	r.Header.Set("Content-Type", contentType)
	e.ServeHTTP(w, r)

	// Asserting on results.
	if w.Code != http.StatusCreated {
		t.Errorf("Got:%v, want:%v", w.Code, http.StatusCreated)
	}
	if !strings.Contains(w.Body.String(), jsonString(todoItem)) {
		t.Errorf("Got:%s, want:%v", w.Body.String(), jsonString(todoItem))
	}
}

func TestEnd2End(t *testing.T) {
	s := NewStore()
	e := echo.New()
	e.Post("/todo", NewAddHandler(s))
	e.Get("/todo", NewGetHandler(s))

	srv := httptest.NewServer(e)
	defer srv.Close()
	addr := fmt.Sprintf("%s/todo", srv.URL)

	// Add.
	addResp, _ := http.Post(addr, contentType, jsonBuffer(todoItem))
	addBody, _ := ioutil.ReadAll(addResp.Body)
	if addResp.StatusCode != http.StatusCreated {
		t.Errorf("Got:%v, want:%v", addResp.StatusCode, http.StatusCreated)
	}
	if !strings.Contains(string(addBody), jsonString(todoItem)) {
		t.Errorf("Got:%s, want:%v", string(addBody), jsonString(todoItem))
	}

	// Add with failed precondition.
	addFPResp, _ := http.Post(
		addr, contentType, bytes.NewBufferString("Foooooo"))
	if addFPResp.StatusCode != http.StatusPreconditionFailed {
		t.Errorf("Got:%v, want:%v",
			addFPResp.StatusCode, http.StatusPreconditionFailed)
	}

	// Get.
	getResp, _ := http.Get(addr)
	getBody, _ := ioutil.ReadAll(getResp.Body)
	if getResp.StatusCode != http.StatusOK {
		t.Errorf("Got:%v, want:%v", getResp.StatusCode, http.StatusOK)
	}
	if !strings.Contains(string(getBody), jsonString(todoItem)) {
		t.Errorf("Got:%s, want:%v", string(getBody), jsonString(todoItem))
	}
}

func jsonBuffer(i *item) *bytes.Buffer {
	b, _ := json.Marshal(i)
	return bytes.NewBuffer(b)
}

func jsonString(i *item) string {
	b, _ := json.Marshal(i)
	return string(b)
}
