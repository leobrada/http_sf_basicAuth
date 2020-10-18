package sf_basic_auth

import (
  "fmt"
  "net/http"
)

type SFBasicAuth struct {
    name string
    // TODO: dst_pdp string, // indicates where the pdp is located
}

func NewSFBasicAuth() SFBasicAuth {
    return SFBasicAuth{name: "basic_auth"}
}

func (sf SFBasicAuth) ApplyFunction(w http.ResponseWriter, req *http.Request) bool {
  var username, password string
  form := `<html>
      <body>
      <form action="/" method="post">
      <label for="fname">Username:</label>
      <input type="text" id="username" name="username"><br><br>
      <label for="lname">Password:</label>
      <input type="password" id="password" name="password"><br><br>
      <input type="submit" value="Submit">
      </form>
      </body>
      </html>
      `

  _, err := req.Cookie("Username")
  if err != nil {
    if req.Method =="POST" {
      if err := req.ParseForm(); err != nil {
        fmt.Println("Parsing Error")
        w.WriteHeader(401)
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        fmt.Fprintf(w, form)
        return false
      }

      nmbr_of_postvalues := len(req.PostForm)
      if nmbr_of_postvalues != 2 {
        fmt.Println("Too many Post Form Values")
        w.WriteHeader(401)
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        fmt.Fprintf(w, form)
        return false
      }

      usernamel, exist := req.PostForm["username"]
      username = usernamel[0]
      if !exist || username != "alex" {
        fmt.Println("username not present or wrong")
        w.WriteHeader(401)
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        fmt.Fprintf(w, form)
        return false
      }

      passwordl, exist := req.PostForm["password"]
      password = passwordl[0]
      if !exist || password != "test" {
        fmt.Println("password not present or wrong")
        w.WriteHeader(401)
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        fmt.Fprintf(w, form)
        return false
      }

      cookie := http.Cookie{
        Name: "Username",
        Value: username,
        MaxAge: 10,
        Path: "/",
      }
      http.SetCookie(w, &cookie)
      http.Redirect(w, req, "https://service1.testbed.informatik.uni-ulm.de", 303)
      return true

    } else {
      fmt.Println("only post methods are accepted in this state")
      w.WriteHeader(401)
      w.Header().Set("Content-Type", "text/html; charset=utf-8")
      fmt.Fprintf(w, form)
      return false
    }
  }
  return true
}
