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

//export default function OrganAutoCard(props) {
const ImagesCard = React.memo((props) => {
  console.log(`! render ImagesCard`);

  const classes = useCardsStyles();
  const bull = <span className={classes.bullet}>•</span>;

  const images = props.images;
  const today = images?.Today;
  const current = images?.Current;
  return (
    <Card className={classes.root}>
      <CardContent>
        <Typography variant="h5" component="h2">
          {bull}Снимки
        </Typography>
        <CardRow descr="Всего" value={images?.ImageCount}></CardRow>
        <Typography variant="h6" component="h2">
          {bull}За сегодня
        </Typography>
        {today?.SingleGraphy ? <CardRow descr="Графия" value={today.SingleGraphy}></CardRow> : <></>}
        {today?.Scopy ? <CardRow descr="Скопия" value={today.Scopy}></CardRow> : <></>}
        {today?.Stitching ? <CardRow descr="Сшивка" value={today.Stitching}></CardRow> : <></>}

        <Typography variant="h6" component="h2">
          {bull}Последний
        </Typography>
        <CardRow descr="Тип" value={current?.Type}></CardRow>
        <CardRow descr="Напряжение" value={current?.Kv}></CardRow>
        <CardRow descr="Ток" value={current?.Mas ?? current?.Ma}></CardRow>
        <CardRow descr="Доза" value={ current?.Dose}></CardRow>
      </CardContent>
    </Card>
  );
});

export default ImagesCard;