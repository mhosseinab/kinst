import React from 'react';
import { Input, Form, Modal, Select } from "antd";

const { Option } = Select;
const { TextArea } = Input;

const LackDataForm = Form.create({ name: 'lack_data_form' })(
  class extends React.Component {
    render() {
      const { visible, onCancel, onCreate, form, data } = this.props;
      const { getFieldDecorator } = form;
      return (
        <Modal
          visible={visible}
          title="اعلام نقص مدرک"
          okText="ثبت"
          cancelText="انصراف"
          onCancel={onCancel}
          okButtonProps={{ type: 'primary' }}
          onOk={onCreate}
        >
          <Form layout="vertical">
            <Form.Item label="نواقص">
              {getFieldDecorator('lack_data_description', {
                // eslint-disable-next-line react/prop-types
                initialValue: data.lack_data_description,
                rules: [
                  { required: true, message: 'لطفا اطلاعات را تکمیل کنید!' },
                ],
              })(
                <TextArea
                  value="salam"
                  defaultValue="salam"
                  rows={5}
                  type="textarea"
                />,
              )}
            </Form.Item>

            <Form.Item label="وضعیت">
              {getFieldDecorator('status', {
                initialValue: data.status,
                rules: [
                  {
                    required: true,
                    message: 'لطفا وضعیت پرونده را انتخاب کنید!',
                  },
                ],
              })(
                <Select defaultValue={data.status}>
                  <Option value="IN_PROGRESS">جاری</Option>
                  <Option value="CLOSED">مختومه</Option>
                  <Option value="SUSPENDED">معوق</Option>
                  <Option value="INCOMPLETE">درخواست ناقص</Option>
                  <Option value="READY_TO_PAY">آماده پرداخت</Option>
                  <Option value="PAYED">پرداخت شده</Option>
                  <Option value="CANCELED_BY_USER">انصراف ذی نفع</Option>
                </Select>,
              )}
            </Form.Item>
          </Form>
        </Modal>
      );
    }
  },
);

export default LackDataForm;
