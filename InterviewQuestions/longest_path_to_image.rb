### From October 2016

# this solution is not finished :(
# I ran out of time :(
# Given a string of a file system, write a method that will find the longest path
# to any given image. So, if our string given is
# "dir1\n dir11\n dir12\n  picture.jpeg\n  dir121\n  file1.txt\ndir2\n file2.gif",
# the longest path goes from /dir1/dir12/picture.jpeg, but don't include the actual
# image in the number returned...

def solution(s)
  if (!s.include?('.jpeg')) && (!s.include?('.png')) && (!s.include?('.gif'))
    return 0
  end

  ans_array = []
  all_words = s.split("\n")

  all_words.reverse.each_with_index do |word, index|
    ans = ''
    if image?(word)
      puts "This is an image: #{word}"
      for i in 1..all_words.length - (index + 1)
        if directory?(all_words.reverse[i + index]) && (all_words.reverse[i + index][/\A */].size + i >= word[/\A */].size)
          ans = '/' + all_words.reverse[i] + ans
        end
      end
      ans = '/' + ans
    end
  end

  return ans_array.max_by(&:length)

end

def image?(word)
  word.include?('.jpeg') || word.include?('.png') || word.include?('.gif')
end

def directory?(word)
  !word.include?('.')
end
