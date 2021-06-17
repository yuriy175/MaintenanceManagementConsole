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
import { ComplexTypeImages } from '../../model/constants'

import {useCardsStyles} from './CommonCard'

const useStyles = makeStyles((theme) => ({
  media: {
    height: 0,
    paddingTop: '100%', 
    backgroundColor: 'gray',
  },
  dimmed:{
    height: 0,
    paddingTop: '100%', 
    backgroundColor: 'gray',
    filter: 'brightness(40%)',
  }
}));

// export default function EquipImageCard() {
const EquipImageCard = React.memo((props) => {
  console.log(`! render EquipImageCard`);

  const classes = useCardsStyles();
  const equipClasses = useStyles();

  const bull = <span className={classes.bullet}>•</span>;

  const equipInfo = props.equipInfo;
  const pathKey = Object.keys(ComplexTypeImages).find(k => 
    {
      return ComplexTypeImages[k].some(v => equipInfo.startsWith(v));
    });
  
  const isConnected = props.isConnected;

  return (
    <Card className={classes.root}>
      <CardMedia
        className={isConnected ? equipClasses.media : equipClasses.dimmed}
        // image="./girl.jpg"
        image={"./" + pathKey}
        title="Аппарат"
      />
      <CardContent>
        <Typography variant="body2" color="textSecondary" component="p">
          Аппарат : {equipInfo}
        </Typography>
      </CardContent>
    </Card>
  );
});

export default EquipImageCard;