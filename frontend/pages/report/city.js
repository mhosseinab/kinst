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
import { request } from '../../store/request';
import { getCookie } from '../../utils/cookie';
import { API_BASE_URL } from '../../utils/const';
import { damageTypes } from '../../utils/damage';

class City extends Component {
  state = {
    data: [],
    token: getCookie('token') || null,
  };

  componentDidMount() {
    this.fetch();
  }

  fetch = () => {
    request.setHeader('Authorization', this.state.token);
    request
      .post(`${API_BASE_URL}/admin/api/v1/search/res/`, {
        aggs: {
          agg_terms_damage: {
            terms: {
              field: 'city.keyword',
              size: 40,
            },
          },
        },
      })
      .then(response => {
        if (response.status !== 200) {
          return;
        }

        const ags = response.data.aggregations.agg_terms_damage;
        const data = Object.keys(ags.buckets).map(function(key) {
          const name = ags.buckets[key].key;
          const value = ags.buckets[key].doc_count;
          return {
            name,
            value,
          };
        });

        this.setState({
          data,
        });
      });
  };

  render() {
    return (
      <div style={{ direction: 'ltr', textAlign: 'left' }} className="">
        <ResponsiveContainer width="100%" height={200}>
          <BarChart
            data={this.state.data}
            margin={{ top: 20, right: 30, left: 20, bottom: 5 }}
          >
            <CartesianGrid strokeDasharray="3 3" />
            <XAxis dataKey="name" />
            <YAxis />
            <Tooltip dataKey="name" />
            {/* <Legend /> */}
            <Bar dataKey="value" name="تعداد" stackId="a" fill="#8884d8" />
          </BarChart>
        </ResponsiveContainer>
      </div>
    );
  }
}

export default City;
