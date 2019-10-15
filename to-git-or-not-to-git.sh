#!/bin/bash

curl https://api.github.com/users/timurnurgali | jq ' .["id"] '
