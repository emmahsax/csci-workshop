15 minute basic interactive Ruby tutorial
=========================================

Here is the link for the entire tutorial. Or, you can do most of it through the `irb` as well. The last two sections use materials the author prepared, however, so I recommend going to the link for those.

* http://tryruby.org/levels/1/challenges/0

## Using the Prompt

First, try typing some math:

```ruby
2 + 6
=> 8
```

Now, try some more advanced mathematics:

```ruby
4 * 10
=> 40

5 - 12
=> -7

40 / 4
=> 10
```

Note that spaces between the operators are not required, but code sure does look nicer with them.

## Say Your Name

Make a string by typing your name surrounded by double quotes:

```ruby
"Emma"
=> "Emma"
```

Now reverse that by typing:

```ruby
"Emma".reverse
=> "ammE"
```

Now find the length of your name:

```ruby
"Emma".length
=> 4
```

Now multiply your name $n$ times to get it to repeat that many times:

```ruby
"Emma" * 5
=> "EmmaEmmaEmmaEmmaEmma"
```

Try reversing something that isn't a string:

```ruby
40.reverse
=> #<NoMethodError: undefined method `reverse' for 40:Fixnum>
```

Okay, so it doesn't like that. You cannot reverse a number in Ruby. However, could you turn 40 into a string to reverse?

```ruby
40.to_s.reverse
=> "04"
```

## Boys are Different From Girls

Numbers are different from strings, just like boys are different from girls. Here are some helpful methods that can be used to go between both:

* to_s: converts values to strings
* to_i: converts values to numbers
* to_a: converts values to arrays

But what are arrays? Let's make an empty one by typing:

```ruby
[]
=> []
```

Since arrays are basically a list that stores information in a sequence, it's like standing in line for popcorn. Add some things to an array, such as:

```ruby
[12, 47, 35]
=> [12, 47, 35]
```

And, the commas are important.

Now we have a list of, let's say, lottery numbers. But what if we want to know what the largest number in the array is?

```ruby
[12, 47, 35].max
=> 47
```

You might've noticed it's pretty irritating to have to retype the array every time. Let's name it something so we can use it:

```ruby
ticket = [12, 47, 35]
=> [12, 47, 35]
```

Let's see if it worked. Type:

```ruby
ticket
=> [12, 47, 35]
```

## Saved, Tucked Away

Now that we know they're saved, we can do stuff with it. Let's say we want to sort it. Ruby has a built in function to help us:

```ruby
ticket.sort!
=> [12, 35, 47]

ticket
=> [12, 35, 47]
```

And ticket is sorted permanently. The reason is because of the exclamation point. Without that, Ruby would just make a new sorted array, and `ticket` would stay unsorted.

## Making a Poem

For the next few exercises, let's save a poem to a variable. This way we can use it multiple times:

```ruby
poem = """
  My toast has flown from my hand
    And my toast has gone to the moon.
    But when I saw it on the television,
    Planting our flag on Halley's comet,
    More still did I want to eat it.
    """
=> "My toast has flown from my hand
  And my toast has gone to the moon.
  But when I saw it on the television,
  Planting our flag on Halley's comet,
  More still did I want to eat it.
  "
```

We can see what this would look like if we print it:

```ruby
print poem
=> "My toast has flown from my hand
	And my toast has gone to the moon.
    But when I saw it on the television,
    Planting our flag on Halley's comet,
    More still did I want to eat it.
   "
```

Let's say we hate the word "toast". We can change that word:

```ruby
poem['toast'] = 'honeydew'
=> "honeydew"

print poem
=> "My honeydew has flown from my hand
	And my toast has gone to the moon.
    But when I saw it on the television,
    Planting our flag on Halley's comet,
    More still did I want to eat it.
   "
```

Try reversing it using the same method as earlier:

```ruby
poem.reverse
=> ".ti tae ot tnaw I did llits eroM
    temoc s'yellaH no galf ruo gnitnalP
    noisivelet eht no ti was I nehw tuB
    noom eht ot enog sah tsaot ym dnA
    dnah ym morf nwolf sah wedyenoh yM
   "
```

I guess that's what you'd expect. But we reversed the characters, too. Let's try reversing just the lines:

```ruby
poem.lines.to_a.reverse
=>  ["More still did I want to eat it.",
     "Planting our flag on Halley's comet,",
     "But when I saw it on the television,",
     "And my toast has gone to the moon.",
     "My toast has flown from my hand"]
```

So what did we do? The `.lines` broke the poem up based on the `\n` character, which is new line. Then we made that into an array, which then we just reversed the elements of the array.

Yay. But, we want one complete poem again, not lines of poems broken down in an array.

```ruby
print poem.lines.to_a.reverse.join
=> "More still did I want to eat it.
    Planting our flag on Halley's comet,
    But when I saw it on the television,
    And my toast has gone to the moon.
    My toast has flown from my hand
   "
```

## A Wee Blank Book

You can make an empty has by typing:

```ruby
book = {}
=> {}
```

A hash is sometimes also called a dictionary. Let's say for our book hash, we'll give books ratings:

```ruby
books["Gravity's Rainbow"] = :splendid
=> :splendid
```

Notice that this rating is not a string. It has a colon in front of it, meaning that it is a symbol. In terms of memory, symbols are cheaper than strings; a computer only stores symbols in memory once, versus storing a string multiple times.

Add some more books:

```ruby
books["Fifty Shades of Grey"] = :abysmal
=> :abysmal

books["Harry Potter and the Deathly Hallows"] = :splendid
=> :splendid

books["The Lion, the Witch, and the Wardrobe"] = :splendid
=> :splendid

books["The Mists of Avalon"] = :mediocre
=> :mediocre

books
=> {"Gravity's Rainbow"=>:splendid,
	"Fifty Shades of Grey"=>:abysmal,
	"Harry Poter and the Deathly Hallows"=>:splendid,
	"The Lion, the Witch, and the Wardrobe"=>:splendid,
	"The Mists of Avalon"=>:mediocre
   }

books.length
=> 5
```

What if you forgot what rating you gave a book? That's okay!

```ruby
books["Gravity's Rainbow"]
=> :splendid
```

Note that hashes don't keep the elements in order. Their job is simply to pair keys and values together.

You can view just the book titles that you've reviewed, or just the ratings given:

```ruby
books.keys
=> ["Gravity's Rainbow", "Fifty Shades of Grey",
    "Harry Poter and the Deathly Hallows", "The Lion, the Witch, and the Wardrobe",
    "The Mists of Avalon"
   ]

books.values
=> [:splendid, :abysmal, :splendid, :splendid, :mediocre]
```

Let's say you want to keep track of how many of each rating you are giving out. You can make a new hash that will keep count:

```ruby
ratings = Hash.new(0)
=> {}
```

This made a new hash, `ratings`, and we'll make each rating name a key, and then a counting number the value:

```ruby
books.values.each {|rate| ratings[rate] += 1}
=> [:splendid, :abysmal, :splendid, :splendid, :mediocre]
```

Let's look at that loop. So for every value in `books`, go through and add that value, `rate` to `ratings` as a key, and then increment the value by one.

Type this to see the full counts:

```ruby
ratings
=> {:splendid=>3, :abysmal=>1, :mediocre=>1}
```

## A Tally

Basically, what we did to make `ratings` and insert information was a block. Let's try another:

```ruby
5.times {print "Hello World!"}
=> "Hello World!Hello World!Hello World!Hello World!Hello World!"
```

## Dealing with Files

Let's say that we want to evaluate some files that we have saved on our computer. For this example, we're going to use a theoretical directory. Let's see what that looks like:

```ruby
Dir.entries "/"
=> [".", "..", "Home", "Libraries", "MouseHole", "Programs", "Tutorials", "comics.txt"]
```

Now there's a lot of stuff here. This is everything in the current directory. Let's say we just wanted things that were `txt` files:

```ruby
Dir["/*.txt"]
=> ["/comics.txt"]
```

There we go! Now let's see what's inside:

```ruby
print File.read("/comics.txt")
=> "Achewood: http://achewood.com/
    Dinosaur Comics: http://qwantz.com/
    Perry Bible Fellowship: http://cheston.com/pbf/archive.html
    Get Your War On: http://mnftiu.cc/
   "
```

But what if you want your own copy of the file to modify? We can do that!

```ruby
FileUtils.cp("/comics.txt", "/Home/comics.txt")
=> nil

Dir["/Home/*.txt"]
=> ["/Home/comics.txt"]
```

## Modifying a File

Let's modify our comics file, and add a comic. This way we can reference back to it.

First, we have to open the file. The `"a"` means that whatever we add will append onto the end of the current file. We're typing it in a block format, so when we use `do`, we also need an `end` at the end.

```ruby
File.open("/Home/comics.txt", "a") do |f|
	f << "Cat and Girl: http://catandgirl.com/"
end
=> #<File:/Home/comics.txt (closed)>
```

The `<<` thing means write the following to the previous. Let's see what it looks like now:

```ruby
print File.read("/Home/comics.txt")
=> "Achewood: http://achewood.com/
    Dinosaur Comics: http://qwantz.com/
    Perry Bible Fellowship: http://cheston.com/pbf/archive.html
    Get Your War On: http://mnftiu.cc/
    Cat and Girl: http://catandgirl.com/
   "
```

And if we wanted to see the precise time we edited the file last:

```ruby
File.mtime("/Home/comics.txt")
=> 2015-05-17 17:10:17 UTC
```

But let's say now that we have comics, we want to use it as an array. That sounds like it could be useful, right?

Let's define a method (this is new!) that will help us.

```ruby
def load_comics (path)
	comics = {}
	File.foreach(path) do |line|
		name, url = line.split(": ")
		comics[name] = url.strip
	end
	comics
end
=> {"Achewood"=>"http://achewoods.com/",
	"Dinosaur Comics"=>"http://qwantz.com/",
	"Perry Bible Fellowship"=>"http://cheston.com/pbf/archive.html",
	"Get Your War On"=>"http://mnftiu.cc/"
   }
```

## Popups

Unfortunately, I don't have the capabilities to mimic the tutorial for the next part. The link is `http://tryruby.org/levels/6/challenges/6` if you want to continue the popup section.
