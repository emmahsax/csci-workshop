## In this example of inheritance, we define the same function
## <override()> in both the parent class and the child class.
## This way, either parents or children (subclasses or base classes)
## can call override, and they'll each use their own version:
##    PARENT override()
##    CHILD override()

class Parent
  def override()
    puts "PARENT override()"
  end
end

class Child < Parent
  def override()
    puts "CHILD override()"
  end
end

dad = Parent.new()
son = Child.new()

dad.override()
son.override()