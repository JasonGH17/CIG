const fs = require('fs-extra');
const path = require('path');

console.time('Postbuild');

Promise.all([
	fs.ensureDir(path.join(__dirname, '../public')),
	fs.copy('./build', path.join(__dirname, '../public')),
]);

console.timeEnd('Postbuild');
