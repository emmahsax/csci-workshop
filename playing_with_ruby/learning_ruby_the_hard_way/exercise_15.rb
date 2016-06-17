## This experiments with opening files when they're both typed in as arguments 
## or when they're writtin as a standard input
## We always close the files when we're done, just like with streams

txt = open(ARGV[0])

puts "Here's your file #{ARGV[0]}:"
print txt.read

txt.close()

print "\nType the filename again:\n> "
file_again = $stdin.gets.chomp

txt_again = open(file_again)

print txt_again.read

txt_again.close()