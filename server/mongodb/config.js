'use strict';
const globalConfig = require('../../env/config'),
    mongoose = require('mongoose'),
    autoIncrement = require('mongoose-auto-increment'),
    glob = require('glob'),
    path = require('path');

mongoose.Promise = global.Promise;

let privateFn = {
    autoIncrementPlugin: null,
    loadModels: ()=> {
        glob.sync('server/mongodb/models/*.js').forEach(function(file) {
            require(path.resolve(file));
        });
    }
};



module.exports = {
    initial:() => {
        //var connection = mongoose.createConnection(globalConfig.mongodbUrl);
        let connection = mongoose.connect(globalConfig.mongodbUrl);

        autoIncrement.initialize(connection);
        privateFn.autoIncrementPlugin = autoIncrement.plugin;

        mongoose.connection.on('connected', function () {
            console.log('Mongoose connected successfully!');
            privateFn.loadModels();
            /*============================   do sth   ========================================*/



            /*============================   do End   ========================================*/

        });

        mongoose.connection.on('error', function(err){
            console.log('Mongoose connection error: ' + err);
        });
    },
    getAutoIncPlugin: ()=>{
        return privateFn.autoIncrementPlugin;
    }
};