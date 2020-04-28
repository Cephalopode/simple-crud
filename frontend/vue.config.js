const serverAddress = 'http://localhost:3000';
module.exports = {
  // Webpack dev server
  devServer: {
    open: process.platform === 'darwin',
    host: '0.0.0.0',
    port: 8080,
    https: false,
    hotOnly: false,
    proxy: {
      '/api': {
        target: serverAddress,
        changeOrigin: true,
        secure: false,
      },
    },
  },
  lintOnSave: false,

  // publicPath: './',
};
