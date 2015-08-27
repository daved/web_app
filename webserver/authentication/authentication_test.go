package authentication

import(  
    "fmt"
    "net/http"
    "net/http/httptest"
    "testing"
    "net/url"
    "bytes"
)

func TestRegister(t *testing.T) {
  v := url.Values{}
  v.Add("email", "test@gmail.com")
  v.Add("password", "root")  
  req, _ := http.NewRequest("POST", "", bytes.NewBufferString(v.Encode()))
  w := httptest.NewRecorder()
  Register(w, req)
  fmt.Println(w)
}