# Chapter 1: Navigation Fundamentals

## Learning Goals
By the end of this chapter, you'll move through files without touching your mouse.

---

## Lesson 1.1: Basic Motions (h, j, k, l)

These replace your arrow keys. Your fingers stay on home row.

| Key | Motion |
|-----|--------|
| `h` | left |
| `j` | down |
| `k` | up |
| `l` | right |

### Practice Task 1.1
1. Open `task-manager/main.go` in Neovim: `nvim task-manager/main.go`
2. Navigate to line 10 using only `j` and `k`
3. Move to the word "Task" using `h` and `l`
4. Practice moving in a square pattern: 5j, 10l, 5k, 10h

---

## Lesson 1.2: Word Motions (w, b, e, ge)

Move by words, not characters.

| Key | Motion |
|-----|--------|
| `w` | next word start |
| `b` | previous word start |
| `e` | next word end |
| `ge` | previous word end |

**Tip**: Capital `W`, `B`, `E` move by WORD (whitespace-delimited).

### Practice Task 1.2
1. With `main.go` open, go to line 1 with `gg`
2. Move forward word-by-word through the import statement using `w`
3. Move backward using `b`
4. Count: How many `w` presses to get from `package` to `sync`?

---

## Lesson 1.3: Line Motions (0, ^, $)

Jump within lines instantly.

| Key | Motion |
|-----|--------|
| `0` | start of line |
| `^` | first non-blank character |
| `$` | end of line |

### Practice Task 1.3
1. Go to line 29 (the `Add` function): `29G`
2. Press `0` - where's your cursor?
3. Press `^` - where's your cursor now?
4. Press `$` - where's your cursor?
5. Repeat on line 33 (indented line) - notice the difference between `0` and `^`

---

## Lesson 1.4: File Navigation (gg, G, {n}G, Ctrl-d, Ctrl-u)

Move through entire files.

| Key | Motion |
|-----|--------|
| `gg` | first line |
| `G` | last line |
| `{n}G` | go to line n |
| `Ctrl-d` | half-page down |
| `Ctrl-u` | half-page up |

### Practice Task 1.4
1. Press `G` to go to end of file
2. Press `gg` to go to beginning
3. Go to line 81 (main function): `81G`
4. Scroll down with `Ctrl-d`, then back up with `Ctrl-u`

---

## Lesson 1.5: Search (/, ?, n, N, *, #)

Find anything instantly.

| Key | Action |
|-----|--------|
| `/pattern` | search forward |
| `?pattern` | search backward |
| `n` | next match |
| `N` | previous match |
| `*` | search word under cursor (forward) |
| `#` | search word under cursor (backward) |

### Practice Task 1.5
1. From top of file (`gg`), search for "func": `/func<Enter>`
2. Press `n` to cycle through all functions
3. Press `N` to go backward
4. Put cursor on word `Task` and press `*` - watch it find all occurrences
5. Press `#` to search backward for `Task`

---

## Lesson 1.6: Jump Motions (%, [[, ]], {, })

Navigate code structure.

| Key | Motion |
|-----|--------|
| `%` | matching bracket/paren |
| `[[` | previous function/section |
| `]]` | next function/section |
| `{` | previous blank line (paragraph) |
| `}` | next blank line (paragraph) |

### Practice Task 1.6
1. Go to line 29, position cursor on the `{` of the function
2. Press `%` - where did you land?
3. Press `%` again to go back
4. Press `]]` repeatedly to jump through functions
5. Press `[[` to go back through functions

---

## Chapter 1 Test

Complete these tasks. Time yourself.

**Target time: Under 2 minutes total**

### Test Tasks (do in order, don't use mouse):

1. Open `task-manager/main.go`
2. Go to the last line of the file
3. Search backward for "Delete"
4. Go to the opening `{` of that function and jump to its closing `}`
5. Go to line 50
6. Move to the word "result" on that line
7. Go to the first line of the file
8. Search forward for "http.Handle" (the static file handler)
9. Go to the beginning of that line

**Record your time in progress.md!**

---

## Commands Summary

```
Basic:      h j k l
Words:      w b e ge (W B E for WORDS)
Line:       0 ^ $
File:       gg G {n}G Ctrl-d Ctrl-u
Search:     /pattern ?pattern n N * #
Structure:  % [[ ]] { }
```

---

**Next Chapter**: Editing Essentials (insert, delete, change, yank)
