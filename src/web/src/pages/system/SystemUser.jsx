import React, { useState } from 'react';
import { Helmet } from '@dr.pogodin/react-helmet';
import { TitleSuffix } from '@/components/Text';
import { TagLeftBlackIconImage } from '@/components/Image';

// 页面配置
const config = {
  title: '用户配置',
  enTitle: 'SYSTEM USER CONFIGURATION'
};

// 页面头部信息
const PageHeader = () => {
  return (
    <>
      <div className="dk-page-header">
        <div className="dk-page-header-title">
          <img src={TagLeftBlackIconImage} alt="tag-left-black" />
          <span>{config.title + ' | ' + config.enTitle}</span>
        </div>
        <div className="dk-page-header-body">
          <p><span style={{ color: '#CC0033' }}>marinerx</span> 是默认用户，不可删除，且具有系统最高权限，所以在系统初始化完成之后一定要修改默认密码，并谨慎管理。</p>
          <p>目前用户的来源已经支持：本地创建，钉钉，飞书，企业微信扫码创建，管理员可以根据自身的需求开启。</p>
          <p>企业员工在第一次扫码的时会默认创建对应用户，并且根据手机号的唯一性创建相关关联，如果该手机号在系统中已经有绑定用户，则不会创建，而是使用存量用户登录。</p>
        </div>
      </div>
    </>
  );
};

const SystemUser = () => {
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
              <div className="dk-page-actions-left">
              </div>
              <div className="dk-page-actions-right">
              </div>
            </div>
            {/* 表格 */}
            <div className="dk-page-table">
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default SystemUser;
