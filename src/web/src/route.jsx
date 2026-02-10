import { HelmetProvider } from '@dr.pogodin/react-helmet';
import { BrowserRouter } from 'react-router';
import { GenerateRoutes } from '@/router/Rules.jsx';

const AdminRoute = () => {
  return (
    <HelmetProvider>
      <BrowserRouter>
        <GenerateRoutes />
      </BrowserRouter>
    </HelmetProvider>
  );
};

export default AdminRoute;
