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

// export default function DosimeterCard() {
const DosimeterCard = React.memo((props) => {
  console.log(`! render DosimeterCard`);

  const classes = useCardsStyles();
  const bull = <span className={classes.bullet}>•</span>;

  const dosimeter = props.dosimeter;
  return (
    <Card className={classes.root}>
      <CardContent>
        <Typography variant="h5" component="h2">
          {bull}Дозиметр
        </Typography>
        <CardRow descr="Состояние" 
          value={dosimeter?.State?.State > 0? "готов" : "Не готов"}
          rightColor={dosimeter?.State?.State > 0 ? "green" : "red"}></CardRow>
        <CardRow descr="Доза" value={dosimeter?.State?.Dose ? 
          parseFloat(dosimeter?.State?.Dose).toFixed(2) + ' сГр' : '' }></CardRow>
      </CardContent>
    </Card>
  );
});

export default DosimeterCard;