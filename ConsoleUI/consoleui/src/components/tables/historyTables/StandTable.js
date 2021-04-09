import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import CommonTable from '../CommonTable'

const columns = [
    { id: 'EquipName', label: 'Комплекс', minWidth: 170 },
    { id: 'DateTime', label: 'Время', minWidth: 100 },
    { id: 'State', label: 'Состояние', minWidth: 100 },
    { id: 'RasterState', label: 'Растр', minWidth: 100 },
    { id: 'Position_Current', label: 'Позиция', minWidth: 100 },
    { id: 'ErrorDescriptions', label: 'Ошибки', minWidth: 100 },
  ];
  
export default function StandTable(props) {
  console.log("render StandTable");

  const rows = props.data;
  return (
    <CommonTable columns={columns} rows={rows}></CommonTable>
  );
}