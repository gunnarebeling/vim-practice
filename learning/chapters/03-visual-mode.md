# Chapter 3: Visual Mode Mastery

## Learning Goals
Select text precisely using visual, visual line, and visual block modes.

---

## Lesson 3.1: Character Visual Mode (v)

Press `v` to start selecting character by character.

| Key | Action |
|-----|--------|
| `v` | start visual mode |
| motions | extend selection |
| `d` | delete selection |
| `y` | yank selection |
| `c` | change selection |
| `Esc` | cancel selection |

### Practice Task 3.1
1. Open `task-manager/main.go`
2. Go to line 90, find `http://localhost:8080`
3. Position cursor on `h` of `http`
4. Press `v` to start visual mode
5. Press `e` repeatedly to select the URL
6. Press `y` to yank it

---

## Lesson 3.2: Line Visual Mode (V)

Press `V` (shift+v) to select entire lines.

| Key | Action |
|-----|--------|
| `V` | start line visual mode |
| `j`/`k` | extend selection by lines |
| `>` | indent selection |
| `<` | unindent selection |

### Practice Task 3.2
1. Go to line 29 (start of `Add` function)
2. Press `V` to select the line
3. Press `j` until you've selected the entire function
4. Press `y` to yank
5. Go to end of file with `G`
6. Press `P` to paste
7. Undo with `u`

---

## Lesson 3.3: Block Visual Mode (Ctrl-v)

Select rectangles of text!

| Key | Action |
|-----|--------|
| `Ctrl-v` | start block visual mode |
| `I` | insert at start of each line |
| `A` | append at end of each line |
| `d` | delete block |
| `c` | change block |

### Practice Task 3.3
1. Create a practice file: `:e learning/block-practice.txt`
2. Insert:
```
item1: apple
item2: banana
item3: cherry
item4: date
```
3. Go to line 1, column 1
4. Press `Ctrl-v`, then `3j` to select 4 lines
5. Press `I`, type `// `, press `Esc`
6. Watch it add `// ` to all lines!
7. Undo with `u`

---

## Lesson 3.4: Visual Mode with Text Objects

| Command | Selects |
|---------|---------|
| `viw` | inner word |
| `vi"` | inside double quotes |
| `vi(` | inside parentheses |
| `vi{` | inside braces |
| `vip` | inner paragraph |

Use `a` instead of `i` to include delimiters.

### Practice Task 3.4
1. In `main.go`, go to line 90
2. Position cursor inside the quotes
3. Press `vi"` - selects string inside quotes
4. Press `Esc`, try `va"` - includes the quotes

---

## Lesson 3.5: Useful Visual Commands

| Key | Action |
|-----|--------|
| `gv` | reselect last selection |
| `o` | jump to other end of selection |
| `u` | lowercase selection |
| `U` | uppercase selection |

### Practice Task 3.5
1. Select some text with `v`
2. Press `Esc`, then `gv` to reselect
3. Press `U` to uppercase
4. Undo with `u`

---

## Chapter 3 Test

**Target time: Under 3 minutes**

### Setup
Create: `:e learning/ch3-test.txt`

```
const users = [
    { name: "alice", role: "admin" },
    { name: "bob", role: "user" },
    { name: "charlie", role: "user" },
    { name: "diana", role: "admin" }
];
```

### Test Tasks:
1. Use `V` to select and delete lines 2-5
2. Undo
3. Use `vi"` to select `alice` and change to `ALICE`
4. Use `Ctrl-v` to add `//` prefix to lines 2-5
5. Undo
6. Yank all 6 lines with `V` and paste at end

**Record your time in progress.md!**

---

## Commands Summary

```
Start:     v (char)  V (line)  Ctrl-v (block)
Actions:   d y c > < u U
Reselect:  gv
Jump:      o (other end)
Objects:   vi{char}  va{char}
Block:     I (insert)  A (append)
```

---

**Next Chapter**: Buffers, Windows & Tabs
