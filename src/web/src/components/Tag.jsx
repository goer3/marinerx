import { Tag } from 'antd';

// 通用标签生成函数
// 参数：value (值), map (映射对象), options (选项)
// 返回：标签组件
const GenerateTag = (value, map, options = {}) => {
  const { variant = 'filled', size = 'small', defaultKey = 'default', textKey = 'text', colorKey = 'color' } = options;

  const config = map[value] || map[defaultKey] || {};
  const text = config[textKey] || value;
  const color = config[colorKey] || 'default';

  return (
    <Tag variant={variant} size={size} color={color}>
      {text}
    </Tag>
  );
};

// 状态标签映射表
const statusMap = {
  1: {
    text: '启用',
    color: 'green'
  },
  2: {
    text: '禁用',
    color: 'red'
  },
  default: {
    text: '未知',
    color: 'default'
  }
};

// 请求方式标签映射表
const methodMap = {
  GET: {
    text: 'GET',
    color: 'green'
  },
  POST: {
    text: 'POST',
    color: 'blue'
  },
  PUT: {
    text: 'PUT',
    color: 'purple'
  },
  DELETE: {
    text: 'DELETE',
    color: 'magenta'
  },
  PATCH: {
    text: 'PATCH',
    color: 'cyan'
  },
  default: {
    text: 'UNKNOWN',
    color: 'default'
  }
};

// 生成状态标签
export const GenerateStatusTag = (status) => {
  return GenerateTag(status, statusMap);
};

// 生成请求方式标签
export const GenerateMethodTag = (method) => {
  return GenerateTag(method, methodMap);
};
