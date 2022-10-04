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
import city from './city.json';
import { API_BASE_URL } from '../../utils/const';
import { getValidJalaliOrNull } from '../../utils/JalaliDate';
import CPDatePicker from '../../components/CP/CPDatePicker/CPDatePicker';
import AdvanceQuery from '../../components/AdvanceQuery';
import { parseJwt } from '../../utils/jwt';
import { userRoleReporter, authToken } from '../../const/user';

const { TreeNode } = TreeSelect;

function isReporter() {
  const cookie = getCookie(authToken);

  if (cookie) {
    const user = parseJwt(cookie);
    if (user.role === userRoleReporter) {
      return true;
    }
  }
  return false;
}

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

  getCityColumnSearch = () => ({
    filterDropdown: ({
      setSelectedKeys,
      selectedKeys,
      confirm,
      clearFilters,
    }) => (
      <div style={{ padding: 8 }}>
        <TreeSelect
          showSearch
          treeData={this.state.cities}
          value={selectedKeys[0]}
          onChange={val => {
            setSelectedKeys([val]);
          }}
          filterTreeNode="title"
          style={{ width: 188, marginBottom: 8, display: 'block' }}
          placeholder="یک شهر را انتخاب کنید"
        />
        <Button
          type="primary"
          onClick={() => this.handleSearch(selectedKeys, confirm, 'city')}
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

  handleSubmitAdvanceQuery = () => {
    const { query, filters } = this.state;

    if (query !== '' && query !== undefined) {
      let data = {};
      if (filters.query === undefined) {
        data.query = query;
      }
      if (filters !== '') {
        data = { ...data, ...filters };
      }
      data = this.ToSnakeCase(data, ['sortOrder', 'sortField']);
      this.fetch(data);
    }
  };

  handleReset = clearFilters => {
    clearFilters();
    this.setState({ searchText: '' });
  };

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

  fetch = (params = {}) => {
    this.setState({ loading: true });
    params = this.ToSnakeCase(params, ['sortOrder', 'sortField']);
    request.setHeader('Authorization', this.state.token);
    request
      .get(`${API_BASE_URL}/admin/api/v1/tavanir/cases/`, params)
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

  handleOk = e => {
    this.setState({
      modalRecord: {},
      modalVisible: false,
    });
  };

  handleCancel = e => {
    this.setState({
      modalRecord: {},
      modalVisible: false,
    });
  };

  getExportQuery() {
    let data = {
      token: this.state.token,
    };

    if (this.state.filters) {
      data = Object.assign(data, this.state.filters);
    }

    data = this.ToSnakeCase(data, ['sortOrder', 'sortField']);

    if (this.state.query) {
      data.query = this.state.query;
    }

    return `?${qs.stringify(data, { arrayFormat: 'brackets' })}`;
  }

  columns = [
    {
      title: 'ID',
      dataIndex: 'id',
      key: 'id',
      sorter: true,
      defaultSortOrder: 'descend',
      render: id => {
        return id;
      },
      ...this.getColumnSearchProps('id'),
    },
    {
      title: 'تکراری',
      key: 'isDuplicate',
      render: d => {
        let c = d.IsDuplicate ? (
          <Icon style={{ fontSize: '1em', color: 'red' }} type="close" />
        ) : (
          <div>-</div>
        );
        return c;
      },
    },
    isReporter()
      ? {}
      : {
          title: 'عملیات',
          render: (_, obj) => {
            return (
              <Link href={`/tavanir/${obj.id}`}>
                <a target="_blank">مشاهده</a>
              </Link>
            );
            // return <a target={"_blank"} href={`/request/${obj.id}`}>مشاهده</a>;
          },
        },
    {
      title: 'تاریخ',
      children: [
        {
          title: 'حادثه',
          dataIndex: 'eventDate',
          key: 'eventDate',
          sorter: true,
          defaultSortOrder: 'descend',
          ...this.getDateColumnSearchProps('eventDate'),
          render: d => {
            return d;
          },
        },
        {
          title: 'کارشناسی',
          dataIndex: 'updated_at',
          key: 'updated_at',
          sorter: true,
          // defaultSortOrder: 'descend',
          render: d => {
            return getValidJalaliOrNull(d);
          },
        },
      ],
    },

    {
      title: 'شرکت',
      dataIndex: 'companyId',
      key: 'companyId',
    },
    {
      title: 'کد رهگیری',
      dataIndex: 'trackingId',
      key: 'trackingId',
      ...this.getColumnSearchProps('trackingId'),
    },
    {
      title: 'شناسه قبض',
      dataIndex: 'billId',
      key: 'billId',
      ...this.getColumnSearchProps('billId'),
    },
    {
      title: 'نام',
      dataIndex: 'userName',
      key: 'userName',
    },
    {
      title: 'کد ملی',
      dataIndex: 'nationalId',
      key: 'nationalId',
      ...this.getColumnSearchProps('nationalId'),
    },
    {
      title: 'وضعیت',
      dataIndex: 'status',
      key: 'status',
      render: d => {
        return tavanirStatusChoices[d];
      },
      filters: getstatusChoicesFilter(),
    },
    {
      title: 'کارشناسی',
      dataIndex: 'expert_status',
      key: 'expert_status',
      render: d => {
        return expertStatusChoices[d];
      },
      filters: getExprtStatusChoicesFilter(),
    },
    {
      title: 'استان',
      dataIndex: 'stateName',
      key: 'stateName',
      filters: getProvinceFilter(),
    },
    {
      title: 'شهر',
      dataIndex: 'cityName',
      key: 'cityName',
      ...this.getCityColumnSearch(),
    },
    {
      title: 'خسارت مورد ادعا',
      dataIndex: 'amount',
      key: 'amount',
      defaultSortOrder: 'descend',
      render: val => {
        return val
          ? val.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ',')
          : '--';
      },
      ...this.getColumnSearchProps('amount'),
    },
    {
      title: 'نوع خسارت',
      key: 'compensationTypeId',
      render: d => {
        return <Tag>{tavanirDamageTypes[d.compensationTypeId]}</Tag>;
      },
      filters: getDamageTypeFilter(),
    },
  ];

  render() {
    const { loading, data } = this.props;

    return (
      <div className={s.root}>
        <Head>
          <title>درخواست خسارت</title>
        </Head>
        {/* <div style={{ marginBottom: 16 }} >
          <Stats />
        </div> */}
        <AdvanceQuery
          amount_field_name="amount"
          onChange={this.handleChangeQuery}
          onSubmit={this.handleSubmitAdvanceQuery}
          onReset={this.handleResetQuery}
        />
        {this.state.role === 'ADMIN' && (
          <a
            target="_blank"
            href={`${API_BASE_URL}/admin/api/v1/export/tavanir/request/${this.getExportQuery()}`}
          >
            <Button className={s['excel-btn']} icon="file-excel" type="success">
              خروجی اکسل
            </Button>
          </a>
        )}

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
