

var env = process.env;


module.exports = {
    isDev: env.isDev === 'true',
    serverPort: env.NODE_PORT || 8080,
    serverIp: env.NODE_IP || '127.0.0.1',
    requireAuth: true,
    mongodbUrl: env.mongodbUrl,
};