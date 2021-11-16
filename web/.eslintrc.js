module.exports = {
  extends: [require.resolve('@umijs/fabric/dist/eslint')],
  globals: {
    ANT_DESIGN_PRO_ONLY_DO_NOT_USE_IN_YOUR_PRODUCTION: true,
    page: true,
    REACT_APP_ENV: true,

    // 自定义变量
    SERVER_ADDRESS: true,
    SERVER_ADDRESS_WS: true,
  },
};
