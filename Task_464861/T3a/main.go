package main

import (
    "bytes"
    "fmt"
    "html/template"
    "io"
    "os"
    "time"
)

// Data struct represents the data passed to the template
type Data struct {
    Title string
    Name  string
}

// TemporaryError represents a temporary error that might be resolved by retrying.
type TemporaryError struct {
    message string
    retry   bool
}

func (e TemporaryError) Error() string {
    return fmt.Sprintf("temporary error: %s (retry: %t)", e.message, e.retry)
}

// PermalinkError represents a permanent error that cannot be resolved by retrying.
type PermalinkError struct {
    message string
}

func (e PermalinkError) Error() string {
    return fmt.Sprintf("permalink error: %s", e.message)
}

func RenderTemplate(templatePath string, data interface{}, w io.Writer) ([]byte, error) {
    tmpl, err := template.ParseFiles(templatePath)
    if err != nil {
        // Depending on the nature of the error, this could be a retry-able error or not.
        if isTemporaryError(err) {
            return nil, TemporaryError{message: err.Error(), retry: true}
        }
        return nil, PermalinkError{message: err.Error()}
    }

    var buf bytes.Buffer
    err = tmpl.Execute(&buf, data)
    if err != nil {
        return nil, PermalinkError{message: err.Error()}
    }

    // Write the rendered template to the provided writer (if any)
    _, err = buf.WriteTo(w)
    if err != nil {
        return nil, PermalinkError{message: err.Error()}
    }

    return buf.Bytes(), nil
}

// Helper function to determine if the error is temporary.
func isTemporaryError(err error) bool {
    // Example condition; adapt to specific errors you want to retry.
    return os.IsNotExist(err)
}

func main() {
    templatePath := "index.html" // Ensure this file exists
    data := Data{
        Title: "Welcome Page",
        Name:  "Bob",
    }

    // Implement error handling logic for retrying and logging.
    renderedOutput, err := renderWithRetry(templatePath, data, os.Stdout, 3, time.Second*2)
    if err != nil {
        handleError(err)
    }

    fmt.Println("Rendered Output:")
    fmt.Println(string(renderedOutput))
}

func renderWithRetry(templatePath string, data interface{}, w io.Writer, maxAttempts int, retryDelay time.Duration) ([]byte, error) {
    var lastErr error
    for attempt := 1; attempt <= maxAttempts; attempt++ {
        renderedOutput, err := RenderTemplate(templatePath, data, w)
        if err == nil {
            return renderedOutput, nil // Success
        }

        if te, ok := err.(TemporaryError); ok && te.retry {
            // Assume a TemporaryError suggests retrying
            fmt.Printf("Attempt %d failed: %v. Retrying in %v...\n", attempt, err, retryDelay)
            lastErr = err
            time.Sleep(retryDelay)
        } else {
            // Non-retryable error
            return nil, err
        }
    }

    // Failed after maximum attempts
    return nil, fmt.Errorf("maximum retry attempts reached: %v", lastErr)
}

func handleError(err error) {
    if pe, ok := err.(PermalinkError); ok {
        // Handle PermalinkError, for example, by logging
        fmt.Printf("Permanent error: %v. Action: %v\n", pe.message, "Refer to logs for details")
    } else {
        // Handle unexpected errors
        fmt.Printf("Unexpected error: %v\n", err)
    }
}
