class Greeter
  attr_accessor :name

  def initialize(name = "World")
    @name = name
  end

  def say_hello
    puts "Hello #{@name}."
  end

  def say_goodbye
    puts "Goodbye #{@name}, come back soon."
  end
end
