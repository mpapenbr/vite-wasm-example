import wasmModule from "./sample.wasm?init";

interface WasmMethods {
  demoAdd: (a: number, b: number) => number;

  getMessages: () => string[];
}

// const go = globalThis.Go; // eslint-disable-line
let instance: WebAssembly.Instance;
let wasmMethods: WasmMethods;
console.log("I will init wasm stuff");
const go = (window as any).Go;
console.log(go);
const impObj = go.importObject;
const initWasm = async () => {
  // const impObj = {};
  instance = await wasmModule(impObj);
  go.run(instance);
};

await initWasm();

wasmMethods = {
  demoAdd: (a: number, b: number) => {
    return (globalThis as any).myAdd(a, b);
  },
  getMessages: () => {
    return (globalThis as any).myLogs();
  },
};

export default wasmMethods;
