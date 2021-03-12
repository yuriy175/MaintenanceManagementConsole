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

export default function OrganAutoCard() {
  console.log(`! render OrganAutoCard`);

  const classes = useCardsStyles();
  const [currEquipState, currEquipDispatch] = useContext(CurrentEquipContext);
  const bull = <span className={classes.bullet}>•</span>;

  return (
    <Card className={classes.root}>
      <CardContent>
        <Typography variant="h5" component="h2">
          {bull}Органоавтоматика
        </Typography>
        <CardRow descr="Название" value={currEquipState.organAuto?.Name} rightWidth={'50%'}></CardRow>
        <CardRow descr="Прекция" value={currEquipState.organAuto?.Projection}></CardRow>
        <CardRow descr="Направление" value={currEquipState.organAuto?.Direction}></CardRow>
        <CardRow descr="Возр. группа" value={ currEquipState.organAuto?.AgeId}></CardRow>
        <CardRow descr="Телосложение" value={currEquipState.organAuto?.Constitution}></CardRow>
      </CardContent>
    </Card>
  );
}