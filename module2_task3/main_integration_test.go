package main

import (
  //nolint:staticcheck
  "io/ioutil"
  "net/http"
  "net/http/httptest"
  "testing"
)

func Test_server(t *testing.T) {
  if testing.Short() {
    t.Skip("Flag `-short` provided: skipping Integration Tests.")
  }
  

  tests := []struct {
    name         string
    URI          string
    responseCode int
    body         string
  }{
    {
      name:         "Home page",
      URI:          "",
      responseCode: 404,
      body:         "404 page not found\n",
    },
    {
      name:         "Hello page",
      URI:          "/hello?name=Holberton",
      responseCode: 200,
      body:         "Hello Holberton!",
    },
    {
      name:         "Health page",
      URI:          "/health",
      responseCode: 200,
      body:         "ALIVE",
    },
    {
      name:         "Hello page",
      URI:          "/hello?name=Holberton",
      responseCode: 200,
      body:         "Hello Holberton!",
    },
    {
      name:         "Home page",
      URI:          "",
      responseCode: 404,
      body:         "404 page not found\n",
    },
    {
      name:         "Hello page with empty name value",
      URI:          "/hello?name=",
      responseCode: 400,
      body:         "",
    },
    {
      name:         "Empty name value",
      URI:          "/hello?name=",
      responseCode: 400,
      body:         "",
    },
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      ts := httptest.NewServer(setupRouter())
      defer ts.Close()

      res, err := http.Get(ts.URL + tt.URI)
      if err != nil {
        t.Fatal(err)
      }

      // Check that the status code is what you expect.
      expectedCode := tt.responseCode
      gotCode := res.StatusCode
      if gotCode != expectedCode {
        t.Errorf("handler returned wrong status code: got %q want %q", gotCode, expectedCode)
      }

      // Check that the response body is what you expect.
      expectedBody := tt.body
      bodyBytes, err := ioutil.ReadAll(res.Body)
      res.Body.Close()
      if err != nil {
        t.Fatal(err)
      }
      gotBody := string(bodyBytes)
      if gotBody != expectedBody {
        t.Errorf("handler returned unexpected body: got %q want %q", gotBody, expectedBody)
      }
    })
  }
}
func Test_setupRouter(t *testing.T) {
  r := setupRouter()

  // Test that the /health endpoint returns a 200 status code
  req, _ := http.NewRequest("GET", "/health", nil)
  rr := httptest.NewRecorder()
  r.ServeHTTP(rr, req)

  if status := rr.Code; status != http.StatusOK {
    t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
  }

  // Test that the /hello endpoint returns a 400 status code with an empty name parameter
  req, _ = http.NewRequest("GET", "/hello?name=", nil)
  rr = httptest.NewRecorder()
  r.ServeHTTP(rr, req)

  if status := rr.Code; status != http.StatusBadRequest {
    t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
  }

  // Test that the /hello endpoint returns a 200 status code with a non-empty name parameter
  req, _ = http.NewRequest("GET", "/hello?name=Holberton", nil)
  rr = httptest.NewRecorder()
  r.ServeHTTP(rr, req)

  if status := rr.Code; status != http.StatusOK {
    t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
  }
}
