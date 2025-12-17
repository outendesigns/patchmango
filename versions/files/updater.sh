#!/bin/bash
echo "Running update script"
CURRENT_VERSION=$(cat /etc/version | cut -d'=' -f2)
echo "$CURRENT_VERSION"