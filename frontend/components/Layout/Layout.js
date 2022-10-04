import React from 'react';
// import LoadingBar from 'react-redux-loading-bar';
import PropTypes from 'prop-types';
import s from './Layout.scss';

class Layout extends React.Component {
  render() {
    const { children } = this.props;
    return <div>{children}</div>;
  }
}

export default Layout;

Layout.propTypes = {
  children: PropTypes.node.isRequired,
};
