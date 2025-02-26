#!/bin/bash

num_files=30
output_dir="test_files"

mkdir -p "$output_dir"

for ((i=1; i<=num_files; i++))
do
    # Generate a random size between 1KB and 1MB
    file_size=$((RANDOM % 1024 + 1))
    
    file_name="$output_dir/file_$i.txt"
    
    # Create the file with random content
    dd if=/dev/urandom of="$file_name" bs=1K count=$file_size status=none
    
    echo "Created $file_name with size $file_size KB"
done

echo "Generated $num_files random-sized text files in the '$output_dir' directory."