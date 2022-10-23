const { writeFileSync, existsSync, readdirSync, readFileSync } = require('fs');

const header = [
    'package utils',
    '',
    'import (',
    '	"github.com/MoonlightPS/Iridium-gidra/gover/gen"',
    ')',
    '',
    'var protoMap = map[int]Message{}',
    '',
    'const ('
];

const body = [
    ')',
    '',
    'func init() {',
];

const footer = [
    '}',
    '',
];

const reg = /\/\/ CmdId: \d+/
const ids = {};

for (const fileName of readdirSync(`${__dirname}/../proto`)) {
    const proto = fileName.slice(0, -6);
    const content = readFileSync(`${__dirname}/../proto/${fileName}`).toString();

    const cmd = reg.exec(content)?.[0];
    if (!cmd) continue;
    ids[cmd.split(' ').pop()] = proto;
}

for (const key in ids) {
    const type = ids[key];
    header.push(`	${type} = ${key}`)
    body.push(`	protoMap[${type}] = &gen.${type}{}`);
}

writeFileSync('packet_gen.go', header.concat(body, footer).join('\n'));