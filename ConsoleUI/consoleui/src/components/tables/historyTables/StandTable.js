import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import CommonTable from './CommonTable'

const columns = [
    { id: 'EquipName', label: 'Комплекс', minWidth: 170 },
    { id: 'DateTime', label: 'Время', minWidth: 100 },
    { id: 'SettingsDB', label: 'Settings БД', minWidth: 100 },
    { id: 'ObservationsDB', label: 'Observations БД', minWidth: 100 },
    { id: 'Version', label: 'Версия Атлас', minWidth: 100 },
    { id: 'XilibVersion', label: 'Версия xilib', minWidth: 100 },
  ];

export default function StandTable(props) {
  console.log("render StandTable");

  const rows = props.data;
  return (
    <CommonTable columns={columns} rows={rows}></CommonTable>
  );
}