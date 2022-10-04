import { Input, Form, Modal, InputNumber, Select } from "antd";

const { Option } = Select;
const { TextArea } = Input;

const AcceptForm = Form.create({ name: 'accept_form' })(
  class extends React.Component {
    render() {
      const { visible, onCancel, onCreate, form, data } = this.props;
      const { getFieldDecorator } = form;
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
                rules: [{ required: true, message: 'لطفا مبلغ مورد تایید را وارد نمایید!' }],
              })(<InputNumber 
                formatter={value => `${value}`.replace(/\B(?=(\d{3})+(?!\d))/g, ',')}
                parser={value => value.replace(/\$\s?|(,*)/g, '')} 
                />
              )}
            </Form.Item>
            <Form.Item label="">
            {getFieldDecorator('expert_status')(
              <Input type="text" value="ACCEPTED" style={{ display: 'none' }} />
            )}
            </Form.Item>
            <Form.Item label="وضعیت">
              {getFieldDecorator('status',{
                initialValue: data.status,
                rules: [{ required: true, message: 'لطفا وضعیت پرونده را انتخاب کنید!' }],
              })(
                <Select>
                  <Option value="IN_PROGRESS">جاری</Option>
                  <Option value="CLOSED">مختومه</Option>
                  <Option value="SUSPENDED">معوق</Option>
                  <Option value="INCOMPLETE">درخواست ناقص</Option>
                  <Option value="READY_TO_PAY">آماده پرداخت</Option>
                  <Option value="PAYED">پرداخت شده</Option>
                  <Option value="CANCELED_BY_USER">انصراف ذی نفع</Option>
                </Select>
              )}
            </Form.Item>
            <Form.Item label="توضیحات">
              {getFieldDecorator('expert_description',{
                initialValue: data.expert_description,
                rules: [
                  {
                    required: true,
                    min: 5,
                    message: 'توضیحات کداقل 5 کاراکتر باشد!',
                  },
                ],
              })(
              <TextArea rows={4} type="textarea" />)}
            </Form.Item>

          </Form>
        </Modal>
      );
    }
  },
);

export default AcceptForm;