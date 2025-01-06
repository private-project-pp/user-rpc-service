#!/bin/bash


if [ "$1" = "ALL" ]; then
    echo "Semua $1"
elif [ "$1" != "" ]; then
    echo "Spesific $1"
else 
    echo "Kosong cak"
fi