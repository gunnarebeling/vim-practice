package learning

// CHAPTER 4 TEST: Buffers, Windows & Tabs
// Complete all tasks to master multi-file workflows
// Open with: nvim learning/ch4-test.go

// ============================================
// SETUP: Open these files as buffers first
// :e task-manager/main.go
// :e task-manager/templates/index.html
// :e learning/progress.md
// ============================================

// ============================================
// TASK 1: Buffer Navigation Basics
// Practice these commands:
//   :ls or :buffers    - List all buffers
//   :b <number>        - Go to buffer by number
//   :b <partial-name>  - Go to buffer by partial name
//   :bn                - Next buffer
//   :bp                - Previous buffer
//   :b#                - Alternate (last) buffer
//
// Challenge: Cycle through all 4 buffers using :bn
// Then jump back to main.go using :b main
// ============================================

func task1_bufferNavigation() {
	// Practice buffer commands here
	// Try jumping between this file and main.go
}

// ============================================
// TASK 2: LazyVim Buffer Keymaps
// LazyVim provides convenient mappings:
//   <S-h>        - Previous buffer (Shift+h)
//   <S-l>        - Next buffer (Shift+l)
//   <leader>bb   - Switch buffer (picker)
//   <leader>bd   - Delete buffer
//   <leader>`    - Switch to alternate buffer
//
// Challenge: Use <S-h> and <S-l> to navigate buffers
// Then use <leader>bb to open the buffer picker
// ============================================

func task2_lazyvimBuffers() {
	// These keymaps are much faster than :bn/:bp
	// Shift+H and Shift+L become second nature
}

// ============================================
// TASK 3: Window Splits
// Split commands:
//   :sp or Ctrl-w s    - Horizontal split
//   :vsp or Ctrl-w v   - Vertical split
//   Ctrl-w w           - Cycle between windows
//   Ctrl-w h/j/k/l     - Move to window (direction)
//   Ctrl-w H/J/K/L     - Move window (reposition)
//   Ctrl-w =           - Equal size all windows
//   Ctrl-w _           - Maximize height
//   Ctrl-w |           - Maximize width
//   Ctrl-w q or :q     - Close window
//
// Challenge:
// 1. Open main.go in a vertical split (:vsp task-manager/main.go)
// 2. Open index.html in a horizontal split (:sp task-manager/templates/index.html)
// 3. Navigate between all 3 windows using Ctrl-w h/j/k/l
// 4. Make all windows equal size with Ctrl-w =
// ============================================

func task3_windowSplits() {
	// Splits let you view multiple files simultaneously
	// Great for comparing code or referencing docs
}

// ============================================
// TASK 4: LazyVim Window Keymaps
// LazyVim window shortcuts:
//   <leader>w      - Window menu (try it!)
//   <leader>-      - Split below
//   <leader>|      - Split right
//   <leader>wd     - Delete window
//   <leader>wm     - Maximize toggle
//
// Challenge: Use <leader>w to explore the window menu
// Create splits using <leader>- and <leader>|
// ============================================

func task4_lazyvimWindows() {
	// The <leader>w menu is very helpful when learning
	// It shows all available window commands
}

// ============================================
// TASK 5: Tabs
// Tab commands:
//   :tabnew            - New tab
//   :tabnew <file>     - New tab with file
//   gt                 - Next tab
//   gT                 - Previous tab
//   :tabclose or :tabc - Close tab
//   :tabonly or :tabo  - Close all other tabs
//   {n}gt              - Go to tab n
//
// LazyVim tab keymaps:
//   <leader><tab>l     - Last tab
//   <leader><tab>f     - First tab
//   <leader><tab><tab> - New tab
//   <leader><tab>]     - Next tab
//   <leader><tab>[     - Previous tab
//   <leader><tab>d     - Close tab
//
// Challenge:
// 1. Create a new tab with main.go (:tabnew task-manager/main.go)
// 2. Create another tab with index.html
// 3. Use gt and gT to navigate between tabs
// 4. Use 1gt to jump to the first tab
// ============================================

func task5_tabs() {
	// Tabs are great for separate contexts
	// Example: Tab 1 for Go code, Tab 2 for templates
}

// ============================================
// TASK 6: Buffer Management
// Useful buffer commands:
//   :bd                - Delete (close) current buffer
//   :bd <number>       - Delete buffer by number
//   :%bd               - Delete all buffers
//   :e <file>          - Edit file (opens in buffer)
//   :w                 - Save current buffer
//   :wa                - Save all buffers
//
// LazyVim extras:
//   <leader>bo         - Delete other buffers
//   <leader>bD         - Delete buffer + window
//
// Challenge: Open 3 files, then use :bd to close one
// Use :ls to verify it's gone
// ============================================

func task6_bufferManagement() {
	// Keep your buffer list clean
	// Use <leader>bo to close everything except current
}

// ============================================
// TASK 7: Practical Workflow
// Combine everything! Common patterns:
//
// Pattern A - Code + Reference side-by-side:
//   :vsp | :b main     (split and show main.go)
//
// Pattern B - Quick file peek:
//   Ctrl-w v           (split)
//   :e <file>          (open file)
//   (read what you need)
//   Ctrl-w q           (close split)
//
// Pattern C - Multiple contexts with tabs:
//   Tab 1: main.go + tests side-by-side
//   Tab 2: templates work
//   Tab 3: documentation
//
// Challenge: Set up Pattern A, then close it
// ============================================

func task7_practicalWorkflow() {
	// Develop your own preferred patterns
	// Speed comes from muscle memory
}

// ============================================
// FINAL CHALLENGE (Under 2 minutes)
//
// Starting from this file:
// 1. Open main.go in a vertical split
// 2. Navigate to the split showing main.go
// 3. Open index.html in a horizontal split below
// 4. Make all windows equal size
// 5. Create a new tab
// 6. In the new tab, open progress.md
// 7. Go back to the first tab
// 8. Close all splits except one
// 9. Delete the extra tab
//
// End state: Single window, single tab, this file
// ============================================

func finalChallenge() {
	// Time yourself!
	// Goal: under 2 minutes
}

// ============================================
// QUICK REFERENCE CARD
// ============================================

type QuickReference struct {
	// BUFFERS:
	//   :ls        List buffers
	//   :b N       Go to buffer N
	//   :bn/:bp    Next/prev buffer
	//   :bd        Delete buffer
	//   <S-h/l>    Prev/next (LazyVim)

	// WINDOWS:
	//   Ctrl-w s   Horizontal split
	//   Ctrl-w v   Vertical split
	//   Ctrl-w q   Close window
	//   Ctrl-w =   Equal size
	//   Ctrl-w hjkl Navigate

	// TABS:
	//   :tabnew    New tab
	//   gt/gT      Next/prev tab
	//   :tabclose  Close tab
}
