import React, {useContext}  from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';

import { CurrentEquipContext } from '../../context/currentEquip-context';
import {useCardsStyles} from './CommonCard'
import CardRow from './CardRow'

export default function GeneratorCard() {
  console.log(`! render GeneratorCard`);

  const classes = useCardsStyles();
  const [currEquipState, currEquipDispatch] = useContext(CurrentEquipContext);
  const bull = <span className={classes.bullet}>•</span>;

  return (
    <Card className={classes.root}>
      <CardContent>
        <Typography variant="h5" component="h2">
          {bull}Генератор
        </Typography>
        <CardRow descr="Состояние" 
          value={currEquipState.generator?.State?.State > 0? "Готов" : "Не готов"}
          rightColor={currEquipState.generator?.State?.State > 0? "green" : "red"}
        ></CardRow>
        <CardRow descr="Ток" value={currEquipState.generator?.State?.Mas ? currEquipState.generator.State.Mas + ' мАс' : ''}></CardRow>
        <CardRow descr="Напряжение" value={currEquipState.generator?.State?.Kv ? currEquipState.generator.State.Kv + ' кВ' : ''}></CardRow>
        <CardRow descr="Логическое. р. м." value={currEquipState.generator?.State?.Workstation}></CardRow>
        <CardRow descr="Нагрев" value={currEquipState.generator?.State?.HeatStatus}></CardRow>
        <CardRow descr="Педаль" value={currEquipState.generator?.State?.PedalPressed}></CardRow>
        <CardRow descr="Ошибки" value={currEquipState.generator?.State?.ErrorDescriptions?.length}></CardRow>
      </CardContent>
    </Card>
  );
}
