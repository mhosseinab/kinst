import React from 'react';
import App from 'next/app';
import { Provider } from 'react-redux';
import LoadingBar from 'react-redux-loading-bar';
import withRedux from 'next-redux-wrapper';
import configureStore from '../store/store';
import AdminLayout from '../components/AdminLayout';
import Layout from '../components/Layout';
import { serverRequestModifier } from '../store/request';

class MyApp extends App {
  static async getInitialProps({ Component, ctx }) {
    if (ctx.req) {
      // this method set cookie to request
      await serverRequestModifier(ctx.req);
    }
    let pageProps = {};
    if (Component.getInitialProps) {
      pageProps = await Component.getInitialProps(ctx);
    }
    return { pageProps };
  }

  render() {
    const { Component, pageProps, store, router } = this.props;
    if (router.pathname === '/login') {
      return (
        <Provider store={store}>
          <Layout>
            <LoadingBar />
            <Component {...pageProps} />
          </Layout>
        </Provider>
      );
    }
    return (
      <Provider store={store}>
        <AdminLayout>
          <Component {...pageProps} />
        </AdminLayout>
      </Provider>
    );
  }
}

export default withRedux(configureStore)(MyApp);
