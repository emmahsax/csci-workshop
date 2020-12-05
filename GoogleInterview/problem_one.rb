# Write a method that takes a number X and chooses two adjacent digits and replaces
# them with a single digit that equals to the rounded-up average of the two digits.
# It should return the largest of these digits.

def solution(x)
  # make the input number into an array of characters
  string_array = x.to_s.split('')
  ans_array = []
  char_1, char_2, avg, copy = nil

  string_array.each_with_index do |i, index|
    next if index == string_array.length - 1

    # gather the two numbers next to each other and find the average
    char_1 = string_array[index].to_f
    char_2 = string_array[index + 1].to_f
    avg = ((char_1 + char_2) / 2).ceil
    copy = ''

    # put on the numbers before the replacing numbers
    for j in 0..index-1
      copy << string_array[j]
    end

    # put on the average number that you are using as a replacement
    copy << avg.to_s

    # put on the numbers after the replacing numbers
    for j in index+2..string_array.length-1
      copy << string_array[j]
    end

    # save the new number you found to the array of potential answers
    ans_array << copy.to_i
  end

  return ans_array.max
end
