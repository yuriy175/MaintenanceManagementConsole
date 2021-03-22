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

// export default function DicomCard() {
const DicomCard = React.memo((props) => {
  console.log(`! render DicomCard`);

  const classes = useCardsStyles();
  const bull = <span className={classes.bullet}>•</span>;

  const dicom = props.dicom;
  return (
    <Card className={classes.root}>
      <CardContent>
        <Typography variant="h5" component="h2">
          {bull}DICOM
        </Typography>
        <Typography variant="h6" component="h2">
          {bull}PACS
        </Typography>
        {dicom?.PACS?.length ? 
          dicom.PACS.map((i, ind) => (
            <CardRow key={ind.toString()}  
              descr={i.Name + '('+ i.IP +')'} 
              value={ 1? "Не готов" : "Готов"} 
              rightColor={0? "green" : "red"}
            ></CardRow>
            ))
            :
            <></>          
        }
        <Typography variant="h6" component="h2">
          {bull}WorkList
        </Typography>
        {dicom?.WorkList?.length ? 
          dicom.WorkList.map((i, ind) => (
            <CardRow key={ind.toString()}  
              descr={i.Name + '('+ i.IP +')'} 
              value={ 1? "Не готов" : "Готов"} 
              rightColor={0? "green" : "red"}
            ></CardRow>
            ))
            :
            <></>          
        }
      </CardContent>
    </Card>
  );
});

export default DicomCard;