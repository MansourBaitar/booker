import {NavLink} from "react-router-dom";
import style from "./navigation.module.scss";
import {FC} from "react";

interface NavigationItemProps {
  path: string;
  icon: string;
  title: string;
}

export const NavigationItem: FC<NavigationItemProps> = (props) => {
  return (
      <NavLink to={props.path} className={(c) => c.isActive ? `${style.active} ${style.navigationItem} ` : style.navigationItem}>
        <div className={style.navigationItemIcon}>
          <i className={props.icon}></i>
        </div>
        <div className={style.navigationItemText}>
          {props.title}
        </div>
      </NavLink>
  )
}