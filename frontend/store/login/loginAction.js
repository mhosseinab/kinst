import {
  GET_LOGIN_ERROR,
  GET_LOGIN_REQUEST,
  GET_LOGIN_SUCCESS,
} from '../constants';
import { request } from '../request';
import { getLoginService } from './loginService';

export const getLoginAction = () => {
  return async dispatch => {
    dispatch({ type: GET_LOGIN_REQUEST });

    const response = await request.post(getLoginService);

    if (response.ok) {
      dispatch({ type: GET_LOGIN_SUCCESS, payload: response.data });
      return response.data;
    }

    dispatch({ type: GET_LOGIN_ERROR });
    return false;
  };
};

export default getLoginAction;
