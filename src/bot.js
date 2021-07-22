require('dotenv').config();

const Discord = require('discord.js');
const client = new Discord.Client();

const pointsCommand = require('./commands/points.js');
const pollingCommand = require('./commands/polling.js');

client.on('ready', () => {
  console.log(`Logged in as ${client.user.tag}!`);
});

client.on('message', msg => {

    if (msg.author.bot) return;

    if (!msg.content.startsWith(';')) return;

    const message = msg.content.substr(1);

    if (message.startsWith('points')) {
        pointsCommand(message);
        msg.reply(pollingCommand());
    }

    if (message.startsWith('polling')) {
        msg.reply(pollingCommand());
    }

  if (message === 'toc toc') {
    msg.reply('To no banheiro!');
  }
});

client.login(process.env.BOT_TOKEN);