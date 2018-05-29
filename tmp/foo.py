#!/usr/bin/env python3
# -*- coding: utf-8 -*-

#########################################################################
# Author: Zhaoting Weng
# Created Time: Thu 03 May 2018 08:05:25 PM CST
# Description:
#########################################################################

import sys

def file_iter(stream, size=1024):
    data = stream.read(size)
    while data:
        yield data
        data = stream.read(size)

def backup(i=None):
    # read from stdin
    if i is None:
        i = sys.stdin
    with open('./output.txt', 'w') as f:
        for data in file_iter(i, size=10):
            f.write(data)

if __name__ == '__main__':
    backup()
