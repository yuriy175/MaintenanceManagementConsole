import React, { useState, useEffect, useRef, useContext } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import FormHelperText from '@material-ui/core/FormHelperText';
import FormControl from '@material-ui/core/FormControl';
import NativeSelect from '@material-ui/core/NativeSelect';

import { AllEquipsContext } from '../../context/allEquips-context';
import { CurrentEquipContext } from '../../context/currentEquip-context';

import * as EquipWorker from '../../workers/equipWorker'
// import * as WebSocket from '../../workers/webSocket'
import {sessionUid} from '../../utilities/utils'
import { useWebSocket } from '../../workers/useWebSocket'

const drawerWidth = 240;

const useStyles = makeStyles((theme) => ({
  appBar: {
    width: `calc(100% - ${drawerWidth}px)`,
    marginLeft: drawerWidth,
  },
  formControl: {
    margin: theme.spacing(1),
    minWidth: 120,
  },
  selectEmpty: {
    // marginTop: theme.spacing(2),
    color: "white",
  },
  optionStyle:{
    backgroundColor: "#3f51b5",
    color:"black",
  }
}));

export default function MainToolBar() {
  console.log(`! render MainToolBar ` + sessionUid);

  const classes = useStyles();
  const [allEquipsState, allEquipsDispatch] = useContext(AllEquipsContext);
  const [currEquipState, currEquipDispatch] = useContext(CurrentEquipContext);
  const [currEquip, setCurrEquip] = useState('none');

  const handleEquipsChange = async (event) => {
    const select = event.target;
    const val = select.options[select.selectedIndex].value;

    await onEquipChanged(val);
  };

  const onEquipChanged = async equipInfo =>
  {
    currEquipDispatch({ type: 'RESET', payload: true });    
    currEquipDispatch({ type: 'SETEQUIPINFO', payload: equipInfo }); 

    // new software & system info come very slowly
    const sysInfo = await EquipWorker.GetPermanentData("SystemInfo", equipInfo);
    currEquipDispatch({ type: 'SETSYSTEM', payload: sysInfo }); 

    await EquipWorker.Activate(equipInfo, currEquipState.equipInfo);
  }
  
  // useEffect(() => {
  //     (async () => {
  //         if(allEquipsState.equips !== null)
  //         {
  //           return;
  //         }

  //         const equips = await EquipWorker.GetAllEquips();
  //         allEquipsDispatch({ type: 'SETEQUIPS', payload: equips ? equips : [] });   
  //         if(!equips || equips?.length === 0)     
  //         {
  //           return;
  //         }

  //         const equipInfo = equips[0];
  //         onEquipChanged(equipInfo);
  //     })();
  // }, [allEquipsState.equips]);


  const webSocket = useWebSocket(
    {
    }
  );


  return (    
    <AppBar position="fixed" className={classes.appBar}>
        <Toolbar>
            <Typography variant="h6" noWrap>
            </Typography>
            <FormControl className={classes.formControl}>
              <NativeSelect
                value={currEquipState.equipInfo}
                onChange={handleEquipsChange}
                name="equips"
                className={classes.selectEmpty}
                variant="outlined"
              >
                {allEquipsState.equips?.map((i, ind) => (
                    <option key={ind.toString()} value={i} className={classes.optionStyle}>{i}</option>
                    ))
                }
              </NativeSelect>
            </FormControl>
        </Toolbar>
    </AppBar>
  );
}