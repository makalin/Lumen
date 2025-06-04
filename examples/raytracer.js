// WebAssembly loader for Lumen ray tracer
const importObject = {
  env: {
    add: (a, b) => a + b,
    mul: (a, b) => a * b,
    div: (a, b) => a / b,
    sqrt: (x) => Math.sqrt(x),
    println: (ptr) => {
      const str = new TextDecoder().decode(new Uint8Array(memory.buffer, ptr));
      const output = document.getElementById('output');
      output.textContent += str + '\n';
      console.log(str);
    }
  }
};

let memory;
let instance;

async function init() {
  try {
    const response = await fetch('raytracer.wasm');
    const bytes = await response.arrayBuffer();
    const result = await WebAssembly.instantiate(bytes, importObject);
    
    instance = result.instance;
    memory = new WebAssembly.Memory({ initial: 256 });
    
    // Run the ray tracer
    instance.exports.main();
    
    // Draw a simple visualization
    const canvas = document.getElementById('canvas');
    const ctx = canvas.getContext('2d');
    
    // Draw a circle representing the sphere
    ctx.beginPath();
    ctx.arc(200, 200, 100, 0, 2 * Math.PI);
    ctx.strokeStyle = '#333';
    ctx.lineWidth = 2;
    ctx.stroke();
    
    // Draw a line representing the ray
    ctx.beginPath();
    ctx.moveTo(50, 200);
    ctx.lineTo(350, 200);
    ctx.strokeStyle = '#0066cc';
    ctx.lineWidth = 1;
    ctx.stroke();
    
    // Draw intersection point
    ctx.beginPath();
    ctx.arc(200, 200, 5, 0, 2 * Math.PI);
    ctx.fillStyle = '#cc0000';
    ctx.fill();
    
  } catch (err) {
    console.error('Failed to load WebAssembly module:', err);
    const output = document.getElementById('output');
    output.textContent = 'Error: ' + err.message;
  }
}

// Initialize when the page loads
window.addEventListener('load', init); 