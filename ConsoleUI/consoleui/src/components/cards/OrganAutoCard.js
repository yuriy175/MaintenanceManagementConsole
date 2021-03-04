import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';

import {useCardsStyles} from './CommonCard'

export default function OrganAutoCard() {
  console.log(`! render OrganAutoCard`);

  const classes = useCardsStyles();
  const bull = <span className={classes.bullet}>•</span>;

  return (
    <Card className={classes.root}>
      <CardContent>
        <Typography variant="h5" component="h2">
          {bull}Органоавтоматика
        </Typography>
        <Typography className={classes.pos} color="textSecondary">
          Название
        </Typography>
        <Typography className={classes.pos} color="textSecondary">
          Прекция
        </Typography>
        <Typography className={classes.pos} color="textSecondary">
          Возр. группа
        </Typography>
        <Typography className={classes.pos} color="textSecondary">
          Телосложение
        </Typography>
      </CardContent>
    </Card>
  );
}