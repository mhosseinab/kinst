import React, { useState } from 'react';
import PropTypes from 'prop-types';
import moment from 'moment';
import momentJalaali from 'moment-jalaali';
// import { SingleDatePicker } from 'react-dates';
import { SingleDatePicker } from 'react-dates-jalali';

import 'react-dates/initialize';
// import 'react-dates/lib/css/_datepicker.css';
import 'react-dates-jalali/lib/css/_datepicker.css';

const CPDatePicker = props => {
  const {
    date,
    id,
    isRTL,
    jalali,
    onDateChange,
    numberOfMonths,
    placeholder,
    orientation,
    disabled,
    required,
    enableOutsideDays,
    horizontalMargin,
    keepOpenOnDateSelect,
    navNext,
    navPrev,
    reopenPickerOnClearDate,
    screenReaderInputMessage,
    showClearDate,
    hideKeyboardShortcutsPanel,
    isOutsideRange,
  } = props;
  const [isFocused, setIsFocused] = useState(false);
  const [selectedDate, setSelectedDate] = useState(null);

  if (jalali === true) {
    moment.locale('fa');
    momentJalaali.locale('fa');
    momentJalaali.loadPersian({
      dialect: 'persian-modern',
      usePersianDigits: false,
    });
  } else {
    moment.locale('en');
  }

  let initialDate = null;
  if (date) {
    initialDate = jalali === true ? momentJalaali(date) : moment(date);
  }
  return (
    <SingleDatePicker
      date={selectedDate || initialDate} // momentPropTypes.momentObj or null
      onDateChange={sDate => {
        if (onDateChange)
          onDateChange(sDate ? `${sDate.toISOString().split('.')[0]}Z` : null); // convert date to standart ISO-8601 format
        setSelectedDate(sDate);
      }} // PropTypes.func.isRequired
      focused={isFocused} // PropTypes.bool
      onFocusChange={({ focused }) => setIsFocused(focused)} // PropTypes.func.isRequired
      id={id || `datepicker_${Math.floor(Math.random() * 1000 + 1)}`}
      anchorDirection={isRTL ? 'right' : 'left'}
      disabled={disabled}
      //   displayFormat={function noRefCheck() {}}
      // displayFormat={
      //   jalali === true
      //     ? () => momentJalaali(selectedDate).format('jYYYY/jMM/jDD')
      //     : undefined
      // }
      enableOutsideDays={enableOutsideDays}
      horizontalMargin={horizontalMargin}
      //   id="date"
      //   initialDate={null}
      //   initialVisibleMonth={null}
      //   isDayBlocked={function noRefCheck() {}}
      //   isDayHighlighted={function noRefCheck() {}}
      isOutsideRange={isOutsideRange}
      isRTL={!!isRTL}
      keepOpenOnDateSelect={keepOpenOnDateSelect}
      //   monthFormat={jalali ? 'jMMMM jYYYY' : 'MMMM YYYY'}
      //   weekDayFormat={jalali ? 'jD' : 'dd'}
      navNext={navNext}
      navPrev={navPrev}
      numberOfMonths={numberOfMonths}
      //   onClose={function noRefCheck() {}}
      //   onNextMonthClick={function noRefCheck() {}}
      //   onPrevMonthClick={function noRefCheck() {}}
      orientation={orientation}
      placeholder={placeholder}
      renderMonth={
        jalali === true
          ? month => momentJalaali(month).format('jMMMM jYYYY')
          : undefined
      }
      reopenPickerOnClearDate={reopenPickerOnClearDate}
      required={required}
      screenReaderInputMessage={screenReaderInputMessage}
      showClearDate={showClearDate}
      hideKeyboardShortcutsPanel={hideKeyboardShortcutsPanel}
    />
  );
};

CPDatePicker.propTypes = {
  id: PropTypes.string,
  date: PropTypes.any,
  jalali: PropTypes.bool,
  isRTL: PropTypes.bool,
  onDateChange: PropTypes.func.isRequired,
  numberOfMonths: PropTypes.number,
  placeholder: PropTypes.string,
  orientation: PropTypes.string,
  disabled: PropTypes.bool,
  required: PropTypes.bool,
  showClearDate: PropTypes.bool,
  enableOutsideDays: PropTypes.bool,
  horizontalMargin: PropTypes.number,
  keepOpenOnDateSelect: PropTypes.bool,
  reopenPickerOnClearDate: PropTypes.bool,
  navNext: PropTypes.any,
  navPrev: PropTypes.any,
  screenReaderInputMessage: PropTypes.string,
  hideKeyboardShortcutsPanel: PropTypes.bool,
  isOutsideRange: PropTypes.func,
};

CPDatePicker.defaultProps = {
  id: null,
  date: null,
  jalali: true,
  isRTL: true,
  numberOfMonths: 1,
  placeholder: 'تقویم',
  orientation: 'horizontal',
  disabled: false,
  required: false,
  enableOutsideDays: false,
  horizontalMargin: 0,
  keepOpenOnDateSelect: false,
  navNext: null,
  navPrev: null,
  reopenPickerOnClearDate: false,
  screenReaderInputMessage: '',
  showClearDate: false,
  hideKeyboardShortcutsPanel: true,
  isOutsideRange: () => false,
};

export default CPDatePicker;
