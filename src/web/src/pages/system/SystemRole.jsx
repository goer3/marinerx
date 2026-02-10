import React, { useState } from 'react';
import { Helmet } from '@dr.pogodin/react-helmet';
import { TitleSuffix } from '@/components/Text';
import { PageHeaderBackgroundImage, TagLeftBlackIconImage } from '@/components/Image';

// 页面配置
const config = {
  title: '角色配置',
  enTitle: 'SYSTEM ROLE CONFIGURATION'
};

// 页面头部信息
const PageHeader = () => {
  return (
    <>
      <div className="dk-page-header" style={{ backgroundImage: `url(${PageHeaderBackgroundImage})` }}>
        <div className="dk-page-header-title">
          <img src={TagLeftBlackIconImage} alt="tag-left-black" />
          <span>{config.title + ' | ' + config.enTitle}</span>
        </div>
        <div className="dk-page-header-body">
          <p>
            系统权限基于 <span style={{ color: '#CC0033' }}>RBAC</span> 角色权限控制模型设计，管理员可以精确的对每个角色的每个接口进行权限控制。
          </p>
          <p>超级管理员是系统的最高权限角色，不可删除，无需单独设置，该角色默认会绕过所有权限控制，所以对于超级管理员的用户管理一定需要谨慎。</p>
        </div>
      </div>
    </>
  );
};

const SystemRole = () => {
  return (
    <>
      <Helmet>
        <title>{config.title + TitleSuffix}</title>
      </Helmet>
      <div>
        <PageHeader />
        <div className="dk-page-content">
          {/* 搜索 */}
          <div className="dk-page-search">
            <div className="dk-page-search-title">
              <span>搜索</span>
            </div>
          </div>
          {/* 主体 */}
          <div className="dk-page-main">
            {/* 操作按钮 */}
            <div className="dk-page-actions">
              <div className="dk-page-actions-left"></div>
              <div className="dk-page-actions-right"></div>
            </div>
            {/* 表格 */}
            <div className="dk-page-table"></div>
          </div>
        </div>
      </div>
    </>
  );
};

export default SystemRole;
