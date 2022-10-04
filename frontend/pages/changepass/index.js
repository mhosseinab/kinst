import React, { Component } from 'react';
import Head from 'next/head';
import s from './index.scss';
import { request } from '../../store/request';
import ChangePasswordForm from './edit_form';
import CPMessage from '../../components/CP/CPMessage';
import { API_BASE_URL } from '../../utils/const';
import { getCookie } from '../../utils/cookie';
import { authToken } from '../../const/user';

class Changepass extends Component {
  state = {
    token: getCookie(authToken) || null,
  };

  saveFormCreate = formCreate => {
    this.formCreate = formCreate;
  };

  handleCreateAccept = e => {
    const { form } = this.formCreate.props;
    form.validateFields((err, values) => {
      if (err) {
        return;
      }

      const u = `${API_BASE_URL}/admin/api/v1/user/change_password/`;
      request.setHeader('Authorization', this.state.token);
      request.post(u, values).then(response => {
        if (response.status !== 200) {
          if (response.data.error_msg != '') {
            CPMessage(response.data.error_msg, 'error');
          } else {
            CPMessage('خطایی رخ داده است', 'error');
          }
          return;
        }

        CPMessage('تغییرات با موفقیت اعمال شد', 'success');
        Router.push('/');
      });
    });
  };

  render() {
    return (
      <div className={s.root}>
        <Head>
          <title>تغییر رمز عبور</title>
        </Head>
        <div style={{ marginBottom: 30 }}>تغییر رمز عبور</div>
        <ChangePasswordForm
          wrappedComponentRef={this.saveFormCreate}
          onCancel={this.handleCancelAccept}
          onCreate={this.handleCreateAccept}
        />
      </div>
    );
  }
}

export default Changepass;
