import React, {useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';

import { CurrentEquipContext } from '../../../context/currentEquip-context';


const useStyles = makeStyles((theme) => ({
  root: {
    display:"flex"
  },
  column:{
    width: "30%",
    marginRight: "12px",
  }
}));

export default function SummaryBDTabPanel(props) {
  console.log("render SummaryBDTabPanel");

  const classes = useStyles();
  const [currEquipState, currEquipDispatch] = useContext(CurrentEquipContext);

  return (
    <div className={classes.root}>
    </div>
  );
}