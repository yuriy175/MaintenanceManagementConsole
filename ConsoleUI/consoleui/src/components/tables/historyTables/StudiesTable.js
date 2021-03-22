import React from 'react';
import CommonTable from './CommonTable'

const columns = [
  { id: 'EquipName', label: 'Комплекс', minWidth: 170 },
  { id: 'DateTime', label: 'Время', minWidth: 100 },
  { id: 'StudyId', label: 'Id исслед', minWidth: 100 },
  { id: 'StudyDicomUid', label: 'Dicom Uid', minWidth: 100 },
  { id: 'StudyName', label: 'Наименование', minWidth: 100 },
  { id: 'State', label: 'Состояние', minWidth: 100 },
];

export default function StudiesTable(props) {
  console.log("render StudiesTable");

  const rows = props.data;
  return (
    <CommonTable columns={columns} rows={rows}></CommonTable>
  );
}