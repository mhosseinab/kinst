import React from 'react';
import Document, { Html, Head, Main, NextScript } from 'next/document';

class MyDocument extends Document {
  render() {
    return (
      <Html>
        <Head>
          <link rel="stylesheet" href="/static/fontiran.css" />
          <link rel="stylesheet" href="/static/css/antd.min.css" />
          <link rel="stylesheet" href="/static/css/antd-rtl.css" />
          <link rel="stylesheet" href="/static/css/main.css" />
          <meta
            name="viewport"
            content="width=device-width, initial-scale=1.0"
          />
          <meta name="theme-color" content="#ff6600" />
          <link rel="apple-touch-icon" href="/static/images/icon_40pt@3x.png" />
          <meta name="apple-mobile-web-app-title" content="baseProject" />
          <meta
            name="apple-mobile-web-app-status-bar-style"
            content="default"
          />
          <meta name="apple-mobile-web-app-capable" content="yes" />
          <meta name="mobile-web-app-capable" content="yes" />
        </Head>
        <body className="rtlLayout">
          <Main />
          <NextScript />
        </body>
      </Html>
    );
  }
}

export default MyDocument;
