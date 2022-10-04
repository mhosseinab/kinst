import React, { Component } from 'react';
import Router from 'next/router';
import { connect } from 'react-redux';
import Cookies from 'universal-cookie';

const cookies = new Cookies();

class Logout extends Component {
  constructor(props) {
    super(props);
  }

  componentDidMount() {
    cookies.remove('token')
    Router.push('/login')
  }

  render() {
    return (
      <div>
        logging out...
    </div>
    )
  }
}

export default connect()(Logout);
