require('dotenv').config();

const Discord = require('discord.js');
const client = new Discord.Client();

const pointsCommand = require('./commands/points.js');

client.on('ready', () => {
  console.log(`Logged in as ${client.user.tag}!`);
});

client.on('message', msg => {

    if (msg.author.bot) return;

    if (!msg.content.startsWith(';')) return;

    const message = msg.content.substr(1);

    if (message.startsWith('points')) {
        
    }

  if (msg.content === 'ping') {
    msg.reply('Pong!');
  }
});

client.login(process.env.BOT_TOKEN);