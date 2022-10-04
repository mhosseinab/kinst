import React from 'react';
import { Input, Form, Modal, Select } from 'antd';
import { missingDocumentTypes } from '../../utils/damage';

const { Option } = Select;
const { TextArea } = Input;

const InPersonForm = Form.create({})(
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
            <Form.Item label="توضیحات">
              {getFieldDecorator('expert_description', {
                rules: [
                  { required: true, message: 'لطفا اطلاعات را تکمیل کنید!' },
                  { min: 5, message: 'توضیحات ناقص است' },
                ],
              })(<TextArea rows={5} type="textarea" />)}
            </Form.Item>
            <Form.Item label="وضعیت">
              {getFieldDecorator('expert_status', {
                initialValue: 'Need_Visit_InPerson',
                rules: [
                  { required: true, message: 'لطفا اطلاعات را تکمیل کنید!' },
                ],
              })(
                <Select defaultValue="Need_Visit_InPerson" disabled>
                  <Option value="Need_Visit_InPerson">
                    نیاز به مراجعه حضوری
                  </Option>
                </Select>,
              )}
            </Form.Item>
          </Form>
        </Modal>
      );
    }
  },
);

export default InPersonForm;
