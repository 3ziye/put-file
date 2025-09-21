#!/bin/bash
curl -u admin:admin123 -F "file=@$(dirname "$0")/test.png" http://localhost:8080/upload