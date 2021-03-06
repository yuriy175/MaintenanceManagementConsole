import React from 'react';
import { makeStyles } from '@material-ui/core/styles';

export const useCardsStyles = makeStyles({
  root: {
    width: "100%",//275,
    //maxWidth: 345,
    marginBottom: 12,
    marginRight: 12,
  },
  bullet: {
    display: 'inline-block',
    margin: '0 2px',
    transform: 'scale(0.8)',
  },
  title: {
    fontSize: 14,
  },
  pos: {
    marginBottom: 12,
  },
});
