## This is designed to play around with both arrays and how to add things to them, as well
## as take another look at while loops.

ten_things = "Apples Oranges Crows Telephone Light Sugar"

puts "Wait there are not 10 things in that list. Let's fix that."

stuff = ten_things.split(' ')
more_stuff = ["Day", "Night", "Song", "Frisbee", "Corn", "Banana", "Girl", "Boy"]

# add items until the length of stuff is 10
while stuff.length < 10
  next_one = more_stuff.shift
  puts "Adding: #{next_one}"
  stuff.push(next_one)
  puts "There are #{stuff.length} items now."
end

puts "There we go: #{stuff}"

puts "Let's do some things with stuff."

puts stuff[1]
puts stuff[-1] # returns last element
puts stuff.pop()
puts stuff.join(' ')
puts stuff[3..5].join("#")
