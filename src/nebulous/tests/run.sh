#!/usr/bin/env bash

for test in $(find /.unit-tests -name "test*.sh" -print | sort); do
    lvl="${test##*/test}"
    lvl="${lvl%.sh}"

    part="${lvl#*.}"
    lvl="${lvl%%.*}"
    [[ "$lvl" != "$part" ]] || unset part

    if "$test" 2>&1 | grep -E -i -q "success.+flag$lvl"; then
        echo -e "\e[32m[+] level$lvl$part pass\e[0m"
    else
        echo -e "\e[31m[!] level$lvl$part fail\e[0m"
    fi
done; unset lvl part test
