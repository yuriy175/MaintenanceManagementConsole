import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import CommonTable from '../CommonTable'

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
  const permanentInfo = props.data?.PermanentInfo?.length > 0 ? props.data?.PermanentInfo[0] : null;
  const hdds = permanentInfo?.HDD?.map(e => 
    { 
      return { Parameter: `Логический ${e.Letter}`, Value: e.TotalSize, DateTime: permanentInfo?.DateTime }
    });
  
  const physicalDisks = permanentInfo?.PhysicalDisks?.map(e => 
    { 
      return { Parameter: `Физический ${e.MediaType} ${e.FriendlyName}`, Value: e.Size_Gb, DateTime: permanentInfo?.DateTime }
    });

  const monitors = permanentInfo?.Monitor?.map(e => 
    { 
      return { Parameter: `Монитор ${e.Device_Name}`, Value: `${e.Width}x${e.Height}`, DateTime: permanentInfo?.DateTime }
    });

  const vga = permanentInfo?.VGA?.map(e => 
    { 
      return { Parameter: `Видеоадаптер ${e.Card_Name}`, Value: `${e.Memory_Gb}`, DateTime: permanentInfo?.DateTime }
    });

  const processor = permanentInfo?.Processor?.Model ? [ { Parameter: "Процессор", Value: permanentInfo?.Processor.Model, DateTime: permanentInfo?.DateTime }] : []
  const motherboard = permanentInfo?.Motherboard?.Model ? [ { Parameter: "Материнская плата", Value: permanentInfo?.Motherboard.Model, DateTime: permanentInfo?.DateTime }] : []
  const memory = permanentInfo?.Memory?.Memory_total_Gb ? [ { Parameter: "Память, Гб", Value: permanentInfo?.Memory.Memory_total_Gb, DateTime: permanentInfo?.DateTime }] : []

  const permanentInfoRows = 
    processor
    .concat(memory)
    .concat(hdds ?? [])
    .concat(motherboard)    
    .concat(physicalDisks ?? [])    
    .concat(monitors ?? [])
    .concat(vga ?? []);

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