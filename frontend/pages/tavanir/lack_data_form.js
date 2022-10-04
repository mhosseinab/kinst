import React from 'react';
import { Input, Form, Modal, Select } from 'antd';
import { missingDocumentTypes } from '../../utils/damage';

const { Option } = Select;
const { TextArea } = Input;

const LackDataForm = Form.create({})(
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
            <Form.Item label="مدارک ناقص">
              {getFieldDecorator('missing_documents', {
                rules: [
                  { required: true, message: 'لطفا اطلاعات را تکمیل کنید!' },
                ],
              })(
                <Select
                  mode="multiple"
                  allowClear
                  placeholder="نوع مدارک ناقص را انتخاب کنید"
                >
                  {Object.keys(missingDocumentTypes).map(key => (
                    <Option value={key}>{missingDocumentTypes[key]}</Option>
                  ))}
                </Select>,
              )}
            </Form.Item>
            <Form.Item label="وضعیت">
              {getFieldDecorator('expert_status', {
                initialValue: 'INCOMPLETE',
                rules: [
                  { required: true, message: 'لطفا اطلاعات را تکمیل کنید!' },
                ],
              })(
                <Select defaultValue="INCOMPLETE" disabled>
                  <Option value="INCOMPLETE">درخواست ناقص</Option>
                </Select>,
              )}
            </Form.Item>
            <Form.Item label="توضیحات">
              {getFieldDecorator('expert_description', {
                rules: [
                  { required: true, message: 'لطفا اطلاعات را تکمیل کنید!' },
                  { min: 5, message: 'توضیحات ناقص است' },
                ],
              })(<TextArea rows={5} type="textarea" />)}
            </Form.Item>
          </Form>
        </Modal>
      );
    }
  },
);

export default LackDataForm;
