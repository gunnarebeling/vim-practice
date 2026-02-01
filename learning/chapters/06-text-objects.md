# Chapter 6: Text Objects & Advanced Motions

## Learning Goals
Master text objects for surgical editing - the key to Vim efficiency.

---

## Lesson 6.1: Inner vs Around

Text objects work with operators (d, c, y, v).

- `i` = inner (inside, excluding delimiters)
- `a` = around (including delimiters)

| Object | Inner `i` | Around `a` |
|--------|-----------|------------|
| word | `diw` | `daw` |
| WORD | `diW` | `daW` |
| sentence | `dis` | `das` |
| paragraph | `dip` | `dap` |

### Practice Task 6.1
1. Open `main.go`, go to line 90
2. Put cursor on any word in the string
3. Press `diw` - deletes word, keeps spaces
4. Undo, try `daw` - deletes word AND space
5. Try on different words to feel the difference

---

## Lesson 6.2: Quote and Bracket Objects

| Object | Inner | Around |
|--------|-------|--------|
| double quotes | `di"` | `da"` |
| single quotes | `di'` | `da'` |
| backticks | `` di` `` | `` da` `` |
| parentheses | `di(` or `dib` | `da(` or `dab` |
| brackets | `di[` | `da[` |
| braces | `di{` or `diB` | `da{` or `daB` |
| angle brackets | `di<` | `da<` |

### Practice Task 6.2
1. In `main.go`, find a string in quotes (line 90)
2. Position cursor inside the quotes
3. `ci"` - change inside quotes, type new text
4. Undo, try `ca"` - changes including quotes
5. Find a function call with `()`, use `ci(` to change arguments
6. Find a struct `{}`, use `ci{` to change contents

---

## Lesson 6.3: Tag Objects (HTML)

| Command | Action |
|---------|--------|
| `dit` | delete inner tag |
| `dat` | delete around tag |
| `cit` | change inner tag |
| `vit` | select inner tag |

### Practice Task 6.3
1. Open `task-manager/templates/index.html`
2. Position cursor inside `<h1>Task Manager</h1>`
3. Press `cit`, type `My Tasks`, press `Esc`
4. Undo
5. Try `dat` to delete the entire tag
6. Undo

---

## Lesson 6.4: Function and Class Objects (Treesitter)

LazyVim with treesitter adds:

| Object | Inner | Around |
|--------|-------|--------|
| function | `dif` | `daf` |
| class | `dic` | `dac` |
| parameter | `dia` | `daa` |
| conditional | `dii` | `dai` |
| loop | `dil` | `dal` |

### Practice Task 6.4
1. In `main.go`, go to any function
2. Press `vaf` to select the entire function
3. Press `Esc`, try `vif` - selects function body only
4. Go to a function parameter
5. Try `via` to select the parameter

---

## Lesson 6.5: Advanced Motions

| Motion | Action |
|--------|--------|
| `f{char}` | find char forward (on line) |
| `F{char}` | find char backward |
| `t{char}` | till char forward (before it) |
| `T{char}` | till char backward |
| `;` | repeat last f/F/t/T |
| `,` | repeat opposite direction |

### Practice Task 6.5
1. On a line with multiple same characters
2. Press `f"` to jump to first quote
3. Press `;` to jump to next quote
4. Press `,` to go back
5. Use `dt"` to delete till the quote
6. Use `ct)` to change till the closing paren

---

## Lesson 6.6: Combining It All

Power combos:
- `ci"` - change inside quotes
- `da(` - delete function arguments including parens
- `yiw` - yank inner word
- `vip` - select paragraph
- `>i{` - indent inside braces
- `gUiw` - uppercase word
- `df,` - delete to and including comma
- `ct.` - change until period

### Practice Task 6.6
1. Create: `:e learning/text-objects-practice.txt`
2. Type:
```
function greet(name, age, city) {
    console.log("Hello, " + name);
    return { name: name, age: age };
}
```
3. Use `ci(` to change function parameters
4. Undo, use `di{` to empty the function body
5. Undo, use `yi"` to yank the string
6. Use `f,` then `;` to navigate commas
7. Use `dt,` to delete up to a comma

---

## Chapter 6 Test

**Target time: Under 3 minutes**

### Setup
Create: `:e learning/ch6-test.txt`
```
const config = {
    apiUrl: "https://api.example.com",
    timeout: 5000,
    headers: {
        "Content-Type": "application/json",
        "Authorization": "Bearer token123"
    }
};

function fetchData(url, options) {
    return fetch(url, options).then(r => r.json());
}
```

### Test Tasks:
1. Change `"https://api.example.com"` to `"http://localhost:3000"` using `ci"`
2. Delete the entire `headers` object (lines 4-7) using `da{` when positioned right
3. Undo
4. Select the entire function using `vaf`
5. Change the function parameters using `ci(`
6. Undo
7. Delete `Bearer token123` using `di"`
8. Undo
9. Yank the word `config` using `yiw`
10. Use `f:` and `;` to navigate through colons

**Record your time in progress.md!**

---

## Commands Summary

```
Inner/Around:  i (inner)  a (around)

Words:         iw aw iW aW
Sentences:     is as
Paragraphs:    ip ap

Quotes:        i" a" i' a' i` a`
Brackets:      i( a( i[ a[ i{ a{ i< a<
Aliases:       ib=i(  iB=i{

Tags:          it at

Treesitter:    if af (function)
               ic ac (class)
               ia aa (parameter)

Find:          f{c} F{c} t{c} T{c}
Repeat:        ; ,
```

---

**Next Chapter**: Macros & Registers
