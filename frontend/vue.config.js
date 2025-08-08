const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
    transpileDependencies: true,
    devServer: {
        proxy: {
            '^/backend': {
                target: 'http://127.0.0.1:8080',
                changeOrigin: true
            }
        }
    }
})
