import React from 'react';
// import { withKnobs, select } from '@storybook/addon-knobs';
import CPDatePicker from './CPDatePicker';

export default {
  component: CPDatePicker,
  title: 'Form/Date Picker',
  //   decorators: [withKnobs],
};

export const singleDayPersian = () => <CPDatePicker />;
export const singleDayGregorian = () => <CPDatePicker jalali={false} />;
