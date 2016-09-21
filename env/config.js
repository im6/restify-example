"use strict";
let env = process.env;

module.exports = {
    isDev: env.isDev === 'true',
    serverPort: env.OPENSHIFT_NODEJS_PORT || 8080,
    serverIp: env.OPENSHIFT_NODEJS_IP || 'localhost',
    requireAuth: env.requireAuth === 'true',
    mongodbUrl: env.OPENSHIFT_MONGODB_DB_URL + 'rest',
    secretKey: env.secretKey
};