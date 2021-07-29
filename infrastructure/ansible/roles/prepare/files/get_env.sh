#!/bin/bash -xe
AWS_DEFAULT_REGION=us-east-2 aws ssm get-parameter --name /kkkj-bot/env | jq '.Parameter.Value' | sed 's/^"//g' | sed 's/"$//g' | sed 's/\\//g' | jq '.[] | .key + "=" + .value' | sed 's/"//g' > /home/bot/kkkj/.env

chown bot:bot /home/bot/kkkj/.env
chmod a-rwx /home/bot/kkkj/.env
chmod ug+r /home/bot/kkkj/.env