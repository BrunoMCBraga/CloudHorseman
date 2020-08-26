#!/usr/bin/env bash

INPUTFILEPATH="/usr/share/cloudhorseman/"
INPUTFILENAME="hosts.txt"
OUTPUTFILEPATH="/usr/share/cloudhorseman/"
OUTPUTFILENAME="output.txt"
STOPFILEPATH="/usr/share/cloudhorseman/"
STOPFILENAME="stop.txt"
NMAPPATH="/usr/bin/nmap"

#So that it never exits
while true; 
do  
    ps aux | grep -ie "nmap" | awk '{print $2}' | xargs kill -9
    until ls "$INPUTFILEPATH" | grep "$INPUTFILENAME";
    do
        sleep 10s
    done

    printf "Input Hosts:\n" > "$OUTPUTFILEPATH$OUTPUTFILENAME"
    cat "$INPUTFILEPATH$INPUTFILENAME" >> "$OUTPUTFILEPATH$OUTPUTFILENAME"
    "$NMAPPATH" -Pn -n -T5 -p80 -iL "$INPUTFILEPATH$INPUTFILENAME" -oN "$OUTPUTFILEPATH$OUTPUTFILENAME"

    until ls "$STOPFILEPATH" | grep "$STOPFILENAME";
    do
        sleep 10s
    done

    rm  "$INPUTFILEPATH$INPUTFILENAME"
    rm "$STOPFILEPATH$STOPFILENAME"

done