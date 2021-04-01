module.exports = {
    publicPath: process.env.NODE_ENV === 'production'
        ? '/dashboard/dist' : '/',
    productionSourceMap: false,
    devServer: {
        port: process.env.PORT ?? '8080',
    }
}