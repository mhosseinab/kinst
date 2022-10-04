import React from 'react';
import PropTypes from 'prop-types';
import { Divider } from 'antd';
import s from './CPDivider.scss';

class CPDivider extends React.Component {
  static propTypes = {
    children: PropTypes.node,
    dashed: PropTypes.bool,
    className: PropTypes.string,
    type: PropTypes.oneOf(['horizontal', 'vertical']),
    orientation: PropTypes.oneOf(['right', 'left', 'center']),
  };

  static defaultProps = {
    children: '',
    dashed: false,
    className: '',
    type: 'horizontal',
    orientation: 'center',
  };

  render() {
    const { dashed, className, type, orientation } = this.props;
    return (
      <Divider
        dashed={dashed}
        className={className}
        type={type}
        orientation={orientation === 'center' ? undefined : orientation}
      >
        {this.props.children}
      </Divider>
    );
  }
}

export default CPDivider;
