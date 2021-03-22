import React from 'react';
import CommonTable from './CommonTable'

const columns = [
    { id: 'EquipName', label: 'Комплекс', minWidth: 170 },
    { id: 'DateTime', label: 'Время', minWidth: 100 },
    { id: 'State', label: 'Состояние', minWidth: 100 },
    { id: 'CPU_Load', label: 'CPU загрузка', minWidth: 100 },
    { id: 'TotalMemory', label: 'Всего памяти', minWidth: 100 },
    { id: 'AvailableSize', label: 'Свободно памяти', minWidth: 100 },
    { id: 'HddName', label: 'Имя диска', minWidth: 100 },
    { id: 'HddTotalSpace', label: 'Всего', minWidth: 100 },  
    { id: 'HddFreeSpace', label: 'Свободно', minWidth: 100 },  
  ];
  
export default function SystemTable(props) {
  console.log("render SystemTable");

  const rows = props.data;
  return (
    <CommonTable columns={columns} rows={rows}></CommonTable>
  );
}