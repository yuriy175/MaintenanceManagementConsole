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
import TableSortLabel from '@material-ui/core/TableSortLabel';
import Button from '@material-ui/core/Button';

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

function descendingComparator(a, b, orderBy) {
  if (b[orderBy] < a[orderBy]) {
    return -1;
  }
  if (b[orderBy] > a[orderBy]) {
    return 1;
  }
  return 0;
}

function getComparator(order, orderBy) {
  return order === 'desc'
    ? (a, b) => descendingComparator(a, b, orderBy)
    : (a, b) => -descendingComparator(a, b, orderBy);
}

function stableSort(array, comparator) {
  const stabilizedThis = array.map((el, index) => [el, index]);
  stabilizedThis.sort((a, b) => {
    const order = comparator(a[0], b[0]);
    if (order !== 0) return order;
    return a[1] - b[1];
  });
  return stabilizedThis.map((el) => el[0]);
}

export default function CommonTable(props) {
  const classes = useStyles();
  const [page, setPage] = React.useState(0);
  const [rowsPerPage, setRowsPerPage] = React.useState(100);
  const [orderBy, setOrderBy] = React.useState(props.defaultSort ?? '');
  const [order, setOrder] = React.useState( props.defaultSortOrder ?? 'asc');

  const handleChangePage = (event, newPage) => {
    setPage(newPage);
  };

  const handleChangeRowsPerPage = (event) => {
    setRowsPerPage(+event.target.value);
    setPage(0);
  };

  const handleRequestSort = (property) => () => {
    const isAsc = orderBy === property && order === 'asc';
    setOrder(isAsc ? 'desc' : 'asc');
    setOrderBy(property);
  };

  const rows = stableSort(props.rows ? props.rows : [], getComparator(order, orderBy)) ?? [];
  if(rows.length < page * rowsPerPage){
    setPage(0);
  }

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
                // <TableCell
                //   key={column.id}
                //   align={column.align}
                //   style={{ minWidth: column.minWidth, maxWidth: column.maxWidth}}  
                // >
                //   {column.label}
                // </TableCell>
                <TableCell
                  key={column.id}
                  align={column.align}
                  style={{ minWidth: column.minWidth, maxWidth: column.maxWidth}}  
                  sortDirection={orderBy === column.id ? order : false}
                >
                  {!column.sortable? 
                    column.label:
                    <TableSortLabel
                      active={orderBy === column.id}
                      direction={orderBy === column.id ? order : 'asc'}
                      onClick={handleRequestSort(column.id)}
                    >
                      {column.label}                    
                    </TableSortLabel>
                  }
                </TableCell>
              ))}
            </TableRow>
          </TableHead>
          <TableBody>
            {rows.slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage).map((row, rowInd) => {
              return (
                <TableRow 
                    selected={selectedRow ? row.EquipName === selectedRow : false} 
                    hover 
                    role="checkbox" 
                    tabIndex={-1} 
                    key={rowInd} 
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
                            inputProps={{ 'aria-label': 'select all desserts', 'data-column' : column.id }}
                          /> : 
                          column.format ? column.format(value) : 
                            column.formatArray ? column.formatArray(value).map(v => <pre>{v}</pre>) : 

                            column.button ? 
                              <Button variant="contained" color="primary" className={classes.commonSpacing} onClick={() => column.onEdit(row)}>
                                    {column.label}
                              </Button> :
                              value}
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