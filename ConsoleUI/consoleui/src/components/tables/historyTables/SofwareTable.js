import React from 'react';
import CommonTable from './CommonTable'

const columns = [
    { id: 'EquipName', label: 'Комплекс', minWidth: 170 },
    { id: 'DateTime', label: 'Время', minWidth: 100 },
    { id: 'SettingsDB', label: 'Settings БД', minWidth: 100 },
    { id: 'ObservationsDB', label: 'Observations БД', minWidth: 100 },
    { id: 'Version', label: 'Версия Атлас', minWidth: 100 },
    { id: 'XilibVersion', label: 'Версия xilib', minWidth: 100 },
    { id: 'ErrorDescriptions', label: 'Ошибки', minWidth: 100, 
      formatArray: (values) =>
      {
        if(!values || !values.length) {
          return values;
        }
        //const errors = values.reduce((accumulator, currentValue) => accumulator + `Code ${currentValue.Code}: ${currentValue. Description}`, '');
        return values.map((currentValue) => `Code ${currentValue.Code}: ${currentValue. Description}`);
      },
      hasErrors: (values) =>
      {
        if(!values || !values.length) {
          return false;
        }

        return true;
      }
    },
  ];

export default function SofwareTable(props) {
  console.log("render SofwareTable");

  const rows = props.data;
  return (
    <CommonTable columns={columns} rows={rows}></CommonTable>
  );
}