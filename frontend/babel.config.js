module.exports = {
  presets: ['next/babel'],
  plugins: [
    '@babel/plugin-proposal-optional-chaining',
    [
      'import',
      {
        libraryName: 'antd',
      },
    ],
  ],
};
