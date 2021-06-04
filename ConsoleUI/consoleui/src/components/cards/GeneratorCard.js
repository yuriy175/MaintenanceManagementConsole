import React, {useContext}  from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';

import { CurrentEquipContext } from '../../context/currentEquip-context';
import {useCardsStyles} from './CommonCard'
import CardRow, {CardErrorRow} from './CardRow'

const GeneratorCard = React.memo((props) => {
//export default function GeneratorCard() {
  console.log(`! render GeneratorCard`);

  const classes = useCardsStyles();
  const bull = <span className={classes.bullet}>•</span>;

  const generator = props.generator;
  const heatStatus = generator?.State?.HeatStatus;
  return (
    <Card className={classes.root}>
      <CardContent>
        <Typography variant="h5" component="h2">
          {bull}Генератор
        </Typography>
        <CardRow descr="Состояние" 
          value={generator?.State?.State > 0? "Готов" : "Не готов"}
          rightColor={generator?.State?.State > 0? "green" : "red"}
        ></CardRow>
        <CardRow descr="Ток" value={generator?.State?.Mas ? generator.State.Mas + ' мАс' : ''}></CardRow>
        <CardRow descr="Напряжение" value={generator?.State?.Kv ? generator.State.Kv + ' кВ' : ''}></CardRow>
        <CardRow descr="Логическое. р. м." value={generator?.State?.Workstation}></CardRow>
        <CardRow descr="Нагрев" 
          value={heatStatus === 2 ? "Высокий" : heatStatus === 3 ? "Критический" : heatStatus === 1 ? "Норм" : ""}
          rightColor={heatStatus === 2 ? "yellow" : heatStatus === 3 ? "red" :  heatStatus === 1 ? "green" : "gray"}
        ></CardRow>
        <CardRow descr="Фокус" value={
          generator?.State?.Focus === 1 ? "Большой" : generator?.State?.Focus === 0 ? "Малый" : ""
        }></CardRow>
        <CardRow descr="Педаль" 
          value={ generator?.State?.PedalPressed === 1 ? "Графия" : generator?.State?.PedalPressed === 2 ? "Скопия" : "Не нажата"} 
          rightColor={generator?.State?.PedalPressed === 1 || generator?.State?.PedalPressed === 2? "green" : undefined }
        ></CardRow>
        <CardRow descr="Ошибки" value={generator?.State?.ErrorDescriptions?.length}></CardRow>
        {generator?.State?.ErrorDescriptions?.length ? 
          generator.State.ErrorDescriptions.map((i, ind) => (
            <CardErrorRow key={ind.toString()}  descr={i.Code} value={i.Description} ></CardErrorRow>
            ))
            :
            <></>          
        }
      </CardContent>
    </Card>
  );
});

export default GeneratorCard;
