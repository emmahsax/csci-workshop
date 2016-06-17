## This is the fourth exercise from:
##  http://learnrubythehardway.org/book/ex4.html
## It's playing around with variables and doing math using the variables.

cars = 100 #total numbers of cars available
space_in_a_car = 4.0 #amount of seats in a car
drivers = 30 #people driving
passengers = 90 #passengers not driving
cars_not_driven = cars - drivers #cars not with a driver yet
cars_driven = drivers #cars with a driver
carpool_capacity = cars_driven * space_in_a_car #how many spaces in cars that have drivers
average_passengers_per_car = passengers / cars_driven #how many passengers per car (average)


puts "There are #{cars} cars available."
puts "There are only #{drivers} drivers available."
puts "There will be #{cars_not_driven} empty cars today."
puts "We can transport #{carpool_capacity} people today."
puts "We have #{passengers} to carpool today."
puts "We need to put about #{average_passengers_per_car} in each car."
