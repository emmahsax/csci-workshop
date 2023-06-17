## This code does the exact same thing as the composition file,
## except it is using modules and mixins instead of classes.
## This way, we don't need to define a new <implicit()> function
## inside of the child class and we don't need an <initialize()>:
##    OTHER implicit()
##    CHILD override()
##    CHILD, BEFORE OTHER altered()
##    OTHER altered()
##    CHILD, AFTER OTHER altered()

module Other
  def override()
    puts "OTHER override()"
  end

  def implicit()
    puts "OTHER implicit()"
  end

  def Other.altered() # must name it this because we call this later...
    puts "OTHER altered()"
  end
end

class Child
  include Other

  def override()
    puts "CHILD override()"
  end

  def altered()
    puts "CHILD, BEFORE OTHER altered()"
    Other.altered() # and we need to differentiate both <altered()>s
    puts "CHILD, AFTER OTHER altered()"
  end
end

son = Child.new()

son.implicit()
son.override()
son.altered()
