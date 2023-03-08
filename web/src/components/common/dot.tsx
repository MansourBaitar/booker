import {FC} from "react";
import {Box} from "@mui/material";
import {ERROR, INFO, PRIMARY, WARNING} from "../../common/colors";

interface DotProps {
  color: string;
  size?: string;
}

export const Dot: FC<DotProps> = (props) => {
  let main;
  switch (props.color) {
    case 'secondary':
      main = PRIMARY.main;
      break;
    case 'error':
      main = ERROR.main;
      break;
    case 'warning':
      main = WARNING.main;
      break;
    case 'info':
      main = INFO.main;
      break;
    case 'success':
      main = "#2ecc71";
      break;
    case 'primary':
    default:
      main = PRIMARY.main;
  }

  return (
      <Box
          sx={{
            width: props.size || 8,
            height: props.size || 8,
            borderRadius: '50%',
            bgcolor: main
          }}
      />
  );
}