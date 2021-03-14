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

export default function SoftwareCard() {
  console.log(`! render SoftwareCard`);

  const [currEquipState, currEquipDispatch] = useContext(CurrentEquipContext);
  const classes = useCardsStyles();
  const bull = <span className={classes.bullet}>•</span>;

  return (
    <Card className={classes.root}>
      <CardContent>
        <Typography variant="h5" component="h2">
          {bull}Software
        </Typography>
        <CardRow descr="Settings БД" 
          value={currEquipState.software?.SettingsDB ?  "Готов" : "Не готов" }
          rightColor={currEquipState.software?.SettingsDB ? "green" : "red"}
        ></CardRow>
        <CardRow descr="Observations БД" 
          value={currEquipState.software?.ObservationsDB ? "Готов" : "Не готов"}
          rightColor={ currEquipState.software?.ObservationsDB ? "green" : "red"}
        ></CardRow>
        <CardRow descr="Версия Атлас" value={currEquipState.software?.Version}></CardRow>
        <CardRow descr="Версия Xilib" value={currEquipState.software?.XilibVersion}></CardRow>
      </CardContent>
    </Card>
  );
}