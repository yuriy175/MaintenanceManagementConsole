import React, { useDebugValue } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TablePagination from '@material-ui/core/TablePagination';
import TableRow from '@material-ui/core/TableRow';
import Checkbox from '@material-ui/core/Checkbox';

const useStyles = makeStyles({
  root: {
    width: '100%',
  },
  container: {
    minHeight: 40,
    maxHeight: '100%',
  },
  errorCell:{
    color: 'white',
    background: 'red',
    margin: '0px',
    wordWrap: 'break-word',
  },
  simpleCell:{
    wordWrap: 'break-word',
  },
  boldCell:{
    wordWrap: 'break-word',
    fontWeight: 'bolder',
    fontSize: 'larger',
  },  
  checkBox:{
    color: 'green',
  },  
});

export default function CommonTable(props) {
  const classes = useStyles();
  const [page, setPage] = React.useState(0);
  const [rowsPerPage, setRowsPerPage] = React.useState(10);

  const handleChangePage = (event, newPage) => {
    setPage(newPage);
  };

  const handleChangeRowsPerPage = (event) => {
    setRowsPerPage(+event.target.value);
    setPage(0);
  };

  const rows = props.rows ? props.rows : [];
  const columns = props.columns;
  const selectedRow = props.selectedRow;
  let onRowClick = props.onRowClick;
  if(!onRowClick){
    onRowClick = (ev, row) => {};
  }
  const isRowBold = props.isRowBold;

  return (
    <Paper className={classes.root}>
      <TableContainer className={classes.container}>
        <Table stickyHeader aria-label="sticky table" size="small">
          <TableHead>
            <TableRow>
              {columns.map((column) => (
                <TableCell
                  key={column.id}
                  align={column.align}
                  style={{ minWidth: column.minWidth, maxWidth: column.maxWidth}}  
                >
                  {column.label}
                </TableCell>
              ))}
            </TableRow>
          </TableHead>
          <TableBody>
            {rows.slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage).map((row) => {
              return (
                <TableRow 
                    selected={selectedRow ? row.EquipName === selectedRow : false} 
                    hover 
                    role="checkbox" 
                    tabIndex={-1} 
                    key={row.code} 
                    onClick={(ev) => onRowClick(ev, row)} >
                  {columns.map((column) => {
                    const value = row[column.id];
                    const checked = column.format ? column.format(row) : value;
                    return (
                      <TableCell key={column.id} align={column.align}
                        className={
                          column.hasErrors && column.hasErrors(value) ? 
                            classes.errorCell : 
                            isRowBold && isRowBold(row) ?
                              classes.boldCell :
                              classes.simpleCell
                        }
                      >
                        <div style={{ maxWidth: column.maxWidth}}>
                        {column.checked ? 
                          <Checkbox
                            style ={{
                              color: checked ? 'green' : 'gray', //"#00e676",
                            }}
                            checked={checked}
                            onChange={(ev) => props.onSelect ? props.onSelect(ev, row) : false}
                            inputProps={{ 'aria-label': 'select all desserts' }}
                          /> : 
                          column.format ? column.format(value) : 
                            column.formatArray ? column.formatArray(value).map(v => <pre>{v}</pre>) : value}
                        </div>  
                      </TableCell>
                    );
                  })}
                </TableRow>
              );
            })}
            
          </TableBody>
        </Table>
      </TableContainer>
      <TablePagination
        rowsPerPageOptions={[10, 25, 100]}
        component="div"
        count={rows?.length}
        rowsPerPage={rowsPerPage}
        page={page}
        onChangePage={handleChangePage}
        onChangeRowsPerPage={handleChangeRowsPerPage}
      />
    </Paper>
  );
}