from ctypes import *
import os

root_dir = os.path.dirname(__file__)
shared_lib = os.path.join(root_dir, 'bind', 'template.so')
lib = cdll.LoadLibrary(shared_lib)

def render_template_string(template_content, values_content):
    # Set argument and return types
    lib.RenderTemplateString.argtypes = [c_char_p, c_char_p]
    lib.RenderTemplateString.restype = c_char_p
    
    # Convert Python strings to bytes
    template_bytes = template_content.encode('utf-8')
    values_bytes = values_content.encode('utf-8')
    
    # Call the Go function and get the result
    result = lib.RenderTemplateString(template_bytes, values_bytes)
    
    # Convert the result back to a Python string
    return result.decode('utf-8')

Template = """{{.Count}} items are made of {{.Material}}"""
Values = """Count: 12\nMaterial: Wool"""

print(render_template_string(Template,Values))