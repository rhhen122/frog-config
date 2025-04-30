<div align="center">
<img src="../media/frog-icon.png" height="100">

# Frog!
## A beautiful config file format.
### Basic Syntax
<a href="http://vimp.rhhen.xyz/Licenses/lookinggood/lice/LICENSE.html"><img src="https://badgen.net/static/license/VIMPPDL%201.0.2/black"></a>
|
<a href="http://go.dev/"><img src="http://badgen.net/static/Go/1.24?icon=https%3A%2F%2Fgo.dev%2Fblog%2Fgo-brand%2FGo-Logo%2FSVG%2FGo-Logo_White.svg"></a>
</div>

### Comments
First we need to know how to make comments:
```
; This is a comment
; We can comment lines we do not want active like so:
"cond" = true
; "other_cond" = true
```
### Types
```
; There are 3 types of values
; String, Int & Bool
"bool" = true           ; A Bool can equal either 'True' or 'False', Bools have special feature which we will get to in a second.
"string" = "Some Text"  ; A String Stores a series of letters
"int" = 5               ; A Int Store a number, it can also store an equation like: 2*2, which when put through the compiler comes out as: 4
```
### Bools (Verb Engine)
```
; Bools have a feature called 'The Verb Engine'
; Which means this is a valid Bool:
"bool" = ha ha fuck yes                 ; This will return 'True' when span through the compiler
"other_bool" = ha ha fuck nah bruh      ; This will return 'False' when span through the compiler
```
Here is a list of all `Verb Engine` Keywords:

Does Nothing:
`bruh`
`fuck`
`shit`
`ha`
`haha`
`he`
`hehe`

Returns True:
`true`
`y`
`yes`
`yup`
`yeah`

Returns False:
`true`
`n`
`no`
`nope`
`nah`

### Formating
```
; Formating Works like this:
"Variable_name_here" = true             ; A Bool
"Variable_name_here" = "A String"       ; A String
"Variable_name_here" = 5*5+2            ; A Int
```

<a href="ad_syntax.md">Next -></a>