import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import CommonTable from '../CommonTable'

// const columns = [
//     { id: 'EquipName', label: 'Комплекс', minWidth: 170 },
//     { id: 'DateTime', label: 'Время', minWidth: 100 },
//     { id: 'SettingsDB', label: 'Settings БД', minWidth: 100 },
//     { id: 'ObservationsDB', label: 'Observations БД', minWidth: 100 },
//     { id: 'Version', label: 'Версия Атлас', minWidth: 100 },
//     { id: 'XilibVersion', label: 'Версия xilib', minWidth: 100 },
//     { id: 'ErrorDescriptions', label: 'Ошибки', minWidth: 100, 
//       formatArray: (values) =>
//       {
//         if(!values || !values.length) {
//           return values;
//         }
//         //const errors = values.reduce((accumulator, currentValue) => accumulator + `Code ${currentValue.Code}: ${currentValue. Description}`, '');
//         return values.map((currentValue) => `Code ${currentValue.Code}: ${currentValue. Description}`);
//       },
//       hasErrors: (values) =>
//       {
//         if(!values || !values.length) {
//           return false;
//         }

//         return true;
//       }
//     },
//   ];

const columnsPermanent = [
  { id: 'Parameter', label: 'Параметр', minWidth: 170 },    
  { id: 'Value', label: 'Значение', minWidth: 100 },
  { id: 'DateTime', label: 'Посл. изменение', minWidth: 100 },
];

const columnsVolatile = [
  { id: 'EquipName', label: 'Комплекс', minWidth: 170 },
  { id: 'DateTime', label: 'Время', minWidth: 100 },
  { id: 'ErrorType', label: 'Источник', minWidth: 100 },
  { id: 'ErrorCode', label: 'Код', minWidth: 100 },
  { id: 'ErrorDescription', label: 'Описание', minWidth: 100 },
];

const useStyles = makeStyles({
  root: {
    width: '100%',
  }
});

export default function SofwareTable(props) {
  console.log("render SofwareTable");

  const classes = useStyles();
  const volatileInfoRows = props.data?.VolatileInfo;
  const permanentInfo = props.data?.PermanentInfo?.length > 0 ? props.data?.PermanentInfo[0] : null;
  const dbs = permanentInfo?.Databases?.map(e => 
    { 
      return { Parameter: `БД ${e.DB_name}`, Value: e.DB_Status, DateTime: permanentInfo?.DateTime }
    });

  const os = permanentInfo?.Sysinfo?.OS ? [ { Parameter: permanentInfo?.Sysinfo?.OS, Value: permanentInfo?.Sysinfo.Version, DateTime: permanentInfo?.DateTime }] : []
  const sql = permanentInfo?.MSSQL?.OS ? [ { Parameter: permanentInfo?.MSSQL?.SQL, Value: permanentInfo?.MSSQL.Version, DateTime: permanentInfo?.DateTime }] : []
  const atlas = permanentInfo?.Atlas ? [ 
    { Parameter: "Версия Атлас", Value: permanentInfo?.Atlas.Atlas_Version, DateTime: permanentInfo?.DateTime },
    { Parameter: "Тип комплекса", Value: permanentInfo?.Atlas.Complex_type, DateTime: permanentInfo?.DateTime },
    { Parameter: "Венрсия Xilib", Value: permanentInfo?.Atlas.XiLibs_Version, DateTime: permanentInfo?.DateTime },
  ] : []

  const permanentInfoRows = 
    os
    .concat(sql)
    .concat(dbs ?? [])    
    .concat(atlas)   

  return (
    <div className={classes.root}>
      |{props.equipName ? 
        <div>
          <Typography variant="h5" component="h2">
              {props.equipName}
          </Typography>
          <CommonTable columns={columnsPermanent} rows={permanentInfoRows}></CommonTable>
          </div>
        : <></>
      }
      <CommonTable columns={columnsVolatile} rows={volatileInfoRows}></CommonTable>
    </div>
  );
}