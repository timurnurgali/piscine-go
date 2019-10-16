#!/bin/bash
find . -not -path '*/\.*' -type f,d | sed '1d' | wc -l