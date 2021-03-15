import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import CommonTable from './CommonTable'

const columns = [
    { id: 'EquipName', label: 'Комплекс', minWidth: 170 },
    { id: 'DateTime', label: 'Время', minWidth: 100 },
    { id: 'PACS', label: 'PACS', minWidth: 100,
      formatArray: (values) =>
      {
        if(!values || !values.length) {
          return values;
        }
        
        return values.map((currentValue) => `${currentValue.Name}(${currentValue.IP}): ${currentValue.State}`);
      },
    },
    { id: 'WorkList', label: 'WorkList', minWidth: 100,
        format: (values) =>
        {
          if(!values || !values.length) {
            return values;
          }
          const errors = values.reduce((accumulator, currentValue) => accumulator + `${currentValue.Name}(${currentValue.IP}): ${currentValue.State}`, '');
          return errors;
        },
    },
  ];

export default function DicomTable(props) {
  console.log("render DicomTable");

  const rows = props.data;
  return (
    <CommonTable columns={columns} rows={rows}></CommonTable>
  );
}