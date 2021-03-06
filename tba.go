package main

import (
    "bytes"
    "crypto/md5"
    "fmt"
    "net/http"
)

type eventParams struct {
    event string
    auth string
    secret string
}

func sendTBARequest(url string, body []byte, params *eventParams) (*http.Response, error) {
    url = fmt.Sprintf("/api/trusted/v1/event/%s/%s", params.event, url)
    sig := fmt.Sprintf("%x", md5.Sum(append([]byte(params.secret + url), body...)))

    url = "https://www.thebluealliance.com" + url
    request, err := http.NewRequest("POST", url, bytes.NewReader(body))
    if err != nil {
        return nil, err
    }
    request.Header.Add("X-TBA-Auth-Id", params.auth);
    request.Header.Add("X-TBA-Auth-Sig", sig);
    client := http.Client{}
    return client.Do(request)
}
