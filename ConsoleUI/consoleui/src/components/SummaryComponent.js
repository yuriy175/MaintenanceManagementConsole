import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
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

  return (
    <div className={classes.root}>
      <div className={classes.column}>
        <EquipImageCard></EquipImageCard>
        <SystemCard></SystemCard>
        <HddCard></HddCard>
      </div>
      <div className={classes.column}>
        <OrganAutoCard></OrganAutoCard>
        <GeneratorCard></GeneratorCard>
        <DetectorCard></DetectorCard>
        <StandCard></StandCard>
        <DosimeterCard></DosimeterCard>
      </div>
      <div className={classes.column}>
        <RemoteAccessCard></RemoteAccessCard>
        <DicomCard></DicomCard>
        <SoftwareCard></SoftwareCard>
      </div>
    </div>
  );
}