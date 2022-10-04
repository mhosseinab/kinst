import React from 'react';
import PropTypes from 'prop-types';
import cs from 'classnames';
import { Input } from 'antd';
import s from './CPInput.scss';

class CPInput extends React.Component {
  static propTypes = {
    label: PropTypes.node,
    placeholder: PropTypes.string,
    type: PropTypes.string,
    addonBefore: PropTypes.string,
    addonAfter: PropTypes.string,
    fullWidth: PropTypes.bool,
    inline: PropTypes.bool,
    size: PropTypes.string,
    className: PropTypes.string,
    direction: PropTypes.string,
    value: PropTypes.string,
    prefix: PropTypes.node,
    suffix: PropTypes.node,
    onChange: PropTypes.func,
    ref: PropTypes.func,
    hintText: PropTypes.string,
    MDInput: PropTypes.bool,
    disabled: PropTypes.bool,
    onBlur: PropTypes.func,
    name: PropTypes.string,
    maxLength: PropTypes.number,
  };

  static defaultProps = {
    placeholder: null,
    type: 'text',
    addonBefore: null,
    addonAfter: null,
    fullWidth: true,
    inline: false,
    size: 'default',
    className: '',
    prefix: '',
    suffix: '',
    value: null,
    label: '',
    direction: null,
    ref: null,
    onChange: () => {},
    hintText: '',
    MDInput: false,
    disabled: false,
    name: null,
    onBlur: () => {},
    maxLength: null,
  };

  render() {
    const {
      size,
      label,
      placeholder,
      type,
      addonBefore,
      addonAfter,
      fullWidth,
      inline,
      onChange,
      className,
      value,
      prefix,
      suffix,
      direction,
      ref,
      hintText,
      MDInput,
      disabled,
      name,
      onBlur,
      maxLength,
    } = this.props;
    return (
      <React.Fragment>
        {MDInput ? (
          <div
            className={cs(s.formGroup, direction, className, {
              fullWidth,
            })}
            ref={ref}
          >
            <input
              disabled={disabled}
              name={name}
              type={type}
              onChange={onChange}
              placeholder={hintText}
              onBlur={onBlur}
              value={value}
              maxLength={maxLength}
            />
            <span className={s.label}>{label}</span>
            <i className={s.bar} />
          </div>
        ) : (
          <div
            className={cs(
              'form-group',
              s.input,
              direction,
              { inline },
              className,
              {
                fullWidth,
              },
            )}
            ref={ref}
          >
            {label && <span className="controlLabel">{label}</span>}
            <Input
              type={type}
              required={false}
              name={name}
              onChange={onChange}
              placeholder={placeholder}
              size={size}
              value={value}
              prefix={prefix}
              suffix={suffix}
              addonBefore={addonBefore}
              addonAfter={addonAfter}
              maxLength={maxLength}
              disabled={disabled}
              onBlur={onBlur}
            />
            <i className="bar" />
          </div>
        )}
      </React.Fragment>
    );
  }
}

export default CPInput;
