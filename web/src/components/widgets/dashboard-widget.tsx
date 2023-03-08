import {FC} from "react";
import {Card, Typography} from "@mui/material";
import {alpha} from "@mui/material/styles";
import Iconify from "../iconify/iconify";
import {FormatShortenNumber} from "../../common/util/formatter/number";
import style from './widget.module.scss';

interface DashboardWidgetProps {
  title: string;
  icon: string;
  total: number;
  color?: Color;
  sx?: any;
  [x:string]: any;
}

interface Color {
  lighter: string;
  light: string;
  main: string;
  dark: string;
  darker: string;
  contrastText: string;
}

const PRIMARY = {
  lighter: '#D1E9FC',
  light: '#76B0F1',
  main: '#2065D1',
  dark: '#103996',
  darker: '#061B64',
  contrastText: '#fff',
};

export const DashboardWidget: FC<DashboardWidgetProps> = (props) => {
  const color = (props.color) ? props.color : PRIMARY;
  return (
      <Card
          sx={{
            py: 5,
            boxShadow: 0,
            borderRadius: 5,
            textAlign: 'center',
            color: color.darker,
            bgcolor: color.lighter,
            ...props.sx,
          }}
          {...props.other}
      >
        <div className={style.styledWidgetIcon}
            style={{ color: color.dark,
              backgroundImage: `linear-gradient(135deg, ${alpha(color.dark, 0)} 0%, ${alpha(color.dark, 0.24)} 100%)`,
            }}
        >
          <Iconify icon={props.icon} width={24} height={24} />
        </div>

        <Typography variant="h3">{FormatShortenNumber(props.total)}</Typography>

        <Typography variant="subtitle2" sx={{ opacity: 0.72 }}>
          {props.title}
        </Typography>
      </Card>
  );
}

DashboardWidget.defaultProps = {
  color: PRIMARY
}