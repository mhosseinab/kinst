import React, { Component } from 'react';
import Router, { useRouter } from 'next/router';
import {
  Descriptions,
  Row,
  Button,
  Col,
  Statistic,
  Icon,
  Spin,
  Divider,
  Alert,
} from 'antd';
import Head from 'next/head';
import momentJalaali from 'moment-jalaali';
import { request } from '../../store/request';
import CPMessage from '../../components/CP/CPMessage/CPMessage';

import s from './index.scss';
import AcceptForm from './accept_form';
import RejectForm from './reject_form';
import StatusForm from './status_form';
import LackDataForm from './lack_data_form';
import InPersonForm from './in_person_form';

import { parseJwt } from '../../utils/jwt';
import { getCookie } from '../../utils/cookie';
import { API_BASE_URL } from '../../utils/const';
import { snakeCase } from '../../utils/snakeCase';
import CPJalaliDate, { getValidJalaliOrNull } from '../../utils/JalaliDate';
import {
  tavanirDamageTypes,
  tavanirStatusChoices,
  expertStatusChoices,
} from '../../utils/damage';
import Changelog from '../changelog';
import SyncLog from '../../components/SyncLog';
import SimilarRequest from './similar';
import Items from './Items';

let endpointURL = `${API_BASE_URL}/admin/api/v1/tavanir`;

class RequestItemDisplay extends Component {
  documentTypes = {
    '3': 'آخرين قبض پرداختي',
    '2': 'تصوير کارت ملي',
    '1': 'سند مالکيت يا اجاره نامه',
    '21': 'تصوير راي قاضي',
    '22': 'گواهي فوت',
    '23': 'شناسنامه ابطال شده',
    '24': 'گزارش پزشک قانوني',
    '25': 'تصوير معاينه جسد',
    '26': 'گواهي انحصار وراثت',
    '41': 'فاکتور فروشگاه',
    '42': 'فاکتور فروشگاه',
    '43': 'فاکتور فروشگاه',
    '44': 'فاکتور فروشگاه',
    '45': 'فاکتور فروشگاه',
    '46': 'فاکتور فروشگاه',
    '47': 'فاکتور فروشگاه',
    '49': 'گزارش معتمدين محل يا نيروي انتظامي',
  };

  constructor(props) {
    super(props);
    this.state = {
      id: props.itemID,
      pk: props.pk,
      data: {},
      token: getCookie('token') || null,
      btnStatusVisible: false,
      userRole: null,
      loading: true,
      acceptVisible: false,
      rejectVisible: false,
      statusVisible: false,
      lackDataVisible: false,
      InPersonFormVisible: false,
    };
  }

  componentDidMount() {
    if (!this.state.token) {
      Router.push('/login');
      return;
    }
    this.fetch();
    this.setUserRole();
  }

