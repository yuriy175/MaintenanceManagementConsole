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
  return (
    <Card className={classes.root}>
      <CardContent>
        <Typography variant="h5" component="h2">
          {bull}CPU
        </Typography>
        {system?.Processor?.Model || system?.Processor?.CPU_Load ? 
          <>
            <CardRow descr="Модель" value={system?.Processor?.Model} rightWidth={'100%'}></CardRow>
            <CardRow descr="Загрузка" value={system?.Processor?.CPU_Load+'%'}></CardRow>
          </> : <></>}

        <Typography variant="h5" component="h2">
          {bull}Память
        </Typography>
        {system?.Memory?.Memory_total_Gb || system?.Memory?.Memory_free_Gb ? 
          <>
            <CardRow descr="Всего" value={system?.Memory?.Memory_total_Gb+' Гб'}></CardRow>
            <CardRow descr="Доступно" value={system?.Memory?.Memory_free_Gb+' Гб'}></CardRow>
          </> : <></>}

        <Typography variant="h5" component="h2">
          {bull}Диски
        </Typography>
        <Typography variant="h6" component="h2">
              {bull}Логические диски
        </Typography>
        {system?.HDD?.map((i, ind) => (
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
        {system?.PhysicalDisks?.map((i, ind) => (
          <div key={ind.toString()} >
            <Typography variant="h6" component="h2">
              {bull}Диск {i.MediaType}
            </Typography>
            <CardRow descr={i.FriendlyName+' Гб'} value={i.Size_Gb+' Гб'}></CardRow>
          </div>
          ))
        } 

        <Typography variant="h5" component="h2">
          {bull}Сеть
        </Typography>
        {system?.Network?.filter(i => i.NIC).map((i, ind) => (
          <CardRow key={ind.toString()} descr={i.NIC} value={i.IP}></CardRow>
          ))
        } 

        <Typography variant="h5" component="h2">
          {bull}Видеоадаптеры
        </Typography>
        {system?.VGA?.map((i, ind) => (
          <div key={ind.toString()} >
            <CardRow descr="Имя" value={i.Card_Name} rightWidth={'100%'}></CardRow>
            <CardRow descr="Память" value={i.Memory_Gb+' Гб'}></CardRow>
            <CardRow descr="Драйвер" value={i.Driver_Version}></CardRow>
          </div>
          ))
        } 

        <Typography variant="h5" component="h2">
          {bull}Мониторы
        </Typography>
        {system?.Monitor?.map((i, ind) => (
          <CardRow key={ind.toString()} descr={i.Device_Name} value={i.Width+'x'+i.Height}></CardRow>
          ))
        } 

        <Typography variant="h5" component="h2">
          {bull}Принтеры
        </Typography>
        {system?.Printer?.map((i, ind) => (
          <div key={ind.toString()} >
            <CardRow descr="Имя" value={i.Printer_Name} rightWidth={'100%'}></CardRow>
            <CardRow descr="Порт" value={i.Port_Name} rightWidth={'100%'}></CardRow>
          </div>
          ))
        } 
      </CardContent>
    </Card>
  );
});

export default SystemCard;