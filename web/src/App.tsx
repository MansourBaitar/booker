import {Router} from "./routes/router";
import { LicenseInfo } from '@mui/x-license-pro';

export const App = () => {
  LicenseInfo.setLicenseKey(
      '',
  );
  return (
        <Router />
  );
}
