import React, {useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';

import { CurrentEquipContext } from '../../context/currentEquip-context';
import {useCardsStyles} from './CommonCard'

export default function HddCard() {
  console.log(`! render HddCard`);

  const [currEquipState, currEquipDispatch] = useContext(CurrentEquipContext);
  const classes = useCardsStyles();
  const bull = <span className={classes.bullet}>•</span>;

  return (
    <Card className={classes.root}>
      <CardContent>        
        {currEquipState.hdd?.map((i, ind) => (
          <div key={ind.toString()} >
            <Typography variant="h5" component="h2">
              {bull}Диск {i.Letter}
            </Typography>
            <Typography className={classes.pos} color="textSecondary">
              Свободно: {i.FreeSize}Гб, Всего: {i.TotalSize}Гб
            </Typography>
          </div>
          ))
        }
      </CardContent>
    </Card>
  );
}