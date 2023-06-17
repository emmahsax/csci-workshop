## Introduction to the basic while loop

def my_while_loop (ending_num)

	i = 0
	numbers = []

	while i <= ending_num do
		puts "Current number i is #{i}"
		numbers.push(i)

		i += 1
		puts "Numbers now: ", numbers
		puts "Next number i is #{i}"
		puts "\n"
	end

	puts "The numbers: "

	numbers.each do |i|
		puts i
	end

	puts "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"
end

puts my_while_loop(2)
puts my_while_loop(8)
