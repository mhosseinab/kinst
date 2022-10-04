import React, { Fragment, useEffect, useState } from 'react';
import { Select, Row, Col, Button, Icon } from 'antd';
import s from './filters.scss';
import cities from '../../utils/city.json';
import {
  statusChoices,
  expertStatusChoices,
  damageTypes,
} from '../../utils/damage';
import CPDatePicker from '../CP/CPDatePicker/CPDatePicker';
import CPMessage from '../CP/CPMessage/CPMessage';

const { Option } = Select;

const Filters = ({
  onChange,
  expert,
  province,
  status,
  city,
  damage_type,
  from,
  to,
}) => {
  const citiesMap = Object.keys(cities);
  const statusValues = Object.values(statusChoices);
  const statusKeys = Object.keys(statusChoices);
  const expertValues = Object.values(expertStatusChoices);
  const expertKeys = Object.keys(expertStatusChoices);
  const damageTypeValues = Object.values(damageTypes);
  const damageTypeKeys = Object.keys(damageTypes);

  const [dates, setDates] = useState({ startDate: '', endDate: '' });
  const [focusedInput, setFocus] = useState('START_DATE');

  const handleRemove = key => {
    onChange(undefined, key);
  };

  const handleToDateChange = value => {
    if (value > from) {
      // console.log(value > from, value, from);
      onChange(value, 'to');
    } else {
      CPMessage('تاریخ وارد شده نمی تواند کوچک تر از تاریخ شروع باشد', 'error');
    }
  };

  return (
    <Fragment>
      {/* <p className={s.title}>فیلتر گزارشات</p> */}
      <Row className={s.root} gutter={[6, 6]}>
        {province && (
          <Col xs={24} md={12} sm={12} lg={6} className={s.state}>
            <p className={s.title}>استان</p>
            <div className={s.selectAndButton}>
              <Select
                showSearch
                className={s.select}
                value={province}
                placeholder="شهر را انتخاب کنید"
                optionFilterProp="children"
                onChange={v => {
                  onChange(v, 'province');
                  onChange(undefined, 'city');
                }}
                filterOption={(input, option) =>
                  option.props.children
                    .toLowerCase()
                    .indexOf(input.toLowerCase()) >= 0
                }
              >
                {citiesMap.map(city => (
                  <Option key={city}>{city}</Option>
                ))}
              </Select>
              {province ? (
                <Button
                  type="danger"
                  onClick={() => {
                    handleRemove('province');
                    handleRemove('city');
                  }}
                >
                  <Icon type="close" />
                  پاک کن
                </Button>
              ) : null}
            </div>
            {province ? (
              <div className={s.state}>
                <p className={s.title}>شهر</p>
                <div className={s.selectAndButton}>
                  <Select
                    showSearch
                    className={s.select}
                    value={city}
                    placeholder="شهر را انتخاب کنید"
                    optionFilterProp="children"
                    onChange={v => onChange(v, 'city')}
                    filterOption={(input, option) =>
                      option.props.children
                        .toLowerCase()
                        .indexOf(input.toLowerCase()) >= 0
                    }
                  >
                    {Object.values(cities[province]).map(city => (
                      <Option key={city}>{city}</Option>
                    ))}
                  </Select>
                  {city ? (
                    <Button type="danger" onClick={() => handleRemove('city')}>
                      <Icon type="close" />
                      پاک کن
                    </Button>
                  ) : null}
                </div>
              </div>
            ) : null}
          </Col>
        )}
        <Col xs={24} md={12} sm={12} lg={6} className={s.state}>
          <p className={s.title}>وضعیت پرونده</p>
          <div className={s.selectAndButton}>
            <Select
              showSearch
              className={s.select}
              name="status"
              placeholder="وضعیت پرونده را انتخاب کنید"
              optionFilterProp="children"
              onChange={v => onChange(v, 'status')}
              value={status}
              filterOption={(input, option) =>
                option.props.children
                  .toLowerCase()
                  .indexOf(input.toLowerCase()) >= 0
              }
            >
              {statusValues.map((city, key) => (
                <Option key={statusKeys[key]}>{city}</Option>
              ))}
            </Select>
            {status ? (
              <Button type="danger" onClick={() => handleRemove('status')}>
                <Icon type="close" />
                پاک کن
              </Button>
            ) : null}
          </div>
        </Col>
        <Col xs={24} md={12} sm={12} lg={6} className={s.state}>
          <p className={s.title}>وضعیت کارشناسی</p>
          <div className={s.selectAndButton}>
            <Select
              showSearch
              className={s.select}
              placeholder="وضعیت کارشناسی را انتخاب کنید"
              optionFilterProp="children"
              onChange={v => onChange(v, 'expert')}
              value={expert}
            >
              {expertValues.map((expert, key) => (
                <Option key={expertKeys[key]}>{expert}</Option>
              ))}
            </Select>
            {expert ? (
              <Button type="danger" onClick={() => handleRemove('expert')}>
                <Icon type="close" />
                پاک کن
              </Button>
            ) : null}
          </div>
        </Col>
        <Col xs={24} md={12} sm={12} lg={6} className={s.state}>
          <p className={s.title}>نوع خسارت</p>
          <div className={s.selectAndButton}>
            <Select
              showSearch
              className={s.select}
              placeholder="نوع خسارت"
              optionFilterProp="children"
              onChange={v => onChange(v, 'damage_type')}
              value={damage_type}
            >
              {damageTypeValues.map((damageType, key) => (
                <Option key={damageTypeKeys[key]}>{damageType}</Option>
              ))}
            </Select>
            {damage_type ? (
              <Button type="danger" onClick={() => handleRemove('damage_type')}>
                <Icon type="close" />
                پاک کن
              </Button>
            ) : null}
          </div>
        </Col>
        <Col xs={24} md={24} sm={24} lg={24} className={s.state}>
          <p className={s.title}>تاریخ گزارشات</p>
          <div className={s.selectAndButton} style={{ alignItems: 'center' }}>
            <p style={{ margin: '0 10px' }}>از تاریخ</p>
            <CPDatePicker
              placeholder="از تاریخ"
              value={from}
              onDateChange={value => onChange(value, 'from')}
            />
            {from && (
              <>
                <p style={{ margin: '0 10px' }}>تا تاریخ</p>
                <CPDatePicker
                  placeholder="تا تاریخ"
                  value={to}
                  onDateChange={handleToDateChange}
                />
              </>
            )}

            {from && to ? (
              <Button
                type="danger"
                onClick={() => {
                  handleRemove('from');
                  handleRemove('to');
                }}
              >
                <Icon type="close" />
                پاک کن
              </Button>
            ) : null}
          </div>
        </Col>
      </Row>
    </Fragment>
  );
};

export default Filters;
