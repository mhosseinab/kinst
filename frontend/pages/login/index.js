import React, { Component } from 'react';
import PropTypes from 'prop-types';
import qs from 'qs';
import Router from 'next/router';
import { connect } from 'react-redux';
import { getHomeAction } from '../../store/home/homeAction';
import { postLoginAction } from '../../store/user/userAction';
import webConfig from '../../webConfig';
import s from './index.scss';
import { Row, Col } from 'antd';
import CPAlert from '../../components/CP/CPAlert';
import CPInput from '../../components/CP/CPInput';
import CPButton from '../../components/CP/CPButton';
import Cookies from 'universal-cookie';

const cookies = new Cookies();

class Login extends Component {
  constructor(props) {
    super(props);
    this.state = {
      username: undefined,
      password: undefined,
    };
  }

  onChange = (name, value) => {
    this.setState({
      [name]: value,
    });
  };

  submitForm = async e => {
    e.preventDefault();
    const { username, password } = this.state;
    const userDTO = {
      username,
      password,
    };

    const response = await this.props.postLoginAction(
      qs.stringify(userDTO, { encode: true }),
    );

    if (response.ok) {
      cookies.set('token', response.data.token);
      Router.replace('/');
      return;
    }
  };

  render() {
    const { loading, error } = this.props;

    const { username, password } = this.state;

    return (
      <Row className={s.root}>
        <Col xs={23} sm={18} md={16} lg={12}>
          <div>
            <div className={s.form}>
              <form method="post">
                <div className={s.imageLogoDiv}>
                  <img
                    alt="logo"
                    src="/static/images/kowsar.png"
                    className={s.imageLogo}
                  />
                </div>
                {/* <h3>ورود به {webConfig.faName}</h3> */}
                {/* <p className={s.description}>
              برای ورود به {webConfig.faName}، نام کاربری و کلمه عبور خود را
              وارد کنید.
            </p> */}
                {error && (
                  <CPAlert
                    showIcon
                    type="error"
                    message="نام کاربری یا رمز عبور صحیح نیست."
                  />
                )}
                <CPInput
                  label="نام کاربری"
                  className={s.userName}
                  direction="ltrInput"
                  name="usernameOrEmail"
                  value={username}
                  onChange={value =>
                    this.onChange('username', value.target.value)
                  }
                  autoFocus
                />
                <CPInput
                  label="رمز عبور"
                  className={s.password}
                  direction="ltrInput"
                  type="password"
                  value={password}
                  name="password"
                  onChange={value =>
                    this.onChange('password', value.target.value)
                  }
                />
                <CPButton
                  className={s.loginBtn}
                  htmlType="submit"
                  onClick={this.submitForm}
                  loading={loading}
                >
                  ورود
                </CPButton>
              </form>
            </div>
          </div>
        </Col>
      </Row>
    );
  }
}

Login.propTypes = {
  data: PropTypes.object,
  error: PropTypes.object,
  loading: PropTypes.bool,
  postLoginAction: PropTypes.func.isRequired,
};

Login.defaultProps = {
  data: null,
  error: null,
  loading: false,
};

const mapState = state => ({
  loading: state.user.postLoginLoading,
  data: state.user.postLoginData,
  error: state.user.postLoginError,
});

const mapDispatch = {
  getHomeAction,
  postLoginAction,
};

export default connect(mapState, mapDispatch)(Login);
