#!/bin/bash

ls --ignore -1dA .* -hint -p | grep -v / | tr '\n' ','