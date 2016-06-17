## In this version of inheritance, we first override the <altered()>
## function in the child class, but then we use <super()> as a way
## of getting the parent version once we call the child version:
##    PARENT altered()
##    CHILD, BEFORE PARENT altered()
##    PARENT altered()
##    CHILD, AFTER PARENT altered()

class Parent
  def altered()
    puts "PARENT altered()"
  end
end

class Child < Parent
  def altered()
    puts "CHILD, BEFORE PARENT altered()"
    super()
    puts "CHILD, AFTER PARENT altered()"
  end

end

dad = Parent.new()
son = Child.new()

dad.altered()
son.altered()