# Vite + React + WASM

This project is a sample on how to use WASM with Vite and React.

The Go sources are included here and serve only for demo purposes

## Steps

**Preparations:** If you want to use TinyGo, install it via `sudo installTinyGo.sh`

### Create wasm

The following step creates the WASM file at `wasmtest/src/wasmdemo/sample.wasm` and the helper `wasmtest/public/wasm_exec.js`

The standard Go variant (produces bigger wasm files)

```console
cd go
make build
```

The TinyGo variant

```console
cd go
make tiny
```

### Start vite

```console
cd wasmtest
yarn dev
```
