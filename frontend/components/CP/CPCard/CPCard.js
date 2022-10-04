import React from 'react';
import PropTypes from 'prop-types';
import { Icon, Collapse, Card } from 'antd';
import cs from 'classnames';
import s from './CPCard.scss';

const { Panel } = Collapse;
const { Meta } = Card;

class CPCard extends React.Component {
  static propTypes = {
    children: PropTypes.node,
    title: PropTypes.node,
    className: PropTypes.string,
    type: PropTypes.oneOf([
      'default',
      'success',
      'info',
      'danger',
      'warning',
      'inner',
    ]),
    more: PropTypes.node,
    close: PropTypes.bool,
    toggle: PropTypes.bool,
    noBorder: PropTypes.bool,
    cover: PropTypes.node,
    actions: PropTypes.array,
    metaAvatar: PropTypes.node,
    metaTitle: PropTypes.string,
    metaDesc: PropTypes.node,
  };

  static defaultProps = {
    children: '',
    title: '',
    className: null,
    type: 'default',
    more: null,
    close: false,
    toggle: false,
    noBorder: false,
    cover: '',
    actions: [],
    metaAvatar: null,
    metaTitle: '',
    metaDesc: '',
  };

  constructor(props) {
    super(props);
    this.state = {
      close: false,
    };
  }

  // handle close button functionality
  handleClose = () => {
    this.setState({ close: !this.state.close });
  };

  // render card extra parts (actions: more, toggle button, close button)
  renderCardActions = () => {
    const { title, close, more } = this.props;
    return (
      <span>
        <span className={s.title}>{title}</span>
        <span className={s.extra}>
          {more && <span className={s.more}>{more}</span>}
          {close && (
            <Icon
              type="close"
              className={s.closeIcon}
              onClick={this.handleClose}
            />
          )}
        </span>
      </span>
    );
  };

  render() {
    const {
      className,
      type,
      children,
      toggle,
      title,
      more,
      noBorder,
      metaAvatar,
      metaDesc,
      metaTitle,
      cover,
      actions,
    } = this.props;
    let style;

    switch (type) {
      case 'success': {
        style = s.success;
        break;
      }
      case 'danger': {
        style = s.danger;
        break;
      }
      case 'info': {
        style = s.info;
        break;
      }
      case 'warning': {
        style = s.warning;
        break;
      }
      default: {
        style = s.default;
        break;
      }
    }

    if (this.state.close) {
      return false;
    }

    return toggle ? (
      <Collapse
        defaultActiveKey={['1']}
        className={cs(s.root, className, style)}
      >
        <Panel header={this.renderCardActions()} key="1">
          {children}
        </Panel>
      </Collapse>
    ) : (
      <Card
        title={title}
        extra={more}
        className={cs(s.card, className, style)}
        bordered={!noBorder}
        cover={cover}
        actions={actions}
      >
        {metaAvatar || metaDesc || metaTitle ? (
          <Meta avatar={metaAvatar} title={metaTitle} description={metaDesc} />
        ) : (
          children
        )}
      </Card>
    );
  }
}

export default CPCard;
export const CPCardTest = CPCard;
