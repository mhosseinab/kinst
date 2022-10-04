import React from 'react';
import PropTypes from 'prop-types';
import { Icon, Menu, Divider } from 'antd';
import Link from './ActiveLink';
import s from './Sidebar.scss';
import { getCookie } from '../../utils/cookie';
import { parseJwt } from '../../utils/jwt';
import {
  userRoleAdmin,
  userRoleTavanir,
  userRoleBranch,
  authToken,
} from '../../const/user';

const { Item, SubMenu } = Menu;

function getUsername() {
  const cookie = getCookie(authToken);

  if (cookie) {
    const user = parseJwt(cookie);
    return user.username;
  }
  return 'مدیریت';
}

function checkRole(roles) {
  const cookie = getCookie(authToken);

  if (cookie) {
    const user = parseJwt(cookie);
    console.log(user);
    if (roles.includes(user.role)) {
      return true;
    }
  }
  return false;
}

const Sidebar = ({ collapsed }) => {
  const items = [
    {
      title: 'خسارت های 1398',
      key: 'sub1',
      icon: 'profile',
      children: [
        { title: 'درخواست ها', icon: 'mail', href: '/' },
        {
          title: 'تغییرات',
          icon: 'diff',
          href: '/changelog',
          roles: [userRoleAdmin, userRoleTavanir, userRoleBranch],
        },
        {
          title: 'آمار',
          icon: 'bar-chart',
          href: '/stats',
          roles: [userRoleAdmin, userRoleTavanir, userRoleBranch],
        },
        { title: 'گزارشات', icon: 'snippets', href: '/report' },
      ],
    },
    { divider: true },
    {
      title: 'خسارت های 1399',
      key: 'sub2',
      icon: 'profile',
      children: [
        { title: 'درخواست ها', icon: 'mail', href: '/tavanir' },
        {
          title: 'تغییرات',
          icon: 'diff',
          href: '/tavanir/changelog',
          roles: [userRoleAdmin, userRoleTavanir, userRoleBranch],
        },
        { title: 'گزارشات', icon: 'snippets', href: '/tavanir/report' },
      ],
    },
    { divider: true },
    { title: 'کاربران', icon: 'user', href: '/users/', roles: [userRoleAdmin] },
    { title: 'ویرایش رمزعبور', icon: 'lock', href: '/changepass/' },
    {
      title: 'خروج',
      icon: 'logout',
      href: '/logout',
    },
  ];

  return (
    <div className={s.sideBar}>
      <div className={s.user}>
        <span className={s.userInfo}>
          <b>{getUsername()}</b>
        </span>
      </div>
      <Menu theme="dark" mode="inline" defaultOpenKeys={['sub1', 'sub2']}>
        {items.map(item => {
          if (item.divider) {
            return <Divider />;
          }
          if (item.children) {
            if (item.roles) {
              return checkRole(item.roles) ? (
                <SubMenu
                  key={item.key}
                  title={
                    <span>
                      <Icon type={item.icon} />
                      <span>{item.title}</span>
                    </span>
                  }
                >
                  {item.children.map(child => {
                    if (child.roles) {
                      return (
                        checkRole(child.roles) && (
                          <Item>
                            <Link
                              activeClassName={s.activeLink}
                              href={child.href}
                            >
                              <a>
                                <Icon type={child.icon} />
                                <span>{child.title}</span>
                              </a>
                            </Link>
                          </Item>
                        )
                      );
                    }
                    return (
                      <Item>
                        <Link activeClassName={s.activeLink} href={child.href}>
                          <a>
                            <Icon type={child.icon} />
                            <span>{child.title}</span>
                          </a>
                        </Link>
                      </Item>
                    );
                  })}
                </SubMenu>
              ) : (
                ''
              );
            }
            return (
              <SubMenu
                key={item.key}
                title={
                  <span>
                    <Icon type={item.icon} />
                    <span>{item.title}</span>
                  </span>
                }
              >
                {item.children.map(child => {
                  if (child.roles) {
                    return (
                      checkRole(child.roles) && (
                        <Item>
                          <Link
                            activeClassName={s.activeLink}
                            href={child.href}
                          >
                            <a>
                              <Icon type={child.icon} />
                              <span>{child.title}</span>
                            </a>
                          </Link>
                        </Item>
                      )
                    );
                  }
                  return (
                    <Item>
                      <Link activeClassName={s.activeLink} href={child.href}>
                        <a>
                          <Icon type={child.icon} />
                          <span>{child.title}</span>
                        </a>
                      </Link>
                    </Item>
                  );
                })}
              </SubMenu>
            );
          }
          if (item.roles) {
            return checkRole(item.roles) ? (
              <Item>
                <Link activeClassName={s.activeLink} href={item.href}>
                  <a>
                    <Icon type={item.icon} />
                    <span>{item.title}</span>
                  </a>
                </Link>
              </Item>
            ) : null;
          }
          return (
            <Item>
              <Link activeClassName={s.activeLink} href={item.href}>
                <a>
                  <Icon type={item.icon} />
                  <span>{item.title}</span>
                </a>
              </Link>
            </Item>
          );
        })}
      </Menu>
    </div>
  );
};

Sidebar.propTypes = {
  collapsed: PropTypes.bool.isRequired,
};

export default Sidebar;
