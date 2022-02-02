### From October 2016

# A zero-indexed array A consisting of N integers is given.
# An equilibrium index of this array is any integer P such that
# 0 <= P < 0 and the sum of elements of lower indices is equal
# to the sum of elements of higher indices. Write a method that,
# given a zero-indexed array A consisting of N integers,
# returns any of its equilibrium indices. It should return -1 if
# no equilibrium index exists.

def solution(a)
  ans = -1
  left_index, right_index, right_sum, left_sum = nil

  a.each_with_index do |element, index|
    left_sum = 0
    right_sum = 0

    for left_index in 0..index-1
      left_sum += a[left_index]
    end

    for right_index in index+1..a.length-1
      right_sum += a[right_index]
    end

    if left_sum == right_sum
      ans = index
    end
  end

  return ans
end
