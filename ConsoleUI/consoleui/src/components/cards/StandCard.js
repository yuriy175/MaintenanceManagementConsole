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

export default function StandCard() {
  console.log(`! render StandCard`);

  const classes = useCardsStyles();
  const [currEquipState, currEquipDispatch] = useContext(CurrentEquipContext);
  const bull = <span className={classes.bullet}>•</span>;

  return (
    <Card className={classes.root}>
      <CardContent>
        <Typography variant="h5" component="h2">
          {bull}Штатив
        </Typography>
        <CardRow descr="Состояние" 
          value={currEquipState.stand?.State?.State > 0? "Готов" : "Не готов" }
          rightColor={currEquipState.stand?.State?.State > 0? "green" : "red"}
        ></CardRow>
        <CardRow descr="Растр" value={currEquipState.stand?.State?.RasterState}></CardRow>
        <CardRow descr="Позиция" value={currEquipState.stand?.State?.Position_Current}></CardRow>
        <CardRow descr="Ошибки" value={currEquipState.stand?.State?.ErrorDescriptions?.length}></CardRow>
      </CardContent>
    </Card>
  );
}