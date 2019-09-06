#!/bin/bash

#########################################################################
# Author: Zhaoting Weng
# Created Time: Fri 06 Sep 2019 10:28:03 AM CST
# Description:
#########################################################################

MYDIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
MYNAME="$(basename "${BASH_SOURCE[0]}")"

if [[ -z $REMOTE_MODE ]]; then
    #shellcheck disable=SC1090
    . "$MYDIR/utils.sh"
fi

main() {
    hello "$@"
}

if [[ -z $REMOTE_MODE ]]; then
    #shellcheck disable=SC1090
    main "$@"
fi
