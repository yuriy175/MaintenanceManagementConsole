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
        <CardRow descr="Загрузка" value={currEquipState.cpu?.CPU_Load+'%'}></CardRow>

        <Typography variant="h5" component="h2">
          {bull}Память
        </Typography>
        <CardRow descr="Всего" value={currEquipState.memory?.TotalMemory+'Мб'}></CardRow>
        <CardRow descr="Доступно" value={currEquipState.memory?.AvailableSize+'Мб'}></CardRow>
      </CardContent>
    </Card>
  );
}