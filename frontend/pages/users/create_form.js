import { Input, Form, Modal, Select } from 'antd';
import city from '../../utils/city.json';

const { Option } = Select;

const CreateForm = Form.create({ name: 'create_form' })(
  class extends React.Component {
    render() {
      const { visible, onCancel, onCreate, form } = this.props;
      const { getFieldDecorator } = form;

      return (
        <Modal
          visible={visible}
          title="کاربر جدید"
          okText="ایجاد"
          cancelText="انصراف"
          onCancel={onCancel}
          onOk={onCreate}
        >
          <Form layout="vertical">
            <Form.Item label="نام کاربری">
              {getFieldDecorator('username', {
                rules: [
                  { required: true, message: 'لطفا نام کاربری وارد کنید!' },
                ],
              })(<Input className="input-ltr english-font" />)}
            </Form.Item>

            <Form.Item label="کلمه عبور">
              {getFieldDecorator('password', {
                rules: [
                  { required: true, message: 'لطفا کلمه عبور وارد کنید!' },
                ],
              })(<Input type="password" className="input-ltr" />)}
            </Form.Item>

            <Form.Item label="سطح دسترسی">
              {getFieldDecorator('role', {
                rules: [
                  { required: true, message: 'لطفا سطح دسترسی را انتخاب کنید' },
                ],
              })(
                <Select>
                  <Option value="ADMIN">مدیر سیستم</Option>
                  <Option value="BRANCH">شعب بیمه</Option>
                  <Option value="TAVANIR">توانیر</Option>
                  <Option value="REPORTER">گزارش گیرنده</Option>
                </Select>,
              )}
            </Form.Item>

            <Form.Item label="استان">
              {getFieldDecorator(
                'province',
                {},
              )(
                <Select mode="multiple">
                  <Option key="011">011 - توزيع نيروي برق تبريز</Option>
                  <Option key="012">
                    012 - توزيع نيروي برق آذربايجان شرقي
                  </Option>
                  <Option key="013">
                    013 - توزيع نيروي برق آذربايجان غربي
                  </Option>
                  <Option key="014">014 - توزيع نيروي برق اردبيل</Option>
                  <Option key="021">
                    021 - توزيع نيروي برق شهرستان اصفهان
                  </Option>
                  <Option key="022">022 - توزيع نيروي برق استان اصفهان</Option>
                  <Option key="023">
                    023 - توزيع نيروي برق چهارمحال بختياري
                  </Option>
                  <Option key="031">031 - توزيع نيروي برق استان مرکزي</Option>
                  <Option key="032">032 - توزيع نيروي برق استان همدان</Option>
                  <Option key="033">033 - توزيع نيروي برق استان لرستان</Option>
                  <Option key="041">041 - توزيع نيروي برق تهران بزرگ</Option>
                  <Option key="042">042 - توزيع نيروي برق استان تهران</Option>
                  <Option key="043">043 - توزيع نيروي برق استان البرز</Option>
                  <Option key="044">044 - توزيع نيروي برق استان قم</Option>
                  <Option key="051">051 - توزيع نيروي برق شهرستان مشهد</Option>
                  <Option key="052">
                    052 - توزيع نيروي برق استان خراسان رضوي
                  </Option>
                  <Option key="053">
                    053 - توزيع نيروي برق استان خراسان شمالي
                  </Option>
                  <Option key="054">
                    054 - توزيع نيروي برق استان خراسان جنوبي
                  </Option>
                  <Option key="061">061 - توزيع نيروي برق شهرستان اهواز</Option>
                  <Option key="062">062 - توزيع نيروي برق استان خوزستان</Option>
                  <Option key="063">
                    063 - توزيع نيروي برق کهکيلويه و بويراحمد
                  </Option>
                  <Option key="071">071 - توزيع نيروي برق استان زنجان</Option>
                  <Option key="072">072 - توزيع نيروي برق استان قزوين</Option>
                  <Option key="081">081 - توزيع نيروي برق استان سمنان</Option>
                  <Option key="091">
                    091 - توزيع نيروي برق استان سيستان و بلوچستان
                  </Option>
                  <Option key="101">
                    101 - توزيع نيروي برق استان کرمانشاه
                  </Option>
                  <Option key="102">102 - توزيع نيروي برق استان کردستان</Option>
                  <Option key="103">103 - توزيع نيروي برق استان ايلام</Option>
                  <Option key="111">111 - توزيع نيروي برق شهرستان شيراز</Option>
                  <Option key="112">112 - توزيع نيروي برق استان فارس</Option>
                  <Option key="113">113 - توزيع نيروي برق استان بوشهر</Option>
                  <Option key="121">
                    121 - توزيع نيروي برق شمال استان کرمان
                  </Option>
                  <Option key="122">
                    122 - توزيع نيروي برق جنوب استان کرمان
                  </Option>
                  <Option key="131">131 - توزيع نيروي برق استان گيلان</Option>
                  <Option key="141">
                    141 - توزيع نيروي برق استان مازندران
                  </Option>
                  <Option key="142">
                    142 - توزيع نيروي برق غرب استان مازندران
                  </Option>
                  <Option key="143">143 - توزيع نيروي برق استان گلستان</Option>
                  <Option key="151">151 - توزيع نيروي برق استان هرمزگان</Option>
                  <Option key="161">161 - توزيع نيروي برق استان يزد</Option>
                </Select>,
              )}
            </Form.Item>
          </Form>
        </Modal>
      );
    }
  },
);

export default CreateForm;
