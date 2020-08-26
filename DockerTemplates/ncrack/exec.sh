#!/usr/bin/env bash

INPUTFILEPATH="/usr/share/cloudhorseman/"
INPUTFILENAME="hosts.txt"
OUTPUTFILEPATH="/usr/share/cloudhorseman/"
OUTPUTFILENAME="output.txt"
STOPFILEPATH="/usr/share/cloudhorseman/"
STOPFILENAME="stop.txt"
USERNAMESLISTSPATH="/usr/share/cloudhorseman/usernames.txt"
PASSWORDSLISTSPATH="/usr/share/cloudhorseman/passwords.txt"
NCRACKPATH="ncrack"

#So that it never exits
while true; 
do  
    ps aux | grep -ie "ncrack" | awk '{print $2}' | xargs kill -9
    until ls "$INPUTFILEPATH" | grep "$INPUTFILENAME";
    do
        sleep 10s
    done

    echo printf "Hosts Users:\n" > "$OUTPUTFILEPATH$OUTPUTFILENAME"
    cat "$INPUTFILEPATH$INPUTFILENAME" >> "$OUTPUTFILEPATH$OUTPUTFILENAME"
    "$NCRACKPATH" -vvvvvvvvvvvvv -iL "$INPUTFILEPATH$INPUTFILENAME" -U "$USERNAMESLISTSPATH" -P "$PASSWORDSLISTSPATH" >> "$OUTPUTFILEPATH$OUTPUTFILENAME" -p ftp CL=1 &

    until ls "$STOPFILEPATH" | grep "$STOPFILENAME";
    do
        sleep 10s
    done

    rm  "$INPUTFILEPATH$INPUTFILENAME"
    rm "$STOPFILEPATH$STOPFILENAME"

done