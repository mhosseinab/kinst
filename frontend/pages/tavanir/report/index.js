import React, { Component } from 'react';
import Head from 'next/head';
import { Table, Tooltip, Button } from 'antd';
import Router from 'next/router';
import Humanize from 'humanize-plus';
import fileDownload from 'js-file-download';
import { request } from '../../../store/request';
import s from './index.scss';
import { API_BASE_URL } from '../../../utils/const';
import { getCookie } from '../../../utils/cookie';
import { parseJwt } from '../../../utils/jwt';
import Filters from '../../../components/Filters/filters';
import { userRoleBranch } from '../../../utils/role';
import { bodyBuilder, addElasticFilters } from '../../../utils/es';
import { companiesTavanir } from '../../../utils/companis';
import Types from '../../../components/Types';
import Activity from './Activity';
import { userRoleReporter, authToken } from '../../../const/user';

class Report extends Component {
  state = {
    province: undefined,
    data: undefined,
    token: getCookie('token') || null,
    status: undefined,
    location_type: undefined,
    loading: false,
    damage_type: undefined,
    company_id: undefined,
    from: undefined,
    to: undefined,
  };

  componentDidUpdate(prevProps) {
    if (this.state.props != this.props) {
      this.setState({ props: this.props });
      this.fetch();
    }
  }

  componentDidMount() {
    const token = getCookie('token') || null;
    const payload = parseJwt(token);
    if (!token) {
      Router.push('/login');
      return;
    }

    // if (payload.role === userRoleBranch) {
    //   Router.push('/access-denied');
    //   return;
    // }

    if (payload.state !== '') {
      this.setState({ company_id: payload.state }, () => {
        this.fetch();
      });
    } else {
      this.fetch();
    }
  }

  handleChangeProvince = value => {
    this.setState({
      province: value,
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

  handleChangeStatus = () => {
    this.setState(prevState => ({
      status: !prevState.status,
    }));
  };

  getElasticQuery = () => {
    const baseQ = bodyBuilder().aggregation(
      'terms',
      'companyId.keyword',
      { min_doc_count: 1, size: 40 },
      'agg_terms_status',
      a => {
        return a
          .aggregation(
            'terms',
            'compensationTypeId.keyword',
            { min_doc_count: 0, size: 40 },
            'damage_types',
            b => {
              return b.aggregation('sum', 'amount', {}, 'sum_damage_amount');
            },
          )
          .aggregation(
            'terms',
            'status.keyword',
            { min_doc_count: 0, size: 40 },
            'damage_status',
            b => {
              return b.aggregation('sum', 'amount', {}, 'sum_damage_amount');
            },
          );
      },
    );
    addElasticFilters(baseQ, this.state);
    return baseQ;
  };

  fetch = () => {
    this.setState({ loading: true });
    request.setHeader('Authorization', this.state.token);
    request
      .post(
        `${API_BASE_URL}/admin/api/v1/search/res_tavanir/`,
        this.getElasticQuery().build(),
      )
      .then(response => {
        if (response.status !== 200) {
          return;
        }

        const rows = [];
        const b = response.data.aggregations.agg_terms_status.buckets;
        b.map(k => {
          const o = {
            code: k.key,
            company_id: companiesTavanir[k.key].replace('توزيع نيروي برق', ''), //companiesTavanir[k.key].replace('توزيع نيروي برق', ''),
            damage_count: k.doc_count,
            sum_all: 0,
            damages: {},
            status: {},
          };

          k.damage_types.buckets.map(d => {
            o.damages[d.key] = {
              doc_count: d.doc_count,
              sum: d.sum_damage_amount.value,
            };
            o.sum_all += d.sum_damage_amount.value;
          });

          k.damage_status.buckets.map(d => {
            o.status[d.key] = {
              doc_count: d.doc_count,
              sum: d.sum_damage_amount.value,
            };
            o.sum_all += d.sum_damage_amount.value;
          });

          rows.push(o);
        });

        this.setState({
          data: rows,
          loading: false,
        });
      });
  };

  handleChange = (value, key) => {
    this.setState({ [key]: value }, () => {
      if (key !== 'from' && key !== 'to') {
        this.fetch();
      }
    });
  };

  handleDownload = () => {
    const timestamp = new Date().getTime();
    request.setHeader('Authorization', this.state.token);
    request.get(`${API_BASE_URL}/admin/api/v1/export/excel/`).then(res => {
      fileDownload(res, `reports_${timestamp}.xlsx`);
    });
  };

  render() {
    const { state, status, from, to, token } = this.state;
    return (
      <div className={s.root}>
        <Head>
          <title>گزارشات</title>
        </Head>
        {/* <Types
          location_type={this.state.location_type}
          onChange={this.handleChange}
        /> */}
        <Filters {...this.state} onChange={this.handleChange} />

        <a
          target="_blank"
          href={`${API_BASE_URL}/admin/api/v1/export/excel/?token=${token}${
            from && to ? `&from=${from}&to=${to}` : ''
          }`}
          download
        >
          <Button style={{ marginBottom: 16 }} type="primary">
            دانلود گزارشات اکسل
          </Button>
        </a>
        <Activity data={this.state.data} loading={this.state.loading} />
      </div>
    );
  }
}

export default Report;
