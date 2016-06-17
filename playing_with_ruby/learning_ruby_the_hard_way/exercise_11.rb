## This file consists of two little forms that the user would complete in terminal
## when this file is run. The first asks some basic questions of the user, and the
## second figures out what year they're born based on the user's age.

puts "~~~~~~~~~~~~~~~~~~FORM 1~~~~~~~~~~~~~~~~~~"

print "How old are you? "
age = gets.chomp
print "How tall are you? "
height = gets.chomp
print "How much do you weigh? "
weight = gets.chomp
print "What's your favorite color? "
color = gets.chomp

puts "Your #{age} old, #{height} tall and #{weight} heavy, and your favorite color is #{color}."

puts "~~~~~~~~~~~~~~~~~~FORM 2~~~~~~~~~~~~~~~~~~"

puts "What is your name? "
name = gets.chomp
age = age.to_i #to_i will make this variable an integer so we can do math with it (previously was a string)
puts "Has your birthday already occurred this year? Enter 'yes' or 'no'. "
this_year = gets.chomp

year_born = 2015 - age

if this_year == "no"
	year_born = year_born - 1
end

puts "#{name} was born in #{year_born}."
