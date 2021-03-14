import React from 'react';
import CommonTable from './CommonTable'

const columns = [
  { id: 'EquipName', label: 'Комплекс', minWidth: 170 },
  { id: 'DateTime', label: 'Время', minWidth: 100 },
  { id: 'Name', label: 'Название', minWidth: 100 },
  { id: 'Projection', label: 'Проекция', minWidth: 100 },
  { id: 'Direction', label: 'Направление', minWidth: 100 },
  { id: 'AgeId', label: 'Возр. группа', minWidth: 100 },
  { id: 'Constitution', label: 'Телосложение', minWidth: 100 },  
];

export default function OrganAutoTable(props) {
  console.log("render OrganAutoTable");

  const rows = props.data;
  return (
    <CommonTable columns={columns} rows={rows}></CommonTable>
  );
}