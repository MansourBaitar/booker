import {FC, useEffect} from "react";
import style from './navigation.module.scss';
import {NavigationItem} from "./navigation-item";
import {useLocation} from "react-router-dom";
import useResponsive from "../../hooks/useResponsive";
import {Box, Drawer} from "@mui/material";

interface NavigationProps {
  open: boolean;
  onClose: () => void;
}

export const Navigation: FC<NavigationProps> = (props) => {
  const { pathname } = useLocation();
  const isDesktop = useResponsive('up', 'lg');
  const NAV_WIDTH = 280;

  useEffect(() => {
    if(props.open) {
      props.onClose();
    }
  }, [pathname]);

  return (
      <>
        <Box component="nav" sx={{ flexShrink: { lg: 0 },  width: { lg: NAV_WIDTH } }}>
        {isDesktop ? (
            <Drawer open variant="permanent" PaperProps={{sx: { width: NAV_WIDTH, bgcolor: 'background.default'} }}>
              <div className={style.wrapper}>
                <header className={style.logo}>
                  LOGO
                </header>
                <nav className={style.navigation}>
                  <NavigationItem title={'Dashboard'} path={'/'} icon={'ri-dashboard-2-line'} />
                  <NavigationItem title={'Requests'} path={'/requests'} icon={'ri-flashlight-line'} />
                  <NavigationItem title={'Holidays'} path={'/holidays'} icon={'ri-sun-line'} />
                  <NavigationItem title={'Absences'} path={'/absences'} icon={'ri-first-aid-kit-line'} />

                </nav>
              </div>
            </Drawer>

        ): (
          <Drawer open={props.open} onClose={props.onClose} ModalProps={{ keepMounted: true }} PaperProps={{ sx: { width: NAV_WIDTH } }}>
            <div className={style.wrapper}>
              <header className={style.logo}>
                LOGO
              </header>
              <nav className={style.navigation}>
                <NavigationItem title={'Dashboard'} path={'/'} icon={'ri-dashboard-2-line'} />
                <NavigationItem title={'Holidays'} path={'/holidays'} icon={'ri-sun-line'} />
                <NavigationItem title={'Absences'} path={'/Absences'} icon={'ri-first-aid-kit-line'} />
              </nav>
            </div>
          </Drawer>
        )}
        </Box>
      </>
  );
};