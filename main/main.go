package main

import (
    "syscall/js"
)

// Fungsi untuk menambahkan dua angka
func add(this js.Value, p []js.Value) interface{} {
    // Mendapatkan dua angka pertama dari parameter dan menambahkannya
    result := p[0].Int() + p[1].Int()
    return result
}

// Fungsi utama untuk menginisialisasi fungsi-fungsi Go dalam JavaScript
func main() {
    // Mengekspos fungsi `add` ke lingkungan JS sebagai "add"
    js.Global().Set("add", js.FuncOf(add))
    // Menjaga program berjalan
    select {}
}
