import React from 'react';
import PropTypes from 'prop-types';
import { Avatar, Badge } from 'antd';
import s from './CPAvatar.scss';

class CPAvatar extends React.Component {
  static propTypes = {
    children: PropTypes.node,
    icon: PropTypes.string,
    shape: PropTypes.string,
    size: PropTypes.string,
    src: PropTypes.string,
    badge: PropTypes.string,
    style: PropTypes.object,
  };

  static defaultProps = {
    children: '',
    icon: '',
    shape: 'circle',
    size: 'default',
    src: 'static/images/avatar.png',
    badge: null,
    style: {},
  };

  render() {
    const { icon, shape, size, src, badge, style, children } = this.props;
    if (badge) {
      if (badge === 'dot') {
        return (
          <Badge dot>
            <Avatar icon={icon} shape={shape} size={size} src={src}>
              {children}
            </Avatar>
          </Badge>
        );
      }
      return (
        <Badge count={badge}>
          <Avatar icon={icon} shape={shape} size={size} src={src}>
            {children}
          </Avatar>
        </Badge>
      );
    }
    return (
      <Avatar icon={icon} shape={shape} size={size} src={src} style={style}>
        {children}
      </Avatar>
    );
  }
}

export default CPAvatar;
