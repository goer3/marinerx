// 开放接口前缀
const OpenApiPrefix = import.meta.env.VITE_BACKEND_URL + '/openapi/v1';

// 认证接口前缀
const ApiPrefix = import.meta.env.VITE_BACKEND_URL + '/api/v1';

// 后端节后汇总
export const API = {
  Open: {
    Health: { URL: OpenApiPrefix + '/health', Method: 'GET', Description: '服务健康检查' },
    Information: { URL: OpenApiPrefix + '/information', Method: 'GET', Description: '服务信息' },
    Version: { URL: OpenApiPrefix + '/version', Method: 'GET', Description: '服务版本' }
  },
  System: {
    User: {
      List: { URL: ApiPrefix + '/system/user/list', Method: 'GET', Description: '获取用户列表' },
      Detail: { URL: ApiPrefix + '/system/user/detail', Method: 'GET', Description: '获取用户详情' }
    },
    Role: {
      List: { URL: ApiPrefix + '/system/role/list', Method: 'GET', Description: '获取角色列表' },
      Detail: { URL: ApiPrefix + '/system/role/detail', Method: 'GET', Description: '获取角色详情' }
    }
  }
};
