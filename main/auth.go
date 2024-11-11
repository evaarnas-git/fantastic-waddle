package main

import (
    "errors"
    "syscall/js"
)

var validToken = "12345" // Token demo untuk verifikasi

// Fungsi verifikasi token autentikasi
func verifyToken(this js.Value, p []js.Value) interface{} {
    token := p[0].String()
    if token == validToken {
        return true
    }
    return false
}

// Register fungsi ke WebAssembly
func registerAuth() {
    js.Global().Set("verifyToken", js.FuncOf(verifyToken))
}