  setUserRole = () => {
    const j = parseJwt(this.state.token);
    this.state.userRole = j.role;
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
      values.status = values.expert_status;
      const u = `${endpointURL}/case/${this.state.id}/`;
      request.put(u, values).then(response => {
        if (response.status !== 200) {
          const { data } = response;
          CPMessage(
            data && data.error_msg ? data.error_msg : 'خطایی رخ داد',
            'error',
          );
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

      values.status = 'REJECTED';
      const u = `${endpointURL}/case/${this.state.id}/`;
      request.put(u, values).then(response => {
        if (response.status !== 200) {
          let { data } = response;
          CPMessage(
            data && data.error_msg ? data.error_msg : 'خطایی رخ داد',
            'error',
          );
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
      values.status = values.expert_status;
      const u = `${endpointURL}/case/${this.state.id}/`;
      request.put(u, values).then(response => {
        if (response.status !== 200) {
          let { data } = response;
          CPMessage(
            data && data.error_msg ? data.error_msg : 'خطایی رخ داد',
            'error',
          );
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

      if (values.missing_documents) {
        values.missing_documents = values.missing_documents.join(',');
      }
      values.status = 'INCOMPLETE';
      console.log('values', values);

      const u = `${endpointURL}/case/${this.state.id}/`;
      request.put(u, values).then(response => {
        if (response.status !== 200) {
          let { data } = response;
          CPMessage(
            data && data.error_msg ? data.error_msg : 'خطایی رخ داد',
            'error',
          );
          return;
        }

        CPMessage('با موفقیت بروز رسانی شد', 'success');
        this.setState({ data: response.data });
      });

      form.resetFields();
      this.setState({ lackDataVisible: false });
    });
  };

  handleCreateInPersonForm = e => {
    const { form } = this.formInPerson.props;
    form.validateFields((err, values) => {
      if (err) {
        return;
      }

      values.status = 'IN_PROGRESS';
      console.log('values', values);

      const u = `${endpointURL}/case/${this.state.id}/`;
      request.put(u, values).then(response => {
        if (response.status !== 200) {
          let { data } = response;
          CPMessage(
            data && data.error_msg ? data.error_msg : 'خطایی رخ داد',
            'error',
          );
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
    const { pk, id, token } = this.state;
    request.setHeader('Authorization', token);
    const u = pk ? `${endpointURL}/id/${pk}/` : `${endpointURL}/case/${id}/`;
    this.setState({ loading: true });
    request.get(u, params).then(response => {
      if (response.status !== 200) {
        const { data } = response;
        CPMessage(
          data && data.error_msg ? data.error_msg : 'خطایی رخ داد',
          'error',
        );
        return;
      }
      this.setState({
        loading: false,
        data: response.data,
        id: response.data.id,
      });
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

  saveInPersonData = data => {
    this.formInPerson = data;
  };

  getStatusText = status => {
    // console.log('status', status);
    // console.log('status', tavanirStatusChoices[status]);
    return tavanirStatusChoices[status];
  };

  getDamageTypeText = damageType => {
    // console.log('status', status);
    // console.log('status', tavanirStatusChoices[status]);
    return tavanirDamageTypes[damageType];
  };

  getImageHtml = (d, key) => {
    return (
      <div>
        <a href={d[key]} target="_blank" rel="noopener noreferrer">
          <img src={d[key]} alt="" />
        </a>
      </div>
    );
  };

  getExpertStatusText = status => {
    return expertStatusChoices[status];
  };

  saveFormRej = formRej => {
    this.formRej = formRej;
  };

  getVisibleButton = r => {
    const { userRole } = this.state;
    switch (userRole) {
      case 'ADMIN':
        return 'block';

      case 'BRANCH':
        if (r === 'BRANCH') return 'block ';
        return 'none';
      default:
        return 'none';
    }
  };

  render() {
    const {
      data,
      token,
      acceptVisible,
      rejectVisible,
      statusVisible,
      InPersonFormVisible,
      lackDataVisible,
    } = this.state;
    const { loading } = this.state;
    // console.log('d', data.documents);
    // console.log('data.id', data.id);
    return (
      <div>
        <Head>
          <title>درخواست خسارت</title>
        </Head>
        {loading ? (
          <div className={s['loading-large']}>
            <Spin size="large" />
          </div>
        ) : (
          <>
            <Row>
              <AcceptForm
                wrappedComponentRef={this.saveFormAcc}
                visible={acceptVisible}
                onCancel={this.handleCancelAccept}
                onCreate={this.handleCreateAccept}
                data={data}
              />

              <RejectForm
                wrappedComponentRef={this.saveFormRej}
                visible={rejectVisible}
                onCancel={this.handleCancelReject}
                onCreate={this.handleCreateReject}
                data={data}
              />

              <StatusForm
                wrappedComponentRef={this.saveFormStatus}
                visible={statusVisible}
                onCancel={this.handleCancelStatus}
                onCreate={this.handleCreateStatus}
                data={data}
              />

              <InPersonForm
                wrappedComponentRef={this.saveInPersonData}
                visible={InPersonFormVisible}
                onCancel={() =>
                  this.setState({
                    InPersonFormVisible: false,
                  })
                }
                onCreate={this.handleCreateInPersonForm}
                data={data}
              />

              <LackDataForm
                wrappedComponentRef={this.saveFormLackData}
                visible={lackDataVisible}
                onCancel={this.handleCancelLackData}
                onCreate={this.handleCreateLackData}
                data={data}
              />

              <div className="noprint">
                <Row type="flex" gutter={[8, 8]}>
                  <Col sm={6} md={4}>
                    <Button
                      style={{ display: this.getVisibleButton('BRANCH') }}
                      type="primary"
                      onClick={this.showModalLackData}
                    >
                      اعلام نقص مدارک
                    </Button>
                  </Col>
                  <Col sm={6} md={4}>
                    <Button
                      style={{ display: this.getVisibleButton('BRANCH') }}
                      type="primary"
                      onClick={() =>
                        this.setState({
                          InPersonFormVisible: true,
                        })
                      }
                    >
                      نیاز به مراجعه حضوری
                    </Button>
                  </Col>
                  <Col sm={6} md={4}>
                    <Button
                      style={{ display: this.getVisibleButton('BRANCH') }}
                      type="primary"
                      onClick={this.showModalAccept}
                    >
                      تایید کارشناس شعبه
                    </Button>
                  </Col>
                  <Col sm={6} md={4}>
                    <Button
                      style={{ display: this.getVisibleButton('BRANCH') }}
                      type="danger"
                      onClick={this.showModalReject}
                    >
                      رد کارشناس شعبه
                    </Button>
                  </Col>
                  <Col sm={6} md={4}>
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
                  {this.getStatusText(data.status)}
                </Descriptions.Item>
                <Descriptions.Item label="نوع خسارت">
                  {this.getDamageTypeText(data.compensationTypeId)}
                </Descriptions.Item>
                <Descriptions.Item label="مبلغ مورد تایید">
                  <Statistic
                    value={
                      data.expert_accepted_amount < 5000000
                        ? data.expert_accepted_amount
                        : data.accepted_amount
                    }
                  />{' '}
                  ریال
                </Descriptions.Item>
              </Descriptions>
              {data.accepted_amount === 0 &&
                data.expert_accepted_amount !== 0 && (
                  <Alert
                    message="مبلغ خسارت نیاز به تایید مدیر دارد"
                    type="warning"
                  />
                )}
              {data.status === 'INCOMPLETE' &&
                data.expert_status !== 'INCOMPLETE_CHANGE' && (
                  <Alert
                    message="پرونده ناقص است | در انتظار تکمیل مدارک توسط مشترک"
                    type="warning"
                  />
                )}
              {data.status === 'INCOMPLETE' &&
                data.expert_status === 'INCOMPLETE_CHANGE' && (
                  <Alert
                    message="پرونده ناقص است | مدارک جدید توسط مشترک بارگذاری شده است"
                    type="success"
                  />
                )}
              {data.status === 'READY_TO_PAY' &&
                data.accepted_amount !== 0 &&
                data.expert_status !== 'INCOMPLETE_CHANGE' && (
                  <Alert
                    message="در انتظار دریافت شماره شبا توسط مشترک"
                    type="warning"
                  />
                )}
              {data.status === 'READY_TO_PAY' &&
                data.expert_status === 'INCOMPLETE_CHANGE' && (
                  <Alert
                    message="شماره شبا توسط مشترک ثبت شده است"
                    type="success"
                  />
                )}
              <Descriptions>
                <Descriptions.Item>
                  <SimilarRequest case_id={data.id} />
                </Descriptions.Item>
              </Descriptions>
              <Descriptions bordered title="کارشناسی" layout="">
                <Descriptions.item label="تاریخ کارشناسی">
                  {getValidJalaliOrNull(data.updated_at)}
                </Descriptions.item>
                <Descriptions.item label="تایید کارشناس">
                  {this.getExpertStatusText(data.expert_status)}
                </Descriptions.item>
                <Descriptions.Item label="مبلغ مورد تایید کارشناس">
                  <Statistic value={data.expert_accepted_amount} /> ریال
                </Descriptions.Item>
                <Descriptions.Item label="نقص مدارک" span={3}>
                  {data.missing_documents.split(',').map(v => (
                    <div>{this.documentTypes[v]}</div>
                  ))}
                </Descriptions.Item>
                <Descriptions.Item label="توضیحات کارشناس" span={3}>
                  {data.expert_description}
                </Descriptions.Item>
              </Descriptions>
              <Descriptions
                bordered
                title="اطلاعات زیاندیده"
                layout="horizontal"
              >
                <Descriptions.Item label="استان">
                  {data.stateName}
                </Descriptions.Item>
                <Descriptions.Item label="شهر">
                  {data.cityName}
                </Descriptions.Item>
                <Descriptions.Item label="شناسه قبض">
                  {data.billId}
                </Descriptions.Item>
                <Descriptions.Item label="کد رهگیری">
                  {data.trackingId}
                </Descriptions.Item>
                <Descriptions.Item label="کد ملی">
                  {data.nationalId}
                </Descriptions.Item>
                <Descriptions.Item label="همراه">
                  {data.mobileNo}
                </Descriptions.Item>
                <Descriptions.Item label="نام">
                  {data.userName}
                </Descriptions.Item>
                <Descriptions.Item label="آدرس">
                  {data.address}
                </Descriptions.Item>
                <Descriptions.Item label="کد پستی">
                  {data.postalCode}
                </Descriptions.Item>
                <Descriptions.Item label="مبلغ خسارت مورد ادعا">
                  <Statistic value={data.amount} /> ریال
                </Descriptions.Item>
                <Descriptions.Item label="تاریخ حادثه">
                  {data.eventDate}
                </Descriptions.Item>
                <Descriptions.Item label="ساعت حادثه">
                  {data.eventTime}
                </Descriptions.Item>
                <Descriptions.Item label="زمان ثبت حادثه	" span={3}>
                  {momentJalaali(data.created_at).format('jYYYY/jM/jD-H:m')}
                </Descriptions.Item>
                <Descriptions.Item label="شماره شبا" span={3}>
                  {data.sheba}
                </Descriptions.Item>
                <Descriptions.Item label="توضیحات" Info>
                  {data.descr}
                </Descriptions.Item>
              </Descriptions>
              <div className="noprint">
                <Descriptions
                  title="مدارک زیاندیده"
                  bordered
                  size="small"
                  layout="vertical"
                  column={6}
                >
                  {data.documents &&
                    data.documents.map((doc, idx) => (
                      <Descriptions.Item
                        label={this.documentTypes[doc.documentTypeId]}
                      >
                        {(doc.fileType === 'image/jpeg' ||
                          doc.fileType === 'image/png' ) && (
                          <a
                            href={`${API_BASE_URL}/admin/api/v1/tavanir/docs/${doc.Id}/?token=${token}`}
                            target="_blank"
                            alt={doc.fileName}
                          >
                            <img
                              src={`${API_BASE_URL}/admin/api/v1/tavanir/docs/${doc.Id}/?token=${token}`}
                              alt={doc.fileName}
                              style={{ maxHeight: 150 }}
                            />
                          </a>
                        )}
                        {doc.fileType == 'application/pdf' && (
                          <center>
                            <a
                              href={`${API_BASE_URL}/admin/api/v1/tavanir/docs/${doc.Id}/?token=${token}`}
                              target="_blank"
                              alt={doc.fileName}
                            >
                              <Icon
                                style={{ fontSize: '2.5em' }}
                                type="file-pdf"
                              />
                            </a>
                          </center>
                        )}
                      </Descriptions.Item>
                    ))}
                </Descriptions>
              </div>
            </Row>
          </>
        )}
        <Divider />
        <div className="noprint">
          {data.id && <Changelog case_id={data.Id} />}
        </div>
        <Divider />
        <div className="noprint">
          {data.id && <SyncLog case_id={data.id} />}
        </div>
        <Divider />
        <Descriptions>
          <Descriptions.Item>
            <small>شناسه داخلی: {data.Id}</small>
          </Descriptions.Item>
          <Descriptions.Item>
            <small>شناسه توانیر: {data.id}</small>
          </Descriptions.Item>
        </Descriptions>
      </div>
    );
  }
}

const Index = () => {
  const router = useRouter();
  let { requestId } = router.query;
  if (requestId.startsWith('id_')) {
    requestId = requestId.substr(3);
    return <RequestItemDisplay pk={requestId} />;
  }
  return <RequestItemDisplay itemID={requestId} />;
};

export default Index;
