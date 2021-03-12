import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';

import {useCardsStyles} from './CommonCard'
import CardRow from './CardRow'

export default function DicomCard() {
  console.log(`! render DicomCard`);

  const classes = useCardsStyles();
  const bull = <span className={classes.bullet}>•</span>;

  return (
    <Card className={classes.root}>
      <CardContent>
        <Typography variant="h5" component="h2">
          {bull}DICOM
        </Typography>
        <CardRow descr="Соединение с PACS" value={ 1? "Не готов" : "Готов"}></CardRow>
        <CardRow descr="PACS IP" value={1? "Не готов" : "Готов"}></CardRow>
        <CardRow descr="Собственный IP" value={1? "Не готов" : "Готов"}></CardRow>
        <CardRow descr="Принтер" value={1? "Не готов" : "Готов"}></CardRow>
      </CardContent>
    </Card>
  );
}