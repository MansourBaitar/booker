import {forwardRef} from "react";
import { Icon } from '@iconify/react';
import {Box} from "@mui/material";

interface IconifyProps {
  icon: any;
  width?: number;
  sx?: any
  [x:string]: any
}

const Iconify = forwardRef(({ icon, width = 20, sx, ...other }: IconifyProps, ref) => (
    <Box ref={ref} component={Icon} icon={icon} sx={{ width, height: width, ...sx }} {...other} />
));

export default Iconify;