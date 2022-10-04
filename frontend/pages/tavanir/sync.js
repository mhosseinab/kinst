import React, { Component } from 'react';
import PropTypes from 'prop-types';
import {
  Spin,
  Table,
  Tag,
  Input,
  Button,
  Icon,
  Highlighter,
  TreeSelect,
} from 'antd';
import Head from 'next/head';
import momentJalaali from 'moment-jalaali';
import { connect } from 'react-redux';
import { showLoading, hideLoading } from 'react-redux-loading-bar';
import Link from 'next/link';
import Router from 'next/router';
import Cookie from 'universal-cookie';
import qs from 'qs';
import { getHomeAction } from '../../store/home/homeAction';
import s from './index.scss';
import { request, serverRequestModifier } from '../../store/request';
import { getCookie } from '../../utils/cookie';
import { snakeCase } from '../../utils/snakeCase';
import {
  tavanirDamageTypes,
  tavanirStatusChoices,
  expertStatusChoices,
} from '../../utils/damage';
import city from '../../utils/city.json';
import { API_BASE_URL } from '../../utils/const';
import { getValidJalaliOrNull } from '../../utils/JalaliDate';
import CPDatePicker from '../../components/CP/CPDatePicker/CPDatePicker';
import AdvanceQuery from '../../components/AdvanceQuery';
import { parseJwt } from '../../utils/jwt';

const { TreeNode } = TreeSelect;

const getProvinceFilter = () => {
  const provinceFilter = [];
  Object.keys(city).map((keyName, i) =>
    provinceFilter.push({
      text: keyName,
      value: keyName,
    }),
  );
  return provinceFilter;
};

const getCityFilter = () => {
  const cityFilter = [];
  const cities = Object.values(city);
  Object.keys(city).map((keyName, i) => {
    let sameName = false;
    const children = cities[i].map((city, n) => {
      if (city === keyName) {
        sameName = true;
      }
      return {
        title: city,
        value: city,
      };
    });
    cityFilter.push({
      title: keyName,
      value: sameName ? `${keyName}a` : keyName,
      children,
      selectable: false,
    });
  });
  return cityFilter;
};

const getDamageTypeFilter = () => {
  const damageTypeFilter = [];
  Object.keys(tavanirDamageTypes).map((keyName, i) =>
    damageTypeFilter.push({
      text: tavanirDamageTypes[keyName],
      value: keyName,
    }),
  );
  return damageTypeFilter;
};

const getstatusChoicesFilter = () => {
  const statusChoicesFilter = [];
  Object.keys(tavanirStatusChoices).map((keyName, i) =>
    statusChoicesFilter.push({
      text: tavanirStatusChoices[keyName],
      value: keyName,
    }),
  );
  return statusChoicesFilter;
};

const getExprtStatusChoicesFilter = () => {
  const exprtStatusChoicesFilter = [];
  Object.keys(expertStatusChoices).map((keyName, i) =>
    exprtStatusChoicesFilter.push({
      text: expertStatusChoices[keyName],
      value: keyName,
    }),
  );
  return exprtStatusChoicesFilter;
};

class Home extends Component {
  state = {
    data: [],
    pagination: {
      total: 0,
    },
    loading: false,
    modalVisible: false,
    modalRecord: {},
    token: getCookie('token') || null,
    cities: getCityFilter(),
    query: '',
    filters: '',
    role: '',
  };

  getColumnSearchProps = dataIndex => ({
    filterDropdown: ({
      setSelectedKeys,
      selectedKeys,
      confirm,
      clearFilters,
    }) => (
      <div style={{ padding: 8 }}>
        <Input
          ref={node => {
            this.searchInput = node;
          }}
          placeholder={`جستجو ${dataIndex}`}
          value={selectedKeys[0]}
          onChange={e =>
            setSelectedKeys(e.target.value ? [e.target.value] : [])
          }
          onPressEnter={() =>
            this.handleSearch(selectedKeys, confirm, dataIndex)
          }
          style={{ width: 188, marginBottom: 8, display: 'block' }}
        />
        <Button
          type="primary"
          onClick={() => this.handleSearch(selectedKeys, confirm, dataIndex)}
          icon="search"
          size="small"
          style={{ width: 90, marginLeft: 8 }}
        >
          جستجو
        </Button>
        <Button
          onClick={() => this.handleReset(clearFilters)}
          size="small"
          style={{ width: 90 }}
        >
          ریست
        </Button>
      </div>
    ),
    filterIcon: filtered => (
      <Icon type="search" style={{ color: filtered ? '#1890ff' : undefined }} />
    ),
    onFilter: (value, record) =>
      record[dataIndex]
        .toString()
        .toLowerCase()
        .includes(value.toLowerCase()),
    onFilterDropdownVisibleChange: visible => {
      if (visible) {
        setTimeout(() => this.searchInput.select());
      }
    },
  });

