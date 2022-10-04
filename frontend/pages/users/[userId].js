import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { Spin, Table, Row, Button, Col } from 'antd';
import Head from 'next/head';
import momentJalaali from 'moment-jalaali';
import { connect } from 'react-redux';
import { showLoading, hideLoading } from 'react-redux-loading-bar';
import Router, { useRouter } from 'next/router';
import Link from 'next/link';
import CPMessage from '../../components/CP/CPMessage';
import { getHomeAction } from '../../store/users/homeAction';
import s from './index.scss';
import CreateForm from './create_form';
import { API_BASE_URL } from '../../utils/const';
import { request } from '../../store/request';
import { getCookie } from '../../utils/cookie';
import { getRoleTitle } from '../../utils/role';
import { parseJwt } from '../../utils/jwt';
import EditForm from './edit_form';

// const API_BASE_URL = process.env.MODE === "production" ? "http://tavanir.example.com/gw/" : "http://0.0.0.0:8080/";

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

class UserItemEdit extends Component {
  constructor(props) {
    super(props);
    this.state = {
      id: props.userID,
      data: {},
      token: getCookie('token') || null,
      btnStatusVisible: false,
      userRole: null,
    };
  }

  state = {
    data: [],
    // id: props.itemID,
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

    if (parseJwt(token).role !== 'ADMIN') {
      Router.push('/access-denied');
      return;
    }

    this.fetch();
  }

  fetch = (params = {}) => {
    this.setState({ loading: true });
    request.setHeader('Authorization', this.state.token);
    request
      .get(`${API_BASE_URL}/admin/api/v1/user/${this.state.id}/`, params)
      .then(response => {
        if (response.status !== 200) {
          if (response.status === 401) {
            Router.push('/login');
            return;
          }

          return;
        }
        response.data.province = response.data.province.split(',');
        this.setState({
          data: response.data,
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

      const u = `${API_BASE_URL}/admin/api/v1/user/${this.state.id}/`;
      request.put(u, values).then(response => {
        if (response.status !== 200) {
          CPMessage('خطایی رخ داده است', 'error');
          return;
        }

        CPMessage('تغییرات با موفقیت اعمال شد', 'success');
        Router.push('/users/');
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
          <title>ویرایش کاربر</title>
        </Head>

        <div>
          <Row>
            <Col span={12}></Col>
            <Col span={12}>
              <EditForm
                wrappedComponentRef={this.saveFormCreate}
                visible={this.state.createVisible}
                onCancel={this.handleCancelAccept}
                onCreate={this.handleCreateAccept}
                data={this.state.data}
              />
            </Col>
          </Row>
        </div>
      </div>
    );
  }
}

const Index = () => {
  const router = useRouter();
  const { userId } = router.query;

  return <UserItemEdit userID={userId} />;
};

export default Index;
