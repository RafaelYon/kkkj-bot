const execute = () => {
    let template = "```[\n{text}]```";
    let text = "";
    //TODO puxar da persistencia
    const data = [
        {
            dev: "Yon",
            ammout: 10
        },
        {
            dev: "Lucas",
            ammout: -10
        },
        {
            dev: "Ta mole",
            ammout: 5
        },
        {
            dev: "Daniel",
            ammout: 5
        },
    ];

    data.forEach(el => {
        text += el.dev.padEnd(15, ' ') + ' => ' + el.ammout + '\n';
    });

    return template.replace('{text}', text);
}

module.exports = execute;