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
const OrganAutoCard = React.memo((props) => {
  console.log(`! render OrganAutoCard`);

  const classes = useCardsStyles();
  const bull = <span className={classes.bullet}>•</span>;

  const organAuto = props.organAuto;
  return (
    <Card className={classes.root}>
      <CardContent>
        <Typography variant="h5" component="h2">
          {bull}Органоавтоматика
        </Typography>
        <CardRow descr="Название" value={organAuto?.Name} rightWidth={'50%'}></CardRow>
        <CardRow descr="Прекция" value={organAuto?.Projection}></CardRow>
        <CardRow descr="Направление" value={organAuto?.Direction}></CardRow>
        {/* <CardRow descr="Возр. группа" value={ organAuto?.AgeId}></CardRow> */}
        <CardRow descr="Возр. группа" 
          value={ 
            organAuto?.AgeId === 1 ? "0 - 0.5" : 
            organAuto?.AgeId === 2 ? "0.5 - 2" : 
            organAuto?.AgeId === 3 ? "2 - 7" :
            organAuto?.AgeId === 4 ? "7 - 12" :
            organAuto?.AgeId === 5 ? "12 - 17" :
            "> 17"} >
        </CardRow>
        <CardRow descr="Телосложение"  value={ 
              organAuto?.Constitution === 1 ? "худой" : 
              organAuto?.Constitution === 3 ? "толстый" : 
              organAuto?.Constitution === 4 ? "очень толстый" :
              "норма"}></CardRow>
      </CardContent>
    </Card>
  );
});

export default OrganAutoCard;