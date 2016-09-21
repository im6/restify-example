

var env = process.env;


module.exports = {
    isDev: env.isDev === 'true',
    serverPort: env.OPENSHIFT_NODEJS_PORT || 3000,
    serverIp: env.OPENSHIFT_NODEJS_IP || 'localhost',
    requireAuth: env.requireAuth === 'true',
    mongodbUrl: env.OPENSHIFT_MONGODB_DB_URL,
    secretKey: env.secretKey
};