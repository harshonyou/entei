# Entei

A dynamically typed, interpreted programming language written in Go. Supports first-class functions, closures, and a rich set of built-in data types — all accessible through an interactive REPL.

> Built by following [Writing An Interpreter In Go](https://interpreterbook.com/) by Thorsten Ball.

______________________________________________________________________

## Features

- **Data types** — integers, booleans, strings, arrays, hashes, null
- **Variable bindings** — `let`
- **Functions** — first-class, anonymous, recursive
- **Closures** — functions capture their enclosing environment
- **Higher-order functions** — pass and return functions freely
- **Conditionals** — `if` / `else` expressions
- **Return statements** — early exit from any function
- **Operators** — arithmetic, comparison, logical prefix, string concatenation
- **Built-ins** — `len`, `first`, `last`, `rest`, `push`, `puts`
- **Interactive REPL**

______________________________________________________________________

## Getting Started

**Requirements:** Go 1.18+

```bash
git clone https://github.com/harshonyou/entei
cd entei
go run main.go
```

```
Hello <you>!
Welcome to Entei, the programming language for everyone!
>>
```

______________________________________________________________________

## Language Tour

### Variable Bindings

```entei
let x = 10;
let name = "Entei";
let active = true;
```

### Arithmetic & Comparison

```entei
let result = (10 + 5) * 2 / 3;    // 10
let bigger = 42 > 7;              // true
let equal  = 5 == 5;              // true
let diff   = 5 != 3;              // true
```

### Strings

```entei
let greet = fn(name) { "Hello, " + name + "!" };
greet("World");                   // Hello, World!
```

### Conditionals

```entei
let abs = fn(x) {
    if (x < 0) { -x } else { x }
};

abs(-7);   // 7
abs(3);    // 3
```

### Functions

```entei
let add = fn(a, b) { a + b };
add(3, 4);   // 7

// Immediately invoked
fn(x) { x * x }(5);   // 25
```

### Recursion

```entei
let fib = fn(n) {
    if (n < 2) { return n; }
    fib(n - 1) + fib(n - 2)
};

fib(10);   // 55
```

### Closures

```entei
let makeAdder = fn(x) {
    fn(y) { x + y }
};

let addTen = makeAdder(10);
addTen(5);    // 15
addTen(42);   // 52
```

### Arrays

```entei
let nums = [1, 2, 3, 4, 5];

nums[0];          // 1
len(nums);        // 5
first(nums);      // 1
last(nums);       // 5
rest(nums);       // [2, 3, 4, 5]
push(nums, 6);    // [1, 2, 3, 4, 5, 6]
```

### Hashes

```entei
let person = {"name": "Alice", "age": 30, "active": true};

person["name"];     // Alice
person["age"];      // 30
person["active"];   // true
```

______________________________________________________________________

## Showcase Program

A single program exercising every feature — closures, higher-order functions, recursion, arrays, hashes, strings, conditionals, and built-ins.

```entei
// --- Higher-order array utilities ---

let map = fn(arr, f) {
    let iter = fn(arr, acc) {
        if (len(arr) == 0) {
            return acc;
        }
        iter(rest(arr), push(acc, f(first(arr))))
    };
    iter(arr, [])
};

let filter = fn(arr, f) {
    let iter = fn(arr, acc) {
        if (len(arr) == 0) {
            return acc;
        }
        let head = first(arr);
        if (f(head)) {
            iter(rest(arr), push(acc, head))
        } else {
            iter(rest(arr), acc)
        }
    };
    iter(arr, [])
};

let reduce = fn(arr, init, f) {
    let iter = fn(arr, result) {
        if (len(arr) == 0) {
            return result;
        }
        iter(rest(arr), f(result, first(arr)))
    };
    iter(arr, init)
};

let sum = fn(arr) {
    reduce(arr, 0, fn(acc, x) { acc + x })
};

// --- Closure: greeting factory ---

let makeGreeter = fn(greeting) {
    fn(name) { greeting + ", " + name + "!" }
};

let hello = makeGreeter("Hello");
let welcome = makeGreeter("Welcome back");

// --- Recursive fibonacci ---

let fib = fn(n) {
    if (n < 2) { return n; }
    fib(n - 1) + fib(n - 2)
};

// --- Data ---

let student = {"name": "Alice", "score": 88, "passed": true};
let scores  = [72, 55, 91, 63, 84, 97, 48];

// --- Computation ---

let passing = filter(scores, fn(s) { s > 60 });
let boosted  = map(passing, fn(s) { s + 5 });
let total    = sum(boosted);
let average  = total / len(boosted);

// --- Output ---

puts(hello(student["name"]));
puts(welcome(student["name"]));
puts(fib(10));
puts(passing);
puts(boosted);
puts(average);
```

**Output:**

```
Hello, Alice!
Welcome back, Alice!
55
[72, 91, 63, 84, 97]
[77, 96, 68, 89, 102]
86
```

______________________________________________________________________

## Built-in Functions

| Function | Description |
|-----------------|----------------------------------------------|
| `len(x)` | Length of string or array |
| `first(arr)` | First element of array |
| `last(arr)` | Last element of array |
| `rest(arr)` | New array with all elements except the first |
| `push(arr, x)` | New array with `x` appended |
| `puts(...)` | Print each argument to stdout |

______________________________________________________________________

## Operators

| Category | Operators |
|-------------|-----------------------------|
| Arithmetic | `+` `-` `*` `/` |
| Comparison | `<` `>` `==` `!=` |
| Prefix | `!` (not) `-` (negate) |
| String | `+` (concatenation) |

______________________________________________________________________

## Architecture

| Package | Role |
|-------------|--------------------------------------------------------|
| `token` | Token types and keyword lookup |
| `lexer` | Source text → token stream |
| `ast` | Abstract Syntax Tree node definitions |
| `parser` | Token stream → AST (Pratt parser) |
| `object` | Runtime value types and environment |
| `evaluator` | Tree-walking interpreter + built-in functions |
| `repl` | Read-Eval-Print loop |
