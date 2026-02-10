import { Helmet } from '@dr.pogodin/react-helmet';
import { TitleSuffix } from '@/components/Text';

// 页面配置
const config = {
  title: '默认页面'
};

const Dashboard = () => {
  return (
    <>
      <Helmet>
        <title>{config.title + TitleSuffix}</title>
      </Helmet>
      <h2>{import.meta.env.VITE_ENV}</h2>
    </>
  );
};

export default Dashboard;
