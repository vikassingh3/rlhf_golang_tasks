package main  
import (  
    "bytes"  
    "fmt"  
    "html/template"  
    "io"  
    "os"  
)  

type Data struct {  
    Title string  
    Name  string  
}  

// RenderTemplate renders the template with the given data and writes the output to the specified writer.  
// It returns the rendered content and any error that occurred during rendering.  
func RenderTemplate(templatePath string, data interface{}, w io.Writer) ([]byte, error) {  
    tmpl, err := template.ParseFiles(templatePath)  
    if err != nil {  
        return nil, err  
    }  

    var buf bytes.Buffer  
    err = tmpl.Execute(&buf, data)  
    if err != nil {  
        return nil, err  
    }  

    // Copy the rendered content to the specified writer  
    _, err = buf.WriteTo(w)  
    if err != nil {  
        return nil, err  
    }  

    return buf.Bytes(), nil  
}  

func main() {  
    templatePath := "index.html"  
    data := Data{  
        Title: "Welcome Page",  
        Name:  "Alice",  
    }  

    // Render the template and handle errors  
    renderedOutput, err := RenderTemplate(templatePath, data, os.Stdout)  
    if err != nil {  
        fmt.Println("Error rendering template:", err)  
        os.Exit(1)  
    }  

    // Do something with the rendered output if needed  
    fmt.Println("Rendered Output:")  
    fmt.Println(string(renderedOutput))  
}  
