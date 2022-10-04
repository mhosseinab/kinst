import momentJalaali from 'moment-jalaali';
import PropTypes from 'prop-types';

const CPJalaliDate = props => {
  const { date, format } = props;

  if (!date) return false;
  return momentJalaali(date).format(format);
};

CPJalaliDate.propTypes = {
  format: PropTypes.string,
};

CPJalaliDate.defaultProps = {
  format: 'jYYYY/jMM/jDD HH:mm:ss',
};

export const getValidJalaliOrNull = (d) => {
  if (!momentJalaali(d).jYear()) {
    return '--';
  }
  return momentJalaali(d).format('jYYYY/jM/jD');
}

export const getValidJalaliDateTimeOrNull = (d) => {
  if (!momentJalaali(d).jYear()) {
    return '--';
  }
  return momentJalaali(d).format('jYYYY/jMM/jDD HH:mm:ss');
}

export default CPJalaliDate;
