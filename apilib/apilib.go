package apilib
//go build api-call
import (
    "bytes"
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

type MyResponse struct {
    Access_token  string
    Expires_in int 
    User User 
}
type User struct {
    Uuid string
    Type string
    Name string
    Created int64
    Modified int64
    Username string
    Email string
    Activated bool
    Picture string
}

type Request struct {
    Grant_type string `json:"grant_type"`
    Username string `json:"username"`
    Password string `json:"password"`
}

func UsergridAuth (username string, password string) bool{
    ret := false
    //url := "http://192.168.1.84:8080/management/token"
    url := "https://api.usergrid.com/jojo/sandbox/token"
    fmt.Println("URL:>", url)

    myRequest := Request{Grant_type:"password", Username: username, Password: password}
    jsonStr, _ := json.Marshal(myRequest)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    //fmt.Println("response Body:", string(body))
    var u MyResponse

    err = json.Unmarshal(body, &u)
    if err != nil{
        panic(err)
    }
    fmt.Println("access token:", u.Access_token)
    fmt.Println(u.Expires_in)
    fmt.Println(fmt.Sprintf("%s", u.User.Name))
    fmt.Println("name:", u.User.Name)
    if u.Expires_in > 0{
        ret = true
    }
    return ret
}

