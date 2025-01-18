#!/bin/bash

if [ $# -eq 0 ]; then
    echo "Usage: $0 <path>"
    exit 1
fi

pdfpath="$1"
txtpath="$2"

echo $1 $2

if [ -f "$pdfpath" ]; then
    echo "Converting File: $pdfpath"
    
    base_name=$(basename "${pdfpath%-statements-8630-.pdf}")
    pdftotext -raw -nopgbrk $pdfpath "$txtpath/$base_name.txt"

elif [ -d "$pdfpath" ]; then
    echo "Converting Dir: $pdfpath"
  
    for file in $pdfpath/*.pdf; do
      base_name=$(basename "${file%-statements-8630-.pdf}")
      pdftotext -raw -nopgbrk $file "$txtpath/$base_name.txt"
    done

else
    echo "Could not find a pdf or dir named: $pdfpath"
fi
