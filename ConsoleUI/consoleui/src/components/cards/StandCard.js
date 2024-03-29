import React, {useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';

import { CurrentEquipContext } from '../../context/currentEquip-context';
import {useCardsStyles} from './CommonCard'
import CardRow, {CardErrorRow} from './CardRow'

const StandCard = React.memo((props) => {
// export default function StandCard() {
  console.log(`! render StandCard`);

  const classes = useCardsStyles();
  const bull = <span className={classes.bullet}>•</span>;

  const stand = props.stand;
  return (
    <Card className={classes.root}>
      <CardContent>
        <Typography variant="h5" component="h2">
          {bull}Штатив
        </Typography>
        <CardRow descr="Состояние" 
          value={stand?.State?.State > 0? "Готов" : "Не готов" }
          rightColor={stand?.State?.State > 0? "green" : "red"}
        ></CardRow>
        <CardRow descr="Режим" value={stand?.State?.Mode}></CardRow>
        <CardRow descr="Растр" value={stand?.State?.RasterState}></CardRow>
        <CardRow descr="Позиция" value={stand?.State?.Position_Current}></CardRow>
        <CardRow descr="Угол наклона трубки" value={stand?.State?.Tube_Incline}></CardRow>
        <CardRow descr="Угол наклона деки" value={stand?.State?.Deck_Incline}></CardRow>
        <CardRow descr="Угол наклона детектора" value={stand?.State?.Camera_Incline}></CardRow>
        <CardRow descr="Фокусное расстояние" value={stand?.State?.Ffd_Current}></CardRow>
        <CardRow descr="Высота стола" value={stand?.State?.Deck_Height}></CardRow>
        <CardRow descr="Высота излучателя" value={stand?.State?.Uarm_Height}></CardRow>
        <CardRow descr="Ошибки" value={stand?.State?.ErrorDescriptions?.length}></CardRow>
        {stand?.State?.ErrorDescriptions?.length ? 
          stand.State.ErrorDescriptions.map((i, ind) => (
            <CardErrorRow key={ind.toString()}  descr={i.Code} value={i.Description} ></CardErrorRow>
            ))
            :
            <></>          
        }
      </CardContent>
    </Card>
  );
});

export default StandCard;