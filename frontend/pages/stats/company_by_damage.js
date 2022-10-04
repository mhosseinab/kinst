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

import { addElasticFilters } from '../../utils/es';

class CompanyByDamage extends Component {
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
          agg_terms_province: {
            terms: {
              field: 'company_id.keyword',
              size: 40,
            },
            aggs: {
              agg_terms_damage_type: {
                terms: {
                  field: 'damage_type.keyword',
                },
              },
            },
          },
        },
      })
      .then(response => {
        if (response.status !== 200) {
          return;
        }

        const ags = response.data.aggregations.agg_terms_province;
        const data = Object.keys(ags.buckets).map(function(key) {
          const name = ags.buckets[key].key;
          const value = ags.buckets[key].doc_count;
          const dt = ags.buckets[key].agg_terms_damage_type.buckets;
          const dm = {};
          for (let index = 0; index < dt.length; index++) {
            const element = dt[index];
            dm[element.key] = element.doc_count;
          }

          return {
            name,
            value,
            dm,
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
        <ResponsiveContainer width="100%" height={400}>
          <BarChart
            data={this.state.data}
            margin={{ top: 20, right: 30, left: 20, bottom: 5 }}
          >
            <CartesianGrid strokeDasharray="3 3" />
            <XAxis dataKey="name" />
            <YAxis />
            <Tooltip />
            <Legend />
            <Bar dataKey="value" name="تعداد" stackId="a" fill="#8884d8" />
            <Bar
              dataKey="dm.instrument_damage"
              name="تجهیزات"
              fill="#FFBB28"
              stackId="dmi"
            />
            <Bar dataKey="dm.firing_damage" name="آتش سوزی" fill="#82ca9d" />
            <Bar dataKey="dm.death_damage" name="فوت" />
            <Bar dataKey="dm.lack_damage" name="نقص عضو" />
            <Bar dataKey="dm.medical_damage" name="پزشکی" />
            <Bar dataKey="dm.explosion_damage" name="انفجار" />
          </BarChart>
        </ResponsiveContainer>
      </div>
    );
  }
}

export default CompanyByDamage;
