import { Input, Form, Modal, Select, Button } from "antd";
import city from '../../utils/city.json';

const { Option } = Select;

const ChangePasswordForm = Form.create({ name: 'change_password_form' })(
  class extends React.Component {

    render() {
      const { onCreate, form } = this.props;
      const { getFieldDecorator } = form;

      return (
        <Form layout="vertical">
          <Form.Item label="کلمه عبور">
            {getFieldDecorator('password', {
              rules: [
                {required: true, message: 'لطفا رمز عبور وارد کنید.' },
                { min: 6, message: "رمز عبور باید حداقل ۶ کاراکتر باشد." },
                { max: 24, message: "رمز عبور باید حداکثر 24 کاراکتر باشد." },
              ],
            })(<Input type="password" className="input-ltr" />)}
          </Form.Item>
          <Button type="primary" onClick={onCreate}>ذخیره</Button>
        </Form>
      );
    }
  },
);

export default ChangePasswordForm;