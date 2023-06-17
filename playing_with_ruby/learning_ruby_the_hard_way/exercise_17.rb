puts "I will attempt to copy the contents of #{ARGV[0]} to #{ARGV[1]}."

# open the first file and read it
file1 = open(ARGV[0])
data = file1.read

# find the length in bytes of the data
length_data = data.length

puts "The input file is #{length_data} bytes long."

puts "Does the output file exist? #{File.exist?(ARGV[1])}"
puts "Ready? Hit RETURN to continue, or CTRL-C to abort."

# sees if the user pressed RETURN
$stdin.gets

# open the second file with the write flag, and write the data into it
file2 = open(ARGV[1], 'w')
file2.write(data)

puts "I'm finished! Type 'cat #{ARGV[1]}' to see the copied file."

# close both of the files
file1.close
file2.close

## This is a one line version of the same copying of one file to another
#open(ARGV[1], 'w').write(open(ARGV[0]).read)
