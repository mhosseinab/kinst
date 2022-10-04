import React, { Component } from 'react';
import Router, { useRouter } from 'next/router';
import momentJalaali from 'moment-jalaali';
import { Descriptions, Row, Button, Col, Statistic } from 'antd';
import Head from 'next/head';
import { request } from '../../store/request';
import CPMessage from '../../components/CP/CPMessage';

import s from './index.scss';
import AcceptForm from './accept_form';
import RejectForm from './reject_form';
import StatusForm from './status_form';
import LackDataForm from './lack_data_form';

import { parseJwt } from '../../utils/jwt';
import { getCookie } from '../../utils/cookie';
import { API_BASE_URL } from '../../utils/const';
import CPJalaliDate, { getValidJalaliOrNull } from '../../utils/JalaliDate';
import Changelog from '../changelog';
import Items from './Items';

class RequestItemDisplay extends Component {
  constructor(props) {
    super(props);
    this.state = {
      id: props.itemID,
      data: {},
      token: getCookie('token') || null,
      btnStatusVisible: false,
      userRole: null,
    };
  }

  setUserRole = () => {
    const j = parseJwt(this.state.token);
    this.state.userRole = j.role;
  };

  componentDidMount() {
    if (!this.state.token) {
      Router.push('/login');
      return;
    }
    this.fetch();
    this.setUserRole();
  }

  state = {
    acceptVisible: false,
    rejectVisible: false,
    statusVisible: false,
    lackDataVisible: false,
  };

  showModalStatus = () => {
    this.setState({
      statusVisible: true,
    });
  };

  showModalReject = () => {
    this.setState({
      rejectVisible: true,
    });
  };

  showModalAccept = () => {
    this.setState({
      acceptVisible: true,
    });
  };

  showModalLackData = () => {
    this.setState({
      lackDataVisible: true,
    });
  };

  handleCreateAccept = e => {
    const { form } = this.formAcc.props;
    form.validateFields((err, values) => {
      if (err) {
        return;
      }

      values.expert_status = 'ACCEPTED';

      const u = `${API_BASE_URL}/admin/api/v1/request/${this.state.id}/`;
      request.put(u, values).then(response => {
        if (response.status !== 200) {
          return;
        }

        CPMessage('با موفقیت بروز رسانی شد', 'success');
        this.setState({ data: response.data });
      });

      form.resetFields();
      this.setState({ acceptVisible: false });
    });
  };

  handleCreateReject = e => {
    const { form } = this.formRej.props;
    form.validateFields((err, values) => {
      if (err) {
        return;
      }

      values.expert_status = 'REJECTED';

      const u = `${API_BASE_URL}/admin/api/v1/request/${this.state.id}/`;
      request.put(u, values).then(response => {
        if (response.status !== 200) {
          return;
        }

        CPMessage('با موفقیت بروز رسانی شد', 'success');
        this.setState({ data: response.data });
      });

      form.resetFields();
      this.setState({ rejectVisible: false });
    });
  };

  handleCancelAccept = e => {
    this.setState({
      acceptVisible: false,
    });
  };

  handleCancelReject = e => {
    this.setState({
      rejectVisible: false,
    });
  };

  handleCreateStatus = e => {
    const { form } = this.formStatus.props;
    form.validateFields((err, values) => {
      if (err) {
        return;
      }

      const u = `${API_BASE_URL}/admin/api/v1/request/${this.state.id}/`;
      request.put(u, values).then(response => {
        if (response.status !== 200) {
          return;
        }

        CPMessage('با موفقیت بروز رسانی شد', 'success');
        this.setState({ data: response.data });
      });

      form.resetFields();
      this.setState({ statusVisible: false });
    });
  };

  handleCreateLackData = e => {
    const { form } = this.formLackData.props;
    form.validateFields((err, values) => {
      if (err) {
        return;
      }

      const u = `${API_BASE_URL}/admin/api/v1/request/${this.state.id}/`;
      request.put(u, values).then(response => {
        if (response.status !== 200) {
          return;
        }

        CPMessage('با موفقیت بروز رسانی شد', 'success');
        this.setState({ data: response.data });
      });

      form.resetFields();
      this.setState({ lackDataVisible: false });
    });
  };

