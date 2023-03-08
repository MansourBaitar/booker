import {FC} from "react";
import style from './auth.module.scss';
import {Outlet} from "react-router-dom";
import DsyncLogoBlack from '../../assets/dsync.svg';

export const AuthPage: FC = (props) => {
  return (
      <>
        <Outlet />
      </>
  );
}