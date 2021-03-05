import React, {useContext}  from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';

import { CurrentEquipContext } from '../../context/currentEquip-context';
import {useCardsStyles} from './CommonCard'

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
        <Typography className={classes.pos} color="textSecondary">
          Состояние - {currEquipState.generator?.State?.State < 1? "Не готов" : "Готов"} 
        </Typography>
        <Typography className={classes.pos} color="textSecondary">
          Ток - {currEquipState.generator?.State?.Mas} мАс
        </Typography>
        <Typography className={classes.pos} color="textSecondary">
          Напряжение - {currEquipState.generator?.State?.Kv} кВ
        </Typography>
        <Typography className={classes.pos} color="textSecondary">
          Логическое. р. м. - {currEquipState.generator?.State?.Workstation}
        </Typography>
        <Typography className={classes.pos} color="textSecondary">
          Нагрев - {currEquipState.generator?.State?.HeatStatus}
        </Typography>
        <Typography className={classes.pos} color="textSecondary">
          Педаль - {currEquipState.generator?.State?.PedalPressed}
        </Typography>
        <Typography className={classes.pos} color="textSecondary">
          Ошибки - {currEquipState.generator?.State?.ErrorDescriptions?.length}
        </Typography>
      </CardContent>
    </Card>
  );
}