  getDateColumnSearchProps = dataIndex => ({
    filterDropdown: ({
      setSelectedKeys,
      selectedKeys,
      confirm,
      clearFilters,
    }) => (
      <div style={{ padding: 8 }}>
        <CPDatePicker
          ref={node => {
            this.searchInput = node;
          }}
          placeholder={`جستجو ${dataIndex}`}
          value={selectedKeys[0]}
          onDateChange={e => setSelectedKeys(e ? [e.split('T')[0]] : [])}
          style={{ width: 188, marginBottom: 8, display: 'block' }}
        />
        <br></br>
        <Button
          type="primary"
          onClick={() => this.handleSearch(selectedKeys, confirm, dataIndex)}
          icon="search"
          size="small"
          style={{ width: 90, marginLeft: 8 }}
        >
          جستجو
        </Button>
        <Button
          onClick={() => this.handleReset(clearFilters)}
          size="small"
          style={{ width: 90 }}
        >
          ریست
        </Button>
      </div>
    ),
    filterIcon: filtered => (
      <Icon type="search" style={{ color: filtered ? '#1890ff' : undefined }} />
    ),
    onFilter: (value, record) =>
      record[dataIndex]
        .toString()
        .toLowerCase()
        .includes(value.toLowerCase()),
    onFilterDropdownVisibleChange: visible => {
      if (visible) {
        setTimeout(() => this.searchInput.select());
      }
    },
    render: text => text,
  });

  ToSnakeCase = (obj, excludes = []) => {
    const camelCaseObj = {};
    for (const key of Object.keys(obj)) {
      if (excludes.includes(key)) {
        camelCaseObj[key] = obj[key];
        continue;
      }
      if (Object.prototype.hasOwnProperty.call(obj, key)) {
        if (key == 'id') {
          camelCaseObj['tavanir_id'] = obj[key];
        } else {
          camelCaseObj[snakeCase(key)] = obj[key];
        }
      }
    }

    return camelCaseObj;
  };

  handleSearch = (selectedKeys, confirm, dataIndex) => {
    confirm();
    this.setState({
      searchText: selectedKeys[0],
      searchedColumn: dataIndex,
    });
  };

  handleChangeQuery = query => {
    if (query !== undefined) {
      this.setState({ query });
    }
  };

  handleResetQuery = () => {
    this.setState({ query: '' }, () => {
      this.fetch({ ...this.state.filters });
    });
  };

  handleReset = clearFilters => {
    clearFilters();
    this.setState({ searchText: '' });
  };

  columns = [
    {
      title: 'ID',
      dataIndex: 'ID',
      key: 'iD',
      sorter: true,
      defaultSortOrder: 'descend',
      render: id => {
        return id;
      },
      ...this.getColumnSearchProps('ID'),
    },
    {
      title: 'CaseID',
      dataIndex: 'CaseID',
      key: 'CaseID',
      render: (_, obj) => {
        return (
          <Link href={`/tavanir/${obj.CaseID}`} target="_blank">
            <a target="_blank">{obj.CaseID}</a>
          </Link>
        );
      },
      ...this.getColumnSearchProps('CaseID'),
    },
    {
      title: 'CreatedAt',
      dataIndex: 'CreatedAt',
      key: 'CreatedAt',
      ...this.getColumnSearchProps('CreatedAt'),
    },
    {
      title: 'IsDone',
      dataIndex: 'IsDone',
      key: 'IsDone',
      render: d => {
        return d ? '✔' : '❌';
      },
      ...this.getColumnSearchProps('IsDone'),
    },
    {
      title: 'Success',
      dataIndex: 'Success',
      key: 'Success',
      render: d => {
        return d ? '✔' : '❌';
      },
      ...this.getColumnSearchProps('Success'),
    },
    {
      title: 'NewStatus',
      dataIndex: 'NewStatus',
      key: 'NewStatus',
      ...this.getColumnSearchProps('NewStatus'),
    },
    {
      title: 'Note',
      dataIndex: 'Note',
      key: 'Note',
    },
  ];

  componentDidMount() {
    const token = getCookie('token') || null;
    const jwt = parseJwt(token);

    this.setState({ token, role: jwt.role });
    this.fetch();
  }

  handleTableChange = (pagination, filters, sorter) => {
    const pager = { ...this.state.pagination };
    pager.current = pagination.current;
    const data = {
      page: (pagination.current - 1) * 10,
      sortOrder: sorter.order,
      sortField: sorter.field,
      ...filters,
    };

    if (this.state.query !== '') {
      data.query = this.state.query;
    }
    this.setState({
      pagination: pager,
      filters: data,
    });
    this.fetch(data);
  };

  fetch = (params = {}) => {
    this.setState({ loading: true });
    params = this.ToSnakeCase(params, ['sortOrder', 'sortField']);
    request.setHeader('Authorization', this.state.token);
    request
      .get(`${API_BASE_URL}/admin/api/v1/tavanir/sync_queue/`, params)
      .then(response => {
        this.setState({ loading: false });
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
          loading: false,
          data: response.data.objects,
          pagination,
        });
      });
  };

  render() {
    return (
      <div className={s.root}>
        <Head>
          <title>Sync Queue</title>
        </Head>
        <div>
          <Table
            dataSource={this.state.data}
            pagination={this.state.pagination}
            onChange={this.handleTableChange}
            columns={this.columns}
            footer={() => `تعداد رکوردها: ${this.state.pagination.total}`}
            loading={this.state.loading}
          />
        </div>
      </div>
    );
  }
}

Home.getInitialProps = async context => {
  const { store, req, res, isServer } = context;
  const cookies = isServer ? new Cookie(req.headers.cookie) : new Cookie();
  const token = cookies.get('token');

  await store.dispatch(showLoading());
  await store.dispatch(getHomeAction());
  await store.dispatch(hideLoading());

  if (!token) {
    if (isServer) {
      res.writeHead(302, {
        Location: '/login',
      });
      res.end();
    } else {
      Router.push('/login');
    }
  }

  return { title: 'Page Title' };
};

Home.propTypes = {
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

export default connect(mapState, mapDispatch)(Home);
