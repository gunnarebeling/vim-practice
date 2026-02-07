# CHAPTER 3 TEST: Visual Mode Mastery
# Complete all tasks in under 3 minutes
# Open with: nvim learning/ch3-test.rb

# ============================================
# TASK 1: Character Visual (v)
# Select "world" and change it to "Ruby"
# ============================================

def greet
  puts "Hello, Ruby!"
end

# ============================================
# TASK 2: Line Visual (V)
# Delete the entire debug_info method (all 4 lines)
# ============================================



# ============================================
# TASK 3: Block Visual (Ctrl-v)
# Remove the "old_" prefix from all 4 variable names
# ============================================

class DataProcessor
  def initialize
    @name = "processor"
    @version = "1.0"
    @status = "active"
    @count = 0
  end
end

# ============================================
# TASK 4: Block Insert (Ctrl-v + I)
# Add "attr_reader :" in front of each attribute name
# Result should be: attr_reader :name, attr_reader :email, etc.
# ============================================

class User
attr_reader :name
attr_reader :email
attr_reader :age
attr_reader :role
end

# ============================================
# TASK 5: Visual + Motion
# Using V and a motion, select and yank the
# entire process method (all 5 lines), then
# paste it below the Robot class
# ============================================

class Robot
  def process(data)
    cleaned = data.strip
    parsed = parse(cleaned)
    validate(parsed)
  end
end

# Paste the process method here:
def process(data)
  cleaned = data.strip
  parsed = parse(cleaned)
  validate(parsed)
end


# ============================================
# TASK 6: Case Change
# Make SCREAMING_CONSTANT lowercase using visual + u
# Make whisper_constant UPPERCASE using visual + U
# ============================================

screaming_constant = "too loud"
WHISPER_CONSTANT = "speak up"

# ============================================
# BONUS: Reselect
# After completing task 6, use gv to reselect
# your last selection, then indent it with >
# ============================================
