import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
// import UglifyJsPlugin from 'uglifyjs-webpack-plugin'
// import CompressionWebpackPlugin from 'compression-webpack-plugin' // gzip压缩
import {resolve} from 'path'

// const productionGzipExtensions = ['js', 'css']

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [vue(),
        // new UglifyJsPlugin({
        //     uglifyOptions: {
        //         compress: {
        //             drop_debugger: true,
        //             drop_console: true, // 生产环境自动删除console
        //         },
        //         warnings: false,
        //     },
        //     sourceMap: false,
        //     parallel: true, // 使用多线程并行运行来提高构建速度。默认并发运行数：os.cpus().length - 1。
        // }),
        // new CompressionWebpackPlugin({
        //     filename: '[path].gz[query]',
        //     algorithm: 'gzip',
        //     test: new RegExp('\\.(' + productionGzipExtensions.join('|') + ')$'),
        //     threshold: 10240,
        //     minRatio: 0.8
        // }),
    ],
    resolve: {
        alias: [
            {find: '@', replacement: resolve(__dirname, 'src')}
        ],
    },
    // css: {
    //     loadOptions: {
    //         sass: {
    //             data: `@import "@/assets/css/_variable.scss;`
    //         },
    //     }
    // },
    // chainWebpack: config => {
    //     config.module
    //         .rule('images')
    //         .use('image-webpack-loader')
    //         .loader('image-webpack-loader')
    //         .options({
    //             bypassOnDebug: true
    //         })
    //         .end()
    // },
    // build: {
    //     sourcemap: false,
    // },
    // server: {
    //     proxy: {
    //         'api': {
    //             ws: false, // 禁用websocket
    //             target: 'http://47.105.189.195:8848/',
    //             changeOrigin: true,
    //         }
    //     },
    //     port: 12322,
    //     disableHostCheck: true,
    // },
})
