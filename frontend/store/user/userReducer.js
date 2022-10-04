import {
  POST_LOGIN_REQUEST,
  POST_LOGIN_SUCCESS,
  POST_LOGIN_FAILURE,
} from '../constants';

const initialState = {
  // Post Login Parameters
  postLoginLoading: false,
  postLoginData: null,
  postLoginError: false,
};

export default (state = initialState, action) => {
  switch (action.type) {
    // Post Login Request Cases
    case POST_LOGIN_REQUEST:
      return {
        ...state,
        postLoginLoading: true,
        postLoginData: null,
        postLoginError: false,
      };
    case POST_LOGIN_SUCCESS:
      return {
        ...state,
        postLoginLoading: false,
        postLoginData: action.payload,
        postLoginError: false,
      };
    case POST_LOGIN_FAILURE:
      return {
        ...state,
        postLoginLoading: false,
        postLoginData: null,
        postLoginError: true,
      };

    default:
      return state;
  }
};
