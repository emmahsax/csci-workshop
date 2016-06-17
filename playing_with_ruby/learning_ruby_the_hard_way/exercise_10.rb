## All of the below play with different things we can do when it comes to tabbing,
## splitting on lines, etc

tabby_cat = "\tI'm tabbed in."
persian_cat = "I'm split\n  on a line."
backslash_cat = "I'm \\ a \\ cat."

fat_cat = """
I'll do a list:
\t* Cat food
\t* Fishies
\t* Catnip\n\t* Grass
"""

skinny_cat = """
I'm making this one up.
If I'm taking notes, then...
\t* I don't like having to backslash everything with an apostrophe
\t* I can split things
\t  on two lines, but they get messy
\t* Writing things with \\ get weird, as well, and I know all programming languages vary too
\t* I don't like using '''
"""

puts tabby_cat
puts persian_cat
puts backslash_cat
puts fat_cat
puts skinny_cat