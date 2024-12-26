package main


import (
    "fmt"
    "net/url"
)

// InvalidParam represents a single invalid query parameter.
type InvalidParam struct {
    ParamName  string
    ParamValue string
    Error      error
}

func validateQueryParams(urlStr string) []InvalidParam {
    var invalidParams []InvalidParam

    // Parse the URL
    u, err := url.Parse(urlStr)
    if err != nil {
        invalidParams = append(invalidParams, InvalidParam{
            ParamName:  "URL",
            ParamValue: urlStr,
            Error:      err,
        })
        return invalidParams
    }

    // Validate each query parameter
    for key, values := range u.Query() {
        for _, value := range values {
            // Add your custom validation logic here
            // For example, checking for empty values or invalid characters.
            if value == "" {
                invalidParams = append(invalidParams, InvalidParam{
                    ParamName:  key,
                    ParamValue: value,
                    Error:      fmt.Errorf("value is empty"),
                })
            }
        }
    }
    return invalidParams
}

func main() {
    // Example URLs with invalid query parameters
    urlStrings := []string{
        "https://example.com?param1=value1&param2=", // Empty value for param2
        "https://example.com?param3=invalid@char",  // Invalid character in value
        "https://example.com?param4",                // Missing value for param4
        "invalid_url",                              // Invalid URL
    }

    for _, urlStr := range urlStrings {
        invalidParams := validateQueryParams(urlStr)
        if len(invalidParams) > 0 {
            fmt.Println("Invalid parameters found in URL:", urlStr)
            for _, param := range invalidParams {
                fmt.Printf("  - %s: %s - %v\n", param.ParamName, param.ParamValue, param.Error)
            }
        } else {
            fmt.Println("All parameters are valid in URL:", urlStr)
        }
    }
}  