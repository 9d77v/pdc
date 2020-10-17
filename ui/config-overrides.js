const {
    override,
    fixBabelImports,
    addLessLoader,
    overrideDevServer
} = require('customize-cra');

const target = "http://domain.local:8080"
const devServerConfig = () => config => {
    return {
        ...config,
        proxy: {
            '/api': {
                target: target,
                changeOrigin: true,
                secure: false
            },
            '/pdc': {
                target: target,
                changeOrigin: true,
                secure: false
            },
            '/ws': {
                target: target,
                changeOrigin: true,
                secure: false,
                ws: true
            }
        }
    };
};

module.exports = {
    webpack: override(
        fixBabelImports('antd', {
            libraryName: 'antd',
            libraryDirectory: 'es',
            style: true,
        }),
        fixBabelImports('antd-mobile', {
            libraryName: 'antd-mobile',
            libraryDirectory: 'es',
            style: true,
        }),
        addLessLoader({
            lessOptions: {
                javascriptEnabled: true,
                '@primary-color': '#85dbf5'
            },
        })
    ),
    devServer: overrideDevServer(devServerConfig())

};