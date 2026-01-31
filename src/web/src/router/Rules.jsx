import React from 'react';
import AdminLayout from '@/layout/Admin.jsx';
import LoginAndErrorLayout from '@/layout/LoginAndError.jsx';
import RouterLazyLoad from '@/router/LazyLoad.jsx';
import { Navigate, useRoutes } from 'react-router';

// 路由对象
export const Rules = [
  {
    path: '/',
    element: <Navigate to="/dashboard" />
  },
  {
    path: '/',
    element: <AdminLayout />,
    children: [
      {
        path: '/dashboard',
        element: RouterLazyLoad(React.lazy(() => import('@/page/default/Default.jsx')))
      },
      {
        path: '/cluster',
        children: [
          {
            path: '/cluster/overview',
            element: RouterLazyLoad(React.lazy(() => import('@/page/default/Default.jsx')))
          },
          {
            path: '/cluster/join',
            element: RouterLazyLoad(React.lazy(() => import('@/page/default/Default.jsx')))
          },
        ]
      },
      {
        path: '/kubernetes',
        children: [
          {
            path: '/kubernetes/overview',
            element: RouterLazyLoad(React.lazy(() => import('@/page/default/Default.jsx')))
          },
          {
            path: '/kubernetes/node',
            element: RouterLazyLoad(React.lazy(() => import('@/page/default/Default.jsx')))
          },
          {
            path: '/kubernetes/namespace',
            element: RouterLazyLoad(React.lazy(() => import('@/page/default/Default.jsx')))
          },
          {
            path: '/kubernetes/pod',
            element: RouterLazyLoad(React.lazy(() => import('@/page/default/Default.jsx')))
          },
          {
            path: '/kubernetes/workload',
            element: RouterLazyLoad(React.lazy(() => import('@/page/default/Default.jsx')))
          },
          {
            path: '/kubernetes/service',
            element: RouterLazyLoad(React.lazy(() => import('@/page/default/Default.jsx')))
          },
          {
            path: '/kubernetes/ingress',
            element: RouterLazyLoad(React.lazy(() => import('@/page/default/Default.jsx')))
          },
          {
            path: '/kubernetes/configmap-and-secret',
            element: RouterLazyLoad(React.lazy(() => import('@/page/default/Default.jsx')))
          },
          {
            path: '/kubernetes/storage',
            element: RouterLazyLoad(React.lazy(() => import('@/page/default/Default.jsx')))
          },
        ]
      },
      {
        path: '/system',
        children: [
          {
            path: '/system/user',
            element: RouterLazyLoad(React.lazy(() => import('@/page/default/Default.jsx')))
          },
          {
            path: '/system/role',
            element: RouterLazyLoad(React.lazy(() => import('@/page/default/Default.jsx')))
          },
          {
            path: '/system/menu',
            element: RouterLazyLoad(React.lazy(() => import('@/page/default/Default.jsx')))
          },
          {
            path: '/system/api',
            element: RouterLazyLoad(React.lazy(() => import('@/page/default/Default.jsx')))
          },
          {
            path: '/system/setting',
            element: RouterLazyLoad(React.lazy(() => import('@/page/default/Default.jsx')))
          }
        ]
      },
      {
        path: '/information',
        element: RouterLazyLoad(React.lazy(() => import('@/page/default/Default.jsx')))
      }
    ]
  },
  {
    path: '/',
    element: <LoginAndErrorLayout />,
    children: [
      {
        path: '/login',
        element: RouterLazyLoad(React.lazy(() => import('@/page/default/Default.jsx')))
      },
      {
        path: '/403',
        element: RouterLazyLoad(React.lazy(() => import('@/page/error/403/403.jsx')))
      },
      {
        path: '/404',
        element: RouterLazyLoad(React.lazy(() => import('@/page/error/404/404.jsx')))
      },
      {
        path: '/500',
        element: RouterLazyLoad(React.lazy(() => import('@/page/error/500/500.jsx')))
      }
    ]
  },
  {
    path: '*',
    element: <Navigate to="/404" />
  }
];

// 生成路由
export const GenerateRoutes = () => {
  return useRoutes(Rules);
};
