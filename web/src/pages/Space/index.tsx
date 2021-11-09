import React from 'react';
import { history } from 'umi';

export default (): React.ReactNode => {
  // 解析url获取orgId以及spaceId（数组解构忽略前两个值）
  const [, , orgId, spaceId] = history.location.pathname.split("/");
  // 请求空间信息
  return (<>我是空间主页{orgId} / {spaceId}</>);
};