  handleCancelStatus = e => {
    this.setState({
      statusVisible: false,
    });
  };

  handleCancelLackData = e => {
    this.setState({
      lackDataVisible: false,
    });
  };

  fetch = (params = {}) => {
    request.setHeader('Authorization', this.state.token);
    const u = `${API_BASE_URL}/admin/api/v1/request/${this.state.id}/`;
    request.get(u, params).then(response => {
      if (response.status !== 200) {
        Router.push('/');
        return;
      }
      this.setState({ data: response.data });
    });
  };

  saveFormAcc = formAcc => {
    this.formAcc = formAcc;
  };

  saveFormStatus = formStatus => {
    this.formStatus = formStatus;
  };

  saveFormLackData = formLackData => {
    this.formLackData = formLackData;
  };

  statusChoices = {
    COMPLETED: 'تکمیل فرم توسط کاربر',
    IN_PROGRESS: 'جاری',
    CLOSED: 'مختومه',
    SUSPENDED: 'معوق',
    INCOMPLETE: 'درخواست ناقص',
    INCOMPLETE_CHANGE: 'تکمیل درخواست ناقص',
    READY_TO_PAY: 'آماده پرداخت',
    PAYED: 'پرداخت شده',
    CANCELED_BY_USER: 'انصراف ذی نفع',
  };

  expertStatus = {
    DEFAULT: 'منتظر تایید و یا در درخواست',
    ACCEPTED: 'تایید شده',
    REJECTED: 'رد شده',
  };

  getStatusText = status => {
    return this.statusChoices[status];
  };

  getImageHtml = (d, key) => {
    return (
      <div>
        <a href={d[key]} target="_blank">
          <img src={d[key]} />
        </a>
      </div>
    );
  };

  getExpertStatusText = status => {
    return this.expertStatus[status];
  };

  saveFormRej = formRej => {
    this.formRej = formRej;
  };

  getVisibleButton = r => {
    switch (this.state.userRole) {
      case 'ADMIN':
        return 'block';
        break;

      case 'BRANCH':
        if (r === 'BRANCH') {
          return 'block ';
        }
    }
    return 'none';
  };

  getItemHtml = (d, label, key) => {
    return d[key] ? (
      <Descriptions.Item label={label}>
        {this.getImageHtml(d, key)}
      </Descriptions.Item>
    ) : null;
  };

