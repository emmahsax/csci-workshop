# Here's some new strange stuff, remember type it exactly.

# each one will print the whole string in one line
days = "Mon Tue Wed Thu Fri Sat Sun"
# except the \n makes a new line
months = "Jan\nFeb\nMar\nApr\nMay\nJun\nJul\nAug"

puts "Here are the days: #{days}"
puts "Here are the months: #{months}"

## BUT... triple quotes enable us to type multiple lines here,
## and they'll print like multiple lines
puts """
There's something going on here.
With the three double-quotes.
We'll be able to type as much as we like.
Even 4 lines if we want, or 5, or 6.

"""
