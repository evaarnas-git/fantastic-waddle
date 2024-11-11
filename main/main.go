package main

import "syscall/js"

// Fungsi penjumlahan (sama seperti sebelumnya)
func add(this js.Value, p []js.Value) interface{} {
    return p[0].Int() + p[1].Int()
}

func main() {
    // Register fungsi dari setiap modul
    js.Global().Set("add", js.FuncOf(add))
    registerAuth()
    registerEncryption()
    registerMath()

    // Menjaga aplikasi berjalan
    select {}
}
