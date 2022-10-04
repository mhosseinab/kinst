import React, { Component } from 'react';
import Link from 'next/link';
import { Table, Tag, Spin } from 'antd';
import momentJalaali from 'moment-jalaali';
import { request } from '../../store/request';
import { getCookie } from '../../utils/cookie';
import { API_BASE_URL } from '../../utils/const';
import {
  tavanirDamageTypes,
  tavanirStatusChoices,
  expertStatusChoices,
} from '../../utils/damage';

class SimilarRequest extends Component {
  endpointURL = null;

  columns = [
    {
      title: 'ID',
      dataIndex: 'id',
      key: 'id',
      render: (_, obj) => {
        return (
          <Link href={`/tavanir/${obj.id}`}>
            <a target="_blank">{obj.id}</a>
          </Link>
        );
      },
    },
    {
      title: 'تاریخ حادثه',
      dataIndex: 'eventDate',
      key: 'eventDate',
    },
    {
      title: 'تاریخ ثبت',
      dataIndex: 'created_at',
      key: 'created_at',
      render: created_at => momentJalaali(created_at).format('jYYYY/jM/jD'),
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
    },
    {
      title: 'نام',
      dataIndex: 'userName',
      key: 'userName',
    },
    {
      title: 'وضعیت',
      dataIndex: 'status',
      key: 'status',
      render: d => {
        return tavanirStatusChoices[d];
      },
    },
    {
      title: 'کارشناسی',
      dataIndex: 'expert_status',
      key: 'expert_status',
      render: d => {
        return expertStatusChoices[d];
      },
    },
    {
      title: 'استان',
      dataIndex: 'stateName',
      key: 'stateName',
    },
    {
      title: 'شهر',
      dataIndex: 'cityName',
      key: 'cityName',
    },
    {
      title: 'خسارت مورد ادعا',
      dataIndex: 'amount',
      key: 'amount',
      render: val => {
        return val
          ? val.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ',')
          : '--';
      },
    },
    {
      title: 'نوع خسارت',
      key: 'compensationTypeId',
      render: d => {
        return <Tag>{tavanirDamageTypes[d.compensationTypeId]}</Tag>;
      },
    },
  ];

  constructor(props) {
    super(props);
    // console.log("props", props)
    this.state = {
      data: [],
      case_id: props.case_id || null,
      token: getCookie('token') || null,
    };
  }

  componentDidMount() {
    this.endpointURL = `${API_BASE_URL}/admin/api/v1/tavanir/similar/case/${this.state.case_id}/`;
    this.fetch();
  }

  fetch = () => {
    this.setState({ loading: true });
    request.setHeader('Authorization', this.state.token);
    request.get(`${this.endpointURL}`).then(response => {
      if (response.status !== 200) {
        return;
      }

      this.setState({
        loading: false,
        data: response.data,
      });
    });
  };

  render() {
    const { loading, data } = this.state;
    return (
      <>
        {loading ? (
          <div>
            <Spin size="small" style={{ margin: '10px' }} />
          </div>
        ) : (
          <div>
            {!(data === undefined || data.length === 0) && (
              <>
                <h3 style={{ color: '#f5222d' }}>
                  پرونده های مشابه <small>(کدملی و شناسه قبض یکسان)</small>
                </h3>
                <Table
                  style={{ whiteSpace: 'pre' }}
                  dataSource={data}
                  columns={this.columns}
                  pagination={false}
                />
              </>
            )}
          </div>
        )}
      </>
    );
  }
}

export default SimilarRequest;
