import React, { useState, useEffect } from 'react';
import { Table } from 'antd';

import { getCookie } from '../../utils/cookie';
import { request } from '../../store/request';
import { API_BASE_URL } from '../../utils/const';

const columns = [
  {
    title: 'ID',
    dataIndex: 'ID',
    key: 'iD',
    sorter: true,
    defaultSortOrder: 'descend',
    render: id => {
      return id;
    },
  },
  {
    title: 'CaseID',
    dataIndex: 'CaseID',
    key: 'CaseID',
  },
  {
    title: 'CreatedAt',
    dataIndex: 'CreatedAt',
    key: 'CreatedAt',
  },
  //   {
  //     title: 'IsDone',
  //     dataIndex: 'IsDone',
  //     key: 'IsDone',
  //     render: d => {
  //       return d ? '✔' : '❌';
  //     },
  //   },
  {
    title: 'Success',
    dataIndex: 'Success',
    key: 'Success',
    render: d => {
      return d ? '✔' : '❌';
    },
  },
  {
    title: 'NewStatus',
    dataIndex: 'NewStatus',
    key: 'NewStatus',
  },
  {
    title: 'Note',
    dataIndex: 'Note',
    key: 'Note',
    render: (text, row) => {
      return row.Success ? '' : (
        text.toLowerCase().includes('invalid state')
        ? 'وضعیت پرونده اجازه این تغییر را نمیدهد'
        : text
      )
    },
  },
];

const SyncLog = case_id => {
  console.log('case_id', case_id);
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const token = getCookie('token') || null;
    setLoading(true);
    request.setHeader('Authorization', token);
    request
      .get(`${API_BASE_URL}/admin/api/v1/tavanir/sync_queue/`, {
        case_id: [case_id.case_id],
      })
      .then(response => {
        setLoading(false);

        if (response.status !== 200) {
          return;
        }
        setData(response.data.objects);
      });
  }, [case_id]);
  return (
    <div>
      <h3>درخواست های همگام سازی با توانیر</h3>
      <Table
        dataSource={data}
        columns={columns}
        loading={loading}
        pagination={false}
      />
    </div>
  );
};

export default SyncLog;
