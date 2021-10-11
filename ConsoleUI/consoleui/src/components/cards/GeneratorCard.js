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
  const generatorState = generator?.State;
  const heatStatus = generator?.State?.HeatStatus;
  
  const mas = generatorState?.Post_mas || generatorState?.Mas;
  const kv = generatorState?.Post_kv || generatorState?.Kv;
  const ma = generatorState?.Post_ma || generatorState?.Ma;
  const ms = generatorState?.Post_time || generatorState?.Ms;

  const scopyMas = generatorState?.Scopy_post_mas || generatorState?.Scopy_mas;
  const scopyKv = generatorState?.Scopy_post_kv || generatorState?.Scopy_kv;
  const scopyMa = generatorState?.Scopy_post_ma || generatorState?.Scopy_ma;
  const scopyMs = generatorState?.Scopy_post_ms || generatorState?.Scopy_ms;

  const isPoints3 = generatorState?.Points_mode === 3;
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
        <CardRow descr="Логическое. р. м." value={generator?.State?.Workstation}></CardRow>
        <CardRow descr="Нагрев" 
          value={heatStatus === 2 ? "Высокий" : heatStatus === 3 ? "Критический" : heatStatus === 1 ? "Норм" : ""}
          rightColor={heatStatus === 2 ? "yellow" : heatStatus === 3 ? "red" :  heatStatus === 1 ? "green" : "gray"}
        ></CardRow>
        <CardRow descr="Фокус" value={
          generator?.State?.Focus === 1 ? "Большой" : generator?.State?.Focus === 0 ? "Малый" : ""
        }></CardRow>
        <CardRow descr="Педаль" 
          value={ 
            generatorState?.PedalPressed === 1 ? "Графия" : 
            generatorState?.PedalPressed === 2 || generatorState?.PedalPressed === 4 ? "Скопия" : 
            generatorState?.PedalPressed === 3 ? "Графия на копии" :
            "Не нажата"} 
          rightColor={generatorState?.PedalPressed >= 1 && generatorState?.PedalPressed <= 4? "green" : undefined }
        ></CardRow>
        <CardRow descr="Техника" value={
          generatorState?.Points_mode === 2 ? "2х точка" : generatorState?.Points_mode === 3 ? "3х точка" : ""
        }></CardRow>
        <Typography variant="h6" component="h2">
          {bull}Графия
        </Typography>  
        <CardRow descr="Напряжение" value={kv ? kv + ' кВ' : ''}></CardRow>      
        {!isPoints3 ? 
          <CardRow descr="Количество электричества" value={mas ? mas + ' мАс' : ''}></CardRow> :
          <div>       
            <CardRow descr="Сила тока" value={ma ? ma + ' мА' : ''}></CardRow>
            <CardRow descr="Время" value={ms ? ms + ' мс' : ''}></CardRow>
          </div>}
        <Typography variant="h6" component="h2">
          {bull}Скопия
        </Typography>
        <CardRow descr="Режим" value={generatorState?.Scopy_mode}></CardRow>
        <CardRow descr="Напряжение" value={scopyKv ? scopyKv + ' кВ' : ''}></CardRow>
        <CardRow descr="Сила тока" value={scopyMa ? scopyMa + ' мА' : ''}></CardRow>
        <CardRow descr="Время" value={scopyMs ? scopyMs + ' мс' : ''}></CardRow>

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
