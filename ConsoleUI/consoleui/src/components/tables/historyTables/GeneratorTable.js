import React from 'react';
import CommonTable from '../CommonTable'

const columns = [
  { id: 'EquipName', label: 'Комплекс', minWidth: 170 },
  { id: 'State', label: 'Состояние', minWidth: 100 },
  { id: 'DateTime', label: 'Время', minWidth: 100 },
  { id: 'Workstation', label: 'Раб. место', minWidth: 100 },
  { id: 'HeatStatus', label: 'Нагрев', minWidth: 100 },
  { id: 'ErrorDescriptions', label: 'Ошибки', minWidth: 100, 
    formatArray: (values) =>
    {
      if(!values || !values.length) {
        return values ?? [];
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
  { id: 'Mas', label: 'Ток', minWidth: 100 },
  { id: 'Kv', label: 'Напряжение', minWidth: 100 },  
];

export default function GeneratorTable(props) {
  console.log("render GeneratorTable");

  const rows = props.data;
  return (
    <CommonTable columns={columns} rows={rows}></CommonTable>
  );
}