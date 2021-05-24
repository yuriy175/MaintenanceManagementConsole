import React, {useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';

import { CurrentEquipContext } from '../../context/currentEquip-context';
import {useCardsStyles} from './CommonCard'
import CardRow from './CardRow'

const SystemCard = React.memo((props) => {
//export default function SystemCard() {
  console.log(`! render SystemCard`);

  const classes = useCardsStyles();
  const bull = <span className={classes.bullet}>•</span>;

  const system = props.system;
  const volatile = props.volatile;
  const logicalDisks = system?.LogicalDisks; // HDD
  const physicalDisks = system?.HardDrives;  // PhysicalDisks
  const lans = system?.Lans; // Network
  const vgas = system?.VideoAdapters; // VGA
  const monitors = system?.Monitors; // Monitor
  const printers = system?.Printers; //Printer
  const processor = system?.Motherboards ? system?.Motherboards[0] : undefined; // Processor
  //
  return (
    <Card className={classes.root}>
      <CardContent>
        <Typography variant="h5" component="h2">
          {bull}CPU
        </Typography>
        {processor?.Cpu || volatile?.Processor?.CPULoad ? 
          <>
            <CardRow descr="Модель" value={processor?.Cpu} rightWidth={'100%'}></CardRow>
            <CardRow descr="Загрузка" value={volatile?.Processor?.CPULoad+'%'}></CardRow>
          </> : <></>}

        <Typography variant="h5" component="h2">
          {bull}Память
        </Typography>
        {volatile?.Memory?.MemoryTotalGb || volatile?.Memory?.MemoryFreeGb ? 
          <>
            <CardRow descr="Всего" value={volatile?.Memory?.MemoryTotalGb+' Гб'}></CardRow>
            <CardRow descr="Доступно" value={volatile?.Memory?.MemoryFreeGb+' Гб'}></CardRow>
          </> : <></>}

        <Typography variant="h5" component="h2">
          {bull}Диски
        </Typography>
        <Typography variant="h6" component="h2">
              {bull}Логические диски
        </Typography>
        {logicalDisks?.map((i, ind) => (
          <div key={ind.toString()} >
            <Typography variant="h6" component="h2">
              {bull}Диск {i.Letter}
            </Typography>
            <CardRow descr="Всего" value={i.TotalSize+' Гб'}></CardRow>
            <CardRow descr="Доступно" value={i.FreeSize+' Гб'}></CardRow>
          </div>
          ))
        } 
        <Typography variant="h6" component="h2">
              {bull}Физические диски
        </Typography>
        {physicalDisks?.map((i, ind) => (
          <div key={ind.toString()} >
            <Typography variant="h6" component="h2">
              {bull}Диск {i.mediatype}
            </Typography>
            <CardRow descr={i.FriendlyName+' Гб'} value={i.SizeGb+' Гб'}></CardRow>
          </div>
          ))
        } 

        <Typography variant="h5" component="h2">
          {bull}Сеть
        </Typography>
        {lans?.filter(i => i.Adapter).map((i, ind) => (
          <CardRow key={ind.toString()} descr={i.Adapter} value={i.Ip}></CardRow>
          ))
        } 

        <Typography variant="h5" component="h2">
          {bull}Видеоадаптеры
        </Typography>
        {vgas?.map((i, ind) => (
          <div key={ind.toString()} >
            <CardRow descr="Имя" value={i.CardName} rightWidth={'100%'}></CardRow>
            <CardRow descr="Память" value={i.MemoryGb+' Гб'}></CardRow>
            <CardRow descr="Драйвер" value={i.DrvDate} rightWidth={'100%'}></CardRow>
          </div>
          ))
        } 

        <Typography variant="h5" component="h2">
          {bull}Мониторы
        </Typography>
        {monitors?.map((i, ind) => (
          // <CardRow key={ind.toString()} descr={i.MonitorName} value={i.Width+'x'+i.Height}></CardRow>
          <CardRow key={ind.toString()} descr={i.MonitorName} value={i.SerialNumber} rightWidth={'100%'}></CardRow>
          ))
        } 

        <Typography variant="h5" component="h2">
          {bull}Принтеры
        </Typography>
        {printers?.map((i, ind) => (
          <div key={ind.toString()} >
            <CardRow descr="Имя" value={i.PrinterName} rightWidth={'100%'}></CardRow>
            <CardRow descr="Порт" value={i.PrinterPort} rightWidth={'100%'}></CardRow>
          </div>
          ))
        } 
      </CardContent>
    </Card>
  );
});

export default SystemCard;