## This is a version of composition that is basically just using
## an object of a class to call methods, instead of direct
## inheritance. This is usually smarter, and less messy:
##    OTHER implicit()
##    CHILD override()
##    CHILD, BEFORE OTHER altered()
##    OTHER altered()
##    CHILD, AFTER OTHER altered()

class Other
  def override()
    puts "OTHER override()"
  end

  def implicit()
    puts "OTHER implicit()"
  end

  def altered()
    puts "OTHER altered()"
  end
end

class Child
  def initialize()
    @other = Other.new()
  end

  def implicit()
    @other.implicit()
  end

  def override()
    puts "CHILD override()"
  end

  def altered()
    puts "CHILD, BEFORE OTHER altered()"
    @other.altered()
    puts "CHILD, AFTER OTHER altered()"
  end
end

son = Child.new()

son.implicit()
son.override()
son.altered()