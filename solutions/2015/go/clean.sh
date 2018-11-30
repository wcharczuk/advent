#! /bin/bash

for each $file in *.go; do 
  dirName = $(echo $file | sed s/.go//g)
  mkdir -p $dirName
  mv $file $dirName/main.go
done


