import React, { Component, Fragment } from 'react';
import Link from 'next/link';
import Head from 'next/head';
import { Table, Badge, Input, Button, Icon } from 'antd';
import s from './index.scss';
import { request } from '../../store/request';
import { API_BASE_URL } from '../../utils/const';
import { getValidJalaliDateTimeOrNull } from '../../utils/JalaliDate';
import { getFieldTitle } from '../../utils/fields';
import { getCookie } from '../../utils/cookie';
import { parseJwt } from '../../utils/jwt';
import { userRoleReporter, authToken } from '../../const/user';

class Changelog extends Component {
  endpointURL = `${API_BASE_URL}/admin/api/v1/tavanir/case_changelog/`;

  constructor(props) {
    super(props);
    // console.log("props", props)
    this.state = {
      data: [],
      token: getCookie('token') || null,
    };
  }

  componentDidMount() {
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

  fetch = params => {
    console.log(params);
    request.setHeader('Authorization', this.state.token);
    request.get(`${this.endpointURL}`, params).then(response => {
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

  isReporter = () => {
    const cookie = getCookie(authToken);

    if (cookie) {
      const user = parseJwt(cookie);
      if (user.role === userRoleReporter) {
        return true;
      }
    }
    return false;
  };

  getChangeHtmlValue = v => {
    if (v.toString().includes('media/storage/')) {
      return (
        <Fragment>
          <a target="_blank" href={v}>
            {v}
          </a>
        </Fragment>
      );
    }
    return <Fragment>{getFieldTitle(v)}</Fragment>;
  };

  getChangeLogHtml = cl => {
    if (!cl) {
      return;
    }
    const obs = [];
    const photo = /Photo/i;

    cl.map(ob => {
      const path = ob.path[0];
      if (path === 'UpdatedAt' || path === 'ExpertUpdatedAt') {
        return;
      }
      obs.push(
        <Fragment>
          <div>
            {getFieldTitle(path)}:<br />
            <Badge color="red" text={this.getChangeHtmlValue(ob.from)} />
            <br />
            <Badge color="green" text={this.getChangeHtmlValue(ob.to)} />
            <br /> <br />
          </div>
        </Fragment>,
      );
    });
    return obs;
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
    // render: text => text,
  });

  handleReset = clearFilters => {
    clearFilters();
    this.setState({ searchText: '' });
  };

  handleSearch = (selectedKeys, confirm, dataIndex) => {
    confirm();
    this.setState({
      searchText: selectedKeys[0],
      searchedColumn: dataIndex,
    });
  };

  render() {
    const columns = [
      {
        title: 'ID',
        dataIndex: 'id',
        key: 'id',
        sorter: true,
        defaultSortOrder: 'descend',
      },
      this.isReporter()
        ? {}
        : {
            title: 'عملیات',
            render: (_, obj) => {
              return (
                <Link href={`/tavanir/id_${obj.case_id}`}>
                  <a>مشاهده</a>
                </Link>
              );
            },
          },
      {
        title: 'شماره درخواست',
        dataIndex: 'case_id',
        key: 'case_id',
        ...this.getColumnSearchProps('case_id'),
      },
      {
        title: 'تاریخ',
        dataIndex: 'created_at',
        key: 'created_at',
        className: 'input-ltr',
        render: d => {
          return getValidJalaliDateTimeOrNull(d);
        },
        // ...this.getColumnSearchProps('created_at'),
      },
      {
        title: 'شناسه کاربر',
        dataIndex: 'user_id',
        key: 'user_id',
        render: d => {
          if (d === 0) {
            return 'زیاندیده';
          }
          return d;
        },
        ...this.getColumnSearchProps('user_id'),
      },
      {
        title: 'تغییرات',
        dataIndex: 'changelog',
        key: 'changelog',
        render: d => {
          return this.getChangeLogHtml(d);
        },
      },
    ];
    return (
      <div className={s.root}>
        <Head>
          <title>تغییرات</title>
        </Head>
        <div>تغییرات</div>
        <Table
          style={{ whiteSpace: 'pre' }}
          dataSource={this.state.data}
          pagination={this.state.pagination}
          onChange={this.handleTableChange}
          columns={columns}
        />
      </div>
    );
  }
}

export default Changelog;
