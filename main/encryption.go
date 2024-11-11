package main

import (
    "crypto/sha256"
    "encoding/hex"
    "syscall/js"
)

// Fungsi untuk membuat hash password
func hashPassword(this js.Value, p []js.Value) interface{} {
    password := p[0].String()
    hash := sha256.Sum256([]byte(password))
    return hex.EncodeToString(hash[:])
}

// Register fungsi ke WebAssembly
func registerEncryption() {
    js.Global().Set("hashPassword", js.FuncOf(hashPassword))
}
