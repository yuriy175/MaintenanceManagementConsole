import React, {useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import CardMedia from '@material-ui/core/CardMedia';

import { CurrentEquipContext } from '../../context/currentEquip-context';
import CardRow from './CardRow'

import {useCardsStyles} from './CommonCard'

const useStyles = makeStyles((theme) => ({
  media: {
    height: 0,
    paddingTop: '100%', 
  },
}));

export default function EquipImageCard() {
  console.log(`! render EquipImageCard`);

  const classes = useCardsStyles();
  const equipClasses = useStyles();
  const [currEquipState, currEquipDispatch] = useContext(CurrentEquipContext);

  const bull = <span className={classes.bullet}>•</span>;

  return (
    <Card className={classes.root}>
      <CardMedia
        className={equipClasses.media}
        image="./girl.jpg"
        title="Аппарат"
      />
      <CardContent>
        <Typography variant="body2" color="textSecondary" component="p">
          Аппарат : {currEquipState.equipInfo}
        </Typography>
      </CardContent>
    </Card>
  );
}