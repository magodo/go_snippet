#!/usr/bin/env python3
# -*- coding: utf-8 -*-

#########################################################################
# Author: Zhaoting Weng
# Created Time: Wed 23 May 2018 11:54:05 AM CST
# Description:
#########################################################################

def foo(*args):
    for arg in args:
        print(arg)

foo("a", *("b", "c"))
