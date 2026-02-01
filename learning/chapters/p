# Chapter 4: Buffers, Windows & Tabs

## Learning Goals
Work with multiple files efficiently using buffers, splits, and tabs.

---

## Lesson 4.1: Understanding Buffers

A buffer is a file loaded into memory. You can have many buffers open.

| Command | Action |
|---------|--------|
| `:e {file}` | edit/open file in new buffer |
| `:ls` or `:buffers` | list all buffers |
| `:b {n}` | go to buffer n |
| `:b {name}` | go to buffer by partial name |
| `:bn` | next buffer |
| `:bp` | previous buffer |
| `:bd` | delete (close) buffer |

### Practice Task 4.1
1. Open main.go: `nvim task-manager/main.go`
2. Open another file: `:e task-manager/templates/index.html`
3. List buffers: `:ls`
4. Switch to main.go: `:b main` (partial match works!)
5. Switch back: `:bn`
6. Close current buffer: `:bd`

---

## Lesson 4.2: LazyVim Buffer Keymaps

LazyVim provides easier buffer navigation.

| Keymap | Action |
|--------|--------|
| `<leader>bb` | Switch buffer (Telescope) |
| `<leader>bd` | Delete buffer |
| `<S-h>` | Previous buffer |
| `<S-l>` | Next buffer |
| `<leader>bo` | Delete other buffers |

**Note**: `<leader>` is Space in LazyVim.

### Practice Task 4.2
1. Open several files with `:e`
2. Press `Space bb` to see buffer picker
3. Use `Shift+h` and `Shift+l` to cycle buffers
4. Close a buffer with `Space bd`

---

## Lesson 4.3: Window Splits

View multiple buffers simultaneously.

| Command | Action |
|---------|--------|
| `:sp {file}` | horizontal split |
| `:vs {file}` | vertical split |
| `Ctrl-w s` | horizontal split current |
| `Ctrl-w v` | vertical split current |
| `Ctrl-w w` | cycle windows |
| `Ctrl-w h/j/k/l` | move to window |
| `Ctrl-w c` | close window |
| `Ctrl-w o` | close other windows |

### Practice Task 4.3
1. Open main.go
2. Vertical split: `:vs task-manager/templates/index.html`
3. Move between windows: `Ctrl-w h` and `Ctrl-w l`
4. Create horizontal split: `Ctrl-w s`
5. Close a window: `Ctrl-w c`
6. Close all but current: `Ctrl-w o`

---

## Lesson 4.4: Window Resizing

| Command | Action |
|---------|--------|
| `Ctrl-w =` | equalize window sizes |
| `Ctrl-w _` | maximize height |
| `Ctrl-w \|` | maximize width |
| `Ctrl-w +/-` | increase/decrease height |
| `Ctrl-w >/<` | increase/decrease width |
| `{n}Ctrl-w _` | set height to n lines |

### Practice Task 4.4
1. Create a vertical split
2. Press `Ctrl-w >` several times to widen left window
3. Press `Ctrl-w =` to equalize
4. Press `Ctrl-w |` to maximize current window width
5. Press `Ctrl-w =` again

---

## Lesson 4.5: LazyVim Window Keymaps

| Keymap | Action |
|--------|--------|
| `<leader>-` | Split below |
| `<leader>\|` | Split right |
| `<leader>wd` | Delete window |
| `<leader>wm` | Toggle maximize |

### Practice Task 4.5
1. Press `Space |` to split right
2. Press `Space -` to split below
3. Press `Space wm` to maximize current
4. Press `Space wm` again to restore

---

## Lesson 4.6: Tabs (Tab Pages)

Tabs hold window layouts. Less common but useful for contexts.

| Command | Action |
|---------|--------|
| `:tabnew {file}` | new tab |
| `:tabn` or `gt` | next tab |
| `:tabp` or `gT` | previous tab |
| `:tabc` | close tab |
| `{n}gt` | go to tab n |

### Practice Task 4.6
1. Open a file
2. Create new tab: `:tabnew task-manager/go.mod`
3. Switch tabs: `gt` and `gT`
4. Close tab: `:tabc`

---

## Chapter 4 Test

**Target time: Under 2 minutes**

### Test Tasks:
1. Open `task-manager/main.go`
2. Create a vertical split with `templates/index.html`
3. In the right window, create a horizontal split with `templates/task.html`
4. Navigate to each window using `Ctrl-w` commands
5. Equalize all window sizes
6. Close the task.html window
7. Switch to the left window and open `go.mod` in a new buffer (not split)
8. Use `:ls` to see your buffers
9. Switch between main.go and go.mod using buffer commands
10. Close all windows except current

**Record your time in progress.md!**

---

## Commands Summary

```
Buffers:   :e :ls :b :bn :bp :bd
LazyVim:   <leader>bb bd  <S-h> <S-l>

Splits:    :sp :vs  Ctrl-w s/v
Navigate:  Ctrl-w h/j/k/l/w
Close:     Ctrl-w c/o
Resize:    Ctrl-w = _ | + - > <
LazyVim:   <leader>- | wd wm

Tabs:      :tabnew :tabn :tabp :tabc  gt gT
```

---

**Next Chapter**: LazyVim Specifics (neo-tree, Telescope, LSP)