  render() {
    const d = this.state.data;
    return (
      <div>
        <Head>
          <title>درخواست خسارت</title>
        </Head>
        <Row>
          <AcceptForm
            wrappedComponentRef={this.saveFormAcc}
            visible={this.state.acceptVisible}
            onCancel={this.handleCancelAccept}
            onCreate={this.handleCreateAccept}
            data={d}
          />

          <RejectForm
            wrappedComponentRef={this.saveFormRej}
            visible={this.state.rejectVisible}
            onCancel={this.handleCancelReject}
            onCreate={this.handleCreateReject}
            data={d}
          />

          <StatusForm
            wrappedComponentRef={this.saveFormStatus}
            visible={this.state.statusVisible}
            onCancel={this.handleCancelStatus}
            onCreate={this.handleCreateStatus}
            data={d}
          />

          <LackDataForm
            wrappedComponentRef={this.saveFormLackData}
            visible={this.state.lackDataVisible}
            onCancel={this.handleCancelLackData}
            onCreate={this.handleCreateLackData}
            data={d}
          />

          <div className="noprint">
            <Row type="flex" gutter={[8, 8]}>
              <Col xs={12} sm={8} md={6}>
                <Button
                  style={{ display: this.getVisibleButton('BRANCH') }}
                  type="primary"
                  onClick={this.showModalLackData}
                >
                  اعلام نقص مدارک
                </Button>
              </Col>
              <Col xs={12} sm={8} md={6}>
                <Button
                  style={{ display: this.getVisibleButton('BRANCH') }}
                  type="primary"
                  onClick={this.showModalAccept}
                >
                  تایید کارشناس شعبه
                </Button>
              </Col>
              <Col xs={12} sm={8} md={6}>
                <Button
                  style={{ display: this.getVisibleButton('BRANCH') }}
                  type="danger"
                  onClick={this.showModalReject}
                >
                  رد کارشناس شعبه
                </Button>
              </Col>
              <Col xs={12} sm={8} md={6}>
                <Button
                  style={{ display: this.getVisibleButton('ADMIN') }}
                  type="primary"
                  onClick={this.showModalStatus}
                >
                  وضعیت پرونده
                </Button>
              </Col>
            </Row>
          </div>
          <Descriptions bordered title="پرونده" layout="">
            <Descriptions.Item label="وضعیت">
              {this.getStatusText(d.status)}
            </Descriptions.Item>
            <Descriptions.Item label="مبلغ مورد تایید">
              <Statistic value={d.accepted_amount} /> ریال
            </Descriptions.Item>
          </Descriptions>
          <Descriptions bordered title="کارشناسی" layout="">
            <Descriptions.item label="تاریخ کارشناسی">
              {getValidJalaliOrNull(d.expert_updated_at)}
            </Descriptions.item>
            <Descriptions.item label="تایید کارشناس">
              {this.getExpertStatusText(d.expert_status)}
            </Descriptions.item>
            <Descriptions.Item label="مبلغ مورد تایید کارشناس">
              <Statistic value={d.expert_accepted_amount} /> ریال
            </Descriptions.Item>
            <Descriptions.Item label="توضیحات کارشناس">
              {d.expert_description}
            </Descriptions.Item>
            <Descriptions.Item label="نقص مدارک">
              {d.lack_data_description}
            </Descriptions.Item>
          </Descriptions>
          <Descriptions bordered title="اطلاعات زیاندیده" layout="horizontal">
            <Descriptions.Item label="استان">{d.province}</Descriptions.Item>
            <Descriptions.Item label="شهر">{d.city}</Descriptions.Item>
            <Descriptions.Item label="شناسه قبض">
              {d.bill_identifier}
            </Descriptions.Item>
            <Descriptions.Item label="کد رهگیری">
              {d.reference_code}
            </Descriptions.Item>
            <Descriptions.Item label="کد ملی">
              {d.national_code}
            </Descriptions.Item>
            <Descriptions.Item label="نوع مشترک">
              {d.subscriber_type === 1 ? 'حقیقی' : 'حقوقی'}
            </Descriptions.Item>
            <Descriptions.Item label="نام">
              {d.firstname} {d.surname}
            </Descriptions.Item>
            <Descriptions.Item label="همراه">
              {d.mobile_number}
            </Descriptions.Item>
            <Descriptions.Item label="آدرس">{d.address}</Descriptions.Item>
            <Descriptions.Item label="کد پستی">
              {d.postal_address}
            </Descriptions.Item>
            <Descriptions.Item label="شماره شبا">{d.sheba}</Descriptions.Item>
            <Descriptions.Item label="مبلغ خسارت مورد ادعا">
              <Statistic value={d.sum_damage_amount} /> ریال
            </Descriptions.Item>
            <Descriptions.Item label="توضیحات">
              {d.description}
            </Descriptions.Item>
            <Descriptions.Item label="تاریخ حادثه">
              {momentJalaali(d.casuality_date).format('jYYYY/jM/jD')}
            </Descriptions.Item>
            <Descriptions.Item label="ساعت حادثه">
              {d.casuality_time}
            </Descriptions.Item>
            <Descriptions.Item label="زمان ثبت حادثه">
              {momentJalaali(d.created_at).format('jYYYY/jM/jD-H:m')}
            </Descriptions.Item>
          </Descriptions>
          <div className="noprint">
            <Descriptions
              title="مدارک زیاندیده"
              bordered
              className={s.smallImage}
              layout="vertical"
            >
              <Descriptions.Item label="تصویر کارت ملی">
                {this.getImageHtml(d, 'id_card_photo')}
              </Descriptions.Item>
              <Descriptions.Item label="آخرین قبض پرداختی">
                {this.getImageHtml(d, 'last_bill_photo')}
              </Descriptions.Item>
              <Descriptions.Item label="سند مالکیت">
                {this.getImageHtml(d, 'other_photo')}
              </Descriptions.Item>
            </Descriptions>
            <Descriptions
              className={s.smallImage}
              layout="vertical"
              bordered
              style={{
                display: d.damage_type === 'firing_damage' ? 'block' : 'none',
              }}
              title="خسارت آتش سوزی"
            >
              <Descriptions.Item label="مبلغ" span={3}>
                {d.firing_damage_amount}
              </Descriptions.Item>
              {this.getItemHtml(
                d,
                'تصویر محل حادثه (1)',
                'firing_place_1_photo',
              )}
              {this.getItemHtml(
                d,
                'تصویر محل حادثه (2)',
                'firing_place_2_photo',
              )}
              {this.getItemHtml(
                d,
                'تصویر گزارش آتش نشانی (1)',
                'firing_station_report_photo',
              )}
              {this.getItemHtml(
                d,
                'تصویر گزارش کلانتری',
                'firing_police_report_photo',
              )}
              {this.getItemHtml(
                d,
                'تصویر پرونده دادگاه در صورت شکایت',
                'firing_court_report_photo',
              )}
              {this.getItemHtml(
                d,
                'تصویر فاکتور فروشگاه یا تعمیرگاه مجاز ۱',
                'firing_invoice_1_photo',
              )}
              {this.getItemHtml(
                d,
                'تصویر فاکتور فروشگاه یا تعمیرگاه مجاز ۲',
                'firing_invoice_2_photo',
              )}
              {this.getItemHtml(
                d,
                'تصویر فاکتور فروشگاه یا تعمیرگاه مجاز ۳',
                'firing_invoice_3_photo',
              )}
              {this.getItemHtml(
                d,
                'تصویر فاکتور فروشگاه یا تعمیرگاه مجاز ۴',
                'firing_invoice_4_photo',
              )}
              {this.getItemHtml(
                d,
                'تصویر فاکتور فروشگاه یا تعمیرگاه مجاز ۵',
                'firing_invoice_5_photo',
              )}
            </Descriptions>

            <Descriptions
              bordered
              className={s.smallImage}
              layout="vertical"
              style={{
                display:
                  d.damage_type === 'instrument_damage' ? 'block' : 'none',
              }}
              title="خسارت تجهیزات"
            >
              {this.getItemHtml(
                d,
                'تصویر فاکتور فروشگاه یا تعمیرگاه مجاز',
                'instrument_invoice_photo',
              )}
              {this.getItemHtml(
                d,
                'تصویر فاکتور فروشگاه یا تعمیرگاه مجاز',
                'instrument_invoice_2_photo',
              )}
              {this.getItemHtml(
                d,
                'تصویر فاکتور فروشگاه یا تعمیرگاه مجاز',
                'instrument_invoice_3_photo',
              )}
              {this.getItemHtml(
                d,
                'تصویر فاکتور فروشگاه یا تعمیرگاه مجاز',
                'instrument_invoice_4_photo',
              )}
              {this.getItemHtml(
                d,
                'تصویر فاکتور فروشگاه یا تعمیرگاه مجاز',
                'instrument_invoice_5_photo',
              )}
              {this.getItemHtml(
                d,
                'تصویر فاکتور فروشگاه یا تعمیرگاه مجاز',
                'instrument_invoice_6_photo',
              )}
              {this.getItemHtml(
                d,
                'تصویر فاکتور فروشگاه یا تعمیرگاه مجاز',
                'instrument_invoice_7_photo',
              )}
              {this.getItemHtml(
                d,
                'تصویر فاکتور فروشگاه یا تعمیرگاه مجاز',
                'instrument_invoice_8_photo',
              )}
              {this.getItemHtml(
                d,
                'تصویر گزارش معتمدین محل یا نیروی انتظامی',
                'instrument_report_photo',
              )}
            </Descriptions>

            <Descriptions
              className={s.smallImage}
              style={{
                display:
                  d.damage_type === 'explosion_damage' ? 'block' : 'none',
              }}
              layout="vertical"
              bordered
              title="خسارت انفجار"
            >
              {this.getItemHtml(
                d,
                'تصویر گزارش آتش نشانی یا مقامات ذیصلاح براساس نوع انفجار',
                'explosion_firestation_report_photo',
              )}
              {this.getItemHtml(
                d,
                'تصویر لیست موارد آسیب دیده',
                'explostion_damaged_items_photo',
              )}
              {this.getItemHtml(
                d,
                'تصویر فاکتور تعمیرات',
                'explosion_invoice_photo',
              )}
              {this.getItemHtml(
                d,
                'تصویر محل حادثه (1)',
                'explosion_place_1_photo',
              )}
              {this.getItemHtml(
                d,
                'تصویر محل حادثه (2)',
                'explosion_place_2_photo',
              )}
            </Descriptions>

            <Descriptions
              className={s.smallImage}
              layout="vertical"
              style={{
                display: d.damage_type === 'medical_damage' ? 'block' : 'none',
              }}
              bordered
              title="خسارت پزشکی"
            >
              {this.getItemHtml(
                d,
                'تصویر اصل صورت حساب بیمارستان',
                'medical_hospital_invoice_photo',
              )}
              {this.getItemHtml(
                d,
                'تصویر مدارک پزشکی و پرونده های بیمارستانی',
                'medical_hospital_document_photo',
              )}
              {this.getItemHtml(
                d,
                'تصویر گزارش مقامات ذیصلاح',
                'medical_report_photo',
              )}
            </Descriptions>

            <Descriptions
              className={s.smallImage}
              layout="vertical"
              style={{
                display: d.damage_type === 'lack_damage' ? 'block' : 'none',
              }}
              bordered
              title="خسارت نقض عضو"
            >
              {this.getItemHtml(
                d,
                'گزارش انتظامی یا بازرس آتش نشانی',
                'lack_report_photo',
              )}
              {this.getItemHtml(
                d,
                'رادیولوژی بعد از حادثه',
                'lack_radiology_photo',
              )}
              {this.getItemHtml(
                d,
                'اولین مرجع درمانی پزشکی معالج',
                'lack_first_reference_photo',
              )}
              {this.getItemHtml(
                d,
                'کپی کارت ملی، شناسنامه مصدوم',
                'lack_id_card_photo',
              )}
              {this.getItemHtml(d, 'گواهی پزشک معالج', 'lack_witness_photo')}
              {this.getItemHtml(
                d,
                'صورت حساب پزشکی پرونده مصدوم',
                'lack_invoice_photo',
              )}
            </Descriptions>

            <Descriptions
              className={s.smallImage}
              layout="vertical"
              style={{
                display: d.damage_type === 'death_damage' ? 'block' : 'none',
              }}
              bordered
              title="خسارت فوت"
            >
              {this.getItemHtml(d, 'تصویر رای قاضی', 'death_judge_vote_photo')}
              {this.getItemHtml(d, 'گواهی فوت', 'death_witness_photo')}
              {this.getItemHtml(
                d,
                'شناسنامه ابطال شده، کارت ملی',
                'death_id_card_photo',
              )}
              {this.getItemHtml(
                d,
                'گزارش پزشک قانونی',
                'death_toxicology_report_photo',
              )}
              {this.getItemHtml(
                d,
                'تصویر معاینه جسد',
                'death_corpse_examination_photo',
              )}
              {this.getItemHtml(d, 'گواهی انحصار وراثت', 'death_probate_photo')}
            </Descriptions>
          </div>
        </Row>
        <div className="noprint">
          <Changelog request_id={this.state.id} />
        </div>
      </div>
    );
  }
}

const Index = () => {
  const router = useRouter();
  const { requestId } = router.query;

  return <RequestItemDisplay itemID={requestId} />;
};

export default Index;
