import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import DetectorCard from './cards/DetectorCard'
import GeneratorCard from './cards/GeneratorCard'

const useStyles = makeStyles((theme) => ({
}));

export default function SummaryComponent(props) {
  console.log("render SummaryComponent");

  const classes = useStyles();

  return (
    <div>
        <GeneratorCard></GeneratorCard>
        <DetectorCard></DetectorCard>
    </div>
  );
}