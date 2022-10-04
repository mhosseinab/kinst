import React from 'react';
import { Input, Form, Modal, Select } from 'antd';
import { NotCoveredReasonTypes } from '../../utils/damage';

const { Option } = Select;
const { TextArea } = Input;

const RejectForm = Form.create({ name: 'reject_form' })(
  class extends React.Component {
    render() {
      const { visible, onCancel, onCreate, form, data } = this.props;
      const { getFieldDecorator } = form;
      return (
        <Modal
          visible={visible}
          title="رد خسارت"
          okText="رد درخواست"
          cancelText="انصراف"
          onCancel={onCancel}
          okButtonProps={{ type: 'danger' }}
          onOk={onCreate}
        >
          <Form.Item label="علت رد درخواست">
            {getFieldDecorator('not_covered_reason', {
              initialValue: 2,
              rules: [
                { required: true, message: 'لطفا اطلاعات را تکمیل کنید!' },
              ],
            })(
              <Select style={{ width: '100%' }}>
                {Object.keys(NotCoveredReasonTypes).map(key => (
                  <Option value={parseInt(key)}>
                    {NotCoveredReasonTypes[key]}
                  </Option>
                ))}
              </Select>,
            )}
          </Form.Item>
          <Form.Item label="وضعیت">
            {getFieldDecorator('expert_status', {
              initialValue: 'CLOSED',
              rules: [
                {
                  required: true,
                  message: 'لطفا وضعیت پرونده را انتخاب کنید!',
                },
              ],
            })(
              <Select>
                <Option value="CLOSED">مختومه</Option>
                <Option value="SUSPENDED">معوق</Option>
                <Option value="CANCELED_BY_USER">انصراف ذی نفع</Option>
              </Select>,
            )}
          </Form.Item>

          <Form layout="vertical">
            <Form.Item label="توضیحات:">
              {getFieldDecorator('expert_description', {
                rules: [
                  {
                    required: true,
                    min: 5,
                    message: 'لطفا توضیحات خود را وارد نمایید',
                  },
                ],
              })(<TextArea rows={4} type="textarea" />)}
            </Form.Item>
          </Form>
        </Modal>
      );
    }
  },
);

export default RejectForm;
