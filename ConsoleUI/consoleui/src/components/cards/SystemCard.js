import React, {useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';

import { CurrentEquipContext } from '../../context/currentEquip-context';
import {useCardsStyles} from './CommonCard'

export default function SystemCard() {
  console.log(`! render SystemCard`);

  const [currEquipState, currEquipDispatch] = useContext(CurrentEquipContext);
  const classes = useCardsStyles();
  const bull = <span className={classes.bullet}>•</span>;

  return (
    <Card className={classes.root}>
      <CardContent>
        <Typography variant="h5" component="h2">
          {bull}CPU
        </Typography>
        <Typography className={classes.pos} color="textSecondary">
          {currEquipState.cpu?.Model} 
        </Typography>
        <Typography className={classes.pos} color="textSecondary">
          Загрузка - {currEquipState.cpu?.CPU_Load}%
        </Typography>

        <Typography variant="h5" component="h2">
          {bull}Память
        </Typography>
        <Typography className={classes.pos} color="textSecondary">
          Всего {currEquipState.memory?.TotalSize} Доступно {currEquipState.memory?.AvailableSize}
        </Typography>
      </CardContent>
    </Card>
  );
}