import {FC} from "react";
import {Stack, Typography} from "@mui/material";
import {Dot} from "./dot";

interface StatusPillProps {
  status: number;
}

export const StatusPill: FC<StatusPillProps> = (props) => {
  let color;
  let title;

  switch (props.status) {
    case 0:
      color = 'warning';
      title = 'Pending';
      break;
    case 1:
      color = 'success';
      title = 'Approved';
      break;
    case 2:
      color = 'error';
      title = 'Rejected';
      break;
    default:
      color = 'primary';
      title = 'None';
  }

  return (
      <Stack direction="row" spacing={1} alignItems="center">
        <Dot color={color} />
        <Typography>{title}</Typography>
      </Stack>
  );
}