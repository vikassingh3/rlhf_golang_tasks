#!/bin/bash

set -e

while [[ "$#" -gt 0 ]]; do
    case $1 in
        --api-key) HEROKU_API_KEY="$2"; shift ;;
        *) echo "Unknown parameter passed: $1"; exit 1 ;;
    esac
    shift
done

if [ -z "$HEROKU_API_KEY" ]; then
    echo "Error: HEROKU_API_KEY is required."
    exit 1
fi

echo "Deploying to Heroku..."
heroku deploy --api-key "$HEROKU_API_KEY"

echo "Deployment complete!"
