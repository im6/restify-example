
'use strict';

const glob = require('glob'),
    path = require('path');

glob.sync('server/resource/mongo/models/*.js').forEach(function(file) {
    console.log(path.resolve(file));
});

