import React from 'react';
import PropTypes from 'prop-types';
import { Icon, Layout } from 'antd';
import LoadingBar from 'react-redux-loading-bar';
import s from './AdminLayout.scss';
import Sidebar from '../Sidebar';

const { Header, Content, Sider } = Layout;

class AdminLayout extends React.Component {
  state = {
    collapsed: false,
  };

  toggle = () => {
    this.setState(prevState => ({
      collapsed: !prevState.collapsed,
    }));
  };

  render() {
    const { collapsed } = this.state;
    const { children } = this.props;
    return (
      <Layout>
        <LoadingBar className={s.loadingBar} />
        <Sider
          trigger={null}
          collapsible
          collapsed={collapsed}
          collapsedWidth={0}
          className={collapsed ? s.triggerStyle : ''}
        >
          <div className="logo" />
          <Sidebar collapsed={collapsed} />
        </Sider>
        <Layout>
          <Header className={s.headerBox}>
            <Icon
              className={s.trigger}
              type={collapsed ? 'menu-unfold' : 'menu-fold'}
              onClick={this.toggle}
            />
          </Header>
          <Content className={s.content}>
            <div>{children}</div>
          </Content>
        </Layout>
      </Layout>
    );
  }
}

AdminLayout.propTypes = {
  children: PropTypes.node.isRequired,
};

export default AdminLayout;
