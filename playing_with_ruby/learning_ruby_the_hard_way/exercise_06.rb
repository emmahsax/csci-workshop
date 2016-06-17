## http://learnrubythehardway.org/book/ex6.html
## putting strings into variables to print them

types_of_people = 10
x = "There are #{types_of_people} types of people."
binary = "binary"
do_not = "don't"
y = "Those who know #{binary} and those who #{do_not}."

puts x
puts y

puts "I said: '#{x}'"
puts "I also said: '#{y}'"

hilarious = true

if hilarious == false
	hilarious1 = "false".capitalize
elsif hilarious == true
	hilarious1 = "true".capitalize
end

joke_evaluation = "Isn't that joke so funny?! #{hilarious1}"

puts joke_evaluation

w = "This is the left side of..."
e = "a string with a right side."

puts w + e
