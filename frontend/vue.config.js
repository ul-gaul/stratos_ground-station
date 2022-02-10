let cssConfig = {};

if (process.env.NODE_ENV === 'production') {
  cssConfig = {
    extract: {
      filename: '[name].css',
      chunkFilename: '[name].css'
    }
  };
}

module.exports = {
  publicPath: process.env.NODE_ENV === 'production'
    ? '/dashboard/dist' : '/',
  productionSourceMap: false,
  css: cssConfig,
  devServer: {
    port: process.env.PORT ?? '8080',
    disableHostCheck: true
  },
  configureWebpack: {
    output: {
      filename: '[name].js',
      hashFunction: 'sha256'
    },
    optimization: {
      splitChunks: false
    },
    performance: {
      hints: false,
    },
  },
  chainWebpack(cfg) {
    const limit = 9999999999999999;
    cfg.module
      .rule('images')
      .test(/\.(png|gif|jpg|webp)(\?.*)?$/i)
      .use('url-loader')
      .loader('url-loader')
      .tap((opts) => Object.assign(opts, { limit }));
    cfg.module
      .rule('fonts')
      .test(/\.(woff2?|eot|ttf|otf|svg)(\?.*)?$/i)
      .use('url-loader')
      .loader('url-loader')
      .options({ limit });
  }
}
