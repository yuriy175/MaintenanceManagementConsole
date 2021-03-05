import React, {useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';

import { CurrentEquipContext } from '../../context/currentEquip-context';
import {useCardsStyles} from './CommonCard'

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
        <Typography className={classes.pos} color="textSecondary">
          Состояние - {currEquipState.stand?.State?.State< 1? "Не готов" : "Готов"} 
        </Typography>
        <Typography className={classes.pos} color="textSecondary">
          Растр - {currEquipState.stand?.State?.RasterState}
        </Typography>
        <Typography className={classes.pos} color="textSecondary">
          Позиция - {currEquipState.stand?.State?.Position_Current}
        </Typography>
        <Typography className={classes.pos} color="textSecondary">
          Ошибки - {currEquipState.stand?.State?.ErrorDescriptions?.length}
        </Typography>
      </CardContent>
    </Card>
  );
}