#! /usr/bin/env bash

# DANGEROUS: by runing this file you are putting on the python servers to run as background process
# so for clean up make sure to kill this manually. 
# Todo: for clean up
# ps aux | grep 'host.py'
# kill PID
# WARN: if not done, ports will be occupied and might not run further.
python3 ./testHost/host.py "server-1" "5000" &
python3 ./testHost/host.py "server-2" "6001" &
python3 ./testHost/host.py "server-3" "8000" &
