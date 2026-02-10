import axios from 'axios';
import { GetToken } from '@/utils/Token';

// 创建 Axios 配置实例，timeout 默认设置为 10000 毫秒，但是可以在请求时覆盖
const instance = axios.create({
  timeout: 10000
});

// 请求拦截器，需要在请求头中添加 Token
instance.interceptors.request.use(
  (config) => {
    const data = GetToken();
    if (data) {
      config.headers.Authorization = `Bearer ${data.token}`;
      config.headers.DeviceId = data.deviceId;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// 响应拦截器，处理响应数据
instance.interceptors.response.use(
  (response) => {
    return response.data;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// config 支持的参数示例：
// headers: {'X-Requested-With': 'XMLHttpRequest'}
// timeout: 5000
// HTTP Basic Auth
// auth: {
//   username: 'admin',
//   password: '123456'
// }
const request = (method, url, data = {}, config = {}) => {
  const cfg = {
    method,
    url,
    ...config
  };

  if (method === 'get') {
    cfg.params = data;
  } else {
    cfg.data = data;
  }

  return instance(cfg);
};

// GET 请求封装
export const GET = (url, params = {}, config = {}) => request('get', url, params, config);

// POST 请求封装
export const POST = (url, data = {}, config = {}) => request('post', url, data, config);

// PUT 请求封装
export const PUT = (url, data = {}, config = {}) => request('put', url, data, config);

// PATCH 请求封装
export const PATCH = (url, data = {}, config = {}) => request('patch', url, data, config);

// DELETE 请求封装
export const DELETE = (url, params = {}, config = {}) => request('delete', url, params, config);
