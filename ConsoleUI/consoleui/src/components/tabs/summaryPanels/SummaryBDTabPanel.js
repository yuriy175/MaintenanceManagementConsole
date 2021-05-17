import React, {useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';

import { CurrentEquipContext } from '../../../context/currentEquip-context';
import CommonTable from '../../tables/CommonTable'

const useStyles = makeStyles((theme) => ({
  root: {
    // display:"flex"
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

  const allDBs = currEquipState.allDBs;

  const getColumn = (key) => { return { id: key, label: key, minWidth: 100, maxWidth: 300 }}

  return (
    <div className={classes.root}>
      {
        allDBs ?
          Object.entries(allDBs).map((tableType, ind) => (
            <>
            <Typography variant="h5" component="h2">
              {tableType[0]}
            </Typography>
            {
              Object.entries(tableType[1]).map((table, ind) => (
                <>
                  <Typography variant="h6" component="h2">
                    {table[0]}
                  </Typography>
                  <CommonTable key={ind.toString()} 
                    columns={table[1].length === 0 ? [] : Object.keys(table[1][0]).map(k => getColumn(k))} 
                    rows={table[1].length === 0 ? [] : table[1]} ></CommonTable>
                </>
              ))
            }
            </>
            ))  : 
          <></>    
      }
    </div>
  );
}