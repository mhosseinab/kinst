import { Input, Form, Modal, InputNumber, Select } from 'antd';

import {tavanirDamageMaxAmount} from '../../utils/damage'; 

const { Option } = Select;
const { TextArea } = Input;
const AcceptForm = Form.create({ name: 'accept_form' })(
  class extends React.Component {
    
    render() {
      const { visible, onCancel, onCreate, form, data } = this.props;
      const { getFieldDecorator } = form;
      const maxAmount = tavanirDamageMaxAmount[data.compensationTypeId]
      return (
        <Modal
          visible={visible}
          title="تایید خسارت"
          okText="تایید"
          cancelText="انصراف"
          onCancel={onCancel}
          onOk={onCreate}
        >
          <Form layout="vertical">
            <Form.Item label="مبلغ مورد تایید کارشناس">
              {getFieldDecorator('expert_accepted_amount', {
                initialValue: data.expert_accepted_amount,
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
            <Form.Item label="">
              {getFieldDecorator('expert_status')(
                <Input
                  type="text"
                  value="ACCEPTED"
                  style={{ display: 'none' }}
                />,
              )}
            </Form.Item>
            <Form.Item label="وضعیت">
              {getFieldDecorator('expert_status', {
                initialValue: "READY_TO_PAY",
                rules: [
                  {
                    required: true,
                    message: 'لطفا وضعیت پرونده را انتخاب کنید!',
                  },
                ],
              })(
                <Select>
                  <Option value="READY_TO_PAY">آماده پرداخت</Option>
                  <Option value="PAYED">پرداخت شده</Option>
                </Select>,
              )}
            </Form.Item>
            <Form.Item label="توضیحات">
              {getFieldDecorator('expert_description', {
                initialValue: data.expert_description,
                rules: [
                  {
                    required: true,
                    min: 5,
                    message: 'توضیحات کداقل 5 کاراکتر باشد!',
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

export default AcceptForm;
