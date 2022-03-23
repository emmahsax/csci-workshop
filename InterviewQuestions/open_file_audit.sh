#!/bin/bash

### From March 2022

# We are running an audit of open files on a small list of Linux machines. We need to SSH into each
# server, gather open file data (using the lsof command), and redirect the output (of that command)
# from each machine into a single file on our local machine. Write a script that performs this action.

# The results file on the computer running this is stored in the result_file_name variable
# The IPs of the servers to connect to are stored in the server_ips variable
# The below code assumes that the username to SSH into is stored in the server_username variable

result_file_name="results.txt"
server_ips=("1.2.3.4" "5.6.7.8")
server_username="ubuntu"

# Create the results file
touch $result_file_name

for ip in ${server_ips[@]}; do
  # Create temp file for just this server
  touch $ip.txt

  # SSH into each server and gather open file data
  echo "SSHing into $ip..."
  ssh $server_username@$ip 'bash -s' <<'ENDSSH'
echo $(sudo lsof) > ./output.txt
ENDSSH

  # Send the contents of each file to the results file for just this box
  scp $server_username@$ip:./output.txt $ip.txt

  # Copy the IP and the data to the official file
  echo "$ip" >> $result_file_name
  echo $(cat $ip.txt) >> $result_file_name

  # Delete the temp file for this server
  rm -f $ip.txt

  # Delete the file on the server
  ssh $server_username@$ip 'bash -s' <<'ENDSSH'
rm -f ./output.txt
ENDSSH
done
