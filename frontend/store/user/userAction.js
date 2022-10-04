import Cookies from 'universal-cookie';
import Router from 'next/router';
import {
  POST_LOGIN_REQUEST,
  POST_LOGIN_SUCCESS,
  POST_LOGIN_FAILURE,
  USER_LOGOUT,
} from '../constants';
import { request, loginConfig } from '../request';
import postLoginService from './userService';

/**
 * Keep user data after otp-login
 * access_token, uuid
 */
const setLoginCookiesAndHeaders = data => {
  const cookies = new Cookies();

  request.setHeader('Authorization', `Bearer ${data?.access_token}`);

  const expires = new Date(new Date().getTime() + data.expires_in * 1000);

  cookies.set('accessToken', data?.access_token, { path: '/', expires });
  cookies.set('refreshToken', data?.refresh_token, { path: '/', expires });
  cookies.set('deviceId', data?.device_id, { path: '/', expires });
  cookies.set('uuid', data?.uuid, { path: '/', expires });
};

export const postLoginAction = body => {
  return async dispatch => {
    dispatch({ type: POST_LOGIN_REQUEST });

    const response = await request.post(postLoginService(), body, {
      headers: loginConfig,
    });

    if (response.ok) {
      setLoginCookiesAndHeaders(response.data);
      dispatch({ type: POST_LOGIN_SUCCESS, payload: response.data });
      return response;
    }

    dispatch({ type: POST_LOGIN_FAILURE, payload: response.data });
    return response;
  };
};

/**
 * Immediate logout without reload page.
 * Browser ONLY.
 */
export function logOut() {
  return async dispatch => {
    const cookies = new Cookies();
    const allCookies = cookies.getAll();
    const array = Object.keys(allCookies);
    array.forEach(item => cookies.remove(item));
    await dispatch({ type: USER_LOGOUT });
    Router.replace('/');
    return true;
  };
}
