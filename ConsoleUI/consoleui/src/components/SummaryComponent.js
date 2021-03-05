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

const useStyles = makeStyles((theme) => ({
}));

export default function SummaryComponent(props) {
  console.log("render SummaryComponent");

  const classes = useStyles();

  return (
    <div>
        <EquipImageCard></EquipImageCard>
        <GeneratorCard></GeneratorCard>
        <DetectorCard></DetectorCard>
        <StandCard></StandCard>
        <SystemCard></SystemCard>
        <HddCard></HddCard>
        <OrganAutoCard></OrganAutoCard>
        <RemoteAccessCard></RemoteAccessCard>
        <DicomCard></DicomCard>
    </div>
  );
}