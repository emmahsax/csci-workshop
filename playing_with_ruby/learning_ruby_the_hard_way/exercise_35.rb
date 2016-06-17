## Using if statements, as well as basic loops and functions to make a simple game where someone tries to survive
## mutliple dangers in a maze of rooms.

def start
  puts "You are in a dark room."
  puts "There is a door to your right and left."
  puts "Which one do you take?"

  print "> "
  choice = $stdin.gets.chomp

  if choice == "left"
    bear_room
  elsif choice == "right"
    cthulhu_room
  else
    dead("You stumble around the room until you starve.")
  end
end

def bear_room
  puts "There is a bear here."
  puts "The bear has a bunch of honey."
  puts "The fat bear is in front of another door."
  puts "How are you going to move the bear?"
  bear_moved = false

  while true
    print "> "
    choice = $stdin.gets.chomp

    if choice.include?("honey")
      dead("The bear looks at you then slaps your face off.")
    elsif choice.include?("taunt") && !bear_moved
      puts "The bear has moved from the door. You can go through it now."
      bear_moved = true
    elsif choice.include?("taunt") && bear_moved
      dead("The bear gets pissed off and chews your leg off.")
    elsif choice.include?("door") && bear_moved
      gold_room
    else
      puts "I got no idea what that means."
    end
  end
end

def cthulhu_room
  puts "Here you see the great evil Cthulhu."
  puts "He, it, whatever stares at you and you go insane."
  puts "Do you flee for your life or eat your head?"

  print "> "
  choice = $stdin.gets.chomp

  if choice.include? "flee"
    start
  elsif choice.include? "head"
    dead("Well that was tasty!")
  else
    cthulhu_room
  end
end


def dead(why)
  puts why, "Good job!"
  exit(0)
end

def gold_room
  puts "This room is full of gold.  How much do you take?"

  print "> "
  choice = $stdin.gets.chomp

  if include_num(choice)
    how_much = choice.to_i
  else
    dead("Dude, learn to type a number.")
  end

  if how_much <= 50
    puts "Nice, you're not greedy, you win!"
    exit(0)
  else
    dead("You greedy bastard!")
  end
end

def include_num (choice)
	return choice.include?("0") || choice.include?("1") ||
     	   choice.include?("2") || choice.include?("3") || 
           choice.include?("4") || choice.include?("4") || 
           choice.include?("6") || choice.include?("7") || 
           choice.include?("8") || choice.include?("9")
end

start
