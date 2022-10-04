import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { Spin, Table, Row, Button, Col } from 'antd';
import Head from 'next/head';
import momentJalaali from 'moment-jalaali';
import { connect } from 'react-redux';
import { showLoading, hideLoading } from 'react-redux-loading-bar';
import Router from 'next/router';
import Link from 'next/link';
import { getHomeAction } from '../../store/users/homeAction';
import s from './index.scss';
import CreateForm from './create_form';
import { API_BASE_URL } from '../../utils/const';
import { request } from '../../store/request';
import { getCookie } from '../../utils/cookie';
import { getRoleTitle, userRoleBranch } from '../../utils/role';
import { parseJwt } from '../../utils/jwt';
import { companies } from '../../utils/companis';

const columns = [
  {
    title: 'ID',
    dataIndex: 'id',
    key: 'id',
    sorter: true,
    defaultSortOrder: 'descend',
  },
  {
    title: 'تاریخ ایجاد',
    dataIndex: 'created_at',
    className: 'direction-ltr',
    key: 'created_at',
    sorter: true,
    defaultSortOrder: 'descend',
    render: d => {
      return momentJalaali(d).format('jYYYY/jM/jD H:m');
    },
  },
  {
    title: 'نام کاربری',
    dataIndex: 'username',
    key: 'username',
    className: 'english-font',
  },
  {
    title: 'دسترسی',
    dataIndex: 'role',
    key: 'role',
    render: d => {
      return getRoleTitle(d);
    },
  },
  {
    title: 'وضعیت',
    dataIndex: 'status',
    key: 'status',
    render: d => {
      switch (d) {
        case 'ACTIVE':
          return 'فعال';
          break;
        case 'IN_ACTIVE':
          return 'غیر فعال';
          break;

        default:
          break;
      }
    },
  },
  {
    title: 'استان',
    dataIndex: 'province',
    key: 'province',
    render: d => {
      const provinceList = [];
      const ds = d.split(',');
      for (let index = 0; index < ds.length; index++) {
        provinceList.push(companies[ds[index]]);
      }
      return provinceList.join(', ');
    },
  },
  {
    title: 'توضیحات',
    dataIndex: 'description',
    key: 'description',
  },
  {
    title: 'آخرین ورود',
    dataIndex: 'last_login',
    key: 'last_login',
    className: 'direction-ltr',
    render: d => {
      if (d) {
        return momentJalaali(d).format('jYYYY/jM/jD H:m');
      }
    },
  },
  {
    title: 'عملیات',
    render: (_, obj) => {
      return <Link href={`/users/${obj.id}`}>ویرایش</Link>;
    },
  },
];

class Users extends Component {
  state = {
    data: [],
    pagination: {
      total: 0,
    },
    loading: false,
    modalCreateShow: false,
    modalRecord: {},
    createVisible: false,
    token: getCookie('token') || null,
  };

  componentDidMount() {
    const token = getCookie('token') || null;
    if (!token) {
      Router.push('/login');
      return;
    }

    if (parseJwt(token).role === userRoleBranch) {
      Router.push('/users/me/');
      return;
    }

    if (parseJwt(token).role !== 'ADMIN') {
      Router.push('/access-denied');
      return;
    }

    this.fetch();
  }

  handleTableChange = (pagination, filters, sorter) => {
    const pager = { ...this.state.pagination };
    pager.current = pagination.current;
    this.setState({
      pagination: pager,
    });
    this.fetch({
      page: (pagination.current - 1) * 10,
      sortOrder: sorter.order,
      sortField: sorter.field,
      ...filters,
    });
  };

  fetch = (params = {}) => {
    this.setState({ loading: true });
    request.setHeader('Authorization', this.state.token);
    request.get(`${API_BASE_URL}/admin/api/v1/user/`, params).then(response => {
      if (response.status !== 200) {
        if (response.status === 401) {
          Router.push('/login');
          return;
        }

        return;
      }
      const pagination = { ...this.state.pagination };
      pagination.total = response.data.meta.total_count;
      this.setState({
        data: response.data.objects,
        pagination,
      });
    });
  };

  showModalCreate = () => {
    this.setState({
      createVisible: true,
    });
  };

  saveFormCreate = formCreate => {
    this.formCreate = formCreate;
  };

  handleCreateAccept = e => {
    const { form } = this.formCreate.props;

    form.validateFields((err, values) => {
      if (err) {
        return;
      }

      const u = `${API_BASE_URL}/admin/api/v1/user/`;
      request.post(u, values).then(response => {
        if (response.status !== 200) {
          if (response.data.Number === 1062) {
            form.setFields({
              username: {
                value: values.username,
                errors: [new Error('این نام کاربری قبلا انتخاب شده.')],
              },
            });
          }
          return;
        }

        form.resetFields();
        this.fetch();
        this.setState({ createVisible: false });
      });
    });
  };

  handleCancelAccept = e => {
    this.setState({
      createVisible: false,
    });
  };

  render() {
    const { loading, data } = this.props;
    return loading ? (
      <Spin />
    ) : (
      <div className={s.root}>
        <Head>
          <title>کاربران</title>
        </Head>

        <CreateForm
          wrappedComponentRef={this.saveFormCreate}
          visible={this.state.createVisible}
          onCancel={this.handleCancelAccept}
          onCreate={this.handleCreateAccept}
        />

        <div>
          <h4>کاربران</h4>
          <Row>
            <Col>
              <div style={{ marginBottom: 16, marginTop: 16 }}>
                <Button type="primary" onClick={this.showModalCreate}>
                  جدید
                </Button>
              </div>
            </Col>
          </Row>
          <Table
            dataSource={this.state.data}
            pagination={this.state.pagination}
            onChange={this.handleTableChange}
            columns={columns}
            onRow={(record, rowIndex) => {
              return {
                onClick: event => {
                  this.setState({
                    modalVisible: true,
                    modalRecord: record,
                  });
                },
              };
            }}
          />
        </div>
      </div>
    );
  }
}

Users.getInitialProps = async ({ store }) => {
  await store.dispatch(showLoading());
  await store.dispatch(getHomeAction());
  await store.dispatch(hideLoading());

  return { title: 'Page Title' };
};

Users.propTypes = {
  data: PropTypes.array.isRequired,
  loading: PropTypes.bool.isRequired,
};

const mapState = state => ({
  loading: state.home.loading,
  data: state.home.data,
  error: state.home.error,
});

const mapDispatch = {
  getHomeAction,
};

export default connect(mapState, mapDispatch)(Users);
