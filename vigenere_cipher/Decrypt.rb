def decypt_with_key(input_file_name, key, output_file_name)
    input_file = File.open(input_file_name, "r")
    output_file = File.open(output_file_name, "w")
    input_text = input_file.read.chomp.downcase
    key = key.downcase

    index = 0
    input_text.each_byte do |c|
        character_ascii = c - 'a'.ord
        key_ascii = (key[index % key.length].ord - 'a'.ord)
        new_character_ascii = (character_ascii - key_ascii) % 26
        new_character = new_character_ascii + 'a'.ord

        output_file.print new_character.chr
        index += 1
    end

    output_file.puts
    input_file.close
    output_file.close
end

# using params from command line:
## argv[0] is input_file_name, argv[1] is key, argv[2] is output_file_name:
## ruby Decrypt.rb input.txt key output.txt
decypt_with_key(ARGV[0], ARGV[1], ARGV[2])
