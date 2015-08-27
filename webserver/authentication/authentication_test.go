package authentication

import(
    "net/http"
    "net/http/httptest"
    "testing"
    "net/url"
    "bytes"
)

func TestRegister(t *testing.T) {
  t.Log("Run TestRegister...")

  v := url.Values{}
  v.Add("email", "test@gmail.com")
  v.Add("password", "root")  
  
  req, _ := http.NewRequest("POST", "", bytes.NewBufferString(v.Encode()))
  w := httptest.NewRecorder()
  Register(w, req)

  t.Log(w)

  if w.Code != http.StatusOK {
    t.Errorf("Home page didn't return %v", http.StatusOK)
  }

  t.Log("Run TestRegister...successful")
}