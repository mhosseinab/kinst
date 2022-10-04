import { Component } from 'react';
import { Collapse, Button, InputNumber, Icon } from 'antd';

import {
  Query,
  Builder,
  BasicConfig,
  Utils as QbUtils,
} from 'react-awesome-query-builder';

// import 'react-awesome-query-builder/css/styles.scss';
import './AdvanceQuery.css';
import Humanize from 'humanize-plus';
// import faIR from 'antd/es/locale/fa_IR';

const { Panel } = Collapse;

const localeSettings = {
  locale: {
    short: 'fa',
    full: 'fa_IR',
  },
  valueLabel: 'مقدار',
  valuePlaceholder: 'مقدار',
  fieldLabel: 'فیلد',
  operatorLabel: 'اپراتور',
  fieldPlaceholder: 'انتخاب فیلد',
  operatorPlaceholder: 'انتخاب اوپراتو',
  deleteLabel: null,
  addGroupLabel: 'اضافه کردن گروه',
  addRuleLabel: 'اضافه کردن شرط',
  delGroupLabel: null,
  valueSourcesPopupTitle: 'Select value source',
  removeRuleConfirmOptions: {
    title: 'از حذف قانون مطمئنید ؟',
    okText: 'بله',
    okType: 'danger',
  },
  removeGroupConfirmOptions: {
    title: 'از حذف گروه مطمئنید ؟',
    okText: 'بله',
    okType: 'danger',
  },
};

// You need to provide your own config. See below 'Config format'
const MakeConfig = (amount_field_name) => {
  console.log("amount_field_name", amount_field_name)
  return {
    ...BasicConfig,
    operators: {
      ...BasicConfig.operators,
      between: {
        ...BasicConfig.operators.between,
        label: 'محدوده',
        valueLabels: [
          { label: 'از', placeholder: 'از مبلغ' },
          { label: 'تا', placeholder: 'تا مبلغ' },
        ],
        textSeparators: ['از', 'تا'],
      },
      equal: {
        ...BasicConfig.operators.equal,
        label: 'مساوی',
      },
    },
    widgets: {
      ...BasicConfig.widgets,
      number: {
        ...BasicConfig.widgets.number,
        customProps: {
          formatter: value => {
            return `${value}`.replace(/\B(?=(\d{3})+(?!\d))/g, ',');
          },
          parser: value => {
            return value.replace(/\$\s?|(,*)/g, '');
          },
        },
      },
    },
    settings: {
      ...BasicConfig.settings,
      ...localeSettings,
      showNot: false,
    },
    fields: {
      [amount_field_name]: {
        label: 'خسارت مورد ادعا',
        type: 'number',
        valueSources: ['value'],
        fieldSettings: {
          min: 0,
        },
        operators: ['between', 'equal'],
      },
    },
  };
};

// You can load query value from your backend storage (for saving see `Query.onChange()`)
const queryValue = { id: QbUtils.uuid(), type: 'group' };

class AdvanceQuery extends Component {
  

  constructor(props) {
    super(props);
    console.log(this.props);
    
    this.config = MakeConfig(props.amount_field_name)
    
    this.state = {
      tree: QbUtils.checkTree(
        QbUtils.loadTree(queryValue),
        this.config,
      ),
    };
  }

  render() {
    let config = this.config;
    return (
      <div style={{ marginBottom: 15 }}>
        <Collapse style={{ marginBottom: 20 }}>
          <Panel header="جستجو پیشرفته">
            <Query
              {...config}
              value={this.state.tree}
              onChange={this.onChange}
              renderBuilder={this.renderBuilder}
            />
            <Button onClick={this.props.onSubmit} type="primary">
              جستجو
            </Button>
            <Button type="error" onClick={this.onReset}>
              پاک کردن
            </Button>
          </Panel>
        </Collapse>
      </div>
    );
  }

  renderBuilder = props => {
    return (
      <div className="query-builder-container">
        <div className="query-builder qb-lite">
          <Builder {...props} />
        </div>
      </div>
    );
  };

  onChange = (immutableTree, config) => {
    this.setState({ tree: immutableTree, config });
    if (typeof this.props.onChange === 'function') {
      this.props.onChange(QbUtils.sqlFormat(immutableTree, config));
    }
  };

  onReset = () => {
    this.setState({
      tree: QbUtils.checkTree(QbUtils.loadTree(queryValue), config),
      config,
    });
    if (typeof this.props.onReset === 'function') {
      this.props.onReset();
    }
  };
}

export default AdvanceQuery;
