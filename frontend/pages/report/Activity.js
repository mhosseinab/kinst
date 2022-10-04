import { Component } from 'react';
import { Tooltip, Table } from 'antd';
import Humanize from 'humanize-plus';
// import ReactExport from 'react-data-export';

// const ExcelFile = ReactExport.ExcelFile;
// const ExcelSheet = ReactExport.ExcelFile.ExcelSheet;
// const ExcelColumn = ReactExport.ExcelFile.ExcelColumn;

class Activity extends Component {
  columns = [
    // {
    //   title: 'کد شرکت',
    //   dataIndex: 'code',
    // },
    {
      title: '',
      children: [
        {
          title: 'شرکت توزیع',
          dataIndex: 'company_id',
        },
      ],
    },
    {
      title: 'نوع خسارت',
      children: [
        {
          dataIndex: 'damages.firing_damage.doc_count',
          title: 'آتش سوزی',
        },
        {
          dataIndex: 'damages.instrument_damage.doc_count',
          title: 'تجهیزات',
        },
        {
          dataIndex: 'damages.explosion_damage.doc_count',
          title: 'انفجار',
        },
        {
          dataIndex: 'damages.death_damage.doc_count',
          title: 'فوت',
        },
        {
          dataIndex: 'damages.lack_damage.doc_count',
          title: 'نقص عضو',
        },
      ],
    },
    {
      title: '',
      children: [
        {
          title: 'جمع',
          dataIndex: 'damage_count',
          // sorter: true,
        },
      ],
    },
    {
      title: 'مبلغ خسارت اعلامی',
      children: [
        {
          dataIndex: 'damages.firing_damage.sum',
          title: 'آتش سوزی',
          render: d => {
            return Humanize.intComma(d);
          },
        },
        {
          dataIndex: 'damages.instrument_damage.sum',
          title: 'تجهیزات',
          render: d => {
            return Humanize.intComma(d);
          },
        },
        {
          dataIndex: 'damages.explosion_damage.sum',
          title: 'انفجار',
          render: d => {
            return Humanize.intComma(d);
          },
        },
        {
          dataIndex: 'damages.death_damage.sum',
          title: 'فوت',
          render: d => {
            return Humanize.intComma(d);
          },
        },
        {
          dataIndex: 'damages.lack_damage.sum',
          title: 'نقص عضو',
          render: d => {
            return Humanize.intComma(d);
          },
        },
      ],
    },
    {
      title: '',
      children: [
        {
          title: 'جمع',
          dataIndex: 'sum_all',
          render: d => {
            return Humanize.intComma(d);
          },
        },
      ],
    },
    {
      title: 'پرداخت شده',
      children: [
        {
          dataIndex: 'status.PAYED.doc_count',
          title: 'تعداد',
        },
        {
          dataIndex: 'status.PAYED.sum',
          title: 'مبلغ',
          render: d => {
            return Humanize.intComma(d);
          },
        },
      ],
    },
    {
      title: 'آماده پرداخت',
      children: [
        {
          dataIndex: 'status.READY_TO_PAY.doc_count',
          title: 'تعداد',
        },
        {
          dataIndex: 'status.READY_TO_PAY.sum',
          title: 'مبلغ',
          render: d => {
            return Humanize.intComma(d);
          },
        },
      ],
    },
    {
      title: 'مختومه',
      children: [
        {
          dataIndex: 'status.CLOSED.doc_count',
          title: 'تعداد',
        },
        {
          dataIndex: 'status.CLOSED.sum',
          title: 'مبلغ',
          render: d => {
            return Humanize.intComma(d);
          },
        },
      ],
    },
    {
      title: 'در حال رسیدگی',
      children: [
        {
          dataIndex: 'status.IN_PROGRESS.doc_count',
          title: 'تعداد',
        },
        {
          dataIndex: 'status.IN_PROGRESS.sum',
          title: 'مبلغ',
          render: d => {
            return Humanize.intComma(d);
          },
        },
      ],
    },
  ];

  render() {
    return (
      <>
        <Table
          loading={this.props.loading}
          pagination={false}
          bordered
          dataSource={this.props.data}
          columns={this.columns}
          className="report-table"
          size="small"
        />
        {/* <ExcelFile>
          <ExcelSheet data={this.props.data} name="گزارش عملکرد"></ExcelSheet>
        </ExcelFile> */}
      </>
    );
  }
}

export default Activity;
