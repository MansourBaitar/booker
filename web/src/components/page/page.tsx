import React, {FC, useState} from "react";
import style from './page.module.scss';
import {Navigation} from "../navigation/navigation";
import {Outlet} from "react-router-dom";
import {NavHeader} from "../header/header";

export const Page: FC = (props) => {
  const [open, setOpen] = useState(false);

  return (
      <div className={style.outer}>
        <header className={style.header}>
          <NavHeader onOpen={() => setOpen(true)} />
        </header>
        <div className={style.elevated}>
          <Navigation open={open} onClose={() => setOpen(false)}/>
          <main className={style.page}>
            <Outlet/>
          </main>
        </div>
    </div>
  );
}
