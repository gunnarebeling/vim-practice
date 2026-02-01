# Chapter 5: LazyVim Specifics

## Learning Goals
Master LazyVim's built-in tools: file explorer, fuzzy finder, and LSP.

---

## Lesson 5.1: Neo-tree File Explorer

| Keymap | Action |
|--------|--------|
| `<leader>e` | Toggle file explorer (focus) |
| `<leader>E` | Toggle file explorer (cwd) |
| `<leader>fe` | Explorer in floating window |

### Neo-tree Navigation
| Key | Action |
|-----|--------|
| `j/k` | navigate up/down |
| `Enter` | open file/toggle folder |
| `l` | open / expand |
| `h` | collapse / go to parent |
| `a` | add file/directory |
| `d` | delete |
| `r` | rename |
| `y` | copy to clipboard |
| `x` | cut |
| `p` | paste |
| `R` | refresh |
| `?` | show help |

### Practice Task 5.1
1. Press `Space e` to open neo-tree
2. Navigate to `task-manager/templates`
3. Press `a` to add a new file, name it `test.html`
4. Press `r` to rename it to `test2.html`
5. Press `d` to delete it
6. Press `Space e` to close neo-tree

---

## Lesson 5.2: Telescope Fuzzy Finder

The most powerful navigation tool in LazyVim.

| Keymap | Action |
|--------|--------|
| `<leader>ff` | Find files |
| `<leader>fg` | Live grep (search content) |
| `<leader>fb` | Browse buffers |
| `<leader>fr` | Recent files |
| `<leader>f"` | Find registers |
| `<leader>fc` | Find config files |
| `<leader><leader>` | Find files (root) |
| `<leader>/` | Grep in project |

### Telescope Controls
| Key | Action |
|-----|--------|
| `Ctrl-j/k` | move down/up in results |
| `Ctrl-n/p` | move down/up (alt) |
| `Enter` | open selection |
| `Ctrl-x` | open in horizontal split |
| `Ctrl-v` | open in vertical split |
| `Ctrl-t` | open in new tab |
| `Esc` | close telescope |

### Practice Task 5.2
1. Press `Space ff` to find files
2. Type `main` to filter, press Enter to open
3. Press `Space fg` to grep
4. Search for `TaskStore`, navigate results, open one
5. Press `Space fr` to see recent files

---

## Lesson 5.3: Which-Key

Press `<leader>` and wait - which-key shows available commands.

### Practice Task 5.3
1. Press `Space` and wait 300ms
2. Explore the menu that appears
3. Press `g` from normal mode and wait - see go-to commands
4. Press `z` and wait - see fold/scroll commands
5. Press `]` and wait - see next commands
6. Press `[` and wait - see prev commands

---

## Lesson 5.4: LSP Features

LazyVim auto-configures LSP. For Go, ensure gopls is installed.

| Keymap | Action |
|--------|--------|
| `gd` | Go to definition |
| `gr` | Go to references |
| `gD` | Go to declaration |
| `gI` | Go to implementation |
| `gy` | Go to type definition |
| `K` | Hover documentation |
| `<leader>ca` | Code actions |
| `<leader>cr` | Rename symbol |
| `<leader>cf` | Format file |
| `]d` | Next diagnostic |
| `[d` | Previous diagnostic |
| `<leader>cd` | Line diagnostics |

### Practice Task 5.4
1. Open `task-manager/main.go`
2. Position cursor on `Task` (line 10)
3. Press `K` to see hover docs
4. Press `gr` to see references
5. Press `gd` on a function call to go to definition
6. Press `Ctrl-o` to jump back
7. Try `Space cr` on a variable to rename it
8. Press `u` to undo the rename

---

## Lesson 5.5: Diagnostics and Trouble

| Keymap | Action |
|--------|--------|
| `<leader>xx` | Toggle Trouble |
| `<leader>xX` | Buffer diagnostics |
| `<leader>cs` | Symbols outline |
| `<leader>xL` | Location list |
| `<leader>xQ` | Quickfix list |

### Practice Task 5.5
1. Open main.go
2. Press `Space xx` to open Trouble (diagnostics list)
3. Navigate through any issues
4. Press `Space cs` to see symbols outline
5. Press `q` to close panels

---

## Lesson 5.6: Git Integration

| Keymap | Action |
|--------|--------|
| `<leader>gg` | LazyGit |
| `<leader>gf` | Git file history |
| `<leader>gb` | Git blame |
| `]h` | Next hunk |
| `[h` | Previous hunk |
| `<leader>ghs` | Stage hunk |
| `<leader>ghr` | Reset hunk |
| `<leader>ghp` | Preview hunk |

### Practice Task 5.6
1. Initialize git if needed: run `git init` in terminal
2. Press `Space gg` to open LazyGit
3. Press `q` to close
4. Make a change to a file
5. Press `Space ghp` to preview the change

---

## Chapter 5 Test

**Target time: Under 3 minutes**

### Test Tasks:
1. Use Telescope to find and open `index.html`
2. Use neo-tree to create a new file `task-manager/test.go`
3. Open the new file from neo-tree
4. Use Telescope grep to find all occurrences of `http`
5. Navigate to main.go using `Space ff`
6. Use `gd` to go to the definition of `TaskStore`
7. Use `gr` to find all references to `Task`
8. Delete the test.go file using neo-tree
9. Use `Space fr` to open a recent file
10. Open which-key and find the terminal toggle command

**Record your time in progress.md!**

---

## Commands Summary

```
Neo-tree:   <leader>e E fe
            j/k Enter h/l a d r

Telescope:  <leader>ff fg fb fr /
            Ctrl-j/k Enter Ctrl-x/v/t

LSP:        gd gr gD gI K
            <leader>ca cr cf cd
            ]d [d

Trouble:    <leader>xx xX cs

Git:        <leader>gg gb gf
            ]h [h <leader>ghs ghr ghp
```

---

**Next Chapter**: Text Objects & Advanced Motions
