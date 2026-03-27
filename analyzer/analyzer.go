package analyzer

import (
    "fmt"
    "net/http"
    "time"
)

type Result struct {
    URL          string
    StatusCode   int
    ResponseTime time.Duration
    ContentType  string
    ContentLength int64
    Headers      map[string]string
}

func Analyze(url string) (*Result, error) {
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    elapsed := time.Since(start)

    headers := map[string]string{
        "Cache-Control":           resp.Header.Get("Cache-Control"),
        "Content-Encoding":        resp.Header.Get("Content-Encoding"),
        "X-Content-Type-Options":  resp.Header.Get("X-Content-Type-Options"),
        "Strict-Transport-Security": resp.Header.Get("Strict-Transport-Security"),
        "X-Frame-Options":         resp.Header.Get("X-Frame-Options"),
    }

    return &Result{
        URL:           url,
        StatusCode:    resp.StatusCode,
        ResponseTime:  elapsed,
        ContentType:   resp.Header.Get("Content-Type"),
        ContentLength: resp.ContentLength,
        Headers:       headers,
    }, nil
}

func PrintStats(r *Result) {
    fmt.Println("\n===== RAW STATS =====")
    fmt.Printf("URL:           %s\n", r.URL)
    fmt.Printf("Status:        %d\n", r.StatusCode)
    fmt.Printf("Response Time: %v\n", r.ResponseTime)
    fmt.Printf("Content-Type:  %s\n", r.ContentType)
    fmt.Printf("Content-Length: %d bytes\n", r.ContentLength)
    fmt.Println("\n--- Headers ---")
    for k, v := range r.Headers {
        if v == "" {
            v = "(missing)"
        }
        fmt.Printf("%-35s %s\n", k+":", v)
    }
}