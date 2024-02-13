#!/bin/bash
set -x  # Enable debugging

# Replace 'your_command_here' with your actual command
cat Report.csv | awk -v oldDate="$(date -u -d '30 days ago' '+%Y-%m-%d')" -F, '{
    split($3, dateTime, "T");
    gsub(/"/, "", dateTime[1]);
    if (dateTime[1] >= oldDate) {
        print $0;
    }
}'