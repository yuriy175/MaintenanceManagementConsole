import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import CommonTable from './CommonTable'

const columnsPermanent = [
    { id: 'Parameter', label: 'Параметр', minWidth: 170 },    
    { id: 'Value', label: 'Значение', minWidth: 100 },
    { id: 'DateTime', label: 'Посл. изменение', minWidth: 100 },
  ];

const columnsVolatile = [
    { id: 'EquipName', label: 'Комплекс', minWidth: 170 },
    { id: 'DateTime', label: 'Время', minWidth: 100 },
    { id: 'Processor_CPU_Load', label: 'CPU загрузка', minWidth: 100 },
    { id: 'Memory_Memory_free_Gb', label: 'Свободно памяти', minWidth: 100 },
    { id: 'HddName', label: 'Имя диска', minWidth: 100 },
    { id: 'HddFreeSpace', label: 'Свободно', minWidth: 100 },  
  ];

  const useStyles = makeStyles({
    root: {
      width: '100%',
    }
  });
  
export default function SystemTable(props) {
  console.log("render SystemTable");

  const classes = useStyles();
  const volatileInfoRows = props.data?.VolatileInfo;
  const permanentInfoRows = props.data?.PermanentInfo;
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