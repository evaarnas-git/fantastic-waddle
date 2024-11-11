use ic_cdk::export::candid::CandidType;
use ic_cdk_macros::query;
use wasmtime::*;

#[derive(CandidType)]
struct GoResult {
    result: i32,
}

#[query]
fn execute_go() -> GoResult {
    let engine = Engine::default();
    let store = Store::new(&engine);

    // Memuat file wasm Go
    let module = Module::from_file(&engine, "src/backend/main.wasm").unwrap();
    let instance = Instance::new(&store, &module, &[]).unwrap();

    // Panggil fungsi `add` dari Go
    let add = instance.get_typed_func::<(i32, i32), i32>("add").unwrap();
    let result = add.call((2, 3)).unwrap();

    GoResult { result }
}
