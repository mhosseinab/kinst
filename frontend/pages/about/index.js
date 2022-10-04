import React, { Component } from 'react';
import Head from 'next/head';
import s from './index.scss';

class About extends Component {
  render() {
    return (
      <div className={s.root}>
        <Head>
          <title>درباره ما</title>
        </Head>
        <div>درباره ما</div>
      </div>
    );
  }
}

export default About;
