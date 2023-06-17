def key_length_helper(input_file_name, key_length, output_file_name)
  key_length = key_length.to_i # change to an integer
  input_file = File.open(input_file_name, "r")
  output_file = File.open(output_file_name, "w")
  input_text = input_file.read.chomp.downcase

  # We'll have one frequency array for each character in the key within another array
  # And each frequency array has room for 26 letter counts, initialized to 0
  arrays = Array.new(key_length) {Array.new(26, 0)}

  index = 0
  input_text.each_byte do |c|
    character_ascii = c - 'a'.ord
    placement = index % key_length

    arrays[placement][character_ascii] += 1
    index += 1
  end

  arrays.each do |array|
    output_file.print array
    output_file.puts
  end

  input_file.close
  output_file.close
end

# using params from command line:
## argv[0] is input_file_name, argv[1] is key_length, argv[2] is output_file_name
## ruby KeyLengthHelper.rb input.txt key_length output.txt
key_length_helper(ARGV[0], ARGV[1], ARGV[2])
