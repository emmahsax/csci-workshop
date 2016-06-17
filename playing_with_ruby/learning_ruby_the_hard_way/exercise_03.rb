## This is the third exercise that can be found here:
##   http://learnrubythehardway.org/book/ex3.html
## It's basic math operations.
## Note that Ruby always follows PEMDAS (% is included in division).

puts "I will now count my chickens:"

puts "Hens: #{25 + 30 / 6}"
puts "Roosters: #{100 - 25 * 3 % 4}"

puts "I will now count the eggs:"

puts 3.0 + 2.0 + 1.0 - 5.0 + 4.0 % 2.0 - 1.0 / 4.0 + 6.0
# you need floating numbers above so that Ruby can work with decimals

puts "Is it true that 3 + 2 < 5 - 7?"

puts 3 + 2 < 5 - 7

puts "What is 3 + 2? #{3 + 2}"
puts "What is 5 - 7? #{5 - 7}"

puts "Oh, that's why it's false."

puts "How about some more."

puts "Is it greater? #{5 > -2}"
puts "Is it greater or equal? #{5 >= -2}"
puts "Is it less or equal? #{5 <= -2}"
