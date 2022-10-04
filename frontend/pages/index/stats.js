import React, { Component } from 'react';
import {
  Legend,
  Tooltip,
  BarChart,
  CartesianGrid,
  XAxis,
  YAxis,
  Bar,
  ResponsiveContainer,
} from 'recharts';
import { Statistic, Row, Col } from 'antd';
import { request } from '../../store/request';
import { getCookie } from '../../utils/cookie';
import { API_BASE_URL } from '../../utils/const';
import { damageTypes, statusChoices } from '../../utils/damage';
import { isUserBranch, getBranchState } from '../../utils/role';

const bodybuilder = require('bodybuilder');

class Stats extends Component {
  getQuery = () => {
    const baseQ = bodybuilder().aggregation(
      'terms',
      'status.keyword',
      'agg_terms_status',
    );
    if (isUserBranch()) {
      baseQ.query('term', 'province', getBranchState());
    }
    return baseQ;
  };

  state = {
    dataAgg: [],
    allCount: 0,
    token: getCookie('token') || null,
  };

  componentDidMount() {
    this.fetch();
  }

  fetch = () => {
    request.setHeader('Authorization', this.state.token);
    request
      .post(`${API_BASE_URL}/admin/api/v1/search/res/`, this.getQuery().build())
      .then(response => {
        if (response.status !== 200) {
          return;
        }

        const ags = response.data.aggregations.agg_terms_status;
        const data = Object.keys(ags.buckets).map(function(key) {
          const name = statusChoices[ags.buckets[key].key];
          const value = ags.buckets[key].doc_count;
          return {
            name,
            value,
          };
        });

        this.setState({
          dataAgg: [
            { name: 'کل', value: response.data.hits.total.value },
            ...data,
          ],
          allCount: response.data.hits.total.value,
        });
      });
  };

  getStatsHtml = d => {
    return Object.keys(d).map(function(key) {
      return (
        <Col xs={24} md={12} lg={3}>
          <Statistic title={d[key].name} value={d[key].value} />
        </Col>
      );
    });
    // return this.state.data;
  };

  render() {
    return (
      <div>
        <Row
          gutter={16}
          style={{ display: 'flex', flexDirection: 'row', flexWrap: 'wrap' }}
        >
          {this.getStatsHtml(this.state.dataAgg)}
        </Row>
      </div>
    );
  }
}

export default Stats;
