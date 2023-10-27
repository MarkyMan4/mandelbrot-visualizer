#!/usr/bin/env bash

x=400.0
zoom_increase=50
iter=1

# x=$(echo "$x * 1.1" | bc)
# echo $x

while (( $(echo "$x <= 20000000000000000.0" | bc -l) ))
do
    filename=mandelbrot_$iter.png
    ./mviz -z $x -o animation/$filename
    x=$(echo "$x + $zoom_increase" | bc)
    zoom_increase=$(echo "$zoom_increase * 1.1" | bc)
    iter=$(($iter + 1))
done

ffmpeg -framerate 20 -i animation/mandelbrot_%d.png -c:v libx264 -pix_fmt yuv420p -q:v 7 animation/output.mp4
