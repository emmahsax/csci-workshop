### This is based on a Ruby tutorial that is from https://www.ruby-lang.org/en/documentation/quickstart/

This is assuming that `irb` is typed into a terminal:

## Working with basic Ruby

`puts` is the basic command that is used to print in Ruby.

Ruby can also be used as a basic calculator with operators such as +, -, \*, and / (and ** is "to the power of"):

```ruby
Math.sqrt(9)
=> 3.0
```

Notice that the results is not just 3 because `sqrt` returns a floating-point number. The dot signifies the receiver of a message.

We can assign the results to a variable:

```ruby
a = 3**2
=> 9
b = 4**2
=> 16
Math.sqrt(a+b)
=> 5.0
```

## Defining simple methods

We can define methods using `def`:

```ruby
def h
  puts "Hello World!"
end
```

The new method is named h, and end is signifying the end of the method. To call it, just type h or h().

Our methods can also take in a parameter:

```ruby
def h1(name)
  puts "Hello #{name}!"
end
```

This is called by typing `h("Emma")` or `h "Emma"`. But let's say that we still want `World` to show up in the message if no name is provided. We can do this by the following pieces of code:

```ruby
def h2(name = "World")
  puts "Hello #{name}!"
end
```

And then this can be called with or without a name!

```ruby
h
=> Hello World!

h "Emma"
=> Hello Emma!
```

## Playing around with classes

The first class will be inside the file `greeter.rb`. When we use that, we can create new greeters by using:

```ruby
g = Greeter.new("Emma")
=> #<Greeter:0x007f8ef10847b0 @name="Emma">

g.say_hello
Hello Emma.
=> nil

g.say_goodbye
Goodbye Emma, come back soon.
=> nil
```

You can use `Greeter.instance_methods` to see all of the instance methods that are "hidden" within the Greeter class. If we want to ignore all of the methods defined by ancestor classes, we can use `Greeter.instance_methods(false)` to get a much cleaner output:

`=> [:initialize, :say_hello, :say_goodbye]`

We can change classes once they're already in use. By adding a simple line into the Greeter class, `attr_accessor :name`, we can make our class respond to a `name`, by changing the name of an object (using `name=`), printing just the name (using `name`), or using the class as it was intended before.

A more advanced version of the greeter, megaGreeter, can be found in the `megaGreeter.rb` file. Using loops and simple checks, such as `@names.repspond_to?("each")`, the megaGreeter is capable of greeting multiple people at once.
