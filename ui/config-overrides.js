const {
    override,
    fixBabelImports,
    addLessLoader
} = require('customize-cra');

module.exports = override(
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
    }),
);