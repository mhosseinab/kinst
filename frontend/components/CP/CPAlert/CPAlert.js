import React from 'react';
import PropTypes from 'prop-types';
import { Alert } from 'antd';
import s from './CPAlert.css';

class CPAlert extends React.Component {
  static propTypes = {
    closable: PropTypes.bool,
    type: PropTypes.string,
    showIcon: PropTypes.bool,
    message: PropTypes.node,
    description: PropTypes.node,
    onClose: PropTypes.func,
    className: PropTypes.string,
  };

  static defaultProps = {
    closable: false,
    type: '',
    message: '',
    description: '',
    showIcon: false,
    onClose: () => {},
    className: '',
  };

  render() {
    const {
      closable,
      type,
      showIcon,
      message,
      description,
      onClose,
      className,
    } = this.props;
    return (
      <Alert
        closable={closable}
        onClose={onClose}
        type={type}
        showIcon={showIcon}
        message={message}
        description={description}
        className={className}
      />
    );
  }
}

export default CPAlert;
