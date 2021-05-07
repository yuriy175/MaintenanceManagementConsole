import React from 'react';
import CommonTable from '../CommonTable'
import {parseLocalString} from '../../../utilities/utils'

const columns = [
    { id: 'EquipName', label: 'Комплекс', minWidth: 170 },
    { id: 'DateTime', label: 'Время', minWidth: 100,
      format: (value) => parseLocalString(value)
    },
    { id: 'Data', label: 'Данные', minWidth: 100, maxWidth: 800,
      hasErrors: (value) => value?.includes('ErrorDescriptions')
    },

    // { id: 'State', label: 'Состояние', minWidth: 100 },
    // { id: 'RasterState', label: 'Растр', minWidth: 100 },
    // { id: 'Position_Current', label: 'Позиция', minWidth: 100 },
    // { id: 'ErrorDescriptions', label: 'Ошибки', minWidth: 100 },
  ];
  
export default function StandTable(props) {
  console.log("render StandTable");

  const rows = props.data;
  return (
    <CommonTable columns={columns} rows={rows}></CommonTable>
  );
}