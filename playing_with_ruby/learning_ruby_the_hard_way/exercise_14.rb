user_name = ARGV.first

puts "Hi #{user_name}."
puts "I'd like to ask you a few questions."
puts "Do you like me, #{user_name}? "
likes = $stdin.gets.chomp

puts "Where do you live #{user_name}? "
lives = $stdin.gets.chomp

puts "What kind of computer do you have? "
computer = $stdin.gets.chomp

#the fancy lives.split.map(&:capitalize).join(' ') capitalizes the first letter in each word
puts """
Alright, so you said '#{likes.downcase}' about liking me.
You live in #{lives.split.map(&:capitalize).join(' ')}. I'm not sure where that is.
And you have a #{computer.capitalize} computer. Nice.
"""