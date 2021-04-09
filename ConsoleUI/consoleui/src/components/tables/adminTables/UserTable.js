import React from 'react';
import CommonTable from '../CommonTable'

const columns = [
  { id: 'Login', label: 'Логин', minWidth: 170 },
  { id: 'Surname', label: 'ФИО', minWidth: 100 },
  { id: 'Email', label: 'Почта', minWidth: 100 },
  { id: 'Role', label: 'Роль', minWidth: 100 },
  { id: 'Disabled', label: '', minWidth: 100 },
];

export default function UserTable(props) {
  console.log("render UserTable");

  const rows = props.data;
  return (
    <CommonTable columns={columns} rows={rows}></CommonTable>
  );
}