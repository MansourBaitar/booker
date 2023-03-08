import {BrowserRouter, Route, Routes} from "react-router-dom";
import {FC} from "react";
import {DashboardPage} from "../pages/dashboard.page";
import {HolidayPage} from "../pages/holiday.page";
import {AbsencePage} from "../pages/absence.page";
import {AuthPage, Page} from "../components";
import {LoginPage} from "../pages/auth/login.page";
import {NotFoundPage} from "../pages/404";
import {AccountPage} from "../pages/account/account.page";
import {RequestPage} from "../pages/request.page";


export const Router: FC = () => {
  return <BrowserRouter>
    <Routes>
      <Route path="/" element={<Page />}>
        <Route index element={<DashboardPage/>} />
        <Route path={"/holidays"} element={<HolidayPage/>} />
        <Route path={"/absences"} element={<AbsencePage/>} />
        <Route path={"/requests"} element={<RequestPage/>} />
        <Route path={"/account/profile"} element={<AccountPage/>} />

      </Route>

      <Route path="/auth">
        <Route path={"/auth/login"} element={<LoginPage/>} />
      </Route>

      {/*Not found routes*/}
      <Route path="*" element={<AuthPage />}>
        <Route path="*" element={<NotFoundPage/>} />
      </Route>
    </Routes>
  </BrowserRouter>
}
