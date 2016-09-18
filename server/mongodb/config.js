'use strict';
const globalConfig = require('../../env/config'),
    mongoose = require('mongoose');

let privateFn = {
};

require('./models');

module.exports = {
    initial:function(){
        mongoose.connect(globalConfig.mongodbUrl);

        mongoose.connection.on('connected', function () {
            console.log('Mongoose connected successfully!');
        });

        mongoose.connection.on('error', function(err){
            console.log('Mongoose connection error: ' + err);
        });
    }
};