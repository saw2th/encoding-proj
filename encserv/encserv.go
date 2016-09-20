package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/saw2th/encoding-proj/enc-client"
    "encoding/hex"
)

// encryption storage
var enc_client_store client.ClientStore

// input for encoding
type Enc_struct struct {
    Content string
    Id string
}

// output from encoding storage
type OutStore struct {
    Key string
}

// input required to retrieve encrypted text
type Ret_struct struct {
     Key string
     Id string
}

// for output from retrieve
type Decoded struct {
    Content string
}


// starts web server
func main() {

    // key value storage
    enc_client_store.Storage = make(map[string][]byte)

    http.HandleFunc("/store", store_handler)
    http.HandleFunc("/retrieve", retrieve_handler)

    http.ListenAndServe(":8080", nil)
}


// web handler for storing text as encrypted text
// requires json of form
// {"Id" : "<id of data>", "Content": "content to encrypt"}
// returns json of form
// {"Key":"<key to retrieve content>"}
func store_handler(w http.ResponseWriter, r *http.Request) {
     var enc Enc_struct

     if r.Body == nil {
            http.Error(w, "Please send a request body", 400)
            return
        }

     err := json.NewDecoder(r.Body).Decode(&enc)
     if err != nil {
            http.Error(w, err.Error(), 400)
            return
     }

     aesKey, err := enc_client_store.Store([]byte(enc.Id), []byte(enc.Content))

     if err != nil {
            http.Error(w, err.Error(), 400)
            return
     }

     aesKeyStr := fmt.Sprintf("%x", aesKey)
     os := OutStore{Key: aesKeyStr}

     // encode back to json
     json.NewEncoder(w).Encode(os) 

}

// web handler for retrieving encrypted text
// requires json of form
// {"Id" : "<id of data>", "Key": "<key returned from /store>"}
// returns json of form
// {"Content": "<content previously stored>"}
func retrieve_handler(w http.ResponseWriter, r *http.Request) {

     var ret Ret_struct

     if r.Body == nil {
            http.Error(w, "Please send a request body", 400)
            return
        }

     err := json.NewDecoder(r.Body).Decode(&ret)
     ret_hex_key, err := hex.DecodeString(string(ret.Key))

     if err != nil {
            http.Error(w, err.Error(), 400)
            return
     }

     rtn, err := enc_client_store.Retrieve([]byte(ret.Id), ret_hex_key)

     if err != nil {
            http.Error(w, err.Error(), 400)
            return
     }

     decoded := Decoded{Content: string(rtn)}

     // encode back to json
     json.NewEncoder(w).Encode(decoded) 

}

