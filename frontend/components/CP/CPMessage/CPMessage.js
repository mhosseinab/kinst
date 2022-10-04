import { message } from 'antd';
// import s from './CPHelper.scss';

const Message = (text, type = 'info', timeout = 5) => {
  switch (type) {
    case 'success': {
      message.success(text, timeout);
      break;
    }
    case 'warning': {
      message.info(text, timeout);
      break;
    }
    case 'error': {
      message.error(text, timeout);
      break;
    }
    default: {
      message.info(text, timeout);
      break;
    }
  }
};

export default Message;
