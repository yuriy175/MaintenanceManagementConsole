import React, {useContext}  from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';

import { CurrentEquipContext } from '../../context/currentEquip-context';
import {useCardsStyles} from './CommonCard'
import CardRow from './CardRow'

// export default function DetectorCard(props) {
const DetectorCard = React.memo((props) => {
  console.log(`! render DetectorCard`);

  const classes = useCardsStyles();
  const bull = <span className={classes.bullet}>•</span>;

  const detectors = props.detectors;
  const aecs = props.aecs;
  return (
    <Card className={classes.root}>
      <CardContent>
        <Typography variant="h5" component="h2">
          {bull}Детекторы
        </Typography>
        {detectors?.map((i, ind) => (
          <div key={ind.toString()} >
            <Typography variant="h6" component="h2">
              {bull}{i.DetectorName} 
            </Typography>
            <CardRow descr="Состояние" 
              value={i.State !== 2? "Не готов" : "Готов"}
              rightColor={i.State !== 2? "red" : "green"}></CardRow>
            <CardRow descr="Поле детектора" value={i.DetectorField}></CardRow>            
            <CardRow descr="Скорость" value={i.DetectorFrameRate ? i?.DetectorFrameRate + ' к/сек' : ''}></CardRow>
            {aecs?
              <>
                <Typography variant="h6" component="h2">
                  {bull}АКЭ
                </Typography>
                <CardRow descr="Состояние" value={aecs.State?.Aec_enabled ? "Вкл" : "Выкл"}></CardRow>
                <CardRow descr="Поля АКЭ" value={aecs.State?.Aec_fields}></CardRow>
                <CardRow descr="Плотность" value={aecs.State?.Aec_density}></CardRow>
              </> : <></>
            }
          </div>
          ))
        }        
      </CardContent>
    </Card>
  );
});

export default DetectorCard;