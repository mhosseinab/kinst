import React, { Component } from 'react';
import Head from 'next/head';
import { Button, Row, Col, Tabs, DatePicker } from 'antd';
import Router from 'next/router';
import s from './index.scss';
// import { request } from '../../store/request';
// import { getCookie } from '../../utils/cookie';
// import { damageTypes } from '../../utils/damage';
// import { API_BASE_URL } from '../../utils/const';
// import CompanyByDamage from './company_by_damage';
import Damages from './damages';
import DamageStatus from './damage_status';
import ExpertStatus from './expert_status';
import Province from './province';
import City from './city';
import DamagesAmount from './damages_amount';
import DamagesAcceptedAmount from './damages_accepted_amount';
import DamagesExpertAcceptedAmount from './damages_export_accepted_amount';
import StatusExpectedAmount from './status_amount';
import StatusAcceptedAmount from './status_accepted_amount';
import StatusExpertAmount from './status_expert_amount';
import Filters from '../../components/Filters/filters';
import { userRoleAdmin, userRoleBranch } from '../../utils/role';
import { getCookie } from '../../utils/cookie';
import { parseJwt } from '../../utils/jwt';

const { TabPane } = Tabs;

class Report extends Component {
  state = {
    province: undefined,
    city: undefined,
    status: undefined,
    expert: undefined,
    company_id: undefined,
  };

  componentDidMount() {
    const token = getCookie('token') || null;
    const payload = parseJwt(token);
    if (!token) {
      Router.push('/login');
      return;
    }

    // if (payload.role === userRoleBranch) {
    //   Router.push('/access-denied');
    // }

    if (payload.stats !== null) {
      this.setState({
        company_id: payload.state,
      });
    }
  }

  handleChange = (value, key) => {
    this.setState({
      [key]: value,
    });
  };

  render() {
    return (
      <div className={s.root}>
        <Head>
          <title>گزارشات</title>
        </Head>
        <Filters {...this.state} onChange={this.handleChange} />

        <Tabs defaultActiveKey="1">
          <TabPane tab="مبالغ" key="2">
            <Row>
              <Col xs={24} sm={24} md={12} lg={8}>
                <h3 style={{ marginBottom: 16 }}>خسارت مورد ادعا</h3>
                <DamagesAmount {...this.state} />
              </Col>
              <Col xs={24} sm={24} md={12} lg={8}>
                <h3 style={{ marginBottom: 16 }}>خسارت مورد تایید مدیر</h3>
                <DamagesAcceptedAmount {...this.state} />
              </Col>
              <Col xs={24} sm={24} md={12} lg={8}>
                <h3 style={{ marginBottom: 16 }}>خسارت مورد تایید کارشناس</h3>
                <DamagesExpertAcceptedAmount {...this.state} />
              </Col>
              <Col xs={24} sm={24} md={12} lg={8}>
                <h3 style={{ marginBottom: 16, marginTop: 16 }}>
                  وضعیت خسارتهای مورد ادعا
                </h3>
                <StatusExpectedAmount {...this.state} />
              </Col>
              <Col xs={24} sm={24} md={12} lg={8}>
                <h3 style={{ marginBottom: 16, marginTop: 16 }}>
                  وضعیت خسارتهای مدیر
                </h3>
                <StatusAcceptedAmount {...this.state} />
              </Col>
              <Col xs={24} sm={24} md={12} lg={8}>
                <h3 style={{ marginBottom: 16, marginTop: 16 }}>
                  وضعیت خسارتهای کارشناس
                </h3>
                <StatusExpertAmount {...this.state} />
              </Col>
            </Row>
          </TabPane>
          <TabPane tab="نمای کلی پرونده ها" key="1">
            <Row>
              <Col xs={24} sm={24} md={12} lg={8}>
                <h3 style={{ marginBottom: 16 }}>انواع خسارت</h3>
                <Damages {...this.state} />
              </Col>
              <Col xs={24} sm={24} md={12} lg={8}>
                <h3 style={{ marginBottom: 16 }}>وضعیت پرونده</h3>
                <DamageStatus {...this.state} />
              </Col>
              <Col xs={24} sm={24} md={12} lg={8}>
                <h3 style={{ marginBottom: 16 }}>وضعیت کارشناسی</h3>
                <ExpertStatus {...this.state} />
              </Col>
            </Row>
            <Row>
              <Col span={24}>
                <h3 style={{ marginBottom: 16, marginTop: 16 }}>
                  خسارتهای ثبت شده به تفکیک استان
                </h3>
                <Province {...this.state} />
              </Col>
              <Col span={24}>
                <h3 style={{ marginBottom: 16, marginTop: 16 }}>
                  خسارتهای ثبت شده به تفکیک شهر
                </h3>
                <City {...this.state} />
              </Col>
            </Row>
          </TabPane>
        </Tabs>
      </div>
    );
  }
}

export default Report;
