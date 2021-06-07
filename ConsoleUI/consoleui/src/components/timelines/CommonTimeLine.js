import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import {parseLocalString} from '../../utilities/utils'
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
  const  rows = props.rows;
  return (
    <List className={classes.root}>
        { rows?.map((i, ind) => (
            <TimeLineItem key={ind.toString()} 
                isImportant={i.Type?.includes('Error')}
                equipName={i.EquipName} 
                title={i.Title} 
                text={i.Description} 
                time={parseLocalString(i.DateTime)}></TimeLineItem>            
            ))
        }
    </List>
  );
}