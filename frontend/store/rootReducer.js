import { combineReducers } from 'redux';
import { loadingBarReducer as loadingBar } from 'react-redux-loading-bar';
import user from './user/userReducer';
import home from './home/homeReducer';
import login from './login/loginReducer';

export default combineReducers({
  home,
  login,
  loadingBar,
  user,
});
