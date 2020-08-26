#!/usr/bin/env bash
STOPFILEPATH="/usr/share/cloudhorseman/"
STOPFILENAME="stop.txt"

#So that it never exits
while true; 
do  
    service ssh start

    until ls "$STOPFILEPATH" | grep "$STOPFILENAME";
    do
        sleep 10s
    done

    service ssh stop
    rm "$STOPFILEPATH$STOPFILENAME"
    break
done