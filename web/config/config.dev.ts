// https://umijs.org/config/
import {defineConfig} from 'umi';

export default defineConfig({
  plugins: [
    // https://github.com/zthxxx/react-dev-inspector
    'react-dev-inspector/plugins/umi/react-inspector',
  ],
  // https://github.com/zthxxx/react-dev-inspector#inspector-loader-props
  inspectorConfig: {
    exclude: [],
    babelPlugins: [],
    babelOptions: {},
  },
  define: {
    SERVER_ADDRESS: 'http://localhost:8080',
    SERVER_ADDRESS_WS: 'ws://localhost:8080'
  }
});
