package main

import "C"
import (
        "bytes"
        "gopkg.in/yaml.v2"
        "text/template"
)

//export RenderTemplateString
func RenderTemplateString(templateStr, valueStr *C.char) *C.char {
        // Convert C strings to Go strings
        templateContent := C.GoString(templateStr)
        valueContent := C.GoString(valueStr)
        
        // Parse YAML values
        values := make(map[string]interface{})
        err := yaml.Unmarshal([]byte(valueContent), &values)
        if err != nil {
                return C.CString("")
        }
        
        // Parse and execute the template
        tmpl, err := template.New("template").Parse(templateContent)
        if err != nil {
                return C.CString("")
        }
        
        var buf bytes.Buffer
        err = tmpl.Execute(&buf, values)
        if err != nil {
                return C.CString("")
        }
        
        // Return the result as a C string
        return C.CString(buf.String())
}

func main() {}