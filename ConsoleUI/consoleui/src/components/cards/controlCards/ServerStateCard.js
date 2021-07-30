import React, {useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';

import {useCardsStyles} from '../CommonCard'
import CardRow, {CardErrorRow} from '../CardRow'

export default function ServerStateCard(props) {
  console.log(`! render ServerStateCard`);

  const classes = useCardsStyles();
  const bull = <span className={classes.bullet}>•</span>;

  const state = props.serverState;
  
  return (
    <Card className={classes.root}>
      <CardContent>       
        <Typography variant="h5" component="h2">
          {bull}Сервер
        </Typography> 
        <CardRow descr="Всего БД" value={state?.DBUsedSize}></CardRow>
        <CardRow descr="Всего диск" value={state?.DiskTotalSpace}></CardRow>
        <CardRow descr="Занято диск" value={state?.DiskUsedSpace}></CardRow>
      </CardContent>
    </Card>
  );
}