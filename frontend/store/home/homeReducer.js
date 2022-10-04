import {
  GET_HOME_ERROR,
  GET_HOME_REQUEST,
  GET_HOME_SUCCESS,
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
    case GET_HOME_REQUEST:
      return { ...state, loading: true, data: null, error: false };
    case GET_HOME_SUCCESS:
      return { ...state, loading: false, data: action.payload, error: false };
    case GET_HOME_ERROR:
      return { ...state, loading: false, data: null, error: true };
    default:
      return state;
  }
};
