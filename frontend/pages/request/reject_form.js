import { Input, Form, Modal, Select } from "antd";
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
          okButtonProps={{type:"danger"}}
          onOk={onCreate}
        >
          <Form layout="vertical">
            <Form.Item label="علت رد کردن پرونده">
              {getFieldDecorator('expert_description',{
                rules: [{ required: true, min: 5, message: 'لطفا توضیحات خود را وارد نمایید' }],
              })(<TextArea rows={4} type="textarea" />)}
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
            
          </Form>
        </Modal>
      );
    }
  },
);

export default RejectForm;