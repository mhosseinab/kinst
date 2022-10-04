import {
  GET_LOGIN_ERROR,
  GET_LOGIN_REQUEST,
  GET_LOGIN_SUCCESS,
} from '../constants';

const initialState = {
  // home parameters
  loading: false,
  data: null,
  error: false,
};

export default (state = initialState, action) => {
  switch (action.type) {
    // home requests
    case GET_LOGIN_REQUEST:
      return { ...state, loading: true, data: null, error: false };
    case GET_LOGIN_SUCCESS:
      return { ...state, loading: false, data: action.payload, error: false };
    case GET_LOGIN_ERROR:
      return { ...state, loading: false, data: null, error: true };

    default:
      return state;
  }
};
