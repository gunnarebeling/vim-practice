# Chapter 2: Editing Essentials

## Learning Goals
Master the core editing commands: insert, delete, change, yank, and put.

---

## Lesson 2.1: Entering Insert Mode

| Key | Action |
|-----|--------|
| `i` | insert before cursor |
| `a` | append after cursor |
| `I` | insert at beginning of line |
| `A` | append at end of line |
| `o` | open line below |
| `O` | open line above |

**Exit insert mode**: `Esc` or `Ctrl-[`

### Practice Task 2.1
1. Open `task-manager/main.go`
2. Go to line 10, press `A`, add a comment: ` // Task represents a todo item`
3. Press `Esc`, go to line 16, press `O` to open line above
4. Type a comment: `// TaskStore manages tasks with thread-safe access`
5. Press `Esc`, undo both changes with `u` twice

---

## Lesson 2.2: Delete Operations

The `d` command + motion = delete.

| Command | Action |
|---------|--------|
| `x` | delete character under cursor |
| `dd` | delete entire line |
| `D` | delete to end of line |
| `dw` | delete word |
| `d$` | delete to end of line |
| `d0` | delete to start of line |

### Practice Task 2.2
1. Create a practice file: `:e learning/scratch.txt`
2. Enter insert mode and type some lines
3. Practice `dd` to delete lines
4. Practice `dw` to delete words
5. Press `u` to undo, `Ctrl-r` to redo
6. Close without saving: `:q!`

---

## Lesson 2.3: Change Operations

`c` = delete + enter insert mode.

| Command | Action |
|---------|--------|
| `cc` | change entire line |
| `C` | change to end of line |
| `cw` | change word |
| `ci"` | change inside quotes |
| `ca(` | change around parentheses |

### Practice Task 2.3
1. Open `task-manager/main.go`, go to line 10
2. Position cursor on `Task`, press `cw`, type `TodoItem`, press `Esc`
3. Undo with `u`
4. Go to line 91, find `":8080"`, position inside quotes, press `ci"`, type `:3000`
5. Undo with `u`

---

## Lesson 2.4: Yank and Put (Copy/Paste)

`y` = yank (copy), `p` = put (paste)

| Command | Action |
|---------|--------|
| `yy` | yank entire line |
| `yw` | yank word |
| `y$` | yank to end of line |
| `p` | put after cursor |
| `P` | put before cursor |

**Note**: Delete (`d`) also yanks! Deleted text can be put.

### Practice Task 2.4
1. In `main.go`, go to line 10 (`type Task struct`)
2. Press `yy` to yank the line
3. Go to line 20, press `p` to paste below
4. Press `u` to undo
5. Try: `dd` then `p` - you just moved a line!
6. Undo with `u`

---

## Lesson 2.5: The Dot Command

| Key | Action |
|-----|--------|
| `u` | undo |
| `Ctrl-r` | redo |
| `.` | repeat last change |

The `.` command is incredibly powerful!

### Practice Task 2.5
1. Search for `Task` with `/Task<Enter>`
2. Press `cw`, type `Item`, press `Esc`
3. Press `n` to go to next "Task"
4. Press `.` to repeat the change!
5. Continue: `n` then `.` to change more
6. Undo all with `u` repeatedly

---

## Chapter 2 Test

**Target time: Under 3 minutes**

### Setup
Create a test file: `:e learning/ch2-test.txt`

Type this content:
```
function calculateTotal(items) {
    let total = 0;
    for (let i = 0; i < items.length; i++) {
        total += items[i].price;
    }
    return total;
}
```

### Test Tasks:
1. Change `calculateTotal` to `sumPrices` (use `cw`)
2. Delete the entire for loop line with `dd`
3. Undo that delete
4. Yank lines 3-5 (`3yy` from line 3) and paste below line 6
5. Delete the duplicated lines
6. Change `price` to `cost` using `cw`
7. Add a comment above the function using `O`
8. Save and close: `:wq`

**Record your time in progress.md!**

---

## Commands Summary

```
Insert:  i a I A o O
Delete:  x dd D dw d{motion}
Change:  cc C cw c{motion}
Yank:    yy yw y{motion}
Put:     p P
Undo:    u Ctrl-r .
```

---

**Next Chapter**: Visual Mode Mastery
