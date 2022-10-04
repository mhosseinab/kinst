import React from 'react';
import { Descriptions } from 'antd';

const Items = ({ children, object, key, ...props }) => {
  console.log(object[key]);

  return null;
};

export default Items;
