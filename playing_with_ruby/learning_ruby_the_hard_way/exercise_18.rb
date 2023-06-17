# this one is like your scripts with ARGV
def print_two(*args) # takes in as many arguments as we want
  puts "arg1: #{args[0]}; arg2: #{args[1]}"
  if args.length >= 3
  	puts "arg3: #{args[2]}"
  end
end

# ok, that *args is actually pointless, we can just do this
def print_two_again(arg1, arg2) # takes in exactly 2 arguments
  puts "arg1: #{arg1}; arg2: #{arg2}"
end

# this just takes one argument
def print_one(arg1)
  puts "arg1: #{arg1}"
end

# this one takes no arguments
def print_none()
  puts "No arguments given."
end

# we call the functions above from inside the main script, like so:
print_two("Emma","Sax")
print_two_again("Aaron","Lemmon")
print_one("Zoe")
print_none()
