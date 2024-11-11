package main

import "syscall/js"

// Fungsi untuk mengalikan dua angka
func multiply(this js.Value, p []js.Value) interface{} {
    result := p[0].Int() * p[1].Int()
    return result
}

// Register fungsi ke WebAssembly
func registerMath() {
    js.Global().Set("multiply", js.FuncOf(multiply))
}
