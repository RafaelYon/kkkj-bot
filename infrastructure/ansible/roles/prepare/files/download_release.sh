#!/bin/bash -xe
wget $(curl https://api.github.com/repos/RafaelYon/kkkj-bot/releases/latest | jq '.assets[0].browser_download_url' | sed 's/"//g') -O /home/bot/kkkj/bot

chown bot:bot /home/bot/kkkj/bot
chmod u+x /home/bot/kkkj/bot