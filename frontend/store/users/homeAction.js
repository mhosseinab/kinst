import {
  GET_HOME_ERROR,
  GET_HOME_REQUEST,
  GET_HOME_SUCCESS,
} from '../constants';
import { request } from '../request';
import { getHomeService } from './homeService';

export const getHomeAction = () => {
  return async dispatch => {
    dispatch({ type: GET_HOME_REQUEST });

    const response = await request.get(getHomeService);

    if (response.ok) {
      dispatch({ type: GET_HOME_SUCCESS, payload: response.data });
      return response.data;
    }

    dispatch({ type: GET_HOME_ERROR });
    return false;
  };
};

export default getHomeAction;
