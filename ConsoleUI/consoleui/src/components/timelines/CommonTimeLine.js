import React, {useState} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import Typography from '@material-ui/core/Typography';
import Divider from '@material-ui/core/Divider';
import Checkbox from '@material-ui/core/Checkbox';

import {parseLocalString, isToday} from '../../utilities/utils'
import TimeLineItem from './TimeLineItem';

const useStyles = makeStyles((theme) => ({
  root: {
    width: '100%',
    backgroundColor: theme.palette.background.paper,
  },
}));

export default function CommonTimeLine(props) {
  const classes = useStyles();
  
  const equipName = props.equipName;
  const showImportantOnly = props.showImportantOnly;
  const  rows = props.rows;
  const todayRows = rows?.filter(i => isToday(i.DateTime)) ?? [];
  const todayCount = todayRows.length;
  const othersCount = rows ? rows.length - todayCount : 0;

  return (
    <List className={classes.root}>
        <Typography
          component="span"
          variant="h5"
          color="textPrimary"
          >
            Сегодня ({todayCount} сообщений)
        </Typography>
        { todayRows?.map((i, ind) =>
          {
            const isImportant=i.Type?.includes('Error');
            const isOffline=i.Type?.includes('Offline');
            return (
              isImportant || !showImportantOnly ?
                <TimeLineItem key={ind.toString()} 
                    isImportant={isImportant}
                    equipName={i.EquipName} 
                    title={i.Title + (isOffline ? '(Offline)' : '')} 
                    text={i.Description} 
                    time={parseLocalString(i.DateTime)}
                    details={i.Details}/> : <></>        
              )
           })
        }
        <Divider />
        <Typography
          component="span"
          variant="h5"
          color="textPrimary"
          >
            Все время ({othersCount} сообщений)
        </Typography>
        { rows?.filter(i => !isToday(i.DateTime))?.map((i, ind) =>
          {
            const isImportant=i.Type?.includes('Error');
            return (
              isImportant || !showImportantOnly ?
                <TimeLineItem key={ind.toString()} 
                    isImportant={isImportant}
                    equipName={i.EquipName} 
                    title={i.Title} 
                    text={i.Description} 
                    time={parseLocalString(i.DateTime)}
                    details={i.Details}/> : <></>        
              )
           })
        }
    </List>
  );
}