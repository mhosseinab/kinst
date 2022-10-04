import React from 'react';
import { InputNumber, Form, Modal, Select } from 'antd';
import {
  tavanirDamageMaxAmount,
  expertStatusChoices,
} from '../../utils/damage';

const { Option } = Select;

const StatusForm = Form.create({ name: 'status_form' })(
  class extends React.Component {
    render() {
      const { visible, onCancel, onCreate, form, data } = this.props;
      const { getFieldDecorator } = form;
      const maxAmount = tavanirDamageMaxAmount[data.compensationTypeId]
      return (
        <Modal
          visible={visible}
          title="وضعیت پرونده"
          okText="تایید"
          cancelText="انصراف"
          onCancel={onCancel}
          onOk={onCreate}
        >
          <Form layout="vertical">
            <Form.Item label="مبلغ مورد تایید">
              {getFieldDecorator('accepted_amount', {
                initialValue: data.accepted_amount,
                rules: [
                  {
                    message: `برای این نوع خسارت حداکثر ${maxAmount} ریال مجاز است`,
                    validator: (rule, value, cb) => value < maxAmount,
                  },
                ],
              })(
                <InputNumber
                  formatter={value =>
                    `${value}`.replace(/\B(?=(\d{3})+(?!\d))/g, ',')
                  }
                  parser={value => value.replace(/\$\s?|(,*)/g, '')}
                />,
              )}
            </Form.Item>
            <Form.Item label="وضعیت">
              {getFieldDecorator('expert_status', {
                initialValue: data.status,
                rules: [
                  {
                    required: true,
                    message: 'لطفا وضعیت پرونده را انتخاب کنید!',
                  },
                ],
              })(
                <Select>
                  {Object.keys(expertStatusChoices).map(key => (
                  <Option value={key}>{expertStatusChoices[key]}</Option>
                ))}
                </Select>,
              )}
            </Form.Item>
          </Form>
        </Modal>
      );
    }
  },
);

export default StatusForm;
