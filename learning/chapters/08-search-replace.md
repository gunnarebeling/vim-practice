# Chapter 8: Search & Replace

## Learning Goals
Master substitution, global commands, and project-wide search/replace.

---

## Lesson 8.1: Basic Substitute

Syntax: `:[range]s/{pattern}/{replacement}/[flags]`

| Command | Action |
|---------|--------|
| `:s/old/new/` | first occurrence on current line |
| `:s/old/new/g` | all occurrences on current line |
| `:%s/old/new/g` | all occurrences in file |
| `:%s/old/new/gc` | all occurrences, confirm each |

### Practice Task 8.1
1. Open `main.go`
2. Try `:s/Task/Item/` on a line with Task
3. Undo with `u`
4. Try `:%s/Task/Item/g` for whole file
5. Undo with `u`
6. Try `:%s/Task/Item/gc` and press `y/n` for each

---

## Lesson 8.2: Ranges

| Range | Meaning |
|-------|---------|
| `.` | current line |
| `$` | last line |
| `%` | entire file (same as 1,$) |
| `1,10` | lines 1-10 |
| `'<,'>` | visual selection |
| `.,.+5` | current line plus next 5 |
| `'a,'b` | between marks a and b |

### Practice Task 8.2
1. In `main.go`, go to line 10
2. `:.,20s/Task/Item/g` - lines 10-20
3. Undo
4. Select lines with `V`, then `:s/Task/Item/g`
5. Undo
6. Set mark `a` with `ma`, go elsewhere, set `mb`
7. `:'a,'bs/Task/Item/g` - between marks

---

## Lesson 8.3: Substitute Flags

| Flag | Meaning |
|------|---------|
| `g` | global (all on line) |
| `c` | confirm each |
| `i` | case insensitive |
| `I` | case sensitive |
| `n` | count matches only |
| `e` | no error if no match |

### Practice Task 8.3
1. `:%s/task/TASK/gi` - case insensitive match
2. Undo
3. `:%s/task/TASK/gn` - just count matches
4. `:%s/task/TASK/gc` - confirm each

---

## Lesson 8.4: Special Characters in Patterns

| Character | Meaning |
|-----------|---------|
| `.` | any character |
| `*` | zero or more of previous |
| `\+` | one or more of previous |
| `\?` | zero or one of previous |
| `\|` | or |
| `\(` `\)` | group |
| `\1` | first group backreference |
| `^` | start of line |
| `$` | end of line |
| `\<` `\>` | word boundaries |

### Practice Task 8.4
1. Create: `:e learning/regex-practice.txt`
2. Add:
```
func Add()
func Delete()
func Toggle()
func NewTaskStore()
```
3. `:%s/func \(\w\+\)/function \1/g` - capture and reuse
4. Undo
5. `:%s/\<Add\>/Insert/g` - whole word only

---

## Lesson 8.5: Replacement Special Characters

| Char | Meaning |
|------|---------|
| `&` | entire match |
| `\1-\9` | captured groups |
| `\u` | uppercase next char |
| `\U` | uppercase until `\e` |
| `\l` | lowercase next char |
| `\L` | lowercase until `\e` |
| `\r` | newline |

### Practice Task 8.5
1. In regex-practice.txt:
2. `:%s/\(\w\+\)()/\U\1\e()/g` - uppercase function names
3. Undo
4. `:%s/func/&tion/g` - func becomes function
5. Undo

---

## Lesson 8.6: Global Command

Syntax: `:g/{pattern}/{command}`

Execute command on all matching lines.

| Command | Action |
|---------|--------|
| `:g/pattern/d` | delete matching lines |
| `:g!/pattern/d` | delete non-matching lines |
| `:g/pattern/m$` | move matching lines to end |
| `:g/^$/d` | delete empty lines |
| `:g/TODO/p` | print all TODOs |

### Practice Task 8.6
1. In `main.go`:
2. `:g/func/p` - print all function lines
3. `:g/http/p` - print all http lines
4. Create test file with empty lines mixed in
5. `:g/^$/d` - delete all empty lines
6. Undo

---

## Lesson 8.7: Project-Wide Search/Replace (LazyVim)

| Keymap | Action |
|--------|--------|
| `<leader>sr` | Search and replace |
| `<leader>/` | Grep project |
| `<leader>sg` | Grep |
| `<leader>sw` | Grep word under cursor |

Using Spectre (LazyVim plugin):
1. `<leader>sr` opens Spectre
2. Type search pattern
3. Type replacement
4. Preview changes
5. Apply all or selectively

### Practice Task 8.7
1. Press `Space sr` to open Spectre
2. Search for `Task`
3. Enter replacement `Todo`
4. Review the preview
5. Press `Esc` to close without applying

---

## Chapter 8 Test

**Target time: Under 4 minutes**

### Setup
Create: `:e learning/ch8-test.txt`
```
var firstName = "john";
var lastName = "doe";
var userEmail = "john@example.com";
var userAge = 25;
var isActive = true;

function getUser() {}
function setUser() {}
function deleteUser() {}
function updateUser() {}
```

### Test Tasks:
1. Change all `var` to `let` using substitute
2. Change all instances of `user` (case insensitive) to `person`
3. Undo both changes
4. Use substitute with confirmation on `User` only
5. Delete all empty lines with global command
6. Use capture groups to change `firstName` to `first_name` (add underscore before capital)
7. Hint: `:%s/\([a-z]\)\([A-Z]\)/\1_\l\2/g`
8. Print all lines containing `function`
9. Change all function names to uppercase
10. Undo all changes

**Record your time in progress.md!**

---

## Commands Summary

```
Substitute:  :[range]s/old/new/[flags]
Ranges:      . $ % 1,10 '<,'> 'a,'b
Flags:       g c i I n e

Special:     . * \+ \? \| \(\) \1
             ^ $ \< \>
Replace:     & \1 \u \U \l \L \r

Global:      :g/pattern/cmd
             :g!/pattern/cmd

LazyVim:     <leader>sr (Spectre)
             <leader>sg sw /
```

---

## Final Notes

Congratulations on completing the curriculum! Key habits:
1. **Stay in normal mode** - it's your home base
2. **Think in text objects** - `ci"` not arrow keys
3. **Use the dot command** - make changes repeatable
4. **Learn one new motion per week**
5. **Use :help** - it's comprehensive

Continue practicing with real projects!
