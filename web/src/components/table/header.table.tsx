import {FC} from "react";
import {TableCell, TableHead, TableRow} from "@mui/material";

interface TableHeaderProps {
  order: any;
  orderBy: any;
  headCellsData: any[];
}

export const TableHeader: FC<TableHeaderProps> =  (props) => {
  return (
      <TableHead>
        <TableRow>
          {props.headCellsData.map((headCell) => (
              <TableCell
                  key={headCell.id}
                  align={headCell.align}
                  padding={headCell.disablePadding ? 'none' : 'normal'}
                  sortDirection={props.orderBy === headCell.id ? props.order : false}
              >
                {headCell.label}
              </TableCell>
          ))}
        </TableRow>
      </TableHead>
  )
}