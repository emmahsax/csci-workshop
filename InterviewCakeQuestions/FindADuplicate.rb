# We have an array of integers where:
# 1. the integers are in the range 1..n
# 2. The array has a length of n+1
# Therefore, the array must have at least one integer that appears at least twice.
# But, it may have several duplicates, and each duplicate may appear more than twice. Write
# a function which finds any integer that appears more than once. Optimize for space!
# Gotchas: this can be done in O(n) space and less than O(n^2) time quite easily

def find_a_duplicate(num_array)
  ans_array = []
  num_hash = Hash.new

  num_array.each do |num|
    if num_hash.has_key?(num)
      num_hash[num] = num_hash[num] + 1
    else
      num_hash[num] = 1
    end
  end

  num_hash = num_hash.select { |key, value| value > 1 }
  ans_array = num_hash.keys

  return ans_array
end
