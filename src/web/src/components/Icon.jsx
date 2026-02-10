import React from 'react';
import { ManOutlined, WomanOutlined, QuestionOutlined } from '@ant-design/icons';

// 右箭头图标
export const ArrowRightIcon = () => {
  return <i className="dk-right-arrow" />;
};

// 左箭头图标
export const ArrowLeftIcon = () => {
  return <i className="dk-left-arrow" />;
};

// 上箭头图标
export const ArrowUpIcon = () => {
  return <i className="dk-up-arrow" />;
};

// 下箭头图标
export const ArrowDownIcon = () => {
  return <i className="dk-down-arrow" />;
};

// 生成性别图标
export const GenerateGenderIcon = (gender) => {
  const icons = {
    1: <ManOutlined style={{ color: '#165dff' }} />,
    2: <WomanOutlined style={{ color: '#ff4d4f' }} />,
    default: <QuestionOutlined style={{ color: '#999999' }} />
  };
  return icons[gender] || icons.default;
};
