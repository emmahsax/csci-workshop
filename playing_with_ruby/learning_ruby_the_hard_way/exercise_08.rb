## Below is a string that will take four variables... we don't know what the
## variables will be yet.
formatter = "%{first} %{second} %{third} %{fourth}"

## This will print in the format that formatter tells us to, but we must
## tell it what we want each of the variables that formatter states to be.
puts formatter % {first: 1, second: 2, third: 3, fourth: 4}
puts formatter % {first: "one", second: "two", third: "three", fourth: "four"}
puts formatter % {first: true, second: false, third: true, fourth: false}
puts formatter % {first: formatter, second: formatter, third: formatter, fourth: formatter}
puts formatter % {first: "a", second: true, third: 1234, fourth: "end of format"}

## This shows an example of a big string concatenation.
puts formatter % {
  first: "I had this thing.",
  second: "That you could type up right.",
  third: "But it didn't sing.",
  fourth: "So I said goodnight."
}
