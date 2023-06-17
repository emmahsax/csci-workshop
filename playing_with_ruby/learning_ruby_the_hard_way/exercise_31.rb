## A little quiz thingy that the user will take and try to survive based on their
## input answers.

puts "You enter a dark room with two doors."
puts "1. Go through door 1."
puts "2. Go throug door 2."
print "> "
door = $stdin.gets.chomp

if door == "1"
  puts "There's a giant bear eating a cheese cake. What do you do?"
  puts "1. Take the cake."
  puts "2. Scream at the bear."
  print "> "
  bear = $stdin.gets.chomp

  if bear == "1"
    puts "The bear gets angry and kills you for dinner. And he gets the cake. Nice try!"
  elsif bear == "2"
    puts "The bear gets frightened and in his terror, he tramples you. Nice try!"
  else
    puts "Well, '%s' is probably smarter. And the bear runs away." % bear
  end

elsif door == "2"
  puts "You stare into the endless abyss at Cthulhu's retina."
  puts "1. Blueberries."
  puts "2. Yellow jacket clothespins."
  puts "3. Understanding revolvers yelling melodies."
  print "> "
  insanity = $stdin.gets.chomp

  if insanity == "1" || insanity == "2"
    puts "Your body survives powered by a mind of jello. Nice try!"
  else
    puts "The insanity rots your eyes into a pool of muck. Nice try!"
  end

else
  puts "You stumble around and fall on a knife and die. Nice try!"
end
