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

	t.Run("GET /hello", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/hello?name=Holberton", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		expectedCode := http.StatusOK
		if status := rr.Code; status != expectedCode {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, expectedCode)
		}

		expectedBody := "Hello Holberton!"
		if rr.Body.String() != expectedBody {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expectedBody)
		}
	})

	t.Run("GET /health", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/health", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		expectedCode := http.StatusOK
		if status := rr.Code; status != expectedCode {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, expectedCode)
		}

		expectedBody := "ALIVE"
		if rr.Body.String() != expectedBody {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expectedBody)
		}
	})

	t.Run("GET /", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		expectedCode := http.StatusNotFound
		if status := rr.Code; status != expectedCode {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, expectedCode)
		}

		expectedBody := "404 page not found\n"
		if rr.Body.String() != expectedBody {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expectedBody)
		}
	})
}
