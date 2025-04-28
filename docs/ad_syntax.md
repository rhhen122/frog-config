<div align="center">
<img src="../media/frog-icon.png" height="100">

# Frog!
## A beautiful config file format.
### Advanced Syntax
<a href="http://vimp.rhhen.xyz/Licenses/lookinggood/lice/LICENSE.html"><img src="https://badgen.net/static/license/VIMPPDL%201.0.2/black"></a>
|
<a href="http://go.dev/"><img src="http://badgen.net/static/Go/1.24?icon=https%3A%2F%2Fgo.dev%2Fblog%2Fgo-brand%2FGo-Logo%2FSVG%2FGo-Logo_White.svg"></a>
</div>

### Lists
```
; Lists work the same as in python
; Like this:
"list" = ["item_1", "item_2", "item_3"]
```
### String Handling
```
; There are many types of strings
; Here are some

; Multiline Strings
"mul_string" = `A Line here!
And another here`

; Intergrated Strings
"string_2_intergrate" = "our string"
"string" = "Look @ this string! -> \* string_2_intergrate \*"
; Or something like this:
"big_string" = "\* string *\ \* string_2_intergrate \*"
```
### Line Spliting
```
; We can split lines like so:
"string" = "our string" \& "int" = 5
```