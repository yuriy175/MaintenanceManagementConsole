import React, {useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';

import { CurrentEquipContext } from '../../context/currentEquip-context';
import {useCardsStyles} from './CommonCard'

export default function DosimeterCard() {
  console.log(`! render DosimeterCard`);

  const classes = useCardsStyles();
  const [currEquipState, currEquipDispatch] = useContext(CurrentEquipContext);
  const bull = <span className={classes.bullet}>•</span>;

  return (
    <Card className={classes.root}>
      <CardContent>
        <Typography variant="h5" component="h2">
          {bull}Дозиметр
        </Typography>
        <Typography className={classes.pos} color="textSecondary">
          Состояние - {currEquipState.dosimeter?.State?.State< 1? "Не готов" : "Готов"} 
        </Typography>
        <Typography className={classes.pos} color="textSecondary">
          Доза - {currEquipState.dosimeter?.State?.Dose} сГр
        </Typography>
      </CardContent>
    </Card>
  );
}