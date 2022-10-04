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
import { priceDataFormater } from '../../utils/price';
import { addElasticFilters, bodyBuilder } from '../../utils/es';

class DamagesAcceptedAmount extends Component {
  state = {
    data: [],
    props: null,
    token: getCookie('token') || null,
  };

  componentDidMount() {
    this.fetch();
  }

  componentDidUpdate(prevProps) {
    if (this.state.props != this.props) {
      this.setState({ props: this.props });
      this.fetch();
    }
  }

  getElasticQuery = () => {
    const baseQ = bodyBuilder().aggregation(
      'terms',
      'damage_type.keyword',
      { min_doc_count: 0 },
      'agg_terms_damage',
      a => {
        return a.aggregation(
          'sum',
          'accepted_amount',
          'aggs_sum_sum_damage_amount',
        );
      },
    );
    addElasticFilters(baseQ, this.props);
    return baseQ;
  };

  fetch = () => {
    request.setHeader('Authorization', this.state.token);
    request
      .post(
        `${API_BASE_URL}/admin/api/v1/search/res/`,
        this.getElasticQuery().build(),
      )
      .then(response => {
        if (response.status !== 200) {
          return;
        }

        const ags = response.data.aggregations.agg_terms_damage;
        const data = Object.keys(ags.buckets).map(function(key) {
          const name = damageTypes[ags.buckets[key].key];
          const { value } = ags.buckets[key].aggs_sum_sum_damage_amount;
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
            <YAxis type="number" tickFormatter={priceDataFormater} />
            <Tooltip
              formatter={value => new Intl.NumberFormat('en').format(value)}
            />
            <Bar dataKey="value" name="مبلغ" stackId="a" fill="#8884d8" />
          </BarChart>
        </ResponsiveContainer>
      </div>
    );
  }
}

export default DamagesAcceptedAmount;
