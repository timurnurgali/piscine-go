#!/bin/bash
find . -type f,d | sed '1d' | wc -l