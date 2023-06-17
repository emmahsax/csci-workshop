## Reading parts of a file (given as an input) using functions.

input_file = ARGV[0]

def print_all (f)
  puts f.read
end

def rewind (f)
  f.seek(0)
end

def print_a_line(line_count, f)
  puts "#{line_count}: #{f.gets.chomp}"
end

# gets.chomp will read the file until it gets to a '\n', in which case the read head will
# move to the next character or byte. The read head doesn't move again until this function
# is called again

current_file = open(input_file)

puts "First let's print the whole file:"

print_all(current_file)

puts "Now let's rewind, kind of like a video tape."

rewind(current_file)

puts "Let's print three lines:"

for i in 1..3 do
  print_a_line(i, current_file)
end
