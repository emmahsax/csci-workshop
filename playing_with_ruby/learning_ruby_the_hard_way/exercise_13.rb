## This script will take in 6 variables

## And will unpack the first three and print them out in strings
first, second, third = ARGV

puts "Your first variable is: #{first}"
puts "Your second variable is: #{second}"
puts "Your third variable is: #{third}"

## Another way to do this is to just grab directly since ARGV is an array of the args

puts "Your fourth variable is: #{ARGV[3]}"
puts "Your fifth variable is: #{ARGV[4]}"
puts "Your sixth variable is: #{ARGV[5]}"

## This way we don't have to unpack, as we do in the fourth line of the file