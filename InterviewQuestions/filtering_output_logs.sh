### From February 2022

# Write a program that gets the raw output logs that exist at the URL
# https://coderbyte.com/api/challenges/logs/web-logs-raw. The logs there
# are a sample of real web server logs. Each line begins with a date,
# e.g. April 10 11:17:35. Your program should do the following:
#
# Loop through each log item, and find the lines that contain the string
# heroku/router. For each of those, echo the request_id to a new line,
# and then in the fwd key has the value of MASKED, then add a [M] to the
# end of the line with a space before it.
#
# Your final output should look something like the following:
#
# b19a87a1-1bbb-000-00000
# b19a87a1-1bbb-000-111118b
# b2413c-3c67-4180-22222 [M]
# 10f93da3-2753-48a3-33333 [M]

#!/bin/bash
curl -s https://coderbyte.com/api/challenges/logs/web-logs-raw -O > /dev/null

heroku_router="heroku/router"
fwd_masked='fwd="MASKED"'
masked_bool=0

while read line; do
  if [[ $line == *$heroku_router* ]]; then
    if [[ $line == *$fwd_masked* ]]; then
      masked_bool=1
    fi

    # On my mac, I use a literal ' ' in the second sed to do the filtering.
    # On a linux machine, we can use '\s' instead.
    temp_line=$(echo "$line" | sed "s/^.*request_id=//g" | sed "s/ .*$//")

    if [[ $masked_bool == 1 ]]; then
      echo "$temp_line [M]"
      masked_bool=0
    else
      echo "$temp_line"
    fi
  fi
done < web-logs-raw

rm web-logs-raw
