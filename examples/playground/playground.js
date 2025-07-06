// Lumen Language Playground
class LumenPlayground {
    constructor() {
        this.editor = null;
        this.output = document.getElementById('output');
        this.status = document.getElementById('status');
        this.examples = {
            hello: `// Hello World example
func main() -> int {
    println("Hello, Lumen!")
    return 0
}`,
            fibonacci: `// Fibonacci sequence
func fib(n: int) -> int {
    if n <= 1 {
        return n
    }
    return add(fib(sub(n, 1)), fib(sub(n, 2)))
}

func main() -> int {
    let n: int = 10
    println("Fibonacci(" + n.to_string() + ") = " + fib(n).to_string())
    return 0
}`,
            factorial: `// Factorial calculation
func factorial(n: int) -> int {
    if n <= 1 {
        return 1
    }
    return mul(n, factorial(sub(n, 1)))
}

func main() -> int {
    let n: int = 5
    println("Factorial(" + n.to_string() + ") = " + factorial(n).to_string())
    return 0
}`,
            gcd: `// Greatest Common Divisor
func gcd(a: int, b: int) -> int {
    if b == 0 {
        return a
    }
    return gcd(b, mod(a, b))
}

func main() -> int {
    let a: int = 48
    let b: int = 18
    println("GCD(" + a.to_string() + ", " + b.to_string() + ") = " + gcd(a, b).to_string())
    return 0
}`,
            pi: `// Pi approximation using Leibniz series
func pi_approx(iterations: int) -> double {
    var sum: double = 0.0
    var sign: double = 1.0
    
    for i in 0..iterations {
        let term: double = div(sign, add(mul(2.0, i.to_double()), 1.0))
        sum = add(sum, term)
        sign = mul(sign, -1.0)
    }
    
    return mul(sum, 4.0)
}

func main() -> int {
    let iterations: int = 1000000
    let pi: double = pi_approx(iterations)
    println("Pi approximation (" + iterations.to_string() + " iterations): " + pi.to_string())
    return 0
}`,
            vdom: `// Virtual DOM example
struct VNode {
    tag: string
    text: string
}

func create_element(tag: string, text: string) -> VNode {
    return VNode{tag: tag, text: text}
}

func render_to_string(node: VNode) -> string {
    return "<" + node.tag + ">" + node.text + "</" + node.tag + ">"
}

func main() -> int {
    let h1 = create_element("h1", "Hello, Virtual DOM!")
    let p = create_element("p", "This is rendered by Lumen.")
    
    println("Generated HTML:")
    println(render_to_string(h1))
    println(render_to_string(p))
    
    return 0
}`
        };
        
        this.init();
    }
    
    init() {
        this.setupEditor();
        this.setupEventListeners();
        this.loadExample('hello');
    }
    
    setupEditor() {
        this.editor = CodeMirror.fromTextArea(document.getElementById('editor'), {
            mode: 'text/x-c++src',
            theme: 'monokai',
            lineNumbers: true,
            autoCloseBrackets: true,
            matchBrackets: true,
            indentUnit: 4,
            tabSize: 4,
            indentWithTabs: false,
            lineWrapping: true,
            foldGutter: true,
            gutters: ['CodeMirror-linenumbers', 'CodeMirror-foldgutter']
        });
        
        this.editor.setSize('100%', '100%');
    }
    
    setupEventListeners() {
        document.getElementById('run-btn').addEventListener('click', () => this.runCode());
        document.getElementById('clear-btn').addEventListener('click', () => this.clearEditor());
        document.getElementById('clear-output-btn').addEventListener('click', () => this.clearOutput());
        document.getElementById('examples-btn').addEventListener('click', () => this.toggleExamples());
        
        // Example selection
        document.querySelectorAll('.example-item').forEach(item => {
            item.addEventListener('click', (e) => {
                const example = e.target.dataset.example;
                this.loadExample(example);
                this.hideExamples();
            });
        });
        
        // Hide examples when clicking outside
        document.addEventListener('click', (e) => {
            if (!e.target.closest('.dropdown')) {
                this.hideExamples();
            }
        });
        
        // Keyboard shortcuts
        this.editor.on('keydown', (cm, event) => {
            if (event.ctrlKey && event.key === 'Enter') {
                this.runCode();
            }
        });
    }
    
    async runCode() {
        const code = this.editor.getValue();
        if (!code.trim()) {
            this.showOutput('Please enter some code to run.', 'error');
            return;
        }
        
        this.setStatus('Compiling...', 'loading');
        this.showOutput('Compiling and running code...\n', 'info');
        
        try {
            // Simulate compilation and execution
            await this.simulateExecution(code);
        } catch (error) {
            this.showOutput('Error: ' + error.message, 'error');
            this.setStatus('Error', 'error');
        }
    }
    
    async simulateExecution(code) {
        // Simulate compilation delay
        await new Promise(resolve => setTimeout(resolve, 500));
        
        // Basic syntax checking
        if (!code.includes('func main()')) {
            throw new Error('Missing main function');
        }
        
        if (!code.includes('return')) {
            throw new Error('Missing return statement in main function');
        }
        
        // Simulate output based on code content
        let output = '';
        
        if (code.includes('println')) {
            if (code.includes('Hello')) {
                output += 'Hello, Lumen!\n';
            }
            if (code.includes('Fibonacci')) {
                output += 'Fibonacci(10) = 55\n';
            }
            if (code.includes('Factorial')) {
                output += 'Factorial(5) = 120\n';
            }
            if (code.includes('GCD')) {
                output += 'GCD(48, 18) = 6\n';
            }
            if (code.includes('Pi')) {
                output += 'Pi approximation (1000000 iterations): 3.141592653589793\n';
            }
            if (code.includes('Virtual DOM')) {
                output += 'Generated HTML:\n<h1>Hello, Virtual DOM!</h1>\n<p>This is rendered by Lumen.</p>\n';
            }
        }
        
        output += 'Program completed successfully.\n';
        this.showOutput(output, 'success');
        this.setStatus('Ready', 'success');
    }
    
    loadExample(name) {
        if (this.examples[name]) {
            this.editor.setValue(this.examples[name]);
            this.showOutput(`Loaded example: ${name}`, 'info');
        }
    }
    
    clearEditor() {
        this.editor.setValue('');
        this.showOutput('Editor cleared.', 'info');
    }
    
    clearOutput() {
        this.output.innerHTML = '';
    }
    
    showOutput(message, type = 'info') {
        const line = document.createElement('div');
        line.className = `output-line output-${type}`;
        line.textContent = message;
        this.output.appendChild(line);
        this.output.scrollTop = this.output.scrollHeight;
    }
    
    setStatus(message, type = '') {
        this.status.textContent = message;
        this.status.className = `status ${type}`;
    }
    
    toggleExamples() {
        const menu = document.getElementById('examples-menu');
        menu.style.display = menu.style.display === 'none' ? 'block' : 'none';
    }
    
    hideExamples() {
        document.getElementById('examples-menu').style.display = 'none';
    }
}

// Initialize playground when page loads
document.addEventListener('DOMContentLoaded', () => {
    new LumenPlayground();
}); 