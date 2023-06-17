## Let's play around with different ways to give arguments to functions
def cheese_and_crackers(cheese_count, boxes_of_crackers)
  puts "You have #{cheese_count} cheeses!"
  puts "You have #{boxes_of_crackers} boxes of crackers!"
  puts "Man that's enough for a party!"
  puts "Get a blanket.\n\n"
end

puts "We can just give the function numbers directly:"
puts "cheese_and_crackers(20, 30):"
cheese_and_crackers(20, 30)

puts "OR, we can use variables from our script:"
amount_of_cheese = 10
amount_of_crackers = 50
puts "cheese_and_crackers(amount_of_cheese, amount_of_crackers)"
cheese_and_crackers(amount_of_cheese, amount_of_crackers)

puts "We can even do math inside too:"
puts "cheese_and_crackers(10 + 20, 5 + 6)"
cheese_and_crackers(10 + 20, 5 + 6)

puts "And we can combine the two, variables and math:"
puts "cheese_and_crackers(amount_of_cheese + 100, amount_of_crackers + 1000)"
cheese_and_crackers(amount_of_cheese + 100, amount_of_crackers + 1000)

puts "See?"
