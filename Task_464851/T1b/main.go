
package main  
import (  
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
    "time"
)

const (
    uploadPath = "./uploads/"
    downloadPath = "./downloads/"
)

func fileHandler(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Check if the request is a POST for upload
        if r.Method == http.MethodPost {
            r.ParseMultipartForm(32 << 20) // 32 MB memory limit
            file, handler, err := r.FormFile("file")
            if err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
            }
            defer file.Close()

            // Create the upload directory if it doesn't exist
            if err := os.MkdirAll(uploadPath, 0755); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            // Generate a unique file name
            filename := time.Now().Format("20060102150405") + "_" + handler.Filename
            dst, err := os.Create(filepath.Join(uploadPath, filename))
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            defer dst.Close()

            // Copy the file to the upload directory
            if _, err := io.Copy(dst, file); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            // Display success message
            fmt.Fprintf(w, "File uploaded successfully: %s\n", filename)
            return
        }

        // Check if the request is a GET for download
        if r.Method == http.MethodGet {
            filename := r.URL.Path[len("/download/"):]
            filePath := filepath.Join(downloadPath, filename)

            // Check if the file exists
            if _, err := os.Stat(filePath); os.IsNotExist(err) {
                http.Error(w, "File not found", http.StatusNotFound)
                return
            }

            // Serve the file
            http.ServeFile(w, r, filePath)
            return
        }

        // Call the next handler if the request is not for upload or download
        next.ServeHTTP(w, r)
    })
}

func main() {
    http.Handle("/", fileHandler(http.FileServer(http.Dir("."))))
    http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, `
            <html>
                <body>
                    <form action="/upload" method="post" enctype="multipart/form-data">
                        <input type="file" name="file" required />
                        <input type="submit" value="Upload" />
                    </form>
                </body>
            </html>
        `)
    })
    http.HandleFunc("/download/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, `
            <html>
                <body>
                    <a href="/download/filename.txt">Download File</a>
                </body>
            </html>
        `)
    })
    fmt.Println("Server is running on port 8080...")
    http.ListenAndServe(":8080", nil)
}

