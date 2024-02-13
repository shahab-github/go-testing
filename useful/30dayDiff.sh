#!/bin/bash

# Get the date 30 days ago in the format YYYY-MM-DD
date_30_days_ago=$(date -u -d '30 days ago' '+%Y-%m-%d')

# Assuming the output is stored in a file named 'output.txt'
input_file="output.txt"

# Process the file and filter rows
awk -v oldDate="$date_30_days_ago" -F, '{
    # Extract the date part from the dateTime field (the 3rd field in the CSV)
    split($3, dateTime, "T");
    gsub(/"/, "", dateTime[1]);  # Remove quotation marks from the date string
    
    if (dateTime[1] >= oldDate) {
        print $0;  # Print the row if the date is within the last 30 days
    }
}' "$input_file"