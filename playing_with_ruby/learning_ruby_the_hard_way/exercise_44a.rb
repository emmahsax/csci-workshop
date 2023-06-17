## In this example of inheritance, we define <implicit()> in the parent class, but
## not in the child class. However, it will still print:
##    PARENT implicit()
##    PARENT implicit()
## because the call from son will turn to the parent version of <implicit()>.
## This works because since we put a function in a base class, then all subclasses
## automatically receive those features.

class Parent
  def implicit()
    puts "PARENT implicit()"
  end
end

class Child < Parent
end

dad = Parent.new()
son = Child.new()

dad.implicit()
son.implicit()
