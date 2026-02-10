// 获取本地保存的登录信息，校验是否正确
export const GetToken = () => {
  const token = localStorage.getItem('token');
  const tokenExpire = localStorage.getItem('token-expire');
  const deviceId = localStorage.getItem('device-id');

  // 判断必要参数是否存在
  if (!token || !tokenExpire || !deviceId) {
    return null;
  }

  // 判断 token 是否过期
  if (Date.now() > Date.parse(tokenExpire)) {
    localStorage.clear();
    return null;
  }

  return { token, deviceId };
};

// 保存登录信息到本地
export const SetToken = (token, tokenExpire) => {
    localStorage.setItem('token', token);
    localStorage.setItem('token-expire', tokenExpire);
}

// 清除本地保存的登录信息
export const ClearToken = () => {
  localStorage.clear();
}
