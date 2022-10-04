import { InputNumber, Form, Modal, Select } from "antd";

const { Option } = Select;

const StatusForm = Form.create({ name: 'status_form' })(
  class extends React.Component {
    render() {
      const { visible, onCancel, onCreate, form, data } = this.props;      
      const { getFieldDecorator } = form;
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
                rules: [{ required: true, message: 'لطفا مبلغ مورد تایید را وارد نمایید!' }],
              })(<InputNumber 
                formatter={value => `${value}`.replace(/\B(?=(\d{3})+(?!\d))/g, ',')}
                parser={value => value.replace(/\$\s?|(,*)/g, '')} 
                />
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
          </Form>
        </Modal>
      );
    }
  },
);

export default StatusForm;