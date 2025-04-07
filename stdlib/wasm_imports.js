const env = {
  print: (ptr) => {
    const msg = readStringFromMemory(ptr);
    console.log(msg);
  },
  println: (ptr) => {
    const msg = readStringFromMemory(ptr);
    console.log(msg + "\n");
  },
};

function readStringFromMemory(ptr) {
  const memory = wasmInstance.exports.memory;
  const bytes = new Uint8Array(memory.buffer, ptr);
  let str = "";
  for (let i = 0; bytes[i] !== 0; i++) {
    str += String.fromCharCode(bytes[i]);
  }
  return str;
}

WebAssembly.instantiateStreaming(fetch("output.wasm"), { env });
