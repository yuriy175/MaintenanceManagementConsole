import React, {useContext} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { CurrentEquipContext } from '../context/currentEquip-context';

import DetectorCard from './cards/DetectorCard'
import GeneratorCard from './cards/GeneratorCard'
import SystemCard from './cards/SystemCard'
import HddCard from './cards/HddCard'
import OrganAutoCard from './cards/OrganAutoCard'
import EquipImageCard from './cards/EquipImageCard'
import DicomCard from './cards/DicomCard'
import RemoteAccessCard from './cards/RemoteAccessCard'
import StandCard from './cards/StandCard'
import DosimeterCard from './cards/DosimeterCard'
import SoftwareCard from './cards/SoftwareCard'

const useStyles = makeStyles((theme) => ({
  root: {
    display:"flex"
  },
  column:{
    width: "30%",
    marginRight: "12px",
  }
}));

export default function SummaryComponent(props) {
  console.log("render SummaryComponent");

  const classes = useStyles();
  const [currEquipState, currEquipDispatch] = useContext(CurrentEquipContext);

  return (
    <div className={classes.root}>
      <div className={classes.column}>
        <EquipImageCard equipInfo={currEquipState.equipInfo}></EquipImageCard>
        <SystemCard system={currEquipState.system}></SystemCard>
        {/* <HddCard></HddCard> */}
      </div>
      <div className={classes.column}>
        <OrganAutoCard organAuto={currEquipState.organAuto}></OrganAutoCard>
        <GeneratorCard generator={currEquipState.generator}></GeneratorCard>
        <DetectorCard detectors={currEquipState.detectors}></DetectorCard>
        <StandCard stand={currEquipState.stand}></StandCard>
        <DosimeterCard dosimeter={currEquipState.dosimeter}></DosimeterCard>
      </div>
      <div className={classes.column}>
        <RemoteAccessCard equipInfo={currEquipState.equipInfo} remoteaccess={currEquipState.remoteaccess}></RemoteAccessCard>
        <DicomCard dicom={currEquipState.dicom}></DicomCard>
        <SoftwareCard software={currEquipState.software}></SoftwareCard>
      </div>
    </div>
  );
}