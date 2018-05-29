#!/bin/bash

#########################################################################
# Author: Zhaoting Weng
# Created Time: Mon 28 May 2018 05:13:43 PM CST
# Description:
#########################################################################

tee >(wc -c >/dev/null) | tee >(md5sum - >/dev/null) | nc localhost 12345
