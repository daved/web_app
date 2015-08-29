package session

import (
  "testing"
  "net/http"
  "net/http/httptest"
  "log"
)

func getRecordedCookie(recorder *httptest.ResponseRecorder, name string) (*http.Cookie, error) {
  r := &http.Response{Header: recorder.HeaderMap}
  for _, cookie := range r.Cookies() {
    if cookie.Name == name {
      return cookie, nil
    }
  }
  return nil, http.ErrNoCookie
}

func TestSetSession(t *testing.T) {
  w := httptest.NewRecorder()
  SetSession("Douglas.Costa@gmail.com", w)

  c, err := getRecordedCookie(w, "session")
  if err != nil {
    t.Errorf("getRecordedCookie failed")
  }

  req, _ := http.NewRequest("GET", "", nil)
  req.AddCookie(c)
  user, err := GetSessionUser(req)
  if err != nil {
    t.Errorf("GetSessionUser failed")
  }

  if user != "Douglas.Costa@gmail.com" {
    t.Errorf("GetSessionUser failed")
  }
}

func TestClearSession(t *testing.T) {
  ClearSession(w)

  req, _ := http.NewRequest("GET", "", nil)
  req.AddCookie(c)

  _, err := GetSessionUser(req)
  if err == nil {
    t.Errorf("GetSessionUser failed")
  }
}