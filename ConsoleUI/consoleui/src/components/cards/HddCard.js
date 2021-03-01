import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';

import {useCardsStyles} from './CommonCard'

export default function HddCard() {
  const classes = useCardsStyles();
  const bull = <span className={classes.bullet}>•</span>;

  return (
    <Card className={classes.root}>
      <CardContent>
        <Typography variant="h5" component="h2">
          {bull}Диски
        </Typography>
        <Typography className={classes.pos} color="textSecondary">
          Органоавтоматика
        </Typography>
        <Typography className={classes.pos} color="textSecondary">
          Состояние
        </Typography>
      </CardContent>
    </Card>
  );
}