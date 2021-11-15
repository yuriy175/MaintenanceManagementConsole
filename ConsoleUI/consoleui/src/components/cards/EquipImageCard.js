import React, {useContext, useEffect} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import CardMedia from '@material-ui/core/CardMedia';

import {MoscowTimeZone} from '../../model/constants'
import {parseLocalString} from '../../utilities/utils'

import { CurrentEquipContext } from '../../context/currentEquip-context';
import CardRow from './CardRow'
import { ComplexTypeImages } from '../../model/constants'

import {useCardsStyles} from './CommonCard'
import {getRussianRegions} from '../../utilities/rusRegions'

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
  const [date, setDate] = React.useState(new Date());
  const [regionsList] = React.useState(getRussianRegions());

  const region = props.region;

  useEffect(() => {    
    const timer = setInterval(()=>
    {
      const currDate = new Date();
      const timezone = region?.Region ? regionsList?.filter(r => 
        {
          return r[1] === region?.Region;
        })?.[0]?.[2] ?? MoscowTimeZone : MoscowTimeZone;
      currDate.setHours( currDate.getHours() + timezone - MoscowTimeZone);
      setDate(currDate);
    }, 2000 )
    return function cleanup() {
        clearInterval(timer)
    }
  }, [props.region]);

  const bull = <span className={classes.bullet}>•</span>;

  const equipInfo = props.equipInfo;
  const pathKey = Object.keys(ComplexTypeImages).find(k => 
    {
      return ComplexTypeImages[k].some(v => equipInfo?.startsWith(v));
    });
  
  const isConnected = props.isConnected;
  const lastSeen = props.lastSeen;
  const hospital = props.hospital;
  const address = props.address;
  const alias = props.alias;
  
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
          Аппарат : { alias ? `${alias} (${equipInfo})`: equipInfo}
          {hospital? <div>ЛПУ: {hospital} </div>: <></>}
          {address? <div>Адрес: {address} </div>: <></>}
          {lastSeen? <div>Посл. сообщение: {parseLocalString(lastSeen)} </div>: <></>}
          {region? <div>Регион: {region.Region ?? 'не указан'}, Время: {date.toLocaleTimeString()}</div>: <></>}
        </Typography>
      </CardContent>
    </Card>
  );
});

export default EquipImageCard;