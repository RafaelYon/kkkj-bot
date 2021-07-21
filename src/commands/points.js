const execute = function(message) {
    // ;points [+1] [@Yon] [Fez Bosta]
    const parameters = message.split(' ', 3);

    const points = parameters[1];
    const author = parameters[2];
    const reason = parameters[3];

    
}

module.exports = execute;