import React, { Fragment } from 'react';
import { Row, Col, Select, Button, Icon } from 'antd';

import s from '../Filters/filters.scss';

const { Option } = Select;

const Types = ({ onChange, location_type }) => {
  return (
    <Fragment>
      <p className={s.title}>نوع گزارشات</p>
      <Row gutter={[6, 6]} className={s.root}>
        <Col lg={6} md={12} sm={12} xs={24} className={s.selectAndButton}>
          <Select
            placeholder="نوع گزارشات"
            className={s.select}
            value={location_type}
            onChange={e => {
              onChange(e, 'location_type');
            }}
          >
            <Option value="actvity">عملکرد</Option>
            <Option value="witdarw">پرداختی</Option>
            <Option value="all">کل پرداختی ها</Option>
          </Select>
          {location_type && (
            <Button
              type="danger"
              onClick={() => onChange(undefined, 'location_type')}
            >
              <Icon type="close" />
              پاک کن
            </Button>
          )}
        </Col>
      </Row>
    </Fragment>
  );
};

export default Types;
