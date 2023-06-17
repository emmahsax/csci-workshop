def add(a, b)
  puts "ADDING #{a}+#{b}:"
  return a + b
end

def subtract(a, b)
  puts "SUBTRACTING #{a}-#{b}:"
  return a - b
end

def multiply(a, b)
  puts "MULTIPLYING #{a}*#{b}:"
  return a * b
end

def divide(a, b)
  puts "DIVIDING #{a}/#{b}:"
  return a / b
end

#~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

puts "Let's do some math with just arithmetic functions!"

# We can return the values from functions as other variables to use later.
age = add(30, 5)
height = subtract(78, 4)
weight = multiply(90, 2)
IQ = divide(100, 2)

puts "Age: #{age}, Height: #{height}, Weight: #{weight}, IQ: #{IQ}"

# A puzzle for the extra credit, type it in anyway.
puts "Here is a puzzle."

what = add(age, subtract(height, multiply(weight, divide(IQ, 2))))

puts "That becomes: #{what}. Could you do it by hand?"
