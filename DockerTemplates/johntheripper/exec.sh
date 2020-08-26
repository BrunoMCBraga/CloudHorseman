#!/usr/bin/env bash

INPUTFILEPATH="/usr/share/cloudhorseman/"
INPUTFILENAME="passwd.txt"
OUTPUTFILEPATH="/usr/share/cloudhorseman/"
OUTPUTFILENAME="output.txt"
STOPFILEPATH="/usr/share/cloudhorseman/"
STOPFILENAME="stop.txt"
WORDLISTSPATH="/usr/share/cloudhorseman/wordlist.txt"
JOHNTHERIPPERPATH="/usr/share/JohnTheRipper/run/john"


#So that it never exits
while true; 
do  
    ps aux | grep -ie "john" | awk '{print $2}' | xargs kill -9
    until ls "$INPUTFILEPATH" | grep "$INPUTFILENAME";
    do
        sleep 10s
    done

    echo printf "Input Hashes:\n" > "$OUTPUTFILEPATH$OUTPUTFILENAME"
    cat "$INPUTFILEPATH$INPUTFILENAME" >> "$OUTPUTFILEPATH$OUTPUTFILENAME"
    "$JOHNTHERIPPERPATH" --wordlist="$WORDLISTSPATH" -format=raw-md5 --rules "$INPUTFILEPATH$INPUTFILENAME" >> "$OUTPUTFILEPATH$OUTPUTFILENAME" &

    until ls "$STOPFILEPATH" | grep "$STOPFILENAME";
    do
        sleep 10s
    done

    rm  "$INPUTFILEPATH$INPUTFILENAME"
    rm "$STOPFILEPATH$STOPFILENAME"

done





    "$NMAPPATH" -Pn -n -T5 -p80 -iL "$INPUTFILEPATH$INPUTFILENAME" -oN "$OUTPUTFILEPATH$TEMPOUTPUTFILENAME"
     mv "$OUTPUTFILEPATH$TEMPOUTPUTFILENAME" "$OUTPUTFILEPATH$OUTPUTFILENAME"

    until ls "$STOPFILEPATH" | grep "$STOPFILENAME";
    do
        sleep 10s
    done

    rm  "$INPUTFILEPATH$INPUTFILENAME"