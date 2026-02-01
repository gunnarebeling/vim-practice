# Chapter 7: Macros & Registers

## Learning Goals
Automate repetitive tasks with macros and understand Vim's register system.

---

## Lesson 7.1: Recording Macros

| Key | Action |
|-----|--------|
| `q{a-z}` | start recording to register |
| `q` | stop recording |
| `@{a-z}` | play macro |
| `@@` | replay last macro |
| `{n}@{a-z}` | play macro n times |

### Practice Task 7.1
1. Create: `:e learning/macro-practice.txt`
2. Type these lines:
```
item1
item2
item3
item4
item5
```
3. Go to line 1
4. Press `qa` to start recording to register `a`
5. Press `I- [ ] ` then `Esc` then `j`
6. Press `q` to stop recording
7. Press `4@a` to run it 4 more times
8. Result: all lines now have `- [ ] ` prefix!

---

## Lesson 7.2: Macro Best Practices

Tips for reusable macros:
1. Start from a consistent position (`0` or `^`)
2. End in position to repeat (`j` to move to next line)
3. Use motions that work on varying content (`w` not `llll`)
4. Test on one line before running many times

### Practice Task 7.2
1. Create lines:
```
firstName
lastName
emailAddress
phoneNumber
```
2. Goal: Convert to `"firstName": "",`
3. Record macro on first line:
   - `0` go to start
   - `i"` insert quote
   - `Esc A": "",` go to end, add rest
   - `Esc j` exit and move down
4. Play `3@a` for remaining lines

---

## Lesson 7.3: Understanding Registers

Vim has many registers for storing text.

| Register | Purpose |
|----------|---------|
| `"` | default register (unnamed) |
| `0` | last yank |
| `1-9` | delete history |
| `a-z` | named registers |
| `A-Z` | append to named register |
| `+` | system clipboard |
| `*` | primary selection |
| `_` | black hole (delete without saving) |
| `/` | last search pattern |
| `:` | last command |
| `.` | last inserted text |
| `%` | current filename |

### Practice Task 7.3
1. Yank a line with `yy`
2. Delete a different line with `dd`
3. Press `"0p` - pastes the yanked line (not deleted)
4. Press `:reg` to see all registers
5. Try `"+yy` to yank to system clipboard
6. Paste from clipboard: `"+p`

---

## Lesson 7.4: Using Named Registers

| Command | Action |
|---------|--------|
| `"ayy` | yank line to register a |
| `"ap` | paste from register a |
| `"Ayy` | append to register a |
| `"bdiw` | delete word to register b |

### Practice Task 7.4
1. Go to a line, press `"ayy` to yank to `a`
2. Go to another line, press `"byy` to yank to `b`
3. Now you have two separate yanks!
4. Press `"ap` to paste from a
5. Press `"bp` to paste from b
6. Use `"Ayy` on another line to append to a
7. Press `"ap` - now pastes both lines

---

## Lesson 7.5: The Expression Register

| Command | Action |
|---------|--------|
| `Ctrl-r =` | insert expression result (insert mode) |
| `"=` | expression register (normal mode) |

### Practice Task 7.5
1. Enter insert mode
2. Press `Ctrl-r =`
3. Type `5 + 3` and press Enter
4. `8` is inserted!
5. Try `Ctrl-r =strftime('%Y-%m-%d')` for date

---

## Lesson 7.6: Editing Macros

Macros are just text in registers - you can edit them!

1. Paste the macro: `"ap`
2. Edit the text
3. Yank it back: `"ayy`

### Practice Task 7.6
1. Record a simple macro to register `a`
2. Open a new line and press `"ap` to paste it
3. You'll see the key sequence as text
4. Edit it (change something)
5. Select and yank back with `"ay$`
6. Test the modified macro with `@a`

---

## Chapter 7 Test

**Target time: Under 4 minutes**

### Setup
Create: `:e learning/ch7-test.txt`
```
apple
banana
cherry
date
elderberry

buyMilk
cleanRoom
doLaundry
feedCat
goShopping
```

### Test Tasks:
1. Record a macro that transforms fruit names to `- fruit: "apple",` format
2. Apply it to all 5 fruit lines
3. Delete the formatted fruit lines using `5dd`
4. Paste them back with `p`
5. Yank "banana" to register `b`, "cherry" to register `c`
6. Paste from register `b`
7. Paste from register `c`
8. Record a macro to transform task names (buyMilk) to `{ task: "buyMilk" },`
9. Apply to all 5 task lines
10. Check `:reg` to see your registers

**Record your time in progress.md!**

---

## Commands Summary

```
Record:    q{a-z}  q (stop)
Play:      @{a-z}  @@  {n}@{a-z}

Registers:
  "        unnamed (default)
  0        last yank
  1-9      delete history
  a-z      named
  A-Z      append to named
  +        system clipboard
  _        black hole

Use:       "{reg}y  "{reg}d  "{reg}p
Clipboard: "+y  "+p
View:      :reg

Expression: Ctrl-r = (insert mode)
```

---

**Next Chapter**: Search & Replace
